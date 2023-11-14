package interceptor

import (
	"context"
	"reflect"
	"time"

	"bitbucket.org/muulin/interlib/configuration/client"
	coreInterceptor "bitbucket.org/muulin/interlib/core/interceptor"

	"bitbucket.org/muulin/interlib/configuration/pb"

	"github.com/94peter/sterna"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type dbMidDI interface {
	log.LoggerDI
	db.MongoDI
	db.RedisDI
	sterna.CommonDI
	SetChannel(string)
	GetChannel() string
}

type cacheData struct {
	di  dbMidDI
	exp time.Time
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func (ri *interConfInterceptor) getDI(ctx context.Context, channel string) (dbMidDI, error) {

	if cache, ok := ri.confCache[channel]; ok && time.Until(cache.exp) > 0 {
		return cache.di, nil
	}

	confByte, err := ri.confSDK.GetChannelConf(ctx, &pb.GetConfRequest{
		ChannelName: channel,
		Version:     "latest",
	})
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(ri.di)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	newValue := reflect.New(val.Type()).Interface()
	sterna.InitConfByByte(confByte, newValue)
	dbdi := newValue.(dbMidDI)
	if _, ok := ri.confCache[channel]; ok {
		ri.confCache[channel].di = dbdi
		ri.confCache[channel].exp = time.Now().Add(time.Hour)
	} else {
		ri.confCache[channel] = &cacheData{
			di:  dbdi,
			exp: time.Now().Add(time.Hour),
		}
	}
	dbdi.SetChannel(channel)
	return dbdi, nil
}

func getContextWitchRsrc(ctx context.Context, di dbMidDI, backDb string) (r *rsrc, err error) {
	uuid := uuid.New().String()
	l := di.NewLogger(uuid)
	dbclt, err := di.NewMongoDBClient(ctx, backDb)
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

func NewInterConfInterceptor(address string, di dbMidDI) (coreInterceptor.Interceptor, error) {
	confSDK, err := client.New(address)
	if err != nil {
		return nil, err
	}
	return &interConfInterceptor{
		confSDK:   confSDK,
		di:        di,
		confCache: map[string]*cacheData{},
	}, nil
}

type interConfInterceptor struct {
	confSDK   client.ConfigurationClient
	di        dbMidDI
	confCache map[string]*cacheData
}

func getChannel(md metadata.MD) string {
	channels := md.Get("X-Channel")
	var channel string
	if len(channels) == 1 {
		channel = channels[0]
	}
	return channel
}

func (ri *interConfInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		if isReflectMethod(info.FullMethod) {
			return handler(srv, ss)
		}
		ctx := ss.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.InvalidArgument, "can not get metadata")
		}
		channel := getChannel(md)
		if channel == "" {
			// for get grpc service list
			handler(srv, &serverStream{
				ServerStream: ss,
				ctx:          ctx,
			})
			return
		}
		di, err := ri.getDI(ctx, channel)
		if err != nil {

			return status.Error(codes.InvalidArgument, "can not get di")
		}
		backDb := di.GetDb() + "-back"
		rsrc, err := getContextWitchRsrc(ctx, di, backDb)
		if err != nil {
			return err
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

func (ri *interConfInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "can not get metadata")
		}

		channel := getChannel(md)
		if channel == "" {
			return nil, status.Error(codes.InvalidArgument, "can not get channel")
		}
		di, err := ri.getDI(ctx, channel)
		if err != nil {
			// for get grpc service list
			return nil, status.Error(codes.InvalidArgument, "can not get di"+err.Error())
		}
		backDb := di.GetDb() + "-back"
		r, err := getContextWitchRsrc(ctx, di, backDb)
		if err != nil {
			return nil, err
		}
		defer r.close()
		return handler(r.setContext(ctx), req)
	})
}

type rsrc struct {
	dbclt db.MongoDBClient
	l     log.Logger
	di    dbMidDI
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

func isReflectMethod(m string) bool {
	return m == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
}
