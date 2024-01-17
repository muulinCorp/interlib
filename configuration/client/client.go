package client

import (
	"github.com/muulinCorp/interlib/configuration/pb"
	"github.com/muulinCorp/interlib/core"
	"golang.org/x/net/context"
)

type ConfigurationClient interface {
	GetChannelConf(ctx context.Context, req *pb.GetConfRequest) ([]byte, error)
}

func New(address string) (ConfigurationClient, error) {
	return &clientImpl{
		address: address,
	}, nil
}

type clientImpl struct {
	address string
}

func (impl *clientImpl) GetChannelConf(ctx context.Context, req *pb.GetConfRequest) ([]byte, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewConfigurationServiceClient(grpc)
	resp, err := clt.GetChannelConf(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
