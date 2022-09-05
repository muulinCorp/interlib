package interceptor

import (
	"context"

	"github.com/94peter/sterna/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getReqUser(ctx context.Context) (auth.ReqUser, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "can not get metadata")
	}

	reqUserStr := md.Get("X-ReqUser")

	if len(reqUserStr) == 1 {
		reqUser := auth.NewEmptyReqUser()
		err := reqUser.Decode(reqUserStr[0])
		if err != nil {
			return nil, status.Error(codes.Internal, "decode error: "+err.Error())
		}
		return reqUser, nil

	}
	return nil, nil
}

func StreamServerAuthInterceptor() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		user, err := getReqUser(ss.Context())
		if err != nil {
			return err
		}
		ctx := context.WithValue(ss.Context(), auth.CtxUserInfoKey, user)
		err = handler(srv, &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		})
		return
	})
}

func UnaryServerAuthInterceptor() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		user, err := getReqUser(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, auth.CtxUserInfoKey, user)
		resp, err := handler(ctx, req)
		return resp, err
	})
}