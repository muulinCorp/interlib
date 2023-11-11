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

type UpdateRawdataStreamClient interface {
	core.MyGrpc
	StartUpdateRawdataStream(resp chan *pb.Response)
	UpdateRawdata(*pb.UpdateRawdataRequest) error
	StopUpdateRawdataStream() error
}

func NewUpdateRawdataStreamClient(address string, l log.Logger) (UpdateRawdataStreamClient, error) {

	return &updateRawdataStreamSdkImpl{
		AutoReConn: core.NewAutoReconn(address),
	}, nil
}

type updateRawdataStreamSdkImpl struct {
	*core.AutoReConn

	updateRawdataStream pb.DeviceService_UpdateRawdataClient
}

func (impl *updateRawdataStreamSdkImpl) StartUpdateRawdataStream(resp chan *pb.Response) {
	var err error
	p := func(myGrpc core.MyGrpc) error {
		clt := pb.NewDeviceServiceClient(impl)
		impl.updateRawdataStream, err = clt.UpdateRawdata(context.Background())
		if err != nil {
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.updateRawdataStream.Recv()
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

func (grpc *updateRawdataStreamSdkImpl) StopUpdateRawdataStream() error {
	if grpc.updateRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.updateRawdataStream.CloseSend()
}

func (grpc *updateRawdataStreamSdkImpl) UpdateRawdata(req *pb.UpdateRawdataRequest) error {
	if grpc.updateRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputUpdateRawdataReq{
		UpdateRawdataRequest: req,
	}
	err := in.Validate()
	if err != nil {
		return err
	}
	return grpc.updateRawdataStream.Send(req)
}

type inputUpdateRawdataReq struct {
	*pb.UpdateRawdataRequest
}

func (req *inputUpdateRawdataReq) Validate() error {
	if req.Data.Mac == "" {
		return errors.New("invalid mac")
	}
	if req.Data.GwID == "" {
		return errors.New("invalid gwid")
	}
	return nil
}
