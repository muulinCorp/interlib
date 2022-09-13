package coreDevice

import (
	pb "bitbucket.org/muulin/interlib/device/core/service"
)

const (
	RouterKey = "gRPC_Core_Device_Router"
)

type Device struct {
	Macaddress string
	VirtualID  uint8
}

type DeviceAry []*Device

func (da DeviceAry) AddDevice(macaddress string, virtualID uint8) {
	da = append(da, &Device{Macaddress: macaddress, VirtualID: virtualID})
}

type RawdataType string

const (
	RawdataType_Controller = RawdataType("controller")
	RawdataType_Sensor     = RawdataType("sensor")
)

func (rt *RawdataType) getRawdataRequestType() pb.DataType {
	switch *rt {
	case RawdataType_Controller:
		return pb.DataType_Controller
	default:
		return pb.DataType_Sensor
	}
}

type SensorValue struct {
	Value float64
	DP    uint32
}

type SensorValuePool map[uint32]*SensorValue

func (pool SensorValuePool) getSensorValueMap() map[uint32]*pb.SensorValue {
	result := make(map[uint32]*pb.SensorValue)
	for k, v := range pool {
		result[k] = &pb.SensorValue{
			Value: v.Value,
			Dp:    v.DP,
		}
	}
	return result
}

type DeviceState string

const (
	// 入庫
	Stock = DeviceState("stock")
	// 配送中
	Sending = DeviceState("sending")
	// 已分配
	Assigned = DeviceState("assigned")
	// 使用中
	Used = DeviceState("used")
	// 待維修
	ToBeRepaired = DeviceState("2bRepaired")
)
