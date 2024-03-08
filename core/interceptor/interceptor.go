package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Interceptor interface {
	StreamServerInterceptor() grpc.StreamServerInterceptor
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
}

func NewServerStream(ctx context.Context, stream grpc.ServerStream) grpc.ServerStream {
	return &serverStream{
		ServerStream: stream,
		ctx:          stream.Context(),
	}
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func IsReflectMethod(m string) bool {
	return m == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
}
