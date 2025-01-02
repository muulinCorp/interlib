package pubsub

import (
	"context"

	"github.com/94peter/log"
)

type Source interface {
	SetLog(log.Logger)
	Run(ctx context.Context)
	Close()
	AddRoutineDataSubscriber(SubscriberRoutineData)
	RemoveRoutineDataSubscriber(s SubscriberRoutineData)
	AddDataChangeSubscriber(SubscriberDataChange)
	RemoveDataChangeSubscriber(SubscriberDataChange)
	AddConnectErrSubscriber(SubscriberConnectErr)
	AddDataWriteSubscriber(s SubscriberDataWrite)
	RemoveDataWriteSubscriber(s SubscriberDataWrite)
	Statue() error
}

type SubscriberRoutineData interface {
	RecieveRoutineData(data *SubscribeData)
}

type Sensor struct {
	Name      string
	Value     any
	IsConnErr bool `json:"isConnErr"`
	Err       string
	ErrDetail string `json:"errDetail"`
}

type SubscribeData struct {
	Timestamp int64
	Sensors   []*Sensor
}

func (data *SubscribeData) ToSensorMap() map[string]*Sensor {
	m := make(map[string]*Sensor)
	for _, s := range data.Sensors {
		m[s.Name] = s
	}
	return m
}

type SubscriberDataChange interface {
	RecieveChangeData(data map[string]any)
	Close()
}

type ConnectErrMsg struct {
	ConnInfo string   `json:"conn_info"`
	Address  uint16   `json:"address"`
	Length   int      `json:"length"`
	Message  string   `json:"err_msg"`
	Fields   []string `json:"fields"`
}

type SubscriberConnectErr interface {
	RecieveConnectErr(*ConnectErrMsg)
}

type SubscriberDataWrite interface {
	RecieveDataWrite(data map[string]any)
	Close()
}
