package client

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/maintenance/pb"
	"google.golang.org/grpc/metadata"

	"context"
)

type MigrationStreamClient interface {
	grpc_tool.Connection
	StartMigrationStream(channel string, resp chan *pb.MigrationResponse)
	Migration(*pb.MigrationMaintenanceRequest) error
	StopMigrationStream() error
}

func NewMigrationStreamClient(address string, timeout time.Duration) MigrationStreamClient {
	return &migrationStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address, timeout),
	}
}

type migrationStreamImpl struct {
	*grpc_tool.AutoReConn

	migrationResp   chan *pb.MigrationResponse
	migrationStream pb.MaintenanceMigrationService_MigrationMaintenanceClient
}

func (impl *migrationStreamImpl) StartMigrationStream(channel string, resp chan *pb.MigrationResponse) {
	var err error
	impl.migrationResp = resp
	p := func(myGrpc grpc_tool.Connection) error {
		md := metadata.New(map[string]string{"X-Channel": channel})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		clt := pb.NewMaintenanceMigrationServiceClient(impl)
		impl.migrationStream, err = clt.MigrationMaintenance(ctx)
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

func (grpc *migrationStreamImpl) StopMigrationStream() error {
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

func (grpc *migrationStreamImpl) Migration(req *pb.MigrationMaintenanceRequest) error {
	if grpc.migrationStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputMigrationMaintenanceReq{
		MigrationMaintenanceRequest: req,
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

type inputMigrationMaintenanceReq struct {
	*pb.MigrationMaintenanceRequest
}

func (req *inputMigrationMaintenanceReq) Validate() error {
	return nil
}
