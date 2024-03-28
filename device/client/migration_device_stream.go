package client

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/94peter/micro-service/grpc_tool"
	"github.com/muulinCorp/interlib/device/pb"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
)

type MigrationDeviceStreamClient interface {
	grpc_tool.Connection
	StartMigrationStream(channel string, resp chan *pb.Response)
	Migration(*pb.MigrationDeviceRequest) error
	StopMigrationStream() error
}

func NewMigrationStreamClient(address string, timeout time.Duration) MigrationDeviceStreamClient {
	return &migrationDeviceStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address, timeout),
	}
}

type migrationDeviceStreamImpl struct {
	*grpc_tool.AutoReConn

	resp            chan *pb.Response
	migrationStream pb.DeviceMigrationService_MigrationDeviceClient
}

func (impl *migrationDeviceStreamImpl) StartMigrationStream(channel string, resp chan *pb.Response) {
	var err error
	impl.resp = resp
	p := func(myGrpc grpc_tool.Connection) error {
		md := metadata.New(map[string]string{"X-Channel": channel})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		clt := pb.NewDeviceMigrationServiceClient(impl)
		impl.migrationStream, err = clt.MigrationDevice(ctx)
		if err != nil {
			return err
		}
		impl.Ready <- true
		for {
			in, err := impl.migrationStream.Recv()
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

func (grpc *migrationDeviceStreamImpl) StopMigrationStream() error {
	if grpc.migrationStream == nil {
		return errors.New("StartMigrationStream first")
	}
	err := grpc.migrationStream.CloseSend()
	if err != nil {
		return err
	}
	grpc.migrationStream = nil
	return nil
}

func (grpc *migrationDeviceStreamImpl) Migration(req *pb.MigrationDeviceRequest) error {
	if grpc.migrationStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputMigrationDeviceReq{
		MigrationDeviceRequest: req,
	}
	err := in.Validate()
	if err != nil {
		return err
	}
	err = grpc.migrationStream.Send(req)
	if err != nil {
		return err
	}
	sendResp := <-grpc.resp
	if sendResp.StatusCode != http.StatusOK {
		return errors.New(sendResp.Message)
	}
	return nil
}

type inputMigrationDeviceReq struct {
	*pb.MigrationDeviceRequest
}

const maxTimezone = 14 * 3600

func (req *inputMigrationDeviceReq) Validate() error {
	if req.Mac == "" {
		return errors.New("invalid mac")
	}
	if req.Gwid == "" {
		return errors.New("invalid gwid")
	}
	if req.Timezone > maxTimezone || req.Timezone < -maxTimezone {
		return errors.New("invalid timezone")
	}
	return nil
}
