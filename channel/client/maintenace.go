package client

import (
	"context"

	"bitbucket.org/muulin/interlib/channel/pb"

	"bitbucket.org/muulin/interlib/core"
	"google.golang.org/grpc/metadata"
)

type MaintenaceGrpcClient interface {
	GetEquipInfo(equipId, sensorId string) (*pb.EquipInfoResponse, error)
	GetEquipIdsByAccount(acc string) ([]string, error)
	EmitEvent(*pb.MaintenanceEventReq) error
}

func NewMaintenaceGrpcClient(address string, channel string) MaintenaceGrpcClient {
	return &maintenaceGrpcClientImpl{
		channel: channel,
		address: address,
	}
}

type maintenaceGrpcClientImpl struct {
	channel string
	address string
}

func (c *maintenaceGrpcClientImpl) GetEquipInfo(equipId, sensorId string) (*pb.EquipInfoResponse, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcClt, err := core.NewMyGrpc(c.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()
	clt := pb.NewMaintenaceServiceClient(grpcClt)

	resp, err := clt.GetEquipInfo(ctx, &pb.EquipInfoRequest{
		EquipId:  equipId,
		SensorId: sensorId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *maintenaceGrpcClientImpl) GetEquipIdsByAccount(acc string) ([]string, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcClt, err := core.NewMyGrpc(c.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()
	clt := pb.NewMaintenaceServiceClient(grpcClt)

	resp, err := clt.GetEquipIdsByAccount(ctx, &pb.GetEquipIdsByAccountRequest{
		Account: acc,
	})
	if err != nil {
		return nil, err
	}
	return resp.EquipIds, nil
}

func (c *maintenaceGrpcClientImpl) EmitEvent(req *pb.MaintenanceEventReq) error {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcClt, err := core.NewMyGrpc(c.address)
	if err != nil {
		return err
	}
	defer grpcClt.Close()
	clt := pb.NewMaintenaceServiceClient(grpcClt)
	_, err = clt.EmitEvent(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
