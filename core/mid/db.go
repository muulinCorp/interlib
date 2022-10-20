package mid

import (
	"net/http"
	"runtime"

	"github.com/94peter/sterna/api"
	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type DBMidDI interface {
	log.LoggerDI
	db.MongoDI
	db.RedisDI
}

type DBMiddle string

func NewDBMid() mid.Middle {
	return &dbMiddle{}
}

func NewGinDBMid(service string) mid.GinMiddle {
	return &dbMiddle{
		service: service,
	}
}

type dbMiddle struct {
	service string
}

func (lm *dbMiddle) GetName() string {
	return "db"
}

func (lm *dbMiddle) outputErr(c *gin.Context, err error) {
	api.GinOutputErr(c, lm.service, err)
}

func (am *dbMiddle) GetMiddleWare() func(f http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// one time scope setup area for middleware
		return func(w http.ResponseWriter, r *http.Request) {
			servDi := util.GetCtxVal(r, CtxServDiKey)
			if servDi == nil {
				api.OutputErr(w, api.NewApiError(http.StatusInternalServerError, "can not get di"))
				return
			}

			if dbdi, ok := servDi.(DBMidDI); ok {
				uuid := uuid.New().String()
				l := dbdi.NewLogger(uuid)

				dbclt, err := dbdi.NewMongoDBClient(r.Context(), "")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				defer dbclt.Close()
				redisClt, err := dbdi.NewRedisClient(r.Context())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				defer redisClt.Close()
				r = util.SetCtxKeyVal(r, db.CtxMongoKey, dbclt)
				r = util.SetCtxKeyVal(r, log.CtxLogKey, l)
				r = util.SetCtxKeyVal(r, db.CtxRedisKey, redisClt)
				f(w, r)
				runtime.GC()
			} else {
				api.OutputErr(w, api.NewApiError(http.StatusInternalServerError, "invalid di"))
				return
			}
		}
	}
}

func (m *dbMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		servDi := util.GetCtxVal(c.Request, CtxServDiKey)
		if servDi == nil {
			m.outputErr(c, api.NewApiError(http.StatusInternalServerError, "can not get di"))
			c.Abort()
			return
		}

		if dbdi, ok := servDi.(DBMidDI); ok {
			uuid := uuid.New().String()
			l := dbdi.NewLogger(uuid)

			dbclt, err := dbdi.NewMongoDBClient(c.Request.Context(), "")
			if err != nil {
				m.outputErr(c, api.NewApiError(http.StatusInternalServerError, err.Error()))
				c.Abort()
				return
			}
			defer dbclt.Close()

			redisClt, err := dbdi.NewRedisClient(c.Request.Context())
			if err != nil {
				m.outputErr(c, api.NewApiError(http.StatusInternalServerError, err.Error()))
				c.Abort()
				return
			}
			defer redisClt.Close()

			c.Request = util.SetCtxKeyVal(c.Request, db.CtxMongoKey, dbclt)
			c.Request = util.SetCtxKeyVal(c.Request, log.CtxLogKey, l)
			c.Request = util.SetCtxKeyVal(c.Request, db.CtxRedisKey, redisClt)

			c.Next()
			runtime.GC()
		} else {
			m.outputErr(c, api.NewApiError(http.StatusInternalServerError, "invalid di"))
			c.Abort()
			return
		}
	}
}
