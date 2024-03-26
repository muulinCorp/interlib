package client

import (
	"context"

	"github.com/muulinCorp/interlib/channel/pb"

	"github.com/muulinCorp/interlib/core"
	"google.golang.org/grpc/metadata"
)

type MaintenaceGrpcClient interface {
	GetEquipInfo(ctx context.Context, equipId string) (*pb.EquipInfoResponse, error)
	GetEquipIdsByAccount(ctx context.Context, acc string) ([]string, error)
	EmitEvent(context.Context, *pb.MaintenanceEventReq) error
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

func (c *maintenaceGrpcClientImpl) GetEquipInfo(ctx context.Context, equipId string) (*pb.EquipInfoResponse, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx = metadata.NewOutgoingContext(ctx, md)
	grpcClt, err := core.NewMyGrpc(ctx, c.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()
	clt := pb.NewMaintenaceServiceClient(grpcClt)

	resp, err := clt.GetEquipInfo(ctx, &pb.EquipInfoRequest{
		EquipId: equipId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *maintenaceGrpcClientImpl) GetEquipIdsByAccount(ctx context.Context, acc string) ([]string, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx = metadata.NewOutgoingContext(ctx, md)
	grpcClt, err := core.NewMyGrpc(ctx, c.address)
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

func (c *maintenaceGrpcClientImpl) EmitEvent(ctx context.Context, req *pb.MaintenanceEventReq) error {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx = metadata.NewOutgoingContext(ctx, md)
	grpcClt, err := core.NewMyGrpc(ctx, c.address)
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
