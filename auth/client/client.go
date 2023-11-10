package client

import (
	"bitbucket.org/muulin/interlib/auth/pb"
	"bitbucket.org/muulin/interlib/core"
	"golang.org/x/net/context"
)

type AuthClient interface {
	GetUserInfo(ctx context.Context, resp *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
	GetAccount(ctx context.Context, id string) (string, error)
	AccountExist(ctx context.Context, acc string) (bool, error)
	CreateInvitation(ctx context.Context, email, name, channel string) (string, error)
	ForgetPwd(ctx context.Context, host, email string) (*pb.ForgetPasswordResponse, error)
}

func New(address string) AuthClient {
	return &authClientImpl{
		address: address,
	}
}

type authClientImpl struct {
	address string
}

func (impl *authClientImpl) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	return clt.GetUserInfo(ctx, req)
}

func (impl *authClientImpl) GetAccount(ctx context.Context, id string) (string, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return "", err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	resp, err := clt.GetAccount(ctx, &pb.GetAccountRequest{
		Id: id,
	})
	if err != nil {
		return "", err
	}
	return resp.Account, nil
}

func (impl *authClientImpl) AccountExist(ctx context.Context, acc string) (bool, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return false, err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	resp, err := clt.IsAccountExist(ctx, &pb.IsAccountExistRequest{
		Account: acc,
	})
	if err != nil {
		return false, err
	}
	return resp.Exists, nil
}

func (impl *authClientImpl) CreateInvitation(ctx context.Context, email, name, channel string) (string, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return "", err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	resp, err := clt.CreateInvitation(ctx, &pb.CreateInvitationRequest{
		Name:    name,
		Email:   email,
		Channel: channel,
	})
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (impl *authClientImpl) ForgetPwd(ctx context.Context, host, email string) (*pb.ForgetPasswordResponse, error) {
	grpc, err := core.NewMyGrpc(impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	return clt.ForgetPassword(ctx, &pb.ForgetPasswordRequest{
		Host:  host,
		Email: email,
	})
}
