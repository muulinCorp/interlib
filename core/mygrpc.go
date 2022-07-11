package core

import "google.golang.org/grpc"

type MyGrpc interface {
	grpc.ClientConnInterface
	Close()
}

func NewMyGrpc(address string) (MyGrpc, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return &myGrpcImpl{
		ClientConn: conn,
	}, nil
}

type myGrpcImpl struct {
	*grpc.ClientConn
}

func (my *myGrpcImpl) Close() {
	my.Close()
}
