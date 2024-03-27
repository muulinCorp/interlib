// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: communicate/proto/communicate.proto

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

// CommunicateServiceClient is the client API for CommunicateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicateServiceClient interface {
	// remote
	Remote(ctx context.Context, in *RemoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// get sensors info
	GetSensors(ctx context.Context, in *GetSensorsRequest, opts ...grpc.CallOption) (*GetSensorsResponse, error)
	// 對單一config做一次性debug
	ConfigDebug(ctx context.Context, in *ConfigDebugRequest, opts ...grpc.CallOption) (*ConfigDebugResponse, error)
}

type communicateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicateServiceClient(cc grpc.ClientConnInterface) CommunicateServiceClient {
	return &communicateServiceClient{cc}
}

func (c *communicateServiceClient) Remote(ctx context.Context, in *RemoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/communicate.CommunicateService/Remote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicateServiceClient) GetSensors(ctx context.Context, in *GetSensorsRequest, opts ...grpc.CallOption) (*GetSensorsResponse, error) {
	out := new(GetSensorsResponse)
	err := c.cc.Invoke(ctx, "/communicate.CommunicateService/GetSensors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicateServiceClient) ConfigDebug(ctx context.Context, in *ConfigDebugRequest, opts ...grpc.CallOption) (*ConfigDebugResponse, error) {
	out := new(ConfigDebugResponse)
	err := c.cc.Invoke(ctx, "/communicate.CommunicateService/ConfigDebug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicateServiceServer is the server API for CommunicateService service.
// All implementations must embed UnimplementedCommunicateServiceServer
// for forward compatibility
type CommunicateServiceServer interface {
	// remote
	Remote(context.Context, *RemoteRequest) (*emptypb.Empty, error)
	// get sensors info
	GetSensors(context.Context, *GetSensorsRequest) (*GetSensorsResponse, error)
	// 對單一config做一次性debug
	ConfigDebug(context.Context, *ConfigDebugRequest) (*ConfigDebugResponse, error)
	mustEmbedUnimplementedCommunicateServiceServer()
}

// UnimplementedCommunicateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicateServiceServer struct {
}

func (UnimplementedCommunicateServiceServer) Remote(context.Context, *RemoteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remote not implemented")
}
func (UnimplementedCommunicateServiceServer) GetSensors(context.Context, *GetSensorsRequest) (*GetSensorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensors not implemented")
}
func (UnimplementedCommunicateServiceServer) ConfigDebug(context.Context, *ConfigDebugRequest) (*ConfigDebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigDebug not implemented")
}
func (UnimplementedCommunicateServiceServer) mustEmbedUnimplementedCommunicateServiceServer() {}

// UnsafeCommunicateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicateServiceServer will
// result in compilation errors.
type UnsafeCommunicateServiceServer interface {
	mustEmbedUnimplementedCommunicateServiceServer()
}

func RegisterCommunicateServiceServer(s grpc.ServiceRegistrar, srv CommunicateServiceServer) {
	s.RegisterService(&CommunicateService_ServiceDesc, srv)
}

func _CommunicateService_Remote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicateServiceServer).Remote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communicate.CommunicateService/Remote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicateServiceServer).Remote(ctx, req.(*RemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicateService_GetSensors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSensorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicateServiceServer).GetSensors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communicate.CommunicateService/GetSensors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicateServiceServer).GetSensors(ctx, req.(*GetSensorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicateService_ConfigDebug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigDebugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicateServiceServer).ConfigDebug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communicate.CommunicateService/ConfigDebug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicateServiceServer).ConfigDebug(ctx, req.(*ConfigDebugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicateService_ServiceDesc is the grpc.ServiceDesc for CommunicateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communicate.CommunicateService",
	HandlerType: (*CommunicateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Remote",
			Handler:    _CommunicateService_Remote_Handler,
		},
		{
			MethodName: "GetSensors",
			Handler:    _CommunicateService_GetSensors_Handler,
		},
		{
			MethodName: "ConfigDebug",
			Handler:    _CommunicateService_ConfigDebug_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communicate/proto/communicate.proto",
}
