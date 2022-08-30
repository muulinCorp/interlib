package mid

import (
	"net/http"
	"reflect"

	"github.com/94peter/sterna"
	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"
)

var (
	serviceDiMap = make(map[string]interface{})
)

const (
	CtxServDiKey = util.CtxKey("ServiceDI")
)

type DIMiddle string

func NewDiMid(proto string, env string, di interface{}) mid.Middle {
	return &diMiddle{
		proto: proto,
		env:   env,
		di:    di,
	}
}

func NewGinDiMid(proto string, env string, di interface{}) mid.GinMiddle {
	return &diMiddle{
		proto: proto,
		env:   env,
		di:    di,
	}
}

type diMiddle struct {
	proto string
	env   string
	di    interface{}
}

func (lm *diMiddle) GetName() string {
	return "di"
}

const path = "/internal/v1/channel/conf/"

func (am *diMiddle) GetMiddleWare() func(f http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// one time scope setup area for middleware
		return func(w http.ResponseWriter, r *http.Request) {
			host := util.GetHost(r)
			var mydi interface{}
			var ok bool
			if mydi, ok = serviceDiMap[host]; !ok {
				val := reflect.ValueOf(am.di)
				if val.Kind() == reflect.Ptr {
					val = reflect.Indirect(val)
				}
				newValue := reflect.New(val.Type()).Interface()
				uri := util.StrAppend(am.proto, "://", host, path, am.env)
				err := sterna.InitConfByUri(uri, newValue)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				serviceDiMap[host] = newValue
				mydi = newValue
			}
			r = util.SetCtxKeyVal(r, CtxServDiKey, mydi)
			f(w, r)
		}
	}
}

func (m *diMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := GetGinHost(c)
		var mydi interface{}
		var ok bool
		if mydi, ok = serviceDiMap[host]; !ok {
			val := reflect.ValueOf(m.di)
			if val.Kind() == reflect.Ptr {
				val = reflect.Indirect(val)
			}
			newValue := reflect.New(val.Type()).Interface()
			uri := util.StrAppend(m.proto, "://", host, path, m.env)
			err := sterna.InitConfByUri(uri, newValue)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			serviceDiMap[host] = newValue
			mydi = newValue
		}
		c.Request = util.SetCtxKeyVal(c.Request, CtxServDiKey, mydi)
		//c.Set(string(CtxServDiKey), mydi)
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
