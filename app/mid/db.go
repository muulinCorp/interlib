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

func NewGinDBMid() mid.GinMiddle {
	return &dbMiddle{}
}

type dbMiddle struct {
}

func (lm *dbMiddle) GetName() string {
	return "db"
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

			// redisClt, err := am.di.NewRedisClient(r.Context())
			// if err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	w.Write([]byte(err.Error()))
			// 	return
			// }
			// defer redisClt.Close()

			// r = util.SetCtxKeyVal(r, db.CtxRedisKey, redisClt)

		}
	}
}

func (m *dbMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		servDi, exist := c.Get(string(CtxServDiKey))
		// servDi := util.GetCtxVal(r, CtxServDiKey)
		if !exist {
			api.GinOutputErr(c, api.NewApiError(http.StatusInternalServerError, "can not get di"))
			return
		}
		if dbdi, ok := servDi.(DBMidDI); ok {
			uuid := uuid.New().String()
			l := dbdi.NewLogger(uuid)
			//c.Set(string(log.CtxLogKey), l)

			dbclt, err := dbdi.NewMongoDBClient(c.Request.Context(), "")
			if err != nil {
				api.GinOutputErr(c, api.NewApiError(http.StatusInternalServerError, err.Error()))
				return
			}
			defer dbclt.Close()
			//c.Set(string(db.CtxMongoKey), dbclt)
			c.Request = util.SetCtxKeyVal(c.Request, db.CtxMongoKey, dbclt)
   			c.Request = util.SetCtxKeyVal(c.Request, log.CtxLogKey, l)

			c.Next()
			runtime.GC()
		} else {
			api.GinOutputErr(c, api.NewApiError(http.StatusInternalServerError, "invalid di"))
			return
		}
	}
}
