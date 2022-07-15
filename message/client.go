package message

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"bitbucket.org/muulin/interlib/core"
	pb "bitbucket.org/muulin/interlib/message/service"
)

type MessageClient interface {
	core.MyGrpc
	MqttPublish(topic string, msg []byte) error
	Push(tokens []string, title, body string, data map[string]interface{}) (errorTokens []string, err error)
	Mail(receivers []*MailReceiver, subject, plaintText, html string, recvHandler func(success bool, name, email, err string)) error
}

func NewGrpcClient(address string) (MessageClient, error) {
	mygrpc, err := core.NewMyGrpc(address)
	if err != nil {
		return nil, err
	}
	return &grpcClt{MyGrpc: mygrpc}, nil
}

type grpcClt struct {
	core.MyGrpc
}

func (grpc *grpcClt) MqttPublish(topic string, msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewMessageServiceClient(grpc)

	resp, err := clt.MqttPublish(ctx, &pb.MqttPublishRequest{
		Topics: []string{topic},
		Msg:    msg,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}

func (grpc *grpcClt) Push(tokens []string, title, body string, data map[string]interface{}) (errorTokens []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewMessageServiceClient(grpc)
	d, _ := json.Marshal(data)
	resp, err := clt.Push(ctx, &pb.PushRequest{
		Tokens: tokens,
		Title:  title,
		Body:   body,
		Data:   d,
	})
	if err != nil {
		return nil, err
	}
	errorTokens = resp.ErrorTokens
	if resp.Error != "" {
		err = errors.New(resp.Error)
	}
	return
}

func (grpc *grpcClt) Mail(receivers []*MailReceiver, subject, plaintText, html string, recvHandler func(success bool, name, email, err string)) error {
	if len(receivers) == 0 {
		return errors.New("no receivers")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clt := pb.NewMessageServiceClient(grpc)
	var pbReceivers []*pb.MailReceiver
	for _, r := range receivers {
		pbReceivers = append(pbReceivers, &pb.MailReceiver{
			Name:  r.Name,
			Email: r.Email,
		})
	}
	stream, err := clt.Mail(ctx, &pb.MailRequest{
		Receiver: pbReceivers,
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
		recvHandler(resp.Success, resp.Name, resp.Mail, resp.Error)
	}
	return nil
}
