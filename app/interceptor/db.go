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
	hosts := md.Get("X-Host")
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

func getContextWitchRsrc(ctx context.Context, di DBMidDI) (context.Context, error) {
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
	return ctx, nil
}

func StreamServerDBInterceptor(clt channel.ChannelClient, env string, di DBMidDI) grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {

		di, err := getDI(ss.Context(), clt, env, di)
		if err != nil {
			return err
		}
		ctx, err := getContextWitchRsrc(ss.Context(), di)
		if err != nil {
			return err
		}
		err = handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
		return
	})
}

func UnaryServerDBInterceptor(clt channel.ChannelClient, env string, di DBMidDI) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		di, err := getDI(ctx, clt, env, di)
		if err != nil {
			return nil, err
		}
		ctx, err = getContextWitchRsrc(ctx, di)
		if err != nil {
			return nil, err
		}
		resp, err := handler(ctx, req)
		return resp, err
	})
}
