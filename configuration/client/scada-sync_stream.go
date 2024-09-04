package client

import (
	"context"
	"io"
	"sync"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/configuration/pb"
)

type ScadaSyncStreamClient interface {
	StartListenSyncStream(context.Context, *pb.SyncConfigReq, chan *pb.SyncConfigResp, chan string)
}

type scadaSyncStreamImpl struct {
	address string
	stream pb.ScadaSyncService_SyncConfigStreamClient
}

func NewScadaSyncStreamClient(address string) ScadaSyncStreamClient {
	return &scadaSyncStreamImpl{
		address: address,
	}
}

func (i *scadaSyncStreamImpl) StartListenSyncStream(
	ctx context.Context, req *pb.SyncConfigReq, 
	respMsg chan *pb.SyncConfigResp, errMsg chan string,
) {
	var err error
	grpcClt, err := grpc_tool.NewConnection(i.address)
	if err != nil {
		errMsg <- err.Error()
		return
	}
	defer grpcClt.Close()

	service := pb.NewScadaSyncServiceClient(grpcClt)
	i.stream, err = service.SyncConfigStream(ctx, req)
	if err != nil {
		errMsg <- err.Error()
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			in, err := i.stream.Recv()
			if err == io.EOF {
				errMsg <- "stream closed"
				return
			}
			if err != nil {
				errMsg <- err.Error()
				return
			}
			respMsg <- in
		}
	}()
	wg.Wait()
}