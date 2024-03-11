package interceptor

import (
	"context"
	"reflect"
	"time"

	"github.com/muulinCorp/interlib/channel"
	"github.com/muulinCorp/interlib/configuration/client"
	coreInterceptor "github.com/muulinCorp/interlib/core/interceptor"

	"github.com/muulinCorp/interlib/configuration/pb"

	"github.com/94peter/di"
	"github.com/94peter/log"
	"github.com/94peter/morm/conn"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type dbMidDI interface {
	log.LoggerDI
	conn.MongoDI
	channel.DI
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

func (ri *interConfInterceptor) getDI(ctx context.Context, ch string) (dbMidDI, error) {

	if cache, ok := ri.confCache[ch]; ok && time.Until(cache.exp) > 0 {
		return cache.di, nil
	}

	confByte, err := ri.confSDK.GetChannelConf(ctx, &pb.GetConfRequest{
		ChannelName: ch,
		Version:     "latest",
	})
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(ri.di)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	newValue := reflect.New(val.Type()).Interface().(channel.DI)
	di.InitConfByByte(confByte, newValue)
	dbdi := newValue.(dbMidDI)
	if _, ok := ri.confCache[ch]; ok {
		ri.confCache[ch].di = dbdi
		ri.confCache[ch].exp = time.Now().Add(time.Hour)
	} else {
		ri.confCache[ch] = &cacheData{
			di:  dbdi,
			exp: time.Now().Add(time.Hour),
		}
	}
	dbdi.SetChannel(ch)
	return dbdi, nil
}

func getContextWitchRsrc(ctx context.Context, di dbMidDI, service string) (r *rsrc, err error) {
	uuid := uuid.New().String()
	l, err := di.NewLogger(service, uuid)
	if err != nil {
		return nil, err
	}
	dbclt, err := di.NewDbConn(ctx, di.GetDb())
	if err != nil {
		return nil, err
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
		if coreInterceptor.IsReflectMethod(info.FullMethod) {
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

		rsrc, err := getContextWitchRsrc(ctx, di, di.GetService())
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
	dbclt conn.MongoDBConn
	l     log.Logger
	di    dbMidDI
}

func (r *rsrc) close() {
	if r.dbclt != nil {
		r.dbclt.Close()
	}
}

func (r *rsrc) setContext(ctx context.Context) context.Context {

	ctx = conn.SetMgoDbConnToCtx(ctx, r.dbclt)
	ctx = log.SetByCtx(ctx, r.l)
	ctx = di.SetDiToCtx(ctx, r.di)
	return ctx
}
