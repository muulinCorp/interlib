package client

import (
	"context"

	"github.com/muulinCorp/interlib/scada-layout/pb"
)

type mockScadaLayoutClientOpt func(*mockScadaLayoutClient)

func WithMockGetFieldsTagsFunc(f func(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getFieldsTagsFunc = f
	}
}

func WithMockGetFieldsTagsReaderFunc(f func(ctx context.Context, fields []string) (FieldTagsReader, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getFieldsTagsReaderFunc = f
	}
}

func WithMockGetReportFieldReaderFunc(f func(ctx context.Context, fields []string) (ReportFieldReader, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getReportFieldReaderFunc = f
	}
}

func WithMockGetFieldListWithIdFunc(f func(ctx context.Context, input *pb.GetFieldListRequest) ([]string, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getFieldListWithIdFunc = f
	}
}

func WithMockGetReportSettingFunc(f func(ctx context.Context) (*pb.GetReportResponse, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getReportSettingFunc = f
	}
}

func WithMockGetScenarioReaderFunc(f func(ctx context.Context) (ScenarioReader, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getScenarioReaderFunc = f
	}
}

func WithMockGetAlarmFieldsFunc(f func(ctx context.Context) ([]*pb.GetAlarmFieldsResponse_FieldDetail, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getAlarmFieldsFunc = f
	}
}

func WithMockGetSmartDefrostFunc(f func(ctx context.Context) (SmartDefrostLayout, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getSmartDefrostFunc = f
	}
}

func WithMockGetElectricityDemandFunc(f func(ctx context.Context) (*pb.GetElectricityDemandResponse, error)) mockScadaLayoutClientOpt {
	return func(client *mockScadaLayoutClient) {
		client.getElectricityDemandFunc = f
	}
}

func NewMockScadaLayoutClient(opts ...mockScadaLayoutClientOpt) ScadaLayoutClient {
	client := &mockScadaLayoutClient{}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

type mockScadaLayoutClient struct {
	getFieldsTagsFunc        func(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error)
	getFieldsTagsReaderFunc  func(ctx context.Context, fields []string) (FieldTagsReader, error)
	getReportFieldReaderFunc func(ctx context.Context, fields []string) (ReportFieldReader, error)
	getFieldListWithIdFunc   func(ctx context.Context, input *pb.GetFieldListRequest) ([]string, error)
	getReportSettingFunc     func(ctx context.Context) (*pb.GetReportResponse, error)
	getScenarioReaderFunc    func(ctx context.Context) (ScenarioReader, error)
	getAlarmFieldsFunc       func(ctx context.Context) ([]*pb.GetAlarmFieldsResponse_FieldDetail, error)
	getSmartDefrostFunc      func(ctx context.Context) (SmartDefrostLayout, error)
	getElectricityDemandFunc func(ctx context.Context) (*pb.GetElectricityDemandResponse, error)
}

func (mock *mockScadaLayoutClient) GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error) {
	if mock.getFieldsTagsFunc != nil {
		return mock.getFieldsTagsFunc(ctx, fields)
	}
	return nil, nil
}

func (mock *mockScadaLayoutClient) GetFieldsTagsReader(ctx context.Context, fields []string) (FieldTagsReader, error) {
	if mock.getFieldsTagsReaderFunc != nil {
		return mock.getFieldsTagsReaderFunc(ctx, fields)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetReportFieldReader(ctx context.Context, fields []string) (ReportFieldReader, error) {
	if mock.getReportFieldReaderFunc != nil {
		return mock.getReportFieldReaderFunc(ctx, fields)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetFieldListWithId(ctx context.Context, input *pb.GetFieldListRequest) ([]string, error) {
	if mock.getFieldListWithIdFunc != nil {
		return mock.getFieldListWithIdFunc(ctx, input)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetReportSetting(ctx context.Context) (*pb.GetReportResponse, error) {
	if mock.getReportSettingFunc != nil {
		return mock.getReportSettingFunc(ctx)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetScenarioReader(ctx context.Context) (ScenarioReader, error) {
	if mock.getScenarioReaderFunc != nil {
		return mock.getScenarioReaderFunc(ctx)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetAlarmFields(ctx context.Context) ([]*pb.GetAlarmFieldsResponse_FieldDetail, error) {
	if mock.getAlarmFieldsFunc != nil {
		return mock.getAlarmFieldsFunc(ctx)
	}
	return nil, nil
}
func (mock *mockScadaLayoutClient) GetSmartDefrost(ctx context.Context) (SmartDefrostLayout, error) {
	if mock.getSmartDefrostFunc != nil {
		return mock.getSmartDefrostFunc(ctx)
	}
	return nil, nil
}

func (mock *mockScadaLayoutClient) GetElectricityDemand(ctx context.Context) (*pb.GetElectricityDemandResponse, error) {
	if mock.getElectricityDemandFunc != nil {
		return mock.getElectricityDemandFunc(ctx)
	}
	return nil, nil
}