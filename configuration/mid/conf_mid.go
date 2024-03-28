package mid

import (
	"net/http"
	"reflect"
	"time"

	"github.com/muulinCorp/interlib/channel"
	"github.com/muulinCorp/interlib/configuration/client"
	"github.com/muulinCorp/interlib/configuration/pb"

	apiErr "github.com/94peter/api-toolkit/errors"
	"github.com/94peter/api-toolkit/mid"
	"github.com/94peter/micro-service/di"
	"github.com/gin-gonic/gin"
)

func NewGinInterConfMid(address string, di channel.DI) (mid.GinMiddle, error) {
	confSDK, err := client.New(address)
	if err != nil {
		return nil, err
	}
	return &interConfMiddle{
		di:        di,
		confSDK:   confSDK,
		confCache: map[string]*cacheData{},
	}, nil
}

func (lm *interConfMiddle) GetName() string {
	return "configuration"
}

type cacheData struct {
	di  channel.DI
	exp time.Time
}

type interConfMiddle struct {
	apiErr.CommonApiErrorHandler
	di      channel.DI
	confSDK client.ConfigurationClient

	confCache map[string]*cacheData
}

func (am *interConfMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		xchannel := c.GetHeader("X-Channel")
		if xchannel == "" {
			xchannel, _ = c.GetQuery("X-Channel")
		}
		isCleaner := c.GetHeader("X-Config-Clear")
		if xchannel == "" {
			c.Next()
			return
		}

		if cache, ok := am.confCache[xchannel]; ok && isCleaner != "true" && time.Until(cache.exp) > 0 {
			di.SetDiToGin(c, cache.di)
		} else {
			confByte, err := am.confSDK.GetChannelConf(c, &pb.GetConfRequest{
				ChannelName: xchannel,
				Version:     "latest",
			})
			if err != nil {
				am.GinApiErrorHandler(c, apiErr.New(http.StatusInternalServerError, err.Error()))
				c.Abort()
				return
			}

			val := reflect.ValueOf(am.di)
			if val.Kind() == reflect.Ptr {
				val = reflect.Indirect(val)
			}
			newValue := reflect.New(val.Type()).Interface().(channel.DI)
			di.InitConfByByte(confByte, newValue)
			if _, ok := am.confCache[xchannel]; ok {
				am.confCache[xchannel].di = newValue
				am.confCache[xchannel].di.SetChannel(xchannel)
				am.confCache[xchannel].exp = time.Now().Add(time.Hour)
			} else {
				am.confCache[xchannel] = &cacheData{
					di:  newValue,
					exp: time.Now().Add(time.Hour),
				}
				am.confCache[xchannel].di.SetChannel(xchannel)
			}
			di.SetDiToGin(c, am.confCache[xchannel].di)
		}
		c.Next()
	}
}
