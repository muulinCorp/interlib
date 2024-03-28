package client

import (
	"errors"
	"net/http"

	"context"

	"github.com/94peter/micro-service/grpc_tool"
	"github.com/muulinCorp/interlib/device/pb"
)

type DeviceClient interface {
	GetVirtualId(ctx context.Context, mac, gwid string) (uint8, error)
	SetTime(ctx context.Context, mac string, virtualId uint8) error
	Remote(ctx context.Context, mac string, virtualId uint8, deviceNo uint8, address uint8, val float64) *pb.RemoteResponse
}

func NewDeviceClient(address string) DeviceClient {
	return &deviceSdkImpl{
		address: address,
	}
}

type deviceSdkImpl struct {
	address string
}

func (grpc *deviceSdkImpl) GetVirtualId(ctx context.Context, mac, gwid string) (uint8, error) {
	var err error

	grpcClt, err := grpc_tool.NewConnection(ctx, grpc.address)
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

func (grpc *deviceSdkImpl) SetTime(ctx context.Context, mac string, virtualId uint8) error {
	var err error

	grpcClt, err := grpc_tool.NewConnection(ctx, grpc.address)
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

func (grpc *deviceSdkImpl) Remote(ctx context.Context, mac string, virtualId uint8, deviceNo uint8, address uint8, val float64) *pb.RemoteResponse {
	var err error

	grpcClt, err := grpc_tool.NewConnection(ctx, grpc.address)
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
