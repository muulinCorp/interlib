package client

import (
	"github.com/muulinCorp/interlib/device/pb"

	"github.com/muulinCorp/interlib/core"
	"golang.org/x/net/context"
)

type DeviceV2Client interface {
	CheckState([]*pb.Device) (map[string]*pb.DeviceState, error)
}

func NewDeviceV2Client(address string) DeviceV2Client {
	return &deviceV2SdkImpl{
		address: address,
	}
}

type deviceV2SdkImpl struct {
	address string
}

func (impl *deviceV2SdkImpl) CheckState(devices []*pb.Device) (map[string]*pb.DeviceState, error) {
	var err error
	mygrpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return nil, err
	}
	defer mygrpc.Close()
	clt := pb.NewDeviceV2ServiceClient(mygrpc)
	resp, err := clt.CheckState(context.Background(), &pb.GetStateRequest{Devices: devices})
	if err != nil {
		return nil, err
	}
	return resp.StateMap, nil
}
