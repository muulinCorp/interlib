package client

import (
	"context"
	"io"
	"sync"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/com2scada/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChangeDataStreamClient interface {
	grpc_tool.AutoReConnInter
	StartChangeDataStream(fields []string, resp chan map[string]float64)
	StopStream() error
}

func NewChangeDataStreamClient(address string) ChangeDataStreamClient {
	return &changeDataStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address),
	}
}

type changeDataStreamImpl struct {
	*grpc_tool.AutoReConn
	stream pb.Com2ScadaService_ChangeDataStreamClient
}

func (impl *changeDataStreamImpl) StartChangeDataStream(fields []string, resp chan map[string]float64) {
	execFunc := func(grpcClt grpc_tool.Connection) error {
		defer grpcClt.Close()
		service := pb.NewCom2ScadaServiceClient(grpcClt)
		var err error
		impl.stream, err = service.ChangeDataStream(context.Background(), &pb.FieldList{Fields: fields})
		if err != nil {
			impl.Reconnect <- true
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.stream.Recv()
			if err == io.EOF {
				impl.Done <- true
				return nil
			}
			if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					impl.Reconnect <- true
					return err
				}
				if status.Code() == codes.NotFound {
					impl.Error <- err
					return err
				}

				impl.Reconnect <- true
				return err
			}
			resp <- in.Values
		}
	}
	// for loop connect
	impl.Start(execFunc)
}

func (impl *changeDataStreamImpl) StopStream() error {
	err := impl.stream.CloseSend()
	if err != nil {
		return err
	}
	return impl.AutoReConn.Close()
}

type RoutineDataStreamClient interface {
	StartRoutineDataStream(fields []string, resp chan map[string]float64) error
	StopStream() error
}

func NewRoutineDataStreamClient(address string) RoutineDataStreamClient {
	return &routineDataStreamImpl{
		address: address,
	}
}

type routineDataStreamImpl struct {
	address string
	stream  pb.Com2ScadaService_RoutineDataStreamClient
}

func (impl *routineDataStreamImpl) StartRoutineDataStream(fields []string, resp chan map[string]float64) error {
	var err error
	ctx := context.Background()
	grpcClt, err := grpc_tool.NewConnection(impl.address)
	if err != nil {
		return err
	}
	defer grpcClt.Close()
	service := pb.NewCom2ScadaServiceClient(grpcClt)
	impl.stream, err = service.RoutineDataStream(ctx, &pb.FieldList{Fields: fields})
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			in, err := impl.stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
			resp <- in.Values
		}
	}()
	wg.Wait()
	return nil
}

func (impl *routineDataStreamImpl) StopStream() error {
	return impl.stream.CloseSend()
}
