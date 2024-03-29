// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: channel/proto/channel_cron.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CronServiceClient is the client API for CronService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CronServiceClient interface {
	// 案場預警檢檢
	WarningCheckingTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cronServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCronServiceClient(cc grpc.ClientConnInterface) CronServiceClient {
	return &cronServiceClient{cc}
}

func (c *cronServiceClient) WarningCheckingTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/channel.CronService/WarningCheckingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronServiceServer is the server API for CronService service.
// All implementations must embed UnimplementedCronServiceServer
// for forward compatibility
type CronServiceServer interface {
	// 案場預警檢檢
	WarningCheckingTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedCronServiceServer()
}

// UnimplementedCronServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCronServiceServer struct {
}

func (UnimplementedCronServiceServer) WarningCheckingTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WarningCheckingTask not implemented")
}
func (UnimplementedCronServiceServer) mustEmbedUnimplementedCronServiceServer() {}

// UnsafeCronServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CronServiceServer will
// result in compilation errors.
type UnsafeCronServiceServer interface {
	mustEmbedUnimplementedCronServiceServer()
}

func RegisterCronServiceServer(s grpc.ServiceRegistrar, srv CronServiceServer) {
	s.RegisterService(&CronService_ServiceDesc, srv)
}

func _CronService_WarningCheckingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServiceServer).WarningCheckingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.CronService/WarningCheckingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServiceServer).WarningCheckingTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CronService_ServiceDesc is the grpc.ServiceDesc for CronService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CronService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.CronService",
	HandlerType: (*CronServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WarningCheckingTask",
			Handler:    _CronService_WarningCheckingTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/proto/channel_cron.proto",
}
