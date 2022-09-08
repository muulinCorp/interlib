package interceptor

import (
	"context"
	"reflect"

	"bitbucket.org/muulin/interlib/channel"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
)

var (
	serviceDiMap = make(map[string]DBMidDI)
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

func getDI(ctx context.Context, clt channel.ChannelClient, env string, di DBMidDI) (DBMidDI, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can not get metadata")
	}
	hosts := md.Get("X-Channel")
	if len(hosts) != 1 {
		return nil, nil
	}

	var mydi DBMidDI
	if mydi, ok = serviceDiMap[hosts[0]]; ok {
		return mydi, nil
	}

	confByte, err := clt.GetConf(hosts[0], env)
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
	serviceDiMap[hosts[0]] = newValue.(DBMidDI)
	return serviceDiMap[hosts[0]], nil
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
	}
	return
}

func StreamServerDBInterceptor(clt channel.ChannelClient, env string, di DBMidDI) grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		ctx := ss.Context()
		di, err := getDI(ctx, clt, env, di)
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

func UnaryServerDBInterceptor(clt channel.ChannelClient, env string, di DBMidDI) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		di, err := getDI(ctx, clt, env, di)
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

type rsrc struct {
	dbclt    db.MongoDBClient
	redisClt db.RedisClient
	l        log.Logger
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
	return ctx
}
