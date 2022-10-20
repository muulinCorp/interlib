package interceptor

import (
	"context"
	"reflect"

	"bitbucket.org/muulin/interlib/core/mid"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
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

func getDI(ctx context.Context, clt db.RedisClient, env string, di DBMidDI) (DBMidDI, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can not get metadata")
	}
	hosts := md.Get("X-DiKey")
	if len(hosts) != 1 {
		return nil, status.Error(codes.InvalidArgument, "missing X-DiKey")
	}

	confByte, err := clt.Get(hosts[0])
	if err != nil {
		return nil, err
	}
	val := reflect.ValueOf(di)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	newValue := reflect.New(val.Type()).Interface()
	err = yaml.Unmarshal(confByte, newValue)
	if err != nil {
		return nil, err
	}

	return newValue.(DBMidDI), nil
}

func getContextWitchRsrc(ctx context.Context, di DBMidDI) (r *rsrc, err error) {
	uuid := uuid.New().String()
	l := di.NewLogger(uuid)
	dbclt, err := di.NewMongoDBClient(ctx, "")
	if err != nil {
		return
	}
	redisClt, err := di.NewRedisClient(ctx)
	if err != nil {
		return
	}
	r = &rsrc{
		l:        l,
		dbclt:    dbclt,
		redisClt: redisClt,
		di:       di,
	}
	return
}

type Interceptor interface {
	StreamServerInterceptor() grpc.StreamServerInterceptor
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
}

func NewRedisDBInterceptor(clt db.RedisClient, env string, di DBMidDI) Interceptor {
	return &redisInterceptor{
		clt: clt,
		env: env,
		di:  di,
	}
}

type redisInterceptor struct {
	clt db.RedisClient
	env string
	di  DBMidDI
}

func (ri *redisInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		ctx := ss.Context()
		di, err := getDI(ctx, ri.clt, ri.env, ri.di)
		if err != nil {
			return err
		}
		rsrc, err := getContextWitchRsrc(ctx, di)
		if err != nil {
			return err
		}

		err = handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          rsrc.setContext(ctx),
		})
		rsrc.close()
		return
	})
}

func (ri *redisInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		di, err := getDI(ctx, ri.clt, ri.env, ri.di)
		if err != nil {
			return nil, err
		}
		r, err := getContextWitchRsrc(ctx, di)
		if err != nil {
			return nil, err
		}
		resp, err := handler(r.setContext(ctx), req)
		r.close()
		return resp, err
	})
}

func NewLocalDBInterceptor(di DBMidDI) Interceptor {
	return &localDBInterceptor{
		di: di,
	}
}

type localDBInterceptor struct {
	di DBMidDI
}

func (i *localDBInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		uuid := uuid.New().String()
		l := i.di.NewLogger(uuid)
		dbclt, err := i.di.NewMongoDBClient(ss.Context(), "")
		if err != nil {
			return err
		}
		defer dbclt.Close()
		redisClt, err := i.di.NewRedisClient(ss.Context())
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

func (i *localDBInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		uuid := uuid.New().String()
		l := i.di.NewLogger(uuid)
		dbclt, err := i.di.NewMongoDBClient(ctx, "")
		if err != nil {
			return nil, err
		}
		defer dbclt.Close()
		redisClt, err := i.di.NewRedisClient(ctx)
		if err != nil {
			return nil, err
		}
		defer redisClt.Close()
		ctx = context.WithValue(ctx, db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		ctx = context.WithValue(ctx, db.CtxRedisKey, redisClt)
		ctx = context.WithValue(ctx, mid.CtxServDiKey, i.di)
		resp, err := handler(ctx, req)
		return resp, err
	})
}

type rsrc struct {
	dbclt    db.MongoDBClient
	redisClt db.RedisClient
	l        log.Logger
	di       DBMidDI
}

func (r *rsrc) close() {
	if r.dbclt != nil {
		r.dbclt.Close()
	}
	if r.redisClt != nil {
		r.redisClt.Close()
	}
}

func (r *rsrc) setContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, db.CtxMongoKey, r.dbclt)
	ctx = context.WithValue(ctx, log.CtxLogKey, r.l)
	ctx = context.WithValue(ctx, db.CtxRedisKey, r.redisClt)
	ctx = context.WithValue(ctx, mid.CtxServDiKey, r.di)
	return ctx
}
