package client

import (
	"context"
	"time"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/report/pb"
)

type ReportClient interface {
	QueryFieldsValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error)
}

func NewReportClient(address string) ReportClient {
	return &reportClientImpl{
		address: address,
	}
}

type reportClientImpl struct {
	address string
}

func (impl *reportClientImpl) QueryFieldsValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewReportServiceClient(grpc)
	end := time.Now()
	start := end.Add(-interval)
	resp, err := clt.QueryFieldsValue(ctx, &pb.QueryFieldsReq{
		Fields: fields,
		Start:  start.Format(time.RFC3339),
		End:    end.Format(time.RFC3339),
	})
	if err != nil {
		return nil, err
	}

	fieldsInfo := make(map[string][]float64)
	for _, entry := range resp.FieldsInfo {
		fieldsInfo[entry.Key] = entry.Value
	}
	return fieldsInfo, nil
}
