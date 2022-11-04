package interceptor

import (
	"context"
	"reflect"

	"bitbucket.org/muulin/interlib"
	"bitbucket.org/muulin/interlib/channel"
	"github.com/94peter/sterna"
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
	IsEmpty() bool
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func getDI(md metadata.MD, clt db.RedisClient, env string, di DBMidDI) (DBMidDI, error) {
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

func getGrpcConf(md metadata.MD, clt db.RedisClient, env string) (interlib.GrpcRouterConf, error) {
	hosts := md.Get("X-GrpcKey")
	if len(hosts) != 1 {
		return nil, nil
	}

	confByte, err := clt.Get(hosts[0])
	if err != nil {
		return nil, err
	}
	grpConf := interlib.GrpcRouterConf{}
	grpConf.InitConfByByte(confByte)

	return grpConf, nil
}

func getContextWitchRsrc(ctx context.Context, di DBMidDI) (r *rsrc, err error) {
	uuid := uuid.New().String()
	l := di.NewLogger(uuid)
	dbclt, err := di.NewMongoDBClient(ctx, "")
	if err != nil {
		return
	}
	r = &rsrc{
		l:     l,
		dbclt: dbclt,
		di:    di,
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
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.InvalidArgument, "can not get metadata")
		}
		di, err := getDI(md, ri.clt, ri.env, ri.di)
		if err != nil {
			return err
		}
		rsrc, err := getContextWitchRsrc(ctx, di)
		if err != nil {
			return err
		}

		grpc, err := getGrpcConf(md, ri.clt, ri.env)
		if err != nil {
			return err
		}
		if grpc != nil {
			ctx = context.WithValue(ctx, interlib.CtxGrpcConfKey, grpc)
		}
		err = handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          rsrc.setContext(ctx),
		})
		if err != nil {
			return err
		}
		rsrc.close()
		return
	})
}

func (ri *redisInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "can not get metadata")
		}
		di, err := getDI(md, ri.clt, ri.env, ri.di)
		if err != nil {
			return nil, err
		}
		r, err := getContextWitchRsrc(ctx, di)
		if err != nil {
			return nil, err
		}
		defer r.close()
		grpc, err := getGrpcConf(md, ri.clt, ri.env)
		if err != nil {
			return nil, err
		}
		if grpc != nil {
			ctx = context.WithValue(ctx, interlib.CtxGrpcConfKey, grpc)
		}
		return handler(r.setContext(ctx), req)
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
		ctx := context.WithValue(ss.Context(), db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		return handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
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
		ctx = context.WithValue(ctx, db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		ctx = context.WithValue(ctx, sterna.CtxServDiKey, i.di)
		return handler(ctx, req)
	})
}

type rsrc struct {
	dbclt db.MongoDBClient
	l     log.Logger
	di    DBMidDI
}

func (r *rsrc) close() {
	if r.dbclt != nil {
		r.dbclt.Close()
	}
}

func (r *rsrc) setContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, db.CtxMongoKey, r.dbclt)
	ctx = context.WithValue(ctx, log.CtxLogKey, r.l)
	ctx = context.WithValue(ctx, sterna.CtxServDiKey, r.di)
	return ctx
}

func NewFixedDBInterceptor(grpc channel.ChannelClient, host, env string, di DBMidDI) Interceptor {
	return &fixedDBInterceptor{
		host:    host,
		env:     env,
		di:      di,
		chaGrpc: grpc,
	}
}

type fixedDBInterceptor struct {
	host, env string
	di        DBMidDI
	chaGrpc   channel.ChannelClient
}

func (f *fixedDBInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		if f.di.IsEmpty() {
			conf, err := f.chaGrpc.GetConf(f.host, f.env)
			if err != nil {
				return err
			}
			sterna.InitConfByByte(conf, f.di)
		}
		uuid := uuid.New().String()
		l := f.di.NewLogger(uuid)
		dbclt, err := f.di.NewMongoDBClient(ss.Context(), "")
		if err != nil {
			return err
		}
		defer dbclt.Close()
		ctx := context.WithValue(ss.Context(), db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		return handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
	})
}

func (f *fixedDBInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if f.di.IsEmpty() {
			conf, err := f.chaGrpc.GetConf(f.host, f.env)
			if err != nil {
				return nil, err
			}
			sterna.InitConfByByte(conf, f.di)
		}
		uuid := uuid.New().String()
		l := f.di.NewLogger(uuid)
		dbclt, err := f.di.NewMongoDBClient(ctx, "")
		if err != nil {
			return nil, err
		}
		defer dbclt.Close()
		ctx = context.WithValue(ctx, db.CtxMongoKey, dbclt)
		ctx = context.WithValue(ctx, log.CtxLogKey, l)
		ctx = context.WithValue(ctx, sterna.CtxServDiKey, f.di)
		return handler(ctx, req)
	})
}
