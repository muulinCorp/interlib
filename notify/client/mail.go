package sdk

import (
	"context"

	"github.com/muulinCorp/interlib/notify/pb"

	"github.com/94peter/microservice/grpc_tool"
)

type MailClient interface {
	SingleMail(ctx context.Context, req *pb.SingleMailRequest) error
}

func NewMailClient(address string) (MailClient, error) {
	return &mailSdkImpl{
		address: address,
	}, nil
}

type mailSdkImpl struct {
	address string
}

func (impl *mailSdkImpl) SingleMail(ctx context.Context, req *pb.SingleMailRequest) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewMailServiceClient(grpc)
	_, err = clt.SingleMail(ctx, req)
	return err
}
