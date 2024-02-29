package client

import (
	"context"

	"github.com/muulinCorp/interlib/communicate/pb"
	"github.com/muulinCorp/interlib/core"
	"github.com/pkg/errors"
)

type CommunicateClient interface {
	Remote(ctx context.Context, data map[string]float64) error
}

func New(address string) CommunicateClient {
	return &clientImpl{
		address: address,
	}
}

type clientImpl struct {
	address string
}

func (impl *clientImpl) Remote(ctx context.Context, data map[string]float64) error {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return errors.Wrap(err, "new grpc fail")
	}
	defer grpc.Close()
	clt := pb.NewCommunicateServiceClient(grpc)
	_, err = clt.Remote(ctx, &pb.RemoteRequest{Values: data})
	if err != nil {
		return errors.Wrap(err, "remote error")
	}
	return nil
}
