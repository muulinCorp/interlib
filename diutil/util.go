package diutil

import (
	"context"
	"errors"
	"net/http"

	"github.com/94peter/sterna"
	"github.com/94peter/sterna/auth"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/util"
)

func RedisCtxHandler(ctx context.Context, dbname string, exec func(redis db.RedisClient) (any, error)) (any, error) {
	di := sterna.GetDiByCtx(ctx)
	if di == nil {
		return nil, errors.New("di not found in request")
	}
	redisDI, ok := di.(db.RedisDI)
	if !ok {
		return nil, errors.New("config not set redis")
	}
	clt, err := redisDI.NewRedisClientDB(ctx, redisDI.GetDB(dbname))
	if err != nil {
		return nil, err
	}
	defer clt.Close()
	return exec(clt)
}

func RedisReqHandler(req *http.Request, dbname string, exec func(redis db.RedisClient) (any, error)) (any, error) {
	return RedisCtxHandler(req.Context(), dbname, exec)
}

func GetJwtDIByRequest(req *http.Request) auth.JwtDI {
	return util.GetCtxVal(req, sterna.CtxServDiKey).(auth.JwtDI)
}
