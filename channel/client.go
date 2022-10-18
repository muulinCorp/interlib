package channel

import (
	"context"

	pb "bitbucket.org/muulin/interlib/channel/service"
	"bitbucket.org/muulin/interlib/core"
)

type ChannelClient interface {
	core.MyGrpc
	GetConfCacheKey(host, env string) (string, error)
	IsExist(host string) (bool, error)
}

func NewGrpcClient(address string) (ChannelClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
}

func (gclt *grpcClt) GetConfCacheKey(host, env string) (string, error) {
	clt := pb.NewChannelConfClient(gclt)

	resp, err := clt.GetConfCacheKey(context.Background(), &pb.GetConfRequest{
		Host: host,
		Env:  env,
	})
	if err != nil {
		return "", err
	}

	return resp.Key, nil
}

func (gclt *grpcClt) IsExist(host string) (bool, error) {
	clt := pb.NewChannelConfClient(gclt)

	resp, err := clt.IsExist(context.Background(), &pb.IsExistRequest{Host: host})
	if err != nil {
		return false, err
	}
	return resp.IsExist, nil
}
