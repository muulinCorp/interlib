package client

import (
	"context"
)

type mockEventClient[T any] struct {
	createWarnFunc  func(ctx context.Context, service string, warns ...*CreateWarn[T]) error
	closeWarnFunc   func(ctx context.Context, service string, keys ...string) error
	extendWarnFunc  func(ctx context.Context, service string, keys ...string) error
	readWarnFunc    func(ctx context.Context, service string, account string, keys ...string) error
	getRealtimeFunc func(ctx context.Context, service string) ([]*RealtimeResp[T], error)
}

func (clt *mockEventClient[T]) CreateWarning(ctx context.Context, service string, warns ...*CreateWarn[T]) error {
	if clt.createWarnFunc != nil {
		return clt.createWarnFunc(ctx, service, warns...)
	}
	return nil
}

func (clt *mockEventClient[T]) CloseWarning(ctx context.Context, service string, keys ...string) error {
	if clt.closeWarnFunc != nil {
		return clt.closeWarnFunc(ctx, service, keys...)
	}
	return nil
}
func (clt *mockEventClient[T]) ExtendWarning(ctx context.Context, service string, keys ...string) error {
	if clt.extendWarnFunc != nil {
		return clt.extendWarnFunc(ctx, service, keys...)
	}
	return nil
}
func (clt *mockEventClient[T]) ReadWarnings(ctx context.Context, service string, account string, keys ...string) error {
	if clt.readWarnFunc != nil {
		return clt.readWarnFunc(ctx, service, account, keys...)
	}
	return nil
}
func (clt *mockEventClient[T]) GetRealtime(ctx context.Context, service string) ([]*RealtimeResp[T], error) {
	if clt.getRealtimeFunc != nil {
		return clt.getRealtimeFunc(ctx, service)
	}
	return nil, nil
}

type mockEventClientOpt[T any] func(*mockEventClient[T])

func WithMockCreateWarnFunc[T any](f func(ctx context.Context, service string, warns ...*CreateWarn[T]) error) mockEventClientOpt[T] {
	return func(clt *mockEventClient[T]) {
		clt.createWarnFunc = f
	}
}

func WithMockCloseWarnFunc[T any](f func(ctx context.Context, service string, keys ...string) error) mockEventClientOpt[T] {
	return func(clt *mockEventClient[T]) {
		clt.closeWarnFunc = f
	}
}

func WithMockExtendWarnFunc[T any](f func(ctx context.Context, service string, keys ...string) error) mockEventClientOpt[T] {
	return func(clt *mockEventClient[T]) {
		clt.extendWarnFunc = f
	}
}

func WithMockReadWarnFunc[T any](f func(ctx context.Context, service string, account string, keys ...string) error) mockEventClientOpt[T] {
	return func(clt *mockEventClient[T]) {
		clt.readWarnFunc = f
	}
}

func WithMockGetRealtimeFunc[T any](f func(ctx context.Context, service string) ([]*RealtimeResp[T], error)) mockEventClientOpt[T] {
	return func(clt *mockEventClient[T]) {
		clt.getRealtimeFunc = f
	}
}

func NewMockEventClient[T any](opts ...mockEventClientOpt[T]) ScadaEventWarnClient[T] {
	client := &mockEventClient[T]{}

	for _, opt := range opts {
		opt(client)
	}
	return client
}
