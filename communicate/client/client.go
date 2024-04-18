package client

import (
	"context"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/communicate/pb"
	"github.com/pkg/errors"
)

type CommunicateClient interface {
	Remote(ctx context.Context, data map[string]float64) error
	GetSensors(ctx context.Context, names []string) (*pb.GetSensorsResponse, error)
	GetSensorsReader(ctx context.Context, names []string) (SensorReader, error)
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
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
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

func (impl *clientImpl) GetSensors(ctx context.Context, names []string) (*pb.GetSensorsResponse, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, errors.Wrap(err, "new grpc fail")
	}
	defer grpc.Close()
	clt := pb.NewCommunicateServiceClient(grpc)

	return clt.GetSensors(ctx, &pb.GetSensorsRequest{Name: names})
}

func (impl *clientImpl) GetSensorsReader(ctx context.Context, names []string) (SensorReader, error) {
	resp, err := impl.GetSensors(ctx, names)
	if err != nil {
		return nil, err
	}
	return NewSensorReader(resp.Sensors), nil
}
