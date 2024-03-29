// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: communicate/proto/communicate.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]float64 `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *RemoteRequest) Reset() {
	*x = RemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteRequest) ProtoMessage() {}

func (x *RemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteRequest.ProtoReflect.Descriptor instead.
func (*RemoteRequest) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteRequest) GetValues() map[string]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type GetSensorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSensorsRequest) Reset() {
	*x = GetSensorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsRequest) ProtoMessage() {}

func (x *GetSensorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsRequest.ProtoReflect.Descriptor instead.
func (*GetSensorsRequest) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{1}
}

func (x *GetSensorsRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

type GetSensorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensors       []*GetSensorsResponse_Sensor `protobuf:"bytes,1,rep,name=sensors,proto3" json:"sensors,omitempty"`
	MissingFields []string                     `protobuf:"bytes,2,rep,name=missing_fields,json=missingFields,proto3" json:"missing_fields,omitempty"`
}

func (x *GetSensorsResponse) Reset() {
	*x = GetSensorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsResponse) ProtoMessage() {}

func (x *GetSensorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsResponse.ProtoReflect.Descriptor instead.
func (*GetSensorsResponse) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{2}
}

func (x *GetSensorsResponse) GetSensors() []*GetSensorsResponse_Sensor {
	if x != nil {
		return x.Sensors
	}
	return nil
}

func (x *GetSensorsResponse) GetMissingFields() []string {
	if x != nil {
		return x.MissingFields
	}
	return nil
}

type ConfigDebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigBinaryData []byte `protobuf:"bytes,1,opt,name=configBinaryData,proto3" json:"configBinaryData,omitempty"`
}

func (x *ConfigDebugRequest) Reset() {
	*x = ConfigDebugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDebugRequest) ProtoMessage() {}

func (x *ConfigDebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDebugRequest.ProtoReflect.Descriptor instead.
func (*ConfigDebugRequest) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigDebugRequest) GetConfigBinaryData() []byte {
	if x != nil {
		return x.ConfigBinaryData
	}
	return nil
}

type ConfigDebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]*ConfigDebugResponse_ValueDetail `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigDebugResponse) Reset() {
	*x = ConfigDebugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDebugResponse) ProtoMessage() {}

func (x *ConfigDebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDebugResponse.ProtoReflect.Descriptor instead.
func (*ConfigDebugResponse) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigDebugResponse) GetValues() map[string]*ConfigDebugResponse_ValueDetail {
	if x != nil {
		return x.Values
	}
	return nil
}

type GetSensorsResponse_Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value       float64   `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	IsConnErr   bool      `protobuf:"varint,3,opt,name=is_conn_err,json=isConnErr,proto3" json:"is_conn_err,omitempty"`
	Error       string    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	ErrorDetail string    `protobuf:"bytes,5,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	IsFilter    bool      `protobuf:"varint,6,opt,name=isFilter,proto3" json:"isFilter,omitempty"`
	ValueList   []float64 `protobuf:"fixed64,7,rep,packed,name=value_list,json=valueList,proto3" json:"value_list,omitempty"`
	Data        []byte    `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSensorsResponse_Sensor) Reset() {
	*x = GetSensorsResponse_Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorsResponse_Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorsResponse_Sensor) ProtoMessage() {}

func (x *GetSensorsResponse_Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorsResponse_Sensor.ProtoReflect.Descriptor instead.
func (*GetSensorsResponse_Sensor) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetSensorsResponse_Sensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSensorsResponse_Sensor) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *GetSensorsResponse_Sensor) GetIsConnErr() bool {
	if x != nil {
		return x.IsConnErr
	}
	return false
}

func (x *GetSensorsResponse_Sensor) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetSensorsResponse_Sensor) GetErrorDetail() string {
	if x != nil {
		return x.ErrorDetail
	}
	return ""
}

func (x *GetSensorsResponse_Sensor) GetIsFilter() bool {
	if x != nil {
		return x.IsFilter
	}
	return false
}

func (x *GetSensorsResponse_Sensor) GetValueList() []float64 {
	if x != nil {
		return x.ValueList
	}
	return nil
}

