package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/muulinCorp/interlib/device/pb"

	"context"

	apiErr "github.com/94peter/api-toolkit/errors"
	"github.com/94peter/micro-service/grpc_tool"
)

type DeviceV1Client interface {
	grpc_tool.Connection
	StartCreateV1Stream(context.Context) error
	CreateV1(*pb.CreateDeviceV1Request) (*pb.CreateDeviceV1Response, error)
	StopCreateV1Stream() error
	CheckExist(context.Context, []*pb.DeviceV1) (map[string]bool, error)
	CheckState(context.Context, []*pb.Device) (map[string]string, error)
	GetDeviceInfo(ctx context.Context, mac, gwid string) (*pb.DeviceInfoResponse, error)

	StartRemoveStream(context.Context) error
	Remove(*pb.RemoveDeviceV1Request) error
	StopRemoveStream() error
}

func NewDeviceClientV1(address string) DeviceV1Client {
	return &sdkV1Impl{
		address: address,
		// AutoReConn: grpc_tool.NewAutoReconn(address),
	}
}

type sdkV1Impl struct {
	address string
	grpc_tool.Connection

	createV1Stream pb.DeviceV1Service_CreateV1Client
	removeV1Stream pb.DeviceV1Service_DeleteClient
}

func (impl *sdkV1Impl) StartCreateV1Stream(ctx context.Context) error {
	var err error
	impl.Connection, err = grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	clt := pb.NewDeviceV1ServiceClient(impl)
	impl.createV1Stream, err = clt.CreateV1(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (grpc *sdkV1Impl) StopCreateV1Stream() error {
	if grpc.createV1Stream == nil {
		return errors.New("StartCreateV1Stream first")
	}
	return grpc.createV1Stream.CloseSend()
}

func (grpc *sdkV1Impl) CreateV1(req *pb.CreateDeviceV1Request) (*pb.CreateDeviceV1Response, error) {
	if grpc.createV1Stream == nil {
		return nil, errors.New("StartCreateV1Stream first")
	}
	if req.Channel == "" {
		return &pb.CreateDeviceV1Response{
			StatusCode: http.StatusBadRequest,
			Message:    "missing channel",
		}, nil
	}

	resp := make(chan *pb.CreateDeviceV1Response)
	chanErr := make(chan error)

	go func(
		stream pb.DeviceV1Service_CreateV1Client,
		chanResp chan *pb.CreateDeviceV1Response,
		chanErr chan error) {
		in, err := stream.Recv()
		if err == io.EOF {
			chanResp <- nil
			chanErr <- nil
			return
		}
		if err != nil {
			chanResp <- nil
			chanErr <- err
			return
		}
		chanResp <- in
		chanErr <- nil
	}(grpc.createV1Stream, resp, chanErr)

	err := grpc.createV1Stream.Send(req)
	if err != nil {
		return nil, err
	}
	return <-resp, <-chanErr
}

func (grpc *sdkV1Impl) CheckExist(ctx context.Context, dvices []*pb.DeviceV1) (map[string]bool, error) {
	var err error
	grpc.Connection, err = grpc_tool.NewConnection(ctx, grpc.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewDeviceV1ServiceClient(grpc)
	resp, err := clt.CheckExist(context.Background(), &pb.CheckExistRequest{
		Devices: dvices,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExistMap, nil
}

func (grpc *sdkV1Impl) CheckState(ctx context.Context, devices []*pb.Device) (map[string]string, error) {
	var err error
	grpc.Connection, err = grpc_tool.NewConnection(ctx, grpc.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewDeviceV1ServiceClient(grpc)
	resp, err := clt.CheckState(context.Background(), &pb.GetStateRequest{Devices: devices})
	if err != nil {
		return nil, err
	}
	return resp.StateMap, nil
}

func (grpc *sdkV1Impl) GetDeviceInfo(ctx context.Context, mac, gwid string) (*pb.DeviceInfoResponse, error) {
	var err error
	grpc.Connection, err = grpc_tool.NewConnection(ctx, grpc.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewDeviceV1ServiceClient(grpc)
	resp, err := clt.GetDeviceInfo(context.Background(), &pb.DeviceV1{Mac: mac, Gwid: gwid})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (impl *sdkV1Impl) StartRemoveStream(ctx context.Context) error {
	var err error
	impl.Connection, err = grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	clt := pb.NewDeviceV1ServiceClient(impl)
	impl.removeV1Stream, err = clt.Delete(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (grpc *sdkV1Impl) StopRemoveStream() error {
	if grpc.removeV1Stream == nil {
		return errors.New("StartRemoveStream first")
	}
	return grpc.removeV1Stream.CloseSend()
}

func (grpc *sdkV1Impl) Remove(req *pb.RemoveDeviceV1Request) error {
	if grpc.removeV1Stream == nil {
		return errors.New("StartRemoveStream first")
	}
	if req.Channel == "" {
		return errors.New("missing channel")
	}

	resp := make(chan *pb.Response)
	chanErr := make(chan error)

	go func(
		stream pb.DeviceV1Service_DeleteClient,
		chanResp chan *pb.Response,
		chanErr chan error) {
		in, err := stream.Recv()
		if err == io.EOF {
			chanResp <- nil
			chanErr <- nil
			return
		}
		if err != nil {
			chanResp <- nil
			chanErr <- err
			return
		}
		chanResp <- in
		chanErr <- nil
	}(grpc.removeV1Stream, resp, chanErr)

	err := grpc.removeV1Stream.Send(req)
	if err != nil {
		return err
	}
	if r := <-resp; r != nil && r.StatusCode != http.StatusOK {
		return apiErr.New(int(r.StatusCode), r.Message)
	}
	return <-chanErr
}
