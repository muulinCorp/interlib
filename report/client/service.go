package client

import (
	"context"
	"time"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/report/pb"
)

type TimeValue struct {
	Time  time.Time
	Value float64
}

type ReportClient interface {
	QueryFieldsValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error)
	QueryFieldsTimeValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]*TimeValue, error)
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
	grpc, err := grpc_tool.NewConnection(impl.address)
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

func (impl *reportClientImpl) QueryFieldsTimeValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]*TimeValue, error) {
	grpc, err := grpc_tool.NewConnection(impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewReportServiceClient(grpc)
	end := time.Now()
	start := end.Add(-interval)
	resp, err := clt.QueryFieldsTimeValue(ctx, &pb.QueryFieldsReq{
		Fields: fields,
		Start:  start.Format(time.RFC3339),
		End:    end.Format(time.RFC3339),
	})
	if err != nil {
		return nil, err
	}

	fieldsInfo := make(map[string][]*TimeValue)
	for _, entry := range resp.FieldsInfo {
		values := make([]*TimeValue, len(entry.Values))
		for i, value := range entry.Values {
			values[i] = &TimeValue{
				Time:  time.Unix(value.Timestamp, 0),
				Value: value.Value,
			}
		}
		fieldsInfo[entry.Key] = values
	}
	return fieldsInfo, nil
}
