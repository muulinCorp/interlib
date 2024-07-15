package client

import (
	"context"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/scada-layout/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ScadaLayoutClient interface {
	GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error)
	GetFieldsTagsReader(ctx context.Context, fields []string) (FieldTagsReader, error)
	GetReportFieldReader(ctx context.Context, fields []string) (ReportFieldReader, error)
	GetFieldListWithId(ctx context.Context, input *pb.GetFieldListRequest) ([]string, error)
	GetReportSetting(ctx context.Context) (*pb.GetReportResponse, error)
	GetScenarioReader(ctx context.Context) (ScenarioReader, error)
	GetAlarmFields(ctx context.Context) ([]*pb.GetAlarmFieldsResponse_FieldDetail, error)
}

type clientImpl struct {
	address string
}

func New(address string) ScadaLayoutClient {
	return &clientImpl{
		address: address,
	}
}

func (impl *clientImpl) GetFieldsTags(ctx context.Context, fields []string) (*pb.GetFieldsTagsResponse, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	return clt.GetFieldsTags(ctx, &pb.FieldList{
		Fields: fields,
	})
}

type FieldTagsReader interface {
	GetFieldTag(f string) *pb.GetFieldsTagsResponse_FieldTag
}

func NewFieldTagsReader(resp *pb.GetFieldsTagsResponse) FieldTagsReader {
	data := make(map[string]*pb.GetFieldsTagsResponse_FieldTag)
	for _, ft := range resp.FieldTags {
		data[ft.Field] = ft
	}
	return &fieldTagsReaderImpl{
		data: data,
	}
}

type fieldTagsReaderImpl struct {
	data map[string]*pb.GetFieldsTagsResponse_FieldTag
}

func (impl *fieldTagsReaderImpl) GetFieldTag(f string) *pb.GetFieldsTagsResponse_FieldTag {
	return impl.data[f]
}

func (impl *clientImpl) GetFieldsTagsReader(ctx context.Context, fields []string) (FieldTagsReader, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetFieldsTags(ctx, &pb.FieldList{
		Fields: fields,
	})
	if err != nil {
		return nil, err
	}
	return NewFieldTagsReader(resp), nil
}

type ReportFieldReader interface {
	GetProjectName() string
	GetField(f string) *pb.GetReportFieldsResponse_Field
}

type reportFieldsReaderImpl struct {
	projectName string
	data        map[string]*pb.GetReportFieldsResponse_Field
}

func (impl *reportFieldsReaderImpl) GetProjectName() string {
	return impl.projectName
}

func (impl *reportFieldsReaderImpl) GetField(f string) *pb.GetReportFieldsResponse_Field {
	return impl.data[f]
}

func NewReportFieldsReader(resp *pb.GetReportFieldsResponse) ReportFieldReader {
	data := make(map[string]*pb.GetReportFieldsResponse_Field)
	for _, ft := range resp.Fields {
		data[ft.Field] = ft
	}
	return &reportFieldsReaderImpl{
		projectName: resp.ProjectName,
		data:        data,
	}
}

func (impl *clientImpl) GetReportFieldReader(ctx context.Context, fields []string) (ReportFieldReader, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetReportFields(ctx, &pb.FieldList{
		Fields: fields,
	})
	if err != nil {
		return nil, err
	}
	return NewReportFieldsReader(resp), nil
}

func (impl *clientImpl) GetFieldListWithId(ctx context.Context, input *pb.GetFieldListRequest) ([]string, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetFieldsWithId(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp.Fields, nil
}

func (impl *clientImpl) GetReportSetting(ctx context.Context) (*pb.GetReportResponse, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetReport(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (impl *clientImpl) GetScenarioReader(ctx context.Context) (ScenarioReader, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetScenario(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return newScenarioReader(resp), nil
}

func (impl *clientImpl) GetAlarmFields(ctx context.Context) ([]*pb.GetAlarmFieldsResponse_FieldDetail, error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()

	clt := pb.NewScadaLayoutServiceClient(grpc)

	resp, err := clt.GetAlarmFields(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Fields, nil
}
