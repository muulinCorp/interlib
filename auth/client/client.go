package client

import (
	"github.com/muulinCorp/interlib/auth/pb"
	"github.com/muulinCorp/interlib/core"
	"golang.org/x/net/context"
)

type AuthClient interface {
	GetTokenInfo(ctx context.Context, resp *pb.GetTokenInfoRequest) (*pb.GetTokenInfoResponse, error)
	GetAccount(ctx context.Context, id string) (string, error)
	AccountExist(ctx context.Context, acc string) (bool, error)
	CreateInvitation(ctx context.Context, email, name, channel string) (string, error)
	ForgetPwd(ctx context.Context, host, email string) (*pb.ForgetPasswordResponse, error)
	GetUserInfo(ctx context.Context, accoutOrEmail string) (*pb.GetUserInfoResponse, error)
}

func New(address string) AuthClient {
	return &authClientImpl{
		address: address,
	}
}

type authClientImpl struct {
	address string
}

func (impl *authClientImpl) GetTokenInfo(ctx context.Context, req *pb.GetTokenInfoRequest) (*pb.GetTokenInfoResponse, error) {
	grpc, err := core.NewMyGrpc(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)

	return clt.GetTokenInfo(ctx, req)
}

func (impl *authClientImpl) GetAccount(ctx context.Context, id string) (string, error) {
	grpc, err := core.NewMyGrpc(ctx, impl.address)
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
	grpc, err := core.NewMyGrpc(ctx, impl.address)
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
	grpc, err := core.NewMyGrpc(ctx, impl.address)
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
	grpc, err := core.NewMyGrpc(ctx, impl.address)
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

func (impl *authClientImpl) GetUserInfo(ctx context.Context, accoutOrEmail string) (*pb.GetUserInfoResponse, error) {
	grpc, err := core.NewMyGrpc(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewAuthServiceClient(grpc)
	return clt.GetUserInfoByAccount(ctx, &pb.GetUserInfoRequest{
		EmailOrAccount: accoutOrEmail,
	})
}
