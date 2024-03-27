package client

import (
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/muulinCorp/interlib/device/pb"

	"github.com/muulinCorp/interlib/core"
	"golang.org/x/net/context"
)

type UpdateRealtimeStreamClient interface {
	core.MyGrpc
	StartUpdateRealtimeStream(resp chan *pb.Response)
	UpdateRealtime(*pb.UpdateRawdataRequest) error
	StopUpdateRealtimeStream() error
}

func NewUpdateRealtimeStreamClient(ctx context.Context, address string) UpdateRealtimeStreamClient {
	return &updateRealtimeStreamSdkImpl{
		AutoReConn: core.NewAutoReconn(ctx, address),
	}
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
