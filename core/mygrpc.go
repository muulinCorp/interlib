package core

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type MyGrpc interface {
	grpc.ClientConnInterface
	Close() error
	IsValid() bool
	WaitUntilReady() bool
}

func NewMyGrpc(address string) (MyGrpc, error) {

	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*10),
	)
	if err != nil {
		return nil, fmt.Errorf("address [%s] error: %s", address, err.Error())
	}
	return &myGrpcImpl{
		ClientConn: conn,
	}, nil
}

type myGrpcImpl struct {
	*grpc.ClientConn
}

func (my *myGrpcImpl) Close() error {
	return my.ClientConn.Close()
}

func (my *myGrpcImpl) IsValid() bool {
	if my.ClientConn == nil {
		return false
	}
	switch my.ClientConn.GetState() {
	case connectivity.Ready:
		return true
	case connectivity.Idle:
		return false
	default:
		return false
	}
}

func (my *myGrpcImpl) WaitUntilReady() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second) //define how long you want to wait for connection to be restored before giving up
	defer cancel()
	return my.WaitForStateChange(ctx, connectivity.Ready)
}

func NewAutoReconn(address string) *AutoReConn {
	return &AutoReConn{
		address:   address,
		Ready:     make(chan bool),
		Done:      make(chan bool),
		Reconnect: make(chan bool),
	}
}

type AutoReConn struct {
	MyGrpc

	address string

	Ready     chan bool
	Done      chan bool
	Reconnect chan bool
}

type GetGrpcFunc func(myGrpc MyGrpc) error

func (my *AutoReConn) Connect() (MyGrpc, error) {
	return NewMyGrpc(my.address)
}

func (my *AutoReConn) IsValid() bool {
	if my.MyGrpc == nil {
		return false
	}
	return my.MyGrpc.IsValid()
}

func (my *AutoReConn) Process(f GetGrpcFunc) {
	var err error
	for {
		defer time.Sleep(time.Second)
		my.MyGrpc, err = my.Connect()
		if err != nil {
			continue
		}
		if err = f(my.MyGrpc); err != nil {
			continue
		}
		break
	}
}
