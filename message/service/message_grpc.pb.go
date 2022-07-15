// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: message/proto/message.proto

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	MqttPublish(ctx context.Context, in *MqttPublishRequest, opts ...grpc.CallOption) (*MqttPublishResponse, error)
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
	Mail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (MessageService_MailClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) MqttPublish(ctx context.Context, in *MqttPublishRequest, opts ...grpc.CallOption) (*MqttPublishResponse, error) {
	out := new(MqttPublishResponse)
	err := c.cc.Invoke(ctx, "/service.MessageService/MqttPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/service.MessageService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Mail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (MessageService_MailClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[0], "/service.MessageService/Mail", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceMailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_MailClient interface {
	Recv() (*MailResponse, error)
	grpc.ClientStream
}

type messageServiceMailClient struct {
	grpc.ClientStream
}

func (x *messageServiceMailClient) Recv() (*MailResponse, error) {
	m := new(MailResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	MqttPublish(context.Context, *MqttPublishRequest) (*MqttPublishResponse, error)
	Push(context.Context, *PushRequest) (*PushResponse, error)
	Mail(*MailRequest, MessageService_MailServer) error
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) MqttPublish(context.Context, *MqttPublishRequest) (*MqttPublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MqttPublish not implemented")
}
func (UnimplementedMessageServiceServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedMessageServiceServer) Mail(*MailRequest, MessageService_MailServer) error {
	return status.Errorf(codes.Unimplemented, "method Mail not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_MqttPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttPublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MqttPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MessageService/MqttPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MqttPublish(ctx, req.(*MqttPublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MessageService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Mail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).Mail(m, &messageServiceMailServer{stream})
}

type MessageService_MailServer interface {
	Send(*MailResponse) error
	grpc.ServerStream
}

type messageServiceMailServer struct {
	grpc.ServerStream
}

func (x *messageServiceMailServer) Send(m *MailResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MqttPublish",
			Handler:    _MessageService_MqttPublish_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _MessageService_Push_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Mail",
			Handler:       _MessageService_Mail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message/proto/message.proto",
}
