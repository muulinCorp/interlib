package core

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type MyGrpc interface {
	grpc.ClientConnInterface
	Close() error
	IsValid() bool
	WaitUntilReady() bool
}

func NewMyGrpc(ctx context.Context, address string) (MyGrpc, error) {

	conn, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
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

func NewAutoReconn(ctx context.Context, address string) *AutoReConn {
	return &AutoReConn{
		ctx:       ctx,
		address:   address,
		Ready:     make(chan bool),
		Done:      make(chan bool),
		Reconnect: make(chan bool),
	}
}

type AutoReConn struct {
	MyGrpc

	ctx context.Context

	address string

	Ready     chan bool
	Done      chan bool
	Reconnect chan bool
}

type GetGrpcFunc func(myGrpc MyGrpc) error

func (my *AutoReConn) Connect() (MyGrpc, error) {
	return NewMyGrpc(my.ctx, my.address)
}

func (my *AutoReConn) IsValid() bool {
	if my.MyGrpc == nil {
		return false
	}
	return my.MyGrpc.IsValid()
}

func (my *AutoReConn) Process(f GetGrpcFunc) {
	var err error
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-my.ctx.Done():
			return
		case <-ticker.C:
			my.MyGrpc, err = my.Connect()
			if err != nil {
				continue
			}
			if err = f(my.MyGrpc); err != nil {
				continue
			}
		}
	}
}

func RunGrpcServ(ctx context.Context, cfg *GrpcConfig) error {
	if cfg.registerServiceFunc == nil {
		return fmt.Errorf("registerServiceFunc must not be nil")
	}
	port := ":" + strconv.Itoa(cfg.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	var serv *grpc.Server
	if len(cfg.interceptors) > 0 {
		var streamInterceptors []grpc.StreamServerInterceptor
		var unaryInterceptors []grpc.UnaryServerInterceptor
		for _, i := range cfg.interceptors {
			streamInterceptors = append(streamInterceptors, i.StreamServerInterceptor())
			unaryInterceptors = append(unaryInterceptors, i.UnaryServerInterceptor())
		}
		serv = grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		)
	} else {
		serv = grpc.NewServer()
	}
	if cfg.ReflectService {
		reflection.Register(serv)
	}
	cfg.registerServiceFunc(serv)
	var grpcWait sync.WaitGroup
	grpcWait.Add(1)
	go func(s *grpc.Server, lis net.Listener, l Log) {
		for {
			l.Infof("app gRPC server is running [%s].", lis.Addr())
			if err := s.Serve(lis); err != nil {
				switch err {
				case grpc.ErrServerStopped:
					grpcWait.Done()
					return
				default:
					l.Fatalf("failed to serve: %v", err)
				}
			}
		}
	}(serv, lis, cfg.Logger)
	<-ctx.Done()
	serv.Stop()
	grpcWait.Wait()
	return nil
}
