// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: channel/proto/channel_conf.proto

package service

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

// ChannelConfClient is the client API for ChannelConf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelConfClient interface {
	// Sends a greeting
	GetConf(ctx context.Context, in *GetConfRequest, opts ...grpc.CallOption) (*GetConfReply, error)
}

type channelConfClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelConfClient(cc grpc.ClientConnInterface) ChannelConfClient {
	return &channelConfClient{cc}
}

func (c *channelConfClient) GetConf(ctx context.Context, in *GetConfRequest, opts ...grpc.CallOption) (*GetConfReply, error) {
	out := new(GetConfReply)
	err := c.cc.Invoke(ctx, "/service.ChannelConf/GetConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelConfServer is the server API for ChannelConf service.
// All implementations must embed UnimplementedChannelConfServer
// for forward compatibility
type ChannelConfServer interface {
	// Sends a greeting
	GetConf(context.Context, *GetConfRequest) (*GetConfReply, error)
	mustEmbedUnimplementedChannelConfServer()
}

// UnimplementedChannelConfServer must be embedded to have forward compatible implementations.
type UnimplementedChannelConfServer struct {
}

func (UnimplementedChannelConfServer) GetConf(context.Context, *GetConfRequest) (*GetConfReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConf not implemented")
}
func (UnimplementedChannelConfServer) mustEmbedUnimplementedChannelConfServer() {}

// UnsafeChannelConfServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelConfServer will
// result in compilation errors.
type UnsafeChannelConfServer interface {
	mustEmbedUnimplementedChannelConfServer()
}

func RegisterChannelConfServer(s grpc.ServiceRegistrar, srv ChannelConfServer) {
	s.RegisterService(&ChannelConf_ServiceDesc, srv)
}

func _ChannelConf_GetConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelConfServer).GetConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ChannelConf/GetConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelConfServer).GetConf(ctx, req.(*GetConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelConf_ServiceDesc is the grpc.ServiceDesc for ChannelConf service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelConf_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ChannelConf",
	HandlerType: (*ChannelConfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConf",
			Handler:    _ChannelConf_GetConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel/proto/channel_conf.proto",
}