// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: channel/proto/channel_consumer.proto

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

type ValueType int32

const (
	ValueType_Sensor     ValueType = 0
	ValueType_Controller ValueType = 1
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0: "Sensor",
		1: "Controller",
	}
	ValueType_value = map[string]int32{
		"Sensor":     0,
		"Controller": 1,
	}
)

func (x ValueType) Enum() *ValueType {
	p := new(ValueType)
	*p = x
	return p
}

func (x ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_channel_proto_channel_consumer_proto_enumTypes[0].Descriptor()
}

func (ValueType) Type() protoreflect.EnumType {
	return &file_channel_proto_channel_consumer_proto_enumTypes[0]
}

func (x ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueType.Descriptor instead.
func (ValueType) EnumDescriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{0}
}

type EquipValueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac         string             `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	VirtualId   uint32             `protobuf:"varint,2,opt,name=virtualId,proto3" json:"virtualId,omitempty"`
	Type        ValueType          `protobuf:"varint,3,opt,name=type,proto3,enum=channel.ValueType" json:"type,omitempty"`
	Values      map[uint32]float64 `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	LastestTime uint32             `protobuf:"varint,5,opt,name=lastestTime,proto3" json:"lastestTime,omitempty"`
	OnlineState string             `protobuf:"bytes,6,opt,name=onlineState,proto3" json:"onlineState,omitempty"`
	NowTime     uint32             `protobuf:"varint,7,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
}

func (x *EquipValueReq) Reset() {
	*x = EquipValueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipValueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipValueReq) ProtoMessage() {}

func (x *EquipValueReq) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipValueReq.ProtoReflect.Descriptor instead.
func (*EquipValueReq) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *EquipValueReq) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *EquipValueReq) GetVirtualId() uint32 {
	if x != nil {
		return x.VirtualId
	}
	return 0
}

func (x *EquipValueReq) GetType() ValueType {
	if x != nil {
		return x.Type
	}
	return ValueType_Sensor
}

func (x *EquipValueReq) GetValues() map[uint32]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *EquipValueReq) GetLastestTime() uint32 {
	if x != nil {
		return x.LastestTime
	}
	return 0
}

func (x *EquipValueReq) GetOnlineState() string {
	if x != nil {
		return x.OnlineState
	}
	return ""
}

func (x *EquipValueReq) GetNowTime() uint32 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

type Frequency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Minutes uint32 `protobuf:"varint,1,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Times   uint32 `protobuf:"varint,2,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *Frequency) Reset() {
	*x = Frequency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frequency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frequency) ProtoMessage() {}

func (x *Frequency) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frequency.ProtoReflect.Descriptor instead.
func (*Frequency) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *Frequency) GetMinutes() uint32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *Frequency) GetTimes() uint32 {
	if x != nil {
		return x.Times
	}
	return 0
}

type ValConds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string  `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ValConds) Reset() {
	*x = ValConds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValConds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValConds) ProtoMessage() {}

func (x *ValConds) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValConds.ProtoReflect.Descriptor instead.
func (*ValConds) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *ValConds) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *ValConds) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Frequency *Frequency  `protobuf:"bytes,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	ValCons   []*ValConds `protobuf:"bytes,3,rep,name=valCons,proto3" json:"valCons,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *Conds) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Conds) GetFrequency() *Frequency {
	if x != nil {
		return x.Frequency
	}
	return nil
}

func (x *Conds) GetValCons() []*ValConds {
	if x != nil {
		return x.ValCons
	}
	return nil
}

type SensorObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Display string `protobuf:"bytes,1,opt,name=display,proto3" json:"display,omitempty"`
	Unit    string `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SensorObj) Reset() {
	*x = SensorObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorObj) ProtoMessage() {}

func (x *SensorObj) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorObj.ProtoReflect.Descriptor instead.
func (*SensorObj) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{4}
}

func (x *SensorObj) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *SensorObj) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *SensorObj) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type WarningConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type   string     `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Conds  []*Conds   `protobuf:"bytes,3,rep,name=conds,proto3" json:"conds,omitempty"`
	Sensor *SensorObj `protobuf:"bytes,4,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *WarningConfigs) Reset() {
	*x = WarningConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarningConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarningConfigs) ProtoMessage() {}

func (x *WarningConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarningConfigs.ProtoReflect.Descriptor instead.
func (*WarningConfigs) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{5}
}

func (x *WarningConfigs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WarningConfigs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *WarningConfigs) GetConds() []*Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *WarningConfigs) GetSensor() *SensorObj {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type PermPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Role    string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *PermPool) Reset() {
	*x = PermPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermPool) ProtoMessage() {}

func (x *PermPool) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermPool.ProtoReflect.Descriptor instead.
func (*PermPool) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{6}
}

func (x *PermPool) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PermPool) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Equip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PermPool []*PermPool       `protobuf:"bytes,3,rep,name=permPool,proto3" json:"permPool,omitempty"`
	Warnings []*WarningConfigs `protobuf:"bytes,4,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *Equip) Reset() {
	*x = Equip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Equip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Equip) ProtoMessage() {}

func (x *Equip) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Equip.ProtoReflect.Descriptor instead.
func (*Equip) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{7}
}

func (x *Equip) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Equip) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Equip) GetPermPool() []*PermPool {
	if x != nil {
		return x.PermPool
	}
	return nil
}

