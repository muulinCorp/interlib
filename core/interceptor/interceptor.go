package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Interceptor interface {
	StreamServerInterceptor() grpc.StreamServerInterceptor
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}