package appDevice

import (
	pb "bitbucket.org/muulin/interlib/device/app/service"
)

const (
	RouterKey = "gRPC_App_Device_Router"
)

type Device struct {
	Mac   string
	Model string
}

type DeviceAry []*Device

func (da DeviceAry) getDevices() (result []*pb.Device) {
	for _, d := range da {
		result = append(result, &pb.Device{
			Mac:   d.Mac,
			Model: d.Model,
		})
	}
	return
}
