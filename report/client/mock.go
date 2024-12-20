package client

import (
	"context"
	"time"
)

type reportClient struct {
	queryFieldsValueByLastInterval     func(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error)
	queryFieldsTimeValueByLastInterval func(ctx context.Context, interval time.Duration, fields []string) (map[string][]*TimeValue, error)
}

type mockReportClientOpt func(*reportClient)

func MockReportClientQueryFieldsValueByLastInterval(f func(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error)) mockReportClientOpt {
	return func(c *reportClient) {
		c.queryFieldsValueByLastInterval = f
	}
}

func MockReportClientQueryFieldsTimeValueByLastInterval(f func(ctx context.Context, interval time.Duration, fields []string) (map[string][]*TimeValue, error)) mockReportClientOpt {
	return func(c *reportClient) {
		c.queryFieldsTimeValueByLastInterval = f
	}
}

func NewMockReportClient(opts ...mockReportClientOpt) ReportClient {
	client := &reportClient{}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *reportClient) QueryFieldsValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]float64, error) {
	if c.queryFieldsValueByLastInterval != nil {
		return c.queryFieldsValueByLastInterval(ctx, interval, fields)
	}
	return nil, nil
}

func (c *reportClient) QueryFieldsTimeValueByLastInterval(ctx context.Context, interval time.Duration, fields []string) (map[string][]*TimeValue, error) {
	if c.queryFieldsTimeValueByLastInterval != nil {
		return c.queryFieldsTimeValueByLastInterval(ctx, interval, fields)
	}
	return nil, nil
}