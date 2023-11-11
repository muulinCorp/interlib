package client

import (
	"errors"
	"io"
	"net/http"

	"bitbucket.org/muulin/interlib/device/pb"

	"bitbucket.org/muulin/interlib/core"
	"github.com/94peter/sterna/log"
	"golang.org/x/net/context"
)

type UpdateRealtimeStreamClient interface {
	core.MyGrpc
	StartUpdateRealtimeStream(resp chan *pb.Response)
	UpdateRealtime(*pb.UpdateRawdataRequest) error
	StopUpdateRealtimeStream() error
}

func NewUpdateRealtimeStreamClient(address string, l log.Logger) (UpdateRealtimeStreamClient, error) {

	return &updateRealtimeStreamSdkImpl{
		AutoReConn: core.NewAutoReconn(address),
	}, nil
}

type updateRealtimeStreamSdkImpl struct {
	*core.AutoReConn

	updateRealtimeStream pb.DeviceService_UpdateRawdataClient
}

func (impl *updateRealtimeStreamSdkImpl) StartUpdateRealtimeStream(resp chan *pb.Response) {
	var err error
	p := func(myGrpc core.MyGrpc) error {
		clt := pb.NewDeviceServiceClient(impl)
		impl.updateRealtimeStream, err = clt.UpdateRealtime(context.Background())
		if err != nil {
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.updateRealtimeStream.Recv()
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
			resp <- &pb.Response{
				StatusCode: http.StatusOK,
				Message:    "ready to send data",
			}
		case <-impl.Reconnect:
			if !impl.WaitUntilReady() {
				resp <- &pb.Response{
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

func (grpc *updateRealtimeStreamSdkImpl) StopUpdateRealtimeStream() error {
	if grpc.updateRealtimeStream == nil {
		return errors.New("StartUpdateRealtimeStream first")
	}
	return grpc.updateRealtimeStream.CloseSend()
}

func (grpc *updateRealtimeStreamSdkImpl) UpdateRealtime(req *pb.UpdateRawdataRequest) error {
	if grpc.updateRealtimeStream == nil {
		return errors.New("StartUpdateRealtimeStream first")
	}
	in := inputUpdateRawdataReq{
		UpdateRawdataRequest: req,
	}
	err := in.Validate()
	if err != nil {
		return err
	}
	return grpc.updateRealtimeStream.Send(req)
}
