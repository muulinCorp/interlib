package comm

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	pb "bitbucket.org/muulin/interlib/comm/service"
	"bitbucket.org/muulin/interlib/core"
	"bitbucket.org/muulin/interlib/types"
	"github.com/94peter/sterna/api"
	"github.com/94peter/sterna/log"
	"google.golang.org/grpc/metadata"
)

type CommClient interface {
	core.MyGrpc
	StartIot627TimingStream(host string, recvHandler func(statusCode int, mac string, virtualID uint8, msg string)) error
	StopIot627TimingStream() error
	StreamIot627Timing(mac string, virtualID uint8, zone string) error
	Iot627Remote(host, mac string, virtualID uint8, key string, val float64) error
	Iot627GetControlValue(host, mac string, virtualID uint8) error
}

func NewGrpcClient(address string, log log.Logger) (CommClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
	iot627TimingStream pb.CommService_Iot627TimingClient
	log                log.Logger
}

func (gclt *grpcClt) StartIot627TimingStream(host string, recvHandler func(statusCode int, mac string, virtualID uint8, msg string)) error {
	clt := pb.NewCommServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	stream, err := clt.Iot627Timing(ctx)
	waitc := make(chan struct{})
	if err != nil {
		return err
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				gclt.log.Fatal(fmt.Sprintf("Failed to receive a note : %v", err))
			}
			recvHandler(int(in.StatusCode), in.Mac, uint8(in.VirtualID), in.Message)
			gclt.log.Info(fmt.Sprintf("Get Message: %v, %s, %s, %s", in.StatusCode, in.Mac, in.VirtualID, in.Message))
		}
	}()
	gclt.iot627TimingStream = stream
	<-waitc
	return nil
}

func (grpc *grpcClt) StopIot627TimingStream() error {
	if grpc.iot627TimingStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.iot627TimingStream.CloseSend()
}

func (grpc *grpcClt) StreamIot627Timing(mac string, virtualID uint8, zone string) error {
	if grpc.iot627TimingStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.iot627TimingStream.Send(
		&pb.Iot627TimingRequest{
			Mac:       mac,
			VirtualID: uint32(virtualID),
			Zone:      zone,
		})
}

func (grpc *grpcClt) Iot627Remote(host, mac string, virtualID uint8, key string, val float64) error {
	clt := pb.NewCommServiceClient(grpc)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	resp, err := clt.Iot627Remote(ctx, &pb.Iot627RemoteRequest{
		Mac:       mac,
		VirtualID: uint32(virtualID),
		Key:       key,
		Value:     val,
	})
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return types.NewErrorWaper(
		api.NewApiError(int(resp.StatusCode), resp.Message),
		fmt.Sprintf("device [%s]-[%d] remote error: %s", resp.Mac, resp.VirtualID, resp.Message))
}

func (grpc *grpcClt) Iot627GetControlValue(host, mac string, virtualID uint8) error {
	clt := pb.NewCommServiceClient(grpc)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	resp, err := clt.Iot627GetControlValue(ctx, &pb.Iot627GetControlValueRequest{
		Mac:       mac,
		VirtualID: uint32(virtualID),
	})
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return types.NewErrorWaper(
		api.NewApiError(int(resp.StatusCode), resp.Message),
		fmt.Sprintf("device [%s]-[%s] remote error: %s", resp.Mac, resp.VirtualID, resp.Message))
}
