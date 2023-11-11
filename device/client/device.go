package client

import (
	"errors"
	"net/http"

	"bitbucket.org/muulin/interlib/device/pb"

	"bitbucket.org/muulin/interlib/core"
	"golang.org/x/net/context"
)

type DeviceClient interface {
	GetVirtualId(mac, gwid string) (uint8, error)
	SetTime(mac string, virtualId uint8) error
	Remote(mac string, virtualId uint8, deviceNo uint8, address uint8, val float64) *pb.RemoteResponse
}

func NewDeviceClient(address string) DeviceClient {
	return &deviceSdkImpl{
		address: address,
	}
}

type deviceSdkImpl struct {
	address string
}

func (grpc *deviceSdkImpl) GetVirtualId(mac, gwid string) (uint8, error) {
	var err error

	grpcClt, err := core.NewMyGrpc(grpc.address)
	if err != nil {
		return 0, err
	}
	defer grpcClt.Close()
	clt := pb.NewDeviceServiceClient(grpcClt)
	resp, err := clt.GetVritualId(context.Background(), &pb.GetVirtualIdRequest{
		Mac:  mac,
		GwID: gwid,
	})
	if err != nil {
		return 0, err
	}
	return uint8(resp.VirtualID), nil
}

func (grpc *deviceSdkImpl) SetTime(mac string, virtualId uint8) error {
	var err error

	grpcClt, err := core.NewMyGrpc(grpc.address)
	if err != nil {
		return err
	}
	defer grpcClt.Close()
	clt := pb.NewDeviceServiceClient(grpcClt)
	r, err := clt.SetTime(context.Background(), &pb.Device{
		Mac:       mac,
		VirtualID: uint32(virtualId),
	})
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Message)
	}
	return nil
}

func (grpc *deviceSdkImpl) Remote(mac string, virtualId uint8, deviceNo uint8, address uint8, val float64) *pb.RemoteResponse {
	var err error

	grpcClt, err := core.NewMyGrpc(grpc.address)
	if err != nil {
		return &pb.RemoteResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	defer grpcClt.Close()
	clt := pb.NewDeviceServiceClient(grpcClt)
	r, err := clt.Remote(context.Background(), &pb.RemoteRequest{
		Device: &pb.Device{
			Mac:       mac,
			VirtualID: uint32(virtualId),
		},
		DeviceNo: uint32(deviceNo),
		Address:  uint32(address),
		Value:    val,
	})
	if err != nil {
		return &pb.RemoteResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return r
}
