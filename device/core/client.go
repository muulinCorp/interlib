package coreDevice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"bitbucket.org/muulin/interlib/core"
	pb "bitbucket.org/muulin/interlib/device/core/service"
	"github.com/94peter/sterna/auth"
	"github.com/94peter/sterna/log"
	"google.golang.org/grpc/metadata"
)

type CoreDeviceClient interface {
	core.MyGrpc
	GetStateMap(devices []string) (map[string]string, error)
	GetInfoMap(devices []string) (map[string]*DeviceInfo, error)
	Remote(deviceID string, device, address uint32, value float64) error
	StartUpdateRawdataStream(recvHandler func(success bool, mac string, virtualID uint8, err string), log log.Logger) error
	StopUpdateRawdataStream() error
	UpdateRawdata(dataType RawdataType, mac string, virtualID uint8, t time.Time, values SensorValuePool) error
	GetValueMap(dataType RawdataType, devices []string, recvHandler func(deviceID string, valuemap map[uint32]float64)) error
	UpdateDeviceState(devics DeviceAry, state DeviceState, comment string, errorHandler func(mac string, virtualID uint8, err string), reqUser auth.ReqUser) error
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

func (grpc *grpcClt) Remote(deviceID string, device, address uint32, value float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewCoreDeviceServiceClient(grpc)
	resp, err := clt.Remote(ctx, &pb.RemoteRequest{
		DeviceID: deviceID,
		Device:   device,
		Address:  address,
		Value:    value,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
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

func (grpc *grpcClt) GetInfoMap(deviceIDs []string) (map[string]*DeviceInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewCoreDeviceServiceClient(grpc)
	resp, err := clt.GetInfoMap(ctx, &pb.GetStateMapRequest{
		DeviceID: deviceIDs,
	})
	if err != nil {
		return nil, err
	}
	result := map[string]*DeviceInfo{}
	for k, v := range resp.InfoMap {
		result[k] = &DeviceInfo{
			Macaddress: v.Mac,
			VirtualID:  uint8(v.VirtualID),
			Model:      v.Model,
			State:      v.State,
		}
	}
	return result, nil
}

func (grpc *grpcClt) GetValueMap(dataType RawdataType, devices []string, recvHandler func(deviceID string, valuemap map[uint32]float64)) error {
	clt := pb.NewCoreDeviceServiceClient(grpc)
	stream, err := clt.GetValueMap(context.Background(), &pb.GetValueMapRequest{
		Type:      dataType.getRawdataRequestType(),
		DeviceIDs: devices,
	})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		recvHandler(resp.DeviceID, resp.ValueMap)
	}
	return nil
}

func (grpc *grpcClt) StartUpdateRawdataStream(recvHandler func(success bool, mac string, virtualID uint8, err string), log log.Logger) error {
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
			recvHandler(in.Success, in.Mac, uint8(in.VirtualID), in.Error)
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

func (grpc *grpcClt) UpdateRawdata(dataType RawdataType, mac string, virtualID uint8, t time.Time, values SensorValuePool) error {
	if grpc.updateRawdataStream == nil {
		return errors.New("StartUpdateRawdataStream first")
	}
	return grpc.updateRawdataStream.Send(
		&pb.UpdateRawdataRequest{
			Type: dataType.getRawdataRequestType(),
			Data: &pb.Rawdata{
				Mac:       mac,
				VirtualID: uint32(virtualID),
				Time:      t.Format(time.RFC3339),
				Values:    values.getSensorValueMap(),
			},
		})
}

func (grpc *grpcClt) UpdateDeviceState(
	devics DeviceAry,
	state DeviceState,
	comment string,
	errorHandler func(mac string, virtualID uint8, err string),
	reqUser auth.ReqUser,
) error {
	if len(devics) == 0 {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	clt := pb.NewCoreDeviceServiceClient(grpc)
	var updateState pb.DeviceState
	switch state {
	case Assigned:
		updateState = pb.DeviceState_Assigned
	case Used:
		updateState = pb.DeviceState_Assigned
	case ToBeRepaired:
		updateState = pb.DeviceState_ToBeRepaired
	case Sending:
		updateState = pb.DeviceState_Reserved
	case Stock:
		updateState = pb.DeviceState_Stock
	default:
		return errors.New("state must be [assigned, used, 2bRepaired]")
	}
	var sendDevices []*pb.CoreDevice
	for _, d := range devics {
		sendDevices = append(sendDevices, &pb.CoreDevice{
			Mac:       d.Macaddress,
			VirtualID: uint32(d.VirtualID),
		})
	}
	stream, err := clt.UpdateDeviceState(ctx, &pb.UpdateDeviceStateRequest{
		State:   updateState,
		Devices: sendDevices,
		Comment: comment,
	})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if !resp.Success {
			errorHandler(resp.Mac, uint8(resp.VirtualID), resp.Error)
		}
	}
	return nil
}
