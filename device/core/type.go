package coreDevice

import (
	pb "bitbucket.org/muulin/interlib/device/core/service"
)

type UpdateRawdataType string

const (
	RawdataType_Controller = UpdateRawdataType("controller")
	RawdataType_Sensor     = UpdateRawdataType("sensor")
)

func (rt *UpdateRawdataType) getRawdataRequestType() pb.UpdateRawdataRequest_DataType {
	switch *rt {
	case RawdataType_Controller:
		return pb.UpdateRawdataRequest_Controller
	default:
		return pb.UpdateRawdataRequest_Sensor
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
