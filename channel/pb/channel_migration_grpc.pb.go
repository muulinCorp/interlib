// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: channel/proto/channel_migration.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChannelMigrationServiceClient is the client API for ChannelMigrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelMigrationServiceClient interface {
	// 搬移案場資訊
	MigrationProject(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationProjectClient, error)
	// 搬移設備資訊
	MigrationEquipment(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationEquipmentClient, error)
	// 搬移設備log
	MigrationEquipmentLog(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationEquipmentLogClient, error)
}

type channelMigrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelMigrationServiceClient(cc grpc.ClientConnInterface) ChannelMigrationServiceClient {
	return &channelMigrationServiceClient{cc}
}

func (c *channelMigrationServiceClient) MigrationProject(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationProjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChannelMigrationService_ServiceDesc.Streams[0], "/channel.ChannelMigrationService/MigrationProject", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelMigrationServiceMigrationProjectClient{stream}
	return x, nil
}

type ChannelMigrationService_MigrationProjectClient interface {
	Send(*MigrationProjectRequest) error
	Recv() (*MigrationResponse, error)
	grpc.ClientStream
}

type channelMigrationServiceMigrationProjectClient struct {
	grpc.ClientStream
}

func (x *channelMigrationServiceMigrationProjectClient) Send(m *MigrationProjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationProjectClient) Recv() (*MigrationResponse, error) {
	m := new(MigrationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *channelMigrationServiceClient) MigrationEquipment(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationEquipmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChannelMigrationService_ServiceDesc.Streams[1], "/channel.ChannelMigrationService/MigrationEquipment", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelMigrationServiceMigrationEquipmentClient{stream}
	return x, nil
}

type ChannelMigrationService_MigrationEquipmentClient interface {
	Send(*MigrationEquipmentRequest) error
	Recv() (*MigrationResponse, error)
	grpc.ClientStream
}

type channelMigrationServiceMigrationEquipmentClient struct {
	grpc.ClientStream
}

func (x *channelMigrationServiceMigrationEquipmentClient) Send(m *MigrationEquipmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationEquipmentClient) Recv() (*MigrationResponse, error) {
	m := new(MigrationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *channelMigrationServiceClient) MigrationEquipmentLog(ctx context.Context, opts ...grpc.CallOption) (ChannelMigrationService_MigrationEquipmentLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChannelMigrationService_ServiceDesc.Streams[2], "/channel.ChannelMigrationService/MigrationEquipmentLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &channelMigrationServiceMigrationEquipmentLogClient{stream}
	return x, nil
}

type ChannelMigrationService_MigrationEquipmentLogClient interface {
	Send(*MigrationLogRequest) error
	Recv() (*MigrationResponse, error)
	grpc.ClientStream
}

type channelMigrationServiceMigrationEquipmentLogClient struct {
	grpc.ClientStream
}

func (x *channelMigrationServiceMigrationEquipmentLogClient) Send(m *MigrationLogRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationEquipmentLogClient) Recv() (*MigrationResponse, error) {
	m := new(MigrationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelMigrationServiceServer is the server API for ChannelMigrationService service.
// All implementations must embed UnimplementedChannelMigrationServiceServer
// for forward compatibility
type ChannelMigrationServiceServer interface {
	// 搬移案場資訊
	MigrationProject(ChannelMigrationService_MigrationProjectServer) error
	// 搬移設備資訊
	MigrationEquipment(ChannelMigrationService_MigrationEquipmentServer) error
	// 搬移設備log
	MigrationEquipmentLog(ChannelMigrationService_MigrationEquipmentLogServer) error
	mustEmbedUnimplementedChannelMigrationServiceServer()
}

// UnimplementedChannelMigrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelMigrationServiceServer struct {
}

func (UnimplementedChannelMigrationServiceServer) MigrationProject(ChannelMigrationService_MigrationProjectServer) error {
	return status.Errorf(codes.Unimplemented, "method MigrationProject not implemented")
}
func (UnimplementedChannelMigrationServiceServer) MigrationEquipment(ChannelMigrationService_MigrationEquipmentServer) error {
	return status.Errorf(codes.Unimplemented, "method MigrationEquipment not implemented")
}
func (UnimplementedChannelMigrationServiceServer) MigrationEquipmentLog(ChannelMigrationService_MigrationEquipmentLogServer) error {
	return status.Errorf(codes.Unimplemented, "method MigrationEquipmentLog not implemented")
}
func (UnimplementedChannelMigrationServiceServer) mustEmbedUnimplementedChannelMigrationServiceServer() {
}

// UnsafeChannelMigrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelMigrationServiceServer will
// result in compilation errors.
type UnsafeChannelMigrationServiceServer interface {
	mustEmbedUnimplementedChannelMigrationServiceServer()
}

func RegisterChannelMigrationServiceServer(s grpc.ServiceRegistrar, srv ChannelMigrationServiceServer) {
	s.RegisterService(&ChannelMigrationService_ServiceDesc, srv)
}

func _ChannelMigrationService_MigrationProject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChannelMigrationServiceServer).MigrationProject(&channelMigrationServiceMigrationProjectServer{stream})
}

type ChannelMigrationService_MigrationProjectServer interface {
	Send(*MigrationResponse) error
	Recv() (*MigrationProjectRequest, error)
	grpc.ServerStream
}

type channelMigrationServiceMigrationProjectServer struct {
	grpc.ServerStream
}

func (x *channelMigrationServiceMigrationProjectServer) Send(m *MigrationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationProjectServer) Recv() (*MigrationProjectRequest, error) {
	m := new(MigrationProjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChannelMigrationService_MigrationEquipment_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChannelMigrationServiceServer).MigrationEquipment(&channelMigrationServiceMigrationEquipmentServer{stream})
}

type ChannelMigrationService_MigrationEquipmentServer interface {
	Send(*MigrationResponse) error
	Recv() (*MigrationEquipmentRequest, error)
	grpc.ServerStream
}

type channelMigrationServiceMigrationEquipmentServer struct {
	grpc.ServerStream
}

func (x *channelMigrationServiceMigrationEquipmentServer) Send(m *MigrationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationEquipmentServer) Recv() (*MigrationEquipmentRequest, error) {
	m := new(MigrationEquipmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChannelMigrationService_MigrationEquipmentLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChannelMigrationServiceServer).MigrationEquipmentLog(&channelMigrationServiceMigrationEquipmentLogServer{stream})
}

type ChannelMigrationService_MigrationEquipmentLogServer interface {
	Send(*MigrationResponse) error
	Recv() (*MigrationLogRequest, error)
	grpc.ServerStream
}

type channelMigrationServiceMigrationEquipmentLogServer struct {
	grpc.ServerStream
}

func (x *channelMigrationServiceMigrationEquipmentLogServer) Send(m *MigrationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *channelMigrationServiceMigrationEquipmentLogServer) Recv() (*MigrationLogRequest, error) {
	m := new(MigrationLogRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChannelMigrationService_ServiceDesc is the grpc.ServiceDesc for ChannelMigrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelMigrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.ChannelMigrationService",
	HandlerType: (*ChannelMigrationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MigrationProject",
			Handler:       _ChannelMigrationService_MigrationProject_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MigrationEquipment",
			Handler:       _ChannelMigrationService_MigrationEquipment_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MigrationEquipmentLog",
			Handler:       _ChannelMigrationService_MigrationEquipmentLog_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "channel/proto/channel_migration.proto",
}
