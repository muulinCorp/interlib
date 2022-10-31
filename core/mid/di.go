package mid

import (
	"net/http"
	"reflect"

	"bitbucket.org/muulin/interlib"
	"github.com/94peter/sterna"
	"github.com/94peter/sterna/api"
	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"
)

type DIMiddle string

func NewDiMid(clt db.RedisClient, env string, di interface{}) mid.Middle {
	return &diMiddle{
		clt:        clt,
		env:        env,
		di:         di,
		routerConf: interlib.GrpcRouterConf{},
	}
}

func NewGinDiMid(clt db.RedisClient, env string, di interface{}, service string) mid.GinMiddle {
	return &diMiddle{
		service:    service,
		clt:        clt,
		env:        env,
		di:         di,
		routerConf: interlib.GrpcRouterConf{},
	}
}

type diMiddle struct {
	service    string
	clt        db.RedisClient
	env        string
	di         interface{}
	routerConf interlib.GrpcRouterConf
}

func (lm *diMiddle) outputErr(c *gin.Context, err error) {
	api.GinOutputErr(c, lm.service, err)
}

func (lm *diMiddle) GetName() string {
	return "di"
}

// const path = "/internal/v1/channel/conf/"

func (am *diMiddle) GetMiddleWare() func(f http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// one time scope setup area for middleware
		return func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-DiKey")
			if key == "" {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("missing X-Dikey"))
				return
			}
			val := reflect.ValueOf(am.di)
			if val.Kind() == reflect.Ptr {
				val = reflect.Indirect(val)
			}
			newValue := reflect.New(val.Type()).Interface()

			confByte, err := am.clt.Get(key)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			sterna.InitConfByByte(confByte, newValue)
			r = util.SetCtxKeyVal(r, sterna.CtxServDiKey, newValue)
			f(w, r)
		}
	}
}

func (am *diMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-DiKey")
		if key == "" {
			am.outputErr(c, api.NewApiError(http.StatusInternalServerError, "missing X-Dikey"))
			c.Abort()
			return
		}

		grpcKey := c.GetHeader("X-GrpcKey")
		if grpcKey == "" {
			am.outputErr(c, api.NewApiError(http.StatusInternalServerError, "missing X-GrpcKey"))
			c.Abort()
			return
		}

		val := reflect.ValueOf(am.di)
		if val.Kind() == reflect.Ptr {
			val = reflect.Indirect(val)
		}
		newValue := reflect.New(val.Type()).Interface()
		confByte, err := am.clt.Get(key)
		if err != nil {
			am.outputErr(c, api.NewApiError(http.StatusInternalServerError, err.Error()))
			c.Abort()
			return
		}
		sterna.InitConfByByte(confByte, newValue)
		c.Request = util.SetCtxKeyVal(c.Request, sterna.CtxServDiKey, newValue)

		grpcConfByte, err := am.clt.Get(grpcKey)
		if err != nil {
			am.outputErr(c, api.NewApiError(http.StatusInternalServerError, err.Error()))
			c.Abort()
			return
		}
		am.routerConf.InitConfByByte(grpcConfByte)
		c.Request = util.SetCtxKeyVal(c.Request, interlib.CtxGrpcConfKey, am.routerConf)
		c.Next()
	}
}

func GetGinHost(c *gin.Context) string {
	host := c.GetHeader("X-Forwarded-Host")
	if host == "" {
		host = c.Request.Host
	}
	return host
}
