package diutil

import (
	"context"
	"errors"
	"net/http"

	"github.com/94peter/cache/conn"
	"github.com/94peter/di"
	"github.com/gin-gonic/gin"
)

func RedisGinHandler(c *gin.Context, dbname string, exec func(clt conn.RedisClient) (any, error)) (any, error) {
	di := di.GetDiFromGin(c)
	if di == nil {
		return nil, errors.New("di not found in request")
	}
	redisDI, ok := di.(conn.RedisDI)
	if !ok {
		return nil, errors.New("config not set redis")
	}
	clt, err := redisDI.NewRedisDbConn(c.Request.Context(), dbname)
	if err != nil {
		return nil, err
	}
	defer clt.Close()
	return exec(clt)
}
func RedisCtxHandler(ctx context.Context, dbname string, exec func(clt conn.RedisClient) (any, error)) (any, error) {
	di := di.GetDiFromCtx(ctx)
	if di == nil {
		return nil, errors.New("di not found in request")
	}
	redisDI, ok := di.(conn.RedisDI)
	if !ok {
		return nil, errors.New("config not set redis")
	}
	clt, err := redisDI.NewRedisDbConn(ctx, dbname)
	if err != nil {
		return nil, err
	}
	defer clt.Close()
	return exec(clt)
}

func RedisReqHandler(req *http.Request, dbname string, exec func(clt conn.RedisClient) (any, error)) (any, error) {
	return RedisCtxHandler(req.Context(), dbname, exec)
}
