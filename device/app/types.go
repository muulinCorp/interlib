package appDevice

import (
	"time"

	pb "bitbucket.org/muulin/interlib/device/app/service"
)

const (
	RouterKey = "gRPC_App_Device_Router"
)

type Device struct {
	Id        string
	Mac       string
	VirtualID uint8
	Model     string
}

type DeviceAry []*Device

func (da DeviceAry) getDevices() (result []*pb.Device) {
	for _, d := range da {
		result = append(result, &pb.Device{
			Id:        d.Id,
			Mac:       d.Mac,
			VirtualID: uint32(d.VirtualID),
			Model:     d.Model,
		})
	}
	return
}

type TxnType string

const (
	TxnTypeSend    = TxnType("sending")
	TxnTypeRecycle = TxnType("recycling")
)

type TxnState string

const (
	TxnStateNew    = TxnState("new")
	TxnStateDone   = TxnState("confirmed")
	TxnStateCancel = TxnState("canceled")
)

type QueryTxnRequest struct {
	Typ       TxnType
	State     TxnState
	StartTime time.Time
	EndTime   time.Time
}

type TxnEditType string

const (
	TxnEditTypeAdd = TxnEditType("add")
	TxnEditTypeDel = TxnEditType("del")
)

type DeviceInfo struct {
	Mac       string
	VirtualID uint8
	Model     string
	Desc      string
}

type TxnAct struct {
	Edit   TxnEditType
	Device *DeviceInfo
}
