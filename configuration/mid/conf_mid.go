package mid

import (
	"net/http"
	"reflect"
	"time"

	"bitbucket.org/muulin/interlib/configuration/client"
	"bitbucket.org/muulin/interlib/configuration/pb"

	"github.com/94peter/sterna"
	apiErr "github.com/94peter/sterna/api/err"
	sternaMid "github.com/94peter/sterna/api/mid"
	"github.com/gin-gonic/gin"
)

func NewGinInterConfMid(address string, di sterna.ChannelDI) (sternaMid.GinMiddle, error) {
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
	di  sterna.ChannelDI
	exp time.Time
}

type interConfMiddle struct {
	di      sterna.CommonDI
	confSDK client.ConfigurationClient

	confCache map[string]*cacheData
}

func (am *interConfMiddle) outputErr(c *gin.Context, err error) {
	apiErr.GinOutputErr(c, am.di.GetServiceName(), err)
}

func (am *interConfMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := c.GetHeader("X-Channel")
		if channel == "" {
			channel, _ = c.GetQuery("X-Channel")
		}
		isCleaner := c.GetHeader("X-Config-Clear")
		if channel == "" {
			c.Next()
			return
		}

		if cache, ok := am.confCache[channel]; ok && isCleaner != "true" && cache.exp.Sub(time.Now()) > 0 {
			c.Set(string(sterna.CtxServDiKey), cache.di)
		} else {
			confByte, err := am.confSDK.GetChannelConf(c, &pb.GetConfRequest{
				ChannelName: channel,
				Version:     "latest",
			})
			if err != nil {
				am.outputErr(c, apiErr.New(http.StatusInternalServerError, err.Error()))
				c.Abort()
				return
			}

			val := reflect.ValueOf(am.di)
			if val.Kind() == reflect.Ptr {
				val = reflect.Indirect(val)
			}
			newValue := reflect.New(val.Type()).Interface()
			sterna.InitConfByByte(confByte, newValue)
			if _, ok := am.confCache[channel]; ok {
				am.confCache[channel].di = newValue.(sterna.ChannelDI)
				am.confCache[channel].di.SetChannel(channel)
				am.confCache[channel].exp = time.Now().Add(time.Hour)
			} else {
				am.confCache[channel] = &cacheData{
					di:  newValue.(sterna.ChannelDI),
					exp: time.Now().Add(time.Hour),
				}
				am.confCache[channel].di.SetChannel(channel)
			}
			c.Set(string(sterna.CtxServDiKey), am.confCache[channel].di)
		}
		c.Next()
	}
}
