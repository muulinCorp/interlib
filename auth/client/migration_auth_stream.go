package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/muulinCorp/interlib/auth/pb"
	"google.golang.org/grpc/metadata"

	"github.com/94peter/microservice/grpc_tool"
	"golang.org/x/net/context"
)

type MigrationUserStreamClient interface {
	grpc_tool.Connection
	StartMigrationStream(channel string, resp chan *pb.MigrationResponse)
	Migration(*pb.MigrationUserInfoRequest) error
	StopMigrationStream() error
}

func NewMigrationStreamClient(address string) MigrationUserStreamClient {
	return &migrationUserStreamImpl{
		AutoReConn: grpc_tool.NewAutoReconn(address),
	}
}

type migrationUserStreamImpl struct {
	*grpc_tool.AutoReConn

	migrationResp   chan *pb.MigrationResponse
	migrationStream pb.AuthMigrationService_MigrationUserInfoClient
}

func (impl *migrationUserStreamImpl) StartMigrationStream(channel string, resp chan *pb.MigrationResponse) {
	var err error
	impl.migrationResp = resp
	p := func(myGrpc grpc_tool.Connection) error {
		md := metadata.New(map[string]string{"X-Channel": channel})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		clt := pb.NewAuthMigrationServiceClient(impl)
		impl.migrationStream, err = clt.MigrationUserInfo(ctx)
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

func (grpc *migrationUserStreamImpl) StopMigrationStream() error {
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

func (grpc *migrationUserStreamImpl) Migration(req *pb.MigrationUserInfoRequest) error {
	if grpc.migrationStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	in := inputMigrationDeviceReq{
		MigrationUserInfoRequest: req,
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

type inputMigrationDeviceReq struct {
	*pb.MigrationUserInfoRequest
}

func (req *inputMigrationDeviceReq) Validate() error {
	if len(req.Perms) == 0 {
		return errors.New("perms is empty")
	}
	return nil
}
