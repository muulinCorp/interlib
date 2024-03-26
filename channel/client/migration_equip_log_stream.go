package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/muulinCorp/interlib/channel/pb"
	"google.golang.org/grpc/metadata"

	"github.com/muulinCorp/interlib/core"
	"golang.org/x/net/context"
)

type MigrationEquipLogStreamClient interface {
	core.MyGrpc
	StartMigrationStream(channel string, resp chan *pb.MigrationResponse)
	Migration(*pb.MigrationLogRequest) error
	StopMigrationStream() error
}

func NewMigrationEquipLogStreamClient(ctx context.Context, address string) MigrationEquipLogStreamClient {
	return &migrationEquipLogStreamImpl{
		AutoReConn: core.NewAutoReconn(ctx, address),
	}
}

type migrationEquipLogStreamImpl struct {
	*core.AutoReConn

	migrationResp   chan *pb.MigrationResponse
	migrationStream pb.ChannelMigrationService_MigrationEquipmentLogClient
}

func (impl *migrationEquipLogStreamImpl) StartMigrationStream(channel string, resp chan *pb.MigrationResponse) {
	var err error
	impl.migrationResp = resp
	p := func(myGrpc core.MyGrpc) error {
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
