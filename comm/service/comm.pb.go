// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: comm/proto/comm.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Iot627TimingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac       string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	VirtualID uint32 `protobuf:"varint,2,opt,name=virtualID,proto3" json:"virtualID,omitempty"`
	Zone      string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *Iot627TimingRequest) Reset() {
	*x = Iot627TimingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_comm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iot627TimingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iot627TimingRequest) ProtoMessage() {}

func (x *Iot627TimingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_comm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iot627TimingRequest.ProtoReflect.Descriptor instead.
func (*Iot627TimingRequest) Descriptor() ([]byte, []int) {
	return file_comm_proto_comm_proto_rawDescGZIP(), []int{0}
}

func (x *Iot627TimingRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Iot627TimingRequest) GetVirtualID() uint32 {
	if x != nil {
		return x.VirtualID
	}
	return 0
}

func (x *Iot627TimingRequest) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

type Iot627RemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac       string  `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	VirtualID uint32  `protobuf:"varint,2,opt,name=virtualID,proto3" json:"virtualID,omitempty"`
	Key       string  `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value     float64 `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Iot627RemoteRequest) Reset() {
	*x = Iot627RemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_comm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iot627RemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iot627RemoteRequest) ProtoMessage() {}

func (x *Iot627RemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_comm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iot627RemoteRequest.ProtoReflect.Descriptor instead.
func (*Iot627RemoteRequest) Descriptor() ([]byte, []int) {
	return file_comm_proto_comm_proto_rawDescGZIP(), []int{1}
}

func (x *Iot627RemoteRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Iot627RemoteRequest) GetVirtualID() uint32 {
	if x != nil {
		return x.VirtualID
	}
	return 0
}

func (x *Iot627RemoteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Iot627RemoteRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Iot627GetControlValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac       string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	VirtualID uint32 `protobuf:"varint,2,opt,name=virtualID,proto3" json:"virtualID,omitempty"`
}

func (x *Iot627GetControlValueRequest) Reset() {
	*x = Iot627GetControlValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_comm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iot627GetControlValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iot627GetControlValueRequest) ProtoMessage() {}

func (x *Iot627GetControlValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_comm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iot627GetControlValueRequest.ProtoReflect.Descriptor instead.
func (*Iot627GetControlValueRequest) Descriptor() ([]byte, []int) {
	return file_comm_proto_comm_proto_rawDescGZIP(), []int{2}
}

func (x *Iot627GetControlValueRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Iot627GetControlValueRequest) GetVirtualID() uint32 {
	if x != nil {
		return x.VirtualID
	}
	return 0
}

type CommResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode uint32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Mac        string `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	VirtualID  uint32 `protobuf:"varint,3,opt,name=virtualID,proto3" json:"virtualID,omitempty"`
	Message    string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommResponse) Reset() {
	*x = CommResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comm_proto_comm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommResponse) ProtoMessage() {}

func (x *CommResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comm_proto_comm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommResponse.ProtoReflect.Descriptor instead.
func (*CommResponse) Descriptor() ([]byte, []int) {
	return file_comm_proto_comm_proto_rawDescGZIP(), []int{3}
}

func (x *CommResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommResponse) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *CommResponse) GetVirtualID() uint32 {
	if x != nil {
		return x.VirtualID
	}
	return 0
}

func (x *CommResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_comm_proto_comm_proto protoreflect.FileDescriptor

var file_comm_proto_comm_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x59, 0x0a, 0x13, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x6d, 0x0a, 0x13, 0x49,
	0x6f, 0x74, 0x36, 0x32, 0x37, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4e, 0x0a, 0x1c, 0x49, 0x6f,
	0x74, 0x36, 0x32, 0x37, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x22, 0x78, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x09,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xf8, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37, 0x54, 0x69,
	0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6f, 0x74, 0x36, 0x32, 0x37, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x45, 0x0a, 0x0c, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x15, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x36, 0x32, 0x37,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comm_proto_comm_proto_rawDescOnce sync.Once
	file_comm_proto_comm_proto_rawDescData = file_comm_proto_comm_proto_rawDesc
)

func file_comm_proto_comm_proto_rawDescGZIP() []byte {
	file_comm_proto_comm_proto_rawDescOnce.Do(func() {
		file_comm_proto_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_comm_proto_comm_proto_rawDescData)
	})
	return file_comm_proto_comm_proto_rawDescData
}

var file_comm_proto_comm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_comm_proto_comm_proto_goTypes = []interface{}{
	(*Iot627TimingRequest)(nil),          // 0: service.Iot627TimingRequest
	(*Iot627RemoteRequest)(nil),          // 1: service.Iot627RemoteRequest
	(*Iot627GetControlValueRequest)(nil), // 2: service.Iot627GetControlValueRequest
	(*CommResponse)(nil),                 // 3: service.CommResponse
}
var file_comm_proto_comm_proto_depIdxs = []int32{
	0, // 0: service.CommService.Iot627Timing:input_type -> service.Iot627TimingRequest
	1, // 1: service.CommService.Iot627Remote:input_type -> service.Iot627RemoteRequest
	2, // 2: service.CommService.Iot627GetControlValue:input_type -> service.Iot627GetControlValueRequest
	3, // 3: service.CommService.Iot627Timing:output_type -> service.CommResponse
	3, // 4: service.CommService.Iot627Remote:output_type -> service.CommResponse
	3, // 5: service.CommService.Iot627GetControlValue:output_type -> service.CommResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comm_proto_comm_proto_init() }
func file_comm_proto_comm_proto_init() {
	if File_comm_proto_comm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comm_proto_comm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iot627TimingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_comm_proto_comm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iot627RemoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_comm_proto_comm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iot627GetControlValueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_comm_proto_comm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comm_proto_comm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comm_proto_comm_proto_goTypes,
		DependencyIndexes: file_comm_proto_comm_proto_depIdxs,
		MessageInfos:      file_comm_proto_comm_proto_msgTypes,
	}.Build()
	File_comm_proto_comm_proto = out.File
	file_comm_proto_comm_proto_rawDesc = nil
	file_comm_proto_comm_proto_goTypes = nil
	file_comm_proto_comm_proto_depIdxs = nil
}
