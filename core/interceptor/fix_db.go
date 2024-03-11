package interceptor

import (
	"context"

	"github.com/94peter/di"
	"github.com/94peter/log"
	"github.com/94peter/morm/conn"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ConnMidDI interface {
	conn.MongoDI
	log.LoggerDI
	di.DI
}

func NewFixDBInterceptor(di ConnMidDI) Interceptor {
	return &fixDBInterceptor{
		ConnMidDI: di,
	}
}

type fixDBInterceptor struct {
	ConnMidDI
}

func (i *fixDBInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		if IsReflectMethod(info.FullMethod) {
			return handler(srv, ss)
		}
		uuid := uuid.New().String()
		l, err := i.NewLogger(i.GetService(), uuid)
		if err != nil {
			return err
		}
		dbclt, err := i.NewDefaultDbConn(ss.Context())
		if err != nil {
			return err
		}
		defer dbclt.Close()

		ctx := conn.SetMgoDbConnToCtx(ss.Context(), dbclt)
		ctx = log.SetByCtx(ctx, l)
		ctx = di.SetDiToCtx(ctx, i.ConnMidDI)
		return handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
	})
}

func (i *fixDBInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		uuid := uuid.New().String()
		l, err := i.NewLogger(i.GetService(), uuid)
		if err != nil {
			return nil, err
		}
		dbclt, err := i.NewDefaultDbConn(ctx)
		if err != nil {
			return nil, err
		}
		defer dbclt.Close()

		ctx = conn.SetMgoDbConnToCtx(ctx, dbclt)
		ctx = log.SetByCtx(ctx, l)
		ctx = di.SetDiToCtx(ctx, i.ConnMidDI)
		return handler(ctx, req)
	})
}
