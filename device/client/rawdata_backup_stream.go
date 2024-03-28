package client

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/94peter/micro-service/grpc_tool"
	"github.com/muulinCorp/interlib/device/pb"

	"golang.org/x/net/context"
)

type BackupRawdataStreamClient interface {
	grpc_tool.Connection
	StartBackupRawdataStream(resp chan *pb.Response)
	BackupRawdata(*pb.UpdateRawdataRequest) error
	StopBackupRawdataStream() error
}

func NewBackupRawdataStreamClient(address string, timeout time.Duration) BackupRawdataStreamClient {

	return &backupRawdataStreamSdkImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address, timeout),
	}
}

type backupRawdataStreamSdkImpl struct {
	*grpc_tool.AutoReConn

	backupRawdataStream pb.DeviceService_BackupRawdataClient
}

func (impl *backupRawdataStreamSdkImpl) StartBackupRawdataStream(resp chan *pb.Response) {
	var err error
	p := func(myGrpc grpc_tool.Connection) error {
		clt := pb.NewDeviceServiceClient(impl)
		impl.backupRawdataStream, err = clt.BackupRawdata(context.Background())
		if err != nil {
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.backupRawdataStream.Recv()
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

func (grpc *backupRawdataStreamSdkImpl) StopBackupRawdataStream() error {
	if grpc.backupRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.backupRawdataStream.CloseSend()
}

func (grpc *backupRawdataStreamSdkImpl) BackupRawdata(req *pb.UpdateRawdataRequest) error {
	if grpc.backupRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputUpdateRawdataReq{
		UpdateRawdataRequest: req,
	}
	err := in.Validate()
	if err != nil {
		return err
	}
	return grpc.backupRawdataStream.Send(req)
}
