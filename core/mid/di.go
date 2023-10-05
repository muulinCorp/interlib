package mid

import (
	"net/http"
	"reflect"

	"github.com/94peter/sterna"
	apiErr "github.com/94peter/sterna/api/err"
	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"
)

type DIMiddle string

func NewDiMid(clt db.RedisClient, env string, di interface{}) mid.Middle {
	return &diMiddle{
		clt: clt,
		env: env,
		di:  di,
	}
}

func NewGinDiMid(clt db.RedisClient, env string, di interface{}, service string) mid.GinMiddle {
	return &diMiddle{
		service: service,
		clt:     clt,
		env:     env,
		di:      di,
	}
}

type diMiddle struct {
	service string
	clt     db.RedisClient
	env     string
	di      interface{}
}

func (lm *diMiddle) outputErr(c *gin.Context, err error) {
	apiErr.GinOutputErr(c, lm.service, err)
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
			am.outputErr(c, apiErr.New(http.StatusInternalServerError, "missing X-Dikey"))
			c.Abort()
			return
		}

		grpcKey := c.GetHeader("X-GrpcKey")
		if grpcKey == "" {
			am.outputErr(c, apiErr.New(http.StatusInternalServerError, "missing X-GrpcKey"))
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
			am.outputErr(c, apiErr.New(http.StatusInternalServerError, err.Error()))
			c.Abort()
			return
		}
		sterna.InitConfByByte(confByte, newValue)
		c.Request = util.SetCtxKeyVal(c.Request, sterna.CtxServDiKey, newValue)

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
