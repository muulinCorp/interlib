package client

import (
	"context"

	"bitbucket.org/muulin/interlib/channel/pb"

	"bitbucket.org/muulin/interlib/core"
	"google.golang.org/grpc/metadata"
)

type ReportGrpcClient interface {
	CountSensorWarning(sensorIds []string) (map[string]int64, error)
	GetSensorReportInfo(sensorIds []string) (*pb.GetSensorsReportInfoResponse, error)
}

func NewReportGrpcClient(address string, channel string) ReportGrpcClient {
	return &reportGrpcClientImpl{
		channel: channel,
		address: address,
	}
}

type reportGrpcClientImpl struct {
	channel string
	address string
}

func (c *reportGrpcClientImpl) CountSensorWarning(sensorIds []string) (map[string]int64, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcClt, err := core.NewMyGrpc(c.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()
	clt := pb.NewReportServiceClient(grpcClt)

	resp, err := clt.CountSensorsWarning(ctx, &pb.SensorIdsRequest{
		SensorIds: sensorIds,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (c *reportGrpcClientImpl) GetSensorReportInfo(sensorIds []string) (*pb.GetSensorsReportInfoResponse, error) {
	var err error
	md := metadata.New(map[string]string{"X-Channel": c.channel})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	grpcClt, err := core.NewMyGrpc(c.address)
	if err != nil {
		return nil, err
	}
	defer grpcClt.Close()

	clt := pb.NewReportServiceClient(grpcClt)
	resp, err := clt.GetSensorReportInfo(ctx, &pb.SensorIdsRequest{
		SensorIds: sensorIds,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}