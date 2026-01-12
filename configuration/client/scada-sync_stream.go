package client

import (
	"context"
	"io"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/configuration/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScadaSyncStreamClient interface {
	grpc_tool.AutoReConnInter
	StartListenSyncStream(context.Context, *pb.SyncConfigReq, chan *pb.SyncConfigResp, chan string)
	StopListenSyncStream() error
}

type scadaSyncStreamImpl struct {
	*grpc_tool.AutoReConn

	stream pb.ScadaSyncService_SyncConfigStreamClient
}

func NewScadaSyncStreamClient(address string) ScadaSyncStreamClient {
	return &scadaSyncStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address),
	}
}

func (i *scadaSyncStreamImpl) StopListenSyncStream() error {
	err := i.stream.CloseSend()
	if err != nil {
		return err
	}
	return i.AutoReConn.Close()
}

func (i *scadaSyncStreamImpl) StartListenSyncStream(
	ctx context.Context, req *pb.SyncConfigReq,
	respMsg chan *pb.SyncConfigResp, errMsg chan string,
) {
	var err error
	p := func(grpcClt grpc_tool.Connection) error {
		defer grpcClt.Close()
		service := pb.NewScadaSyncServiceClient(grpcClt)
		i.stream, err = service.SyncConfigStream(context.Background(), req)
		if err != nil {
			i.Reconnect <- true
			return err
		}
		i.Ready <- true
		for {
			in, err := i.stream.Recv()
			if err == io.EOF {
				i.Done <- true
				return nil
			}
			if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					i.Reconnect <- true
					return err
				}
				if status.Code() == codes.NotFound {
					i.Error <- err
					return err
				}

				i.Reconnect <- true
				return err
			}
			respMsg <- in
		}
	}

	go i.Process(p)
	for {
		select {
		case <-ctx.Done():
			i.Close()
			return
		case <-i.Ready:
			continue
		case <-i.Reconnect:
			if !i.WaitUntilReady() {
				continue
			}
		case <-i.Done:
			i.Close()
			return
		case err := <-i.Error:
			i.Close()
			errMsg <- err.Error()
			return
		}
	}
}
