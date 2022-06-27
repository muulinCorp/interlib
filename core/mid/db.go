package mid

import (
	"net/http"
	"runtime"

	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/94peter/sterna/util"

	"github.com/google/uuid"
)

type DBMidDI interface {
	log.LoggerDI
	db.MongoDI
	db.RedisDI
}

type DBMiddle string

func NewDBMid(di DBMidDI, name string) mid.Middle {
	return &dbMiddle{
		name: name,
		di:   di,
	}
}

type dbMiddle struct {
	name string
	di   DBMidDI
}

func (lm *dbMiddle) GetName() string {
	return lm.name
}

func (am *dbMiddle) GetMiddleWare() func(f http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// one time scope setup area for middleware
		return func(w http.ResponseWriter, r *http.Request) {
			uuid := uuid.New().String()
			l := am.di.NewLogger(uuid)

			dbclt, err := am.di.NewMongoDBClient(r.Context(), "")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			defer dbclt.Close()

			// redisClt, err := am.di.NewRedisClient(r.Context())
			// if err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	w.Write([]byte(err.Error()))
			// 	return
			// }
			// defer redisClt.Close()

			// r = util.SetCtxKeyVal(r, db.CtxRedisKey, redisClt)
			r = util.SetCtxKeyVal(r, db.CtxMongoKey, dbclt)
			r = util.SetCtxKeyVal(r, log.CtxLogKey, l)
			f(w, r)

			runtime.GC()
		}
	}
}