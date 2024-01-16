package interlib

import (
	"net/http"

	"github.com/94peter/log"

	apiErr "github.com/94peter/api-toolkit/errors"
	"github.com/gin-gonic/gin"
)

var (
	ServerErrorHandler = func(c *gin.Context, service string, err error) {
		if err == nil {
			return
		}

		l := log.GetByGinCtx(c)
		if l != nil {
			l.WarnPkg(err)
		}
		if apiErr, ok := err.(apiErr.ApiError); ok {
			c.AbortWithStatusJSON(apiErr.GetStatus(),
				map[string]interface{}{
					"status":  apiErr.GetStatus(),
					"error":   apiErr.Error(),
					"service": service,
				})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"title":   err.Error(),
					"service": service,
				})
		}
	}
)
