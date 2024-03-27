// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: config-sync/proto/config-sync.proto

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

// ConfigSyncServiceClient is the client API for ConfigSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigSyncServiceClient interface {
	// 通知服務已完成重新啟動
	ServiceStarted(ctx context.Context, in *ServiceStartedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 取得目前版本的config檔
	GetConfigFiles(ctx context.Context, in *GetConfigFilesReq, opts ...grpc.CallOption) (ConfigSyncService_GetConfigFilesClient, error)
}

type configSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigSyncServiceClient(cc grpc.ClientConnInterface) ConfigSyncServiceClient {
	return &configSyncServiceClient{cc}
}

func (c *configSyncServiceClient) ServiceStarted(ctx context.Context, in *ServiceStartedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/config_sync.ConfigSyncService/ServiceStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configSyncServiceClient) GetConfigFiles(ctx context.Context, in *GetConfigFilesReq, opts ...grpc.CallOption) (ConfigSyncService_GetConfigFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConfigSyncService_ServiceDesc.Streams[0], "/config_sync.ConfigSyncService/GetConfigFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &configSyncServiceGetConfigFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigSyncService_GetConfigFilesClient interface {
	Recv() (*GetConfigFilesResp, error)
	grpc.ClientStream
}

type configSyncServiceGetConfigFilesClient struct {
	grpc.ClientStream
}

func (x *configSyncServiceGetConfigFilesClient) Recv() (*GetConfigFilesResp, error) {
	m := new(GetConfigFilesResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigSyncServiceServer is the server API for ConfigSyncService service.
// All implementations must embed UnimplementedConfigSyncServiceServer
// for forward compatibility
type ConfigSyncServiceServer interface {
	// 通知服務已完成重新啟動
	ServiceStarted(context.Context, *ServiceStartedRequest) (*emptypb.Empty, error)
	// 取得目前版本的config檔
	GetConfigFiles(*GetConfigFilesReq, ConfigSyncService_GetConfigFilesServer) error
	mustEmbedUnimplementedConfigSyncServiceServer()
}

// UnimplementedConfigSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigSyncServiceServer struct {
}

func (UnimplementedConfigSyncServiceServer) ServiceStarted(context.Context, *ServiceStartedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStarted not implemented")
}
func (UnimplementedConfigSyncServiceServer) GetConfigFiles(*GetConfigFilesReq, ConfigSyncService_GetConfigFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfigFiles not implemented")
}
func (UnimplementedConfigSyncServiceServer) mustEmbedUnimplementedConfigSyncServiceServer() {}

// UnsafeConfigSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigSyncServiceServer will
// result in compilation errors.
type UnsafeConfigSyncServiceServer interface {
	mustEmbedUnimplementedConfigSyncServiceServer()
}

func RegisterConfigSyncServiceServer(s grpc.ServiceRegistrar, srv ConfigSyncServiceServer) {
	s.RegisterService(&ConfigSyncService_ServiceDesc, srv)
}

func _ConfigSyncService_ServiceStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStartedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigSyncServiceServer).ServiceStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_sync.ConfigSyncService/ServiceStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigSyncServiceServer).ServiceStarted(ctx, req.(*ServiceStartedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigSyncService_GetConfigFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetConfigFilesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigSyncServiceServer).GetConfigFiles(m, &configSyncServiceGetConfigFilesServer{stream})
}

type ConfigSyncService_GetConfigFilesServer interface {
	Send(*GetConfigFilesResp) error
	grpc.ServerStream
}

type configSyncServiceGetConfigFilesServer struct {
	grpc.ServerStream
}

func (x *configSyncServiceGetConfigFilesServer) Send(m *GetConfigFilesResp) error {
	return x.ServerStream.SendMsg(m)
}

// ConfigSyncService_ServiceDesc is the grpc.ServiceDesc for ConfigSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config_sync.ConfigSyncService",
	HandlerType: (*ConfigSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceStarted",
			Handler:    _ConfigSyncService_ServiceStarted_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfigFiles",
			Handler:       _ConfigSyncService_GetConfigFiles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "config-sync/proto/config-sync.proto",
}