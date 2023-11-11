package Client

import (
	"bitbucket.org/muulin/interlib/configuration/pb"
	"bitbucket.org/muulin/interlib/core"
	"golang.org/x/net/context"
)

type Client interface {
	GetChannelConf(ctx context.Context, req *pb.GetConfRequest) ([]byte, error)
}

func New(address string) (Client, error) {
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
