// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: channel/proto/channel_maintenace.proto

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

// MaintenaceServiceClient is the client API for MaintenaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaintenaceServiceClient interface {
	// 設得設備資訊
	GetEquipInfo(ctx context.Context, in *EquipInfoRequest, opts ...grpc.CallOption) (*EquipInfoResponse, error)
	// 取得使用者設備Id清單
	GetEquipIdsByAccount(ctx context.Context, in *GetEquipIdsByAccountRequest, opts ...grpc.CallOption) (*GetEquipIdsByAccountResponse, error)
	// 維修事件
	EmitEvent(ctx context.Context, in *MaintenanceEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type maintenaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaintenaceServiceClient(cc grpc.ClientConnInterface) MaintenaceServiceClient {
	return &maintenaceServiceClient{cc}
}

func (c *maintenaceServiceClient) GetEquipInfo(ctx context.Context, in *EquipInfoRequest, opts ...grpc.CallOption) (*EquipInfoResponse, error) {
	out := new(EquipInfoResponse)
	err := c.cc.Invoke(ctx, "/channel.MaintenaceService/GetEquipInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenaceServiceClient) GetEquipIdsByAccount(ctx context.Context, in *GetEquipIdsByAccountRequest, opts ...grpc.CallOption) (*GetEquipIdsByAccountResponse, error) {
	out := new(GetEquipIdsByAccountResponse)
	err := c.cc.Invoke(ctx, "/channel.MaintenaceService/GetEquipIdsByAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenaceServiceClient) EmitEvent(ctx context.Context, in *MaintenanceEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/channel.MaintenaceService/EmitEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintenaceServiceServer is the server API for MaintenaceService service.
// All implementations must embed UnimplementedMaintenaceServiceServer
// for forward compatibility
type MaintenaceServiceServer interface {
	// 設得設備資訊
	GetEquipInfo(context.Context, *EquipInfoRequest) (*EquipInfoResponse, error)
	// 取得使用者設備Id清單
	GetEquipIdsByAccount(context.Context, *GetEquipIdsByAccountRequest) (*GetEquipIdsByAccountResponse, error)
	// 維修事件
	EmitEvent(context.Context, *MaintenanceEventReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedMaintenaceServiceServer()
}

// UnimplementedMaintenaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMaintenaceServiceServer struct {
}

func (UnimplementedMaintenaceServiceServer) GetEquipInfo(context.Context, *EquipInfoRequest) (*EquipInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEquipInfo not implemented")
}
func (UnimplementedMaintenaceServiceServer) GetEquipIdsByAccount(context.Context, *GetEquipIdsByAccountRequest) (*GetEquipIdsByAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEquipIdsByAccount not implemented")
}
func (UnimplementedMaintenaceServiceServer) EmitEvent(context.Context, *MaintenanceEventReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitEvent not implemented")
}
func (UnimplementedMaintenaceServiceServer) mustEmbedUnimplementedMaintenaceServiceServer() {}

// UnsafeMaintenaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaintenaceServiceServer will
// result in compilation errors.
type UnsafeMaintenaceServiceServer interface {
	mustEmbedUnimplementedMaintenaceServiceServer()
}

func RegisterMaintenaceServiceServer(s grpc.ServiceRegistrar, srv MaintenaceServiceServer) {
	s.RegisterService(&MaintenaceService_ServiceDesc, srv)
}

func _MaintenaceService_GetEquipInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenaceServiceServer).GetEquipInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.MaintenaceService/GetEquipInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenaceServiceServer).GetEquipInfo(ctx, req.(*EquipInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintenaceService_GetEquipIdsByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEquipIdsByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenaceServiceServer).GetEquipIdsByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.MaintenaceService/GetEquipIdsByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenaceServiceServer).GetEquipIdsByAccount(ctx, req.(*GetEquipIdsByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintenaceService_EmitEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaintenanceEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenaceServiceServer).EmitEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.MaintenaceService/EmitEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenaceServiceServer).EmitEvent(ctx, req.(*MaintenanceEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MaintenaceService_ServiceDesc is the grpc.ServiceDesc for MaintenaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaintenaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.MaintenaceService",
	HandlerType: (*MaintenaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEquipInfo",
			Handler:    _MaintenaceService_GetEquipInfo_Handler,
		},
		{
			MethodName: "GetEquipIdsByAccount",
			Handler:    _MaintenaceService_GetEquipIdsByAccount_Handler,
		},
		{
			MethodName: "EmitEvent",
			Handler:    _MaintenaceService_EmitEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/proto/channel_maintenace.proto",
}
