package core

import (
	"context"
	"runtime"

	"github.com/94peter/api-toolkit/errors"
	"github.com/94peter/api-toolkit/mid"
	"github.com/94peter/di"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/muulinCorp/interlib/core/interceptor"
	pkgErr "github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ModelCfg interface {
	Close() error
	Init(uuid string, di di.DI) error
	Copy() ModelCfg
}

type ctxType string

const cfgKey = "model_config"

func GetFromGinCtx[T ModelCfg](ctx *gin.Context) (T, bool) {
	var result T
	val, ok := ctx.Get(cfgKey)
	if !ok {
		return result, false
	}
	return val.(T), true
}

func setToGinCtx[T ModelCfg](ctx *gin.Context, cfg T) {
	ctx.Set(cfgKey, cfg)
}

func GetFromCtx[T ModelCfg](ctx context.Context) (T, bool) {
	var result T
	val := ctx.Value(ctxType(cfgKey))
	if val == nil {
		return result, false
	}
	return val.(T), true
}
func setToCtx[T ModelCfg](ctx context.Context, cfg T) context.Context {
	return context.WithValue(ctx, ctxType(cfgKey), cfg)
}

type ModelCfgMgr interface {
	mid.GinMiddle
	interceptor.Interceptor
}

type modelCfgMgr[T ModelCfg] struct {
	cfg T
	errors.CommonApiErrorHandler
}

func NewFixModelCfgGinMid[T ModelCfg](cfg T) ModelCfgMgr {
	return &modelCfgMgr[T]{
		cfg: cfg,
	}
}

func (m *modelCfgMgr[T]) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := m.cfg.Copy()
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
		if err := data.Init(uuid.New().String(), servDi); err != nil {
			m.GinApiErrorHandler(c, err)
			c.Abort()
			return
		}

		setToGinCtx(c, data)
		c.Next()
		data.Close()
		runtime.GC()
	}
}

func (m *modelCfgMgr[T]) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		if interceptor.IsReflectMethod(info.FullMethod) {
			return handler(srv, ss)
		}
		ctx := ss.Context()
		data := m.cfg.Copy()
		servDi := di.GetDiFromCtx(ctx)
		if servDi == nil {
			return pkgErr.New("can not get di")
		}
		if err := servDi.IsConfEmpty(); err != nil {
			return err
		}
		if err := data.Init(uuid.New().String(), servDi); err != nil {
			return err
		}

		defer func() {
			data.Close()
			runtime.GC()
		}()

		return handler(srv, interceptor.NewServerStream(
			setToCtx(ctx, data), ss))

	})
}

func (m *modelCfgMgr[T]) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		data := m.cfg.Copy()
		servDi := di.GetDiFromCtx(ctx)
		if servDi == nil {
			return nil, pkgErr.New("can not get di")
		}
		if err := servDi.IsConfEmpty(); err != nil {
			return nil, err
		}
		if err := data.Init(uuid.New().String(), servDi); err != nil {
			return nil, err
		}
		defer func() {
			data.Close()
			runtime.GC()
		}()
		return handler(setToCtx(ctx, data), req)
	})
}
