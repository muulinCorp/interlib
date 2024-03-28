package client

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/muulinCorp/interlib/channel/pb"
	"google.golang.org/grpc/metadata"

	"github.com/94peter/micro-service/grpc_tool"
	"golang.org/x/net/context"
)

type MigrationEquipLogStreamClient interface {
	grpc_tool.Connection
	StartMigrationStream(channel string, resp chan *pb.MigrationResponse)
	Migration(*pb.MigrationLogRequest) error
	StopMigrationStream() error
}

func NewMigrationEquipLogStreamClient(address string, timeout time.Duration) MigrationEquipLogStreamClient {
	return &migrationEquipLogStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address, timeout),
	}
}

type migrationEquipLogStreamImpl struct {
	*grpc_tool.AutoReConn

	migrationResp   chan *pb.MigrationResponse
	migrationStream pb.ChannelMigrationService_MigrationEquipmentLogClient
}

func (impl *migrationEquipLogStreamImpl) StartMigrationStream(channel string, resp chan *pb.MigrationResponse) {
	var err error
	impl.migrationResp = resp
	p := func(myGrpc grpc_tool.Connection) error {
		md := metadata.New(map[string]string{"X-Channel": channel})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		clt := pb.NewChannelMigrationServiceClient(impl)
		impl.migrationStream, err = clt.MigrationEquipmentLog(ctx)
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
			resp <- &pb.MigrationResponse{
				StatusCode: http.StatusOK,
				Message:    "ready to send data",
			}
		case <-impl.Reconnect:
			if !impl.WaitUntilReady() {
				resp <- &pb.MigrationResponse{
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

func (grpc *migrationEquipLogStreamImpl) StopMigrationStream() error {
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

func (grpc *migrationEquipLogStreamImpl) Migration(req *pb.MigrationLogRequest) error {
	if grpc.migrationStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputMigrationEquipLogReq{
		MigrationLogRequest: req,
	}
	err := in.Validate()
	if err != nil {
		return err
	}
	err = grpc.migrationStream.Send(req)
	if err != nil {
		return err
	}
	sendResp := <-grpc.migrationResp
	if sendResp.StatusCode != http.StatusOK {
		return errors.New(sendResp.Message)
	}
	return nil
}

type inputMigrationEquipLogReq struct {
	*pb.MigrationLogRequest
}

func (req *inputMigrationEquipLogReq) Validate() error {
	if req.Timestamp == 0 {
		return errors.New("missing timestamp")
	}
	return nil
}
