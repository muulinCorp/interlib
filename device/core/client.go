package coreDevice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"bitbucket.org/muulin/interlib/core"
	pb "bitbucket.org/muulin/interlib/device/core/service"
	"github.com/94peter/sterna/log"
)

type CoreDeviceClient interface {
	core.MyGrpc
	GetStateMap(devices []string) (map[string]string, error)
	StartUpdateRawdataStream(recvHandler func(success bool, mac string, err string), log log.Logger) error
	StopUpdateRawdataStream() error
	UpdateRawdata(dataType UpdateRawdataType, mac string, t time.Time, values SensorValuePool) error
}

func NewGrpcClient(address string) (CoreDeviceClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
	updateRawdataStream pb.CoreDeviceService_UpdateRawdataClient
}

func (grpc *grpcClt) GetStateMap(deviceIDs []string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewCoreDeviceServiceClient(grpc)
	resp, err := clt.GetStateMap(ctx, &pb.GetStateMapRequest{
		DeviceID: deviceIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.StateMap, nil
}

func (grpc *grpcClt) StartUpdateRawdataStream(recvHandler func(success bool, mac string, err string), log log.Logger) error {
	clt := pb.NewCoreDeviceServiceClient(grpc)
	stream, err := clt.UpdateRawdata(context.Background())
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
				log.Fatal(fmt.Sprintf("Failed to receive a note : %v", err))
			}
			recvHandler(in.Success, in.Mac, in.Error)
			log.Info(fmt.Sprintf("Get Message: %v, %s, %s", in.Success, in.Mac, in.Error))
		}
	}()
	grpc.updateRawdataStream = stream
	<-waitc
	return nil
}

func (grpc *grpcClt) StopUpdateRawdataStream() error {
	if grpc.updateRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.updateRawdataStream.CloseSend()
}

func (grpc *grpcClt) UpdateRawdata(dataType UpdateRawdataType, mac string, t time.Time, values SensorValuePool) error {
	if grpc.updateRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.updateRawdataStream.Send(
		&pb.UpdateRawdataRequest{
			Type: dataType.getRawdataRequestType(),
			Data: &pb.Rawdata{
				Mac:    mac,
				Time:   t.Format(time.RFC3339),
				Values: values.getSensorValueMap(),
			},
		})
}
