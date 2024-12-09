package client

import "context"

type mockCom2ScadaClientOpt func(*mockCom2ScadaClient)

func MockCom2ScadaWithStartRoutRoutineDataStream(f func(ctx context.Context, fields []string, resp chan map[string]float64) error) mockCom2ScadaClientOpt {
	return func(c *mockCom2ScadaClient) {
		c.startRoutineDataStream = f
	}
}

func NewMockCom2ScadaClient(opts ...mockCom2ScadaClientOpt) RoutineDataStreamClient {
	client := &mockCom2ScadaClient{}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

type mockCom2ScadaClient struct {
	startRoutineDataStream func(ctx context.Context, fields []string, resp chan map[string]float64) error
	stopRoutinDataStream   func() error
}

func (client *mockCom2ScadaClient) StartRoutineDataStream(fields []string, resp chan map[string]float64) error {
	if client.startRoutineDataStream != nil {
		return client.startRoutineDataStream(context.Background(), fields, resp)
	}
	return nil
}

func (client *mockCom2ScadaClient) StopStream() error {
	if client.stopRoutinDataStream != nil {
		return client.stopRoutinDataStream()
	}
	return nil
}
