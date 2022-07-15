package interceptor

import (
	"context"

	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type DBMidDI interface {
	log.LoggerDI
	db.MongoDI
	db.RedisDI
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func StreamServerDBInterceptor(di DBMidDI) grpc.ServerOption {
	return grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		uuid := uuid.New().String()
		l := di.NewLogger(uuid)
		dbclt, err := di.NewMongoDBClient(ss.Context(), "")
		if err != nil {
			return err
		}
		defer dbclt.Close()
		redisClt, err := di.NewRedisClient(ss.Context())
		if err != nil {
			return err
		}
		defer redisClt.Close()
		ctx := context.WithValue(ss.Context(), db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		ctx = context.WithValue(ctx, db.CtxRedisKey, redisClt)
		err = handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
		return
	})
}

func UnaryServerDBInterceptor(di DBMidDI) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		uuid := uuid.New().String()
		l := di.NewLogger(uuid)
		dbclt, err := di.NewMongoDBClient(ctx, "")
		if err != nil {
			return nil, err
		}
		defer dbclt.Close()
		redisClt, err := di.NewRedisClient(ctx)
		if err != nil {
			return nil, err
		}
		defer redisClt.Close()
		ctx = context.WithValue(ctx, db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		ctx = context.WithValue(ctx, db.CtxRedisKey, redisClt)
		resp, err := handler(ctx, req)
		return resp, err
	})
}
