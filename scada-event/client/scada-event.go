package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/94peter/microservice/grpc_tool"
	"github.com/muulinCorp/interlib/scada-event/pb"
)

type ScadaBasicEventClient interface {
	CreateEvent(ctx context.Context, service, typ, summary, detail string) error
}

func NewScadaBasicEventClient(address string) ScadaBasicEventClient {
	return &scadaBasicEventClientImpl{
		address: address,
	}
}

type scadaBasicEventClientImpl struct {
	address string
}

func (impl *scadaBasicEventClientImpl) CreateEvent(ctx context.Context, service, typ, summary, detail string) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()

	clt := pb.NewScadaEventServiceClient(grpc)
	_, err = clt.CreateEvent(ctx, &pb.CreateEventReq{
		Service:   service,
		Type:      typ,
		Summarize: summary,
		Detail:    detail,
	})
	return err
}

type ScadaEventWarnClient[T any] interface {
	CreateWarning(ctx context.Context, service string, warns ...*CreateWarn[T]) error
	CloseWarning(ctx context.Context, service string, keys ...string) error
	ExtendWarning(ctx context.Context, service string, keys ...string) error
	ReadWarnings(ctx context.Context, service string, account string, keys ...string) error
	GetRealtime(ctx context.Context, service string) ([]*realtimeResp[T], error)
}

func NewScadaEventClient[T any](address string) ScadaEventWarnClient[T] {
	return &scadaEventClientImpl[T]{
		marshalFunc:   json.Marshal,
		unmarshalFunc: json.Unmarshal,
		address:       address,
	}
}

type scadaEventClientImpl[T any] struct {
	marshalFunc   func(v any) ([]byte, error)
	unmarshalFunc func(data []byte, v any) error
	address       string
}

type FieldInfo struct {
	Name  string
	Field string
}

func NewCreateWarn[T any](key string, name string, field *FieldInfo, data T, detail map[string]any) *CreateWarn[T] {
	return &CreateWarn[T]{
		Key:    key,
		Name:   name,
		Field:  field,
		Data:   data,
		Detail: detail,
	}
}

type CreateWarn[T any] struct {
	Key    string
	Name   string
	Field  *FieldInfo
	Data   T
	Detail map[string]any
}

func (impl *scadaEventClientImpl[T]) CreateWarning(ctx context.Context, service string, warns ...*CreateWarn[T]) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewScadaEventServiceClient(grpc)
	req := &pb.CreateWarningsReq{
		Service: service,
	}
	req.Warnings = make([]*pb.CreateWarningsReq_Warning, len(warns))

	for i, warn := range warns {
		data, err := impl.marshalFunc(warn.Data)
		if err != nil {
			return err
		}
		detail, err := impl.marshalFunc(warn.Detail)
		if err != nil {
			return err
		}
		req.Warnings[i] = &pb.CreateWarningsReq_Warning{
			Key:    warn.Key,
			Name:   warn.Name,
			Field:  &pb.FieldInfo{Name: warn.Field.Name, Field: warn.Field.Field},
			Data:   data,
			Detail: detail,
		}
	}
	_, err = clt.CreateWarning(ctx, req)
	return err
}

func (impl *scadaEventClientImpl[T]) CloseWarning(ctx context.Context, service string, keys ...string) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewScadaEventServiceClient(grpc)
	req := &pb.CloseWarningReq{
		Service: service,
		Key:     keys,
	}
	_, err = clt.CloseWarning(ctx, req)
	return err
}

func (impl *scadaEventClientImpl[T]) ExtendWarning(ctx context.Context, service string, keys ...string) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewScadaEventServiceClient(grpc)
	req := &pb.ExtendWarningReq{
		Service: service,
		Key:     keys,
	}
	_, err = clt.ExtendWarning(ctx, req)
	return err
}

func (impl *scadaEventClientImpl[T]) ReadWarnings(ctx context.Context, service string, account string, keys ...string) error {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return err
	}
	defer grpc.Close()
	clt := pb.NewScadaEventServiceClient(grpc)
	req := &pb.ReadWarningsReq{
		Service: service,
		Key:     keys,
		Account: account,
	}
	_, err = clt.ReadWarnings(ctx, req)
	return err
}

type readerInfo struct {
	Account string
	Time    time.Time
}
type realtimeResp[T any] struct {
	Key        string
	CreateTime time.Time
	Data       T
	Readers    []*readerInfo
}

func (impl *scadaEventClientImpl[T]) GetRealtime(ctx context.Context, service string) ([]*realtimeResp[T], error) {
	grpc, err := grpc_tool.NewConnection(ctx, impl.address)
	if err != nil {
		return nil, err
	}
	defer grpc.Close()
	clt := pb.NewScadaEventServiceClient(grpc)
	resp, err := clt.GetRealtime(ctx, &pb.GetRealtimeReq{
		Service: service,
	})
	if err != nil {
		return nil, err
	}
	realtime := make([]*realtimeResp[T], len(resp.Warnings))
	for i, entry := range resp.Warnings {
		var realtimeObj T
		err = impl.unmarshalFunc(entry.Detail, &realtimeObj)
		if err != nil {
			return nil, err
		}
		readerInfoSlice := make([]*readerInfo, len(entry.Readers))
		for j, reader := range entry.Readers {
			readerInfoSlice[j] = &readerInfo{
				Account: reader.Account,
				Time:    time.Unix(reader.Timestamp, 0),
			}
		}
		realtime[i] = &realtimeResp[T]{
			Key:        entry.Key,
			Data:       realtimeObj,
			CreateTime: time.Unix(entry.CreateTime, 0),
			Readers:    readerInfoSlice,
		}
	}
	return realtime, nil
}
