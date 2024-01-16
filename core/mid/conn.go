package mid

import (
	"runtime"

	"github.com/94peter/api-toolkit/errors"
	"github.com/94peter/api-toolkit/mid"
	"github.com/94peter/di"
	pkgErr "github.com/pkg/errors"

	"github.com/94peter/log"
	"github.com/94peter/morm/conn"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type ConnMidDI interface {
	log.LoggerDI
	conn.MongoOptsDI
	di.DI
}

type ConnMiddle string

func NewGinConnMid() mid.GinMiddle {
	return &connMiddle{}
}

type connMiddle struct {
	errors.CommonApiErrorHandler
}

func (m *connMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		servDi := di.GetDiFromGin(c)
		if servDi == nil {
			m.GinApiErrorHandler(c, pkgErr.New("can not get di"))
			c.Abort()
			return
		}
		if err := servDi.IsConfEmpty(); err != nil {
			m.GinApiErrorHandler(c, err)
			c.Abort()
			return
		}

		if dbdi, ok := servDi.(ConnMidDI); ok {
			uuid := uuid.New().String()
			l, err := dbdi.NewLogger(dbdi.GetService(), uuid)
			if err != nil {
				m.GinApiErrorHandler(c, err)
				c.Abort()
				return
			}

			dbclt, err := dbdi.NewDefaultDbConnWithOpts(c.Request.Context())
			if err != nil {
				m.GinApiErrorHandler(c, err)
				c.Abort()
				return
			}
			defer dbclt.Close()

			conn.SetMgoDbConnToGin(c, dbclt)
			log.SetByGinCtx(c, l)
			c.Next()
			runtime.GC()
		} else {
			m.GinApiErrorHandler(c, pkgErr.New("invalid di"))
			c.Abort()
			return
		}
	}
}