func (x *GetSensorsResponse_Sensor) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConfigDebugResponse_ValueDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	ErrMsg string  `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *ConfigDebugResponse_ValueDetail) Reset() {
	*x = ConfigDebugResponse_ValueDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_proto_communicate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDebugResponse_ValueDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDebugResponse_ValueDetail) ProtoMessage() {}

func (x *ConfigDebugResponse_ValueDetail) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_proto_communicate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDebugResponse_ValueDetail.ProtoReflect.Descriptor instead.
func (*ConfigDebugResponse_ValueDetail) Descriptor() ([]byte, []int) {
	return file_communicate_proto_communicate_proto_rawDescGZIP(), []int{4, 1}
}

func (x *ConfigDebugResponse_ValueDetail) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ConfigDebugResponse_ValueDetail) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_communicate_proto_communicate_proto protoreflect.FileDescriptor

var file_communicate_proto_communicate_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8a, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xda, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0xda, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x82, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x67, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0b, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x32, 0xf9, 0x01, 0x0a, 0x12, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x1e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communicate_proto_communicate_proto_rawDescOnce sync.Once
	file_communicate_proto_communicate_proto_rawDescData = file_communicate_proto_communicate_proto_rawDesc
)

func file_communicate_proto_communicate_proto_rawDescGZIP() []byte {
	file_communicate_proto_communicate_proto_rawDescOnce.Do(func() {
		file_communicate_proto_communicate_proto_rawDescData = protoimpl.X.CompressGZIP(file_communicate_proto_communicate_proto_rawDescData)
	})
	return file_communicate_proto_communicate_proto_rawDescData
}

var file_communicate_proto_communicate_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_communicate_proto_communicate_proto_goTypes = []interface{}{
	(*RemoteRequest)(nil),                   // 0: communicate.RemoteRequest
	(*GetSensorsRequest)(nil),               // 1: communicate.GetSensorsRequest
	(*GetSensorsResponse)(nil),              // 2: communicate.GetSensorsResponse
	(*ConfigDebugRequest)(nil),              // 3: communicate.ConfigDebugRequest
	(*ConfigDebugResponse)(nil),             // 4: communicate.ConfigDebugResponse
	nil,                                     // 5: communicate.RemoteRequest.ValuesEntry
	(*GetSensorsResponse_Sensor)(nil),       // 6: communicate.GetSensorsResponse.Sensor
	nil,                                     // 7: communicate.ConfigDebugResponse.ValuesEntry
	(*ConfigDebugResponse_ValueDetail)(nil), // 8: communicate.ConfigDebugResponse.ValueDetail
	(*emptypb.Empty)(nil),                   // 9: google.protobuf.Empty
}
var file_communicate_proto_communicate_proto_depIdxs = []int32{
	5, // 0: communicate.RemoteRequest.values:type_name -> communicate.RemoteRequest.ValuesEntry
	6, // 1: communicate.GetSensorsResponse.sensors:type_name -> communicate.GetSensorsResponse.Sensor
	7, // 2: communicate.ConfigDebugResponse.values:type_name -> communicate.ConfigDebugResponse.ValuesEntry
	8, // 3: communicate.ConfigDebugResponse.ValuesEntry.value:type_name -> communicate.ConfigDebugResponse.ValueDetail
	0, // 4: communicate.CommunicateService.Remote:input_type -> communicate.RemoteRequest
	1, // 5: communicate.CommunicateService.GetSensors:input_type -> communicate.GetSensorsRequest
	3, // 6: communicate.CommunicateService.ConfigDebug:input_type -> communicate.ConfigDebugRequest
	9, // 7: communicate.CommunicateService.Remote:output_type -> google.protobuf.Empty
	2, // 8: communicate.CommunicateService.GetSensors:output_type -> communicate.GetSensorsResponse
	4, // 9: communicate.CommunicateService.ConfigDebug:output_type -> communicate.ConfigDebugResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_communicate_proto_communicate_proto_init() }
func file_communicate_proto_communicate_proto_init() {
	if File_communicate_proto_communicate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communicate_proto_communicate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteRequest); i {
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
		file_communicate_proto_communicate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsRequest); i {
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
		file_communicate_proto_communicate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsResponse); i {
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
		file_communicate_proto_communicate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDebugRequest); i {
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
		file_communicate_proto_communicate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDebugResponse); i {
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
		file_communicate_proto_communicate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSensorsResponse_Sensor); i {
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
		file_communicate_proto_communicate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDebugResponse_ValueDetail); i {
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
			RawDescriptor: file_communicate_proto_communicate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communicate_proto_communicate_proto_goTypes,
		DependencyIndexes: file_communicate_proto_communicate_proto_depIdxs,
		MessageInfos:      file_communicate_proto_communicate_proto_msgTypes,
	}.Build()
	File_communicate_proto_communicate_proto = out.File
	file_communicate_proto_communicate_proto_rawDesc = nil
	file_communicate_proto_communicate_proto_goTypes = nil
	file_communicate_proto_communicate_proto_depIdxs = nil
}
