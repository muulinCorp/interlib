package client

import (
	"errors"
	"io"
	"net/http"
	"time"

	"context"

	"github.com/94peter/micro-service/grpc_tool"
	"github.com/muulinCorp/interlib/device/pb"
)

type GetVirtualIdStreamClient interface {
	grpc_tool.Connection

	StartGetVirtualIdStream(resp chan *pb.GetVirtualIdStreamResponse)
	GetVirtualReq(mac, gwid string) error
	StopGetVirtualIdStream() error
}

func NewVirutalIdStreamClient(address string, timeout time.Duration) GetVirtualIdStreamClient {
	return &virtualIDStream{
		AutoReConn: grpc_tool.NewAutoReconn(address, timeout),
	}
}

type virtualIDStream struct {
	*grpc_tool.AutoReConn
	getVirtualIdStream pb.DeviceService_GetVritualIdStreamClient
}

func (impl *virtualIDStream) StartGetVirtualIdStream(resp chan *pb.GetVirtualIdStreamResponse) {
	var err error
	p := func(myGrpc grpc_tool.Connection) error {
		clt := pb.NewDeviceServiceClient(impl)
		impl.getVirtualIdStream, err = clt.GetVritualIdStream(context.Background())
		if err != nil {
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.getVirtualIdStream.Recv()
			if err == io.EOF {
				impl.Done <- true
				return nil
			}
			if err != nil {
				impl.Reconnect <- true
				return err
			}
			resp <- in
		}
	}
	go impl.Process(p)
	for {
		select {
		case <-impl.Ready:
			resp <- &pb.GetVirtualIdStreamResponse{
				StatusCode: http.StatusOK,
				Message:    "ready to send data",
			}
		case <-impl.Reconnect:
			if !impl.WaitUntilReady() {
				resp <- &pb.GetVirtualIdStreamResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "failed to establish a connection within the defined timeout",
				}
			}
			go impl.Process(p)
		case <-impl.Done:
			impl.Close()
			return
		}
	}
}

func (grpc *virtualIDStream) GetVirtualReq(mac, gwid string) error {
	if grpc.getVirtualIdStream == nil {
		return errors.New("StartGetVirtualIdStream first")
	}
	return grpc.getVirtualIdStream.Send(&pb.GetVirtualIdRequest{
		Mac:  mac,
		GwID: gwid,
	})
}

func (grpc *virtualIDStream) StopGetVirtualIdStream() error {
	if grpc.getVirtualIdStream == nil {
		return errors.New("StartGetVirtualIdStream first")
	}
	return grpc.getVirtualIdStream.CloseSend()
}
