package appDevice

import (
	"context"
	"errors"
	"io"
	"time"

	"bitbucket.org/muulin/interlib/core"
	pb "bitbucket.org/muulin/interlib/device/app/service"
	"github.com/94peter/sterna/auth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppDeviceClient interface {
	core.MyGrpc
	CreateSendTxn(
		host string, devices DeviceAry, desc string,
		recvHandler func(suc bool, device *Device, err string),
		reqUser auth.ReqUser,
	) error
	ModifySendTxn(
		host string, id string, acts []*TxnAct, comment string, reqUser auth.ReqUser,
	) error
	CancelSendTxn(
		host string, id string, comment string, reqUser auth.ReqUser,
	) error
	ListTxn(host string, q *QueryTxnRequest, handle func(rep *pb.QueryTxnResponse)) error
	TxnAddComment(host string, id string, comment string, reqUser auth.ReqUser) error
	MigrationTxn(host string, targetHost string, id string, reqUser auth.ReqUser) error
	RemoveTxn(host string, id string, reqUser auth.ReqUser) error
	ConfirmRecycle(host string, id string, comment string, reqUser auth.ReqUser) error
}

func NewGrpcClient(address string) (AppDeviceClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
}

func (gclt *grpcClt) CreateSendTxn(
	host string, devices DeviceAry, desc string,
	recvHandler func(suc bool, device *Device, err string),
	reqUser auth.ReqUser,
) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	stream, err := clt.CreateSendTxn(ctx, &pb.CreateSendTxnRequest{
		Time:        timestamppb.New(time.Now()),
		Devices:     devices.getDevices(),
		Description: desc,
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
		recvHandler(resp.Success, &Device{
			Mac:       resp.Device.Mac,
			VirtualID: resp.Device.VirtualID,
			Model:     resp.Device.Model,
		}, resp.Error)
	}
	return nil
}

func (gclt *grpcClt) ModifySendTxn(
	host string, id string, acts []*TxnAct, comment string, reqUser auth.ReqUser,
) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	var reqActs []*pb.ActTxn
	for _, act := range acts {
		if act.Edit == TxnEditTypeAdd {
			reqActs = append(reqActs, &pb.ActTxn{
				Act: pb.EditType_Add,
				Device: &pb.Device{
					Mac:       act.Device.Mac,
					VirtualID: act.Device.VirtualID,
					Model:     act.Device.Model,
				},
			})
		}
	}
	resp, err := clt.ModifySendTxn(ctx, &pb.ModifySendTxnRequest{
		TransID: id,
		MacList: reqActs,
		Comment: comment,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) CancelSendTxn(
	host string, id string, comment string, reqUser auth.ReqUser,
) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	resp, err := clt.CancelSendTxn(ctx, &pb.ApplyTxn{
		Id:  id,
		Msg: comment,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) ListTxn(host string, q *QueryTxnRequest, handle func(rep *pb.QueryTxnResponse)) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	var txnType pb.TxnType
	if q.Typ == TxnTypeSend {
		txnType = pb.TxnType_Sending
	} else {
		txnType = pb.TxnType_Recycling
	}
	var txnState pb.TxnState
	switch q.State {
	case TxnStateDone:
		txnState = pb.TxnState_Confirmed
	case TxnStateCancel:
		txnState = pb.TxnState_Canceled
	default:
		txnState = pb.TxnState_New
	}
	if q.State == TxnStateNew {
		txnState = pb.TxnState_New
	}
	stream, err := clt.ListTxn(ctx, &pb.QueryTxnRequest{
		Type:      txnType,
		State:     txnState,
		StartTime: q.StartTime.Format(time.RFC3339),
		EndTime:   q.EndTime.Format(time.RFC3339),
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
		handle(resp)
	}
	return nil
}

func (gclt *grpcClt) TxnAddComment(host string, id string, comment string, reqUser auth.ReqUser) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	resp, err := clt.TxnAddComment(ctx, &pb.TxnAddCommentRequest{
		TxnID: id,
		Msg:   comment,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) MigrationTxn(host string, targetHost string, id string, reqUser auth.ReqUser) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	resp, err := clt.MigrationTxn(ctx, &pb.MigrationTxnRequest{
		Id:     id,
		Target: targetHost,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) RemoveTxn(host string, id string, reqUser auth.ReqUser) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	resp, err := clt.RemoveTxn(ctx, &pb.RemoveTxnRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) ConfirmRecycle(host string, id string, comment string, reqUser auth.ReqUser) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-ReqUser", reqUser.Encode())
	resp, err := clt.ConfirmRecycle(ctx, &pb.ApplyTxn{
		Id:  id,
		Msg: comment,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (gclt *grpcClt) GetDevicesByEquips(
	host string,
	equipIDs []string,
	handle func(string, []*Device),
) error {
	clt := pb.NewAppDeviceServiceClient(gclt)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host)
	stream, err := clt.GetDevicesByEquips(ctx, &pb.GetDevicesByEquipsRequest{
		EquipIDs: equipIDs,
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
		var mydevices []*Device
		for _, d := range resp.Devices {
			mydevices = append(mydevices, &Device{
				Mac:       d.MacAddress,
				VirtualID: d.VirtualID,
				Model:     d.Model,
			})
		}
		handle(resp.EquipID, mydevices)
	}
	return nil
}
