package message

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"bitbucket.org/muulin/interlib/core"
	"bitbucket.org/muulin/interlib/message/plugin"
	pb "bitbucket.org/muulin/interlib/message/service"
	"google.golang.org/grpc/metadata"
)

type MessageClient interface {
	core.MyGrpc
	MqttPublish(host, topic string, msg []byte) error
	Push(host string, msg *plugin.PushMessage) (errorTokens []string, err error)
	Mail(host string, msg *plugin.MailMessage) error
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

func (grpc *grpcClt) MqttPublish(host string, topic string, msg []byte) error {
	ctx, cancel := context.WithTimeout(
		metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host),
		time.Second)
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

func (grpc *grpcClt) Push(host string, msg *plugin.PushMessage) (errorTokens []string, err error) {
	ctx, cancel := context.WithTimeout(
		metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host),
		time.Second)
	defer cancel()
	if len(msg.Receivers) == 0 {
		return nil, nil
	}
	clt := pb.NewMessageServiceClient(grpc)
	var tokens []string
	for _, r := range msg.Receivers {
		tokens = append(tokens, r.Address)
	}
	data, _ := json.Marshal(msg.Content.Data)
	resp, err := clt.Push(ctx, &pb.PushRequest{
		Tokens:    tokens,
		Title:     msg.Title,
		Body:      msg.Content.Body,
		Data:      data,
		Variables: msg.Variables,
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

func (grpc *grpcClt) Mail(host string, msg *plugin.MailMessage) error {
	if len(msg.Receivers) == 0 {
		return errors.New("no receivers")
	}

	ctx, cancel := context.WithTimeout(
		metadata.AppendToOutgoingContext(context.Background(), "X-Channel", host),
		time.Second)
	defer cancel()
	clt := pb.NewMessageServiceClient(grpc)
	var pbReceivers []*pb.MessageReceiver
	for _, r := range msg.Receivers {
		pbReceivers = append(pbReceivers, &pb.MessageReceiver{
			Name:    r.Name,
			Address: r.Address,
		})
	}
	resp, err := clt.Mail(ctx, &pb.MailRequest{
		Receiver:   pbReceivers,
		Subject:    msg.Title,
		PlaintText: msg.Content.Plaint,
		Html:       msg.Content.Html,
		Variables:  msg.Variables,
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.Error)
	}
	return nil
}
