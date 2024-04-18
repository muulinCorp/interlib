package client

import (
	"context"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/scada-layout/pb"
)

type ScadaLayoutClient interface {
	GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error)
}

type clientImpl struct {
	address string
}

func New(address string) ScadaLayoutClient {
	return &clientImpl{
		address: address,
	}
}

func (impl *clientImpl) GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	return clt.GetFieldsTags(ctx, &pb.FieldList{
		Fields: fields,
	})
}