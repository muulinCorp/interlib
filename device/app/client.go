package appDevice

import (
	"context"
	"io"

	"bitbucket.org/muulin/interlib/core"
	pb "bitbucket.org/muulin/interlib/device/app/service"
	"github.com/94peter/sterna/auth"
	"google.golang.org/grpc/metadata"
)

type AppDeviceClient interface {
	core.MyGrpc
	AssignDevices(channel string, devices DeviceAry, recvHandler func(suc bool, mac string, err string), reqUser auth.ReqUser) error
}

func NewGrpcClient(address string) (AppDeviceClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
}

func (gclt *grpcClt) AssignDevices(host string, devices DeviceAry, recvHandler func(suc bool, mac string, err string), reqUser auth.ReqUser) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Host", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	stream, err := clt.AssignDevices(ctx, &pb.AssignDevicesRequest{
		Devices: devices.getDevices(),
	})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		recvHandler(resp.Success, resp.Mac, resp.Error)
	}
	return nil
}
