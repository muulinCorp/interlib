package sdk

import (
	"bitbucket.org/muulin/interlib/core"
	"bitbucket.org/muulin/interlib/notify/pb"
	"golang.org/x/net/context"
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
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewMailServiceClient(grpc)
	_, err = clt.SingleMail(ctx, req)
	return err
}
