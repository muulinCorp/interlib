package client

import (
	"context"
	"io"
	"sync"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/com2scada/pb"
)

type ChangeDataStreamClient interface {
	StartChangeDataStream(fields []string, resp chan map[string]float64) error
	StopStream() error
}

func NewChangeDataStreamClient(address string) ChangeDataStreamClient {
	return &changeDataStreamImpl{
		address: address,
	}
}

type changeDataStreamImpl struct {
	address string
	stream  pb.Com2ScadaService_ChangeDataStreamClient
}

func (impl *changeDataStreamImpl) StartChangeDataStream(fields []string, resp chan map[string]float64) error {
	var err error
	ctx := context.Background()
	grpcClt, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpcClt.Close()
	service := pb.NewCom2ScadaServiceClient(grpcClt)
	impl.stream, err = service.ChangeDataStream(ctx, &pb.FieldList{Fields: fields})
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

func (impl *changeDataStreamImpl) StopStream() error {
	return impl.stream.CloseSend()
}