func (x *Equip) GetWarnings() []*WarningConfigs {
	if x != nil {
		return x.Warnings
	}
	return nil
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{8}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WarningCheckingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Equips  []*Equip `protobuf:"bytes,2,rep,name=equips,proto3" json:"equips,omitempty"`
}

func (x *WarningCheckingReq) Reset() {
	*x = WarningCheckingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_channel_consumer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarningCheckingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarningCheckingReq) ProtoMessage() {}

func (x *WarningCheckingReq) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_channel_consumer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarningCheckingReq.ProtoReflect.Descriptor instead.
func (*WarningCheckingReq) Descriptor() ([]byte, []int) {
	return file_channel_proto_channel_consumer_proto_rawDescGZIP(), []int{9}
}

func (x *WarningCheckingReq) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *WarningCheckingReq) GetEquips() []*Equip {
	if x != nil {
		return x.Equips
	}
	return nil
}

var File_channel_proto_channel_consumer_proto protoreflect.FileDescriptor

var file_channel_proto_channel_consumer_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a,
	0x0d, 0x45, 0x71, 0x75, 0x69, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63,
	0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65,
	0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x09, 0x46,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x56, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x07, 0x76,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x22, 0x4d, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x4f, 0x62, 0x6a, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x38,
	0x0a, 0x08, 0x50, 0x65, 0x72, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x6d, 0x50, 0x6f,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x65, 0x72,
	0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x12, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x06, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x73, 0x2a, 0x27, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x10, 0x01, 0x32, 0x9d, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x0c, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_channel_proto_channel_consumer_proto_rawDescOnce sync.Once
	file_channel_proto_channel_consumer_proto_rawDescData = file_channel_proto_channel_consumer_proto_rawDesc
)

func file_channel_proto_channel_consumer_proto_rawDescGZIP() []byte {
	file_channel_proto_channel_consumer_proto_rawDescOnce.Do(func() {
		file_channel_proto_channel_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_channel_consumer_proto_rawDescData)
	})
	return file_channel_proto_channel_consumer_proto_rawDescData
}

var file_channel_proto_channel_consumer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_channel_proto_channel_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_channel_proto_channel_consumer_proto_goTypes = []interface{}{
	(ValueType)(0),             // 0: channel.ValueType
	(*EquipValueReq)(nil),      // 1: channel.EquipValueReq
	(*Frequency)(nil),          // 2: channel.Frequency
	(*ValConds)(nil),           // 3: channel.ValConds
	(*Conds)(nil),              // 4: channel.Conds
	(*SensorObj)(nil),          // 5: channel.SensorObj
	(*WarningConfigs)(nil),     // 6: channel.WarningConfigs
	(*PermPool)(nil),           // 7: channel.PermPool
	(*Equip)(nil),              // 8: channel.Equip
	(*Project)(nil),            // 9: channel.Project
	(*WarningCheckingReq)(nil), // 10: channel.WarningCheckingReq
	nil,                        // 11: channel.EquipValueReq.ValuesEntry
	(*emptypb.Empty)(nil),      // 12: google.protobuf.Empty
}
var file_channel_proto_channel_consumer_proto_depIdxs = []int32{
	0,  // 0: channel.EquipValueReq.type:type_name -> channel.ValueType
	11, // 1: channel.EquipValueReq.values:type_name -> channel.EquipValueReq.ValuesEntry
	2,  // 2: channel.Conds.frequency:type_name -> channel.Frequency
	3,  // 3: channel.Conds.valCons:type_name -> channel.ValConds
	4,  // 4: channel.WarningConfigs.conds:type_name -> channel.Conds
	5,  // 5: channel.WarningConfigs.sensor:type_name -> channel.SensorObj
	7,  // 6: channel.Equip.permPool:type_name -> channel.PermPool
	6,  // 7: channel.Equip.warnings:type_name -> channel.WarningConfigs
	9,  // 8: channel.WarningCheckingReq.project:type_name -> channel.Project
	8,  // 9: channel.WarningCheckingReq.equips:type_name -> channel.Equip
	1,  // 10: channel.ComsumerService.EquipRawdata:input_type -> channel.EquipValueReq
	10, // 11: channel.ComsumerService.WarningChecking:input_type -> channel.WarningCheckingReq
	12, // 12: channel.ComsumerService.EquipRawdata:output_type -> google.protobuf.Empty
	12, // 13: channel.ComsumerService.WarningChecking:output_type -> google.protobuf.Empty
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_channel_proto_channel_consumer_proto_init() }
func file_channel_proto_channel_consumer_proto_init() {
	if File_channel_proto_channel_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_channel_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipValueReq); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frequency); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValConds); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorObj); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarningConfigs); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermPool); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Equip); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_channel_proto_channel_consumer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarningCheckingReq); i {
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
			RawDescriptor: file_channel_proto_channel_consumer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_channel_consumer_proto_goTypes,
		DependencyIndexes: file_channel_proto_channel_consumer_proto_depIdxs,
		EnumInfos:         file_channel_proto_channel_consumer_proto_enumTypes,
		MessageInfos:      file_channel_proto_channel_consumer_proto_msgTypes,
	}.Build()
	File_channel_proto_channel_consumer_proto = out.File
	file_channel_proto_channel_consumer_proto_rawDesc = nil
	file_channel_proto_channel_consumer_proto_goTypes = nil
	file_channel_proto_channel_consumer_proto_depIdxs = nil
}