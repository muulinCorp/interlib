package coreDevice

import (
	pb "bitbucket.org/muulin/interlib/device/core/service"
)

const (
	RouterKey = "gRPC_Core_Device_Router"
)

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
	// 已分配
	Assigned = DeviceState("assigned")
	// 使用中
	Used = DeviceState("used")
	// 待維修
	ToBeRepaired = DeviceState("2bRepaired")
)
