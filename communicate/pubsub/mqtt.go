package pubsub

import (
	"context"
	"encoding/json"

	"github.com/94peter/log"
	"github.com/94peter/mqtt"
	"github.com/94peter/mqtt/config"
	"github.com/94peter/mqtt/trans"
	"github.com/pkg/errors"
)

func NewMqttSource(routineTopic, changeDataTopic, connErrTopic, writeDataTopic string) (Source, error) {
	mqttConf, err := config.GetConfigFromEnvWithoutTopic()
	if err != nil {
		return nil, err
	}
	mqttConf.AddTopics(routineTopic, changeDataTopic, connErrTopic, writeDataTopic)
	return &mqttSource{
		routineTopic:    routineTopic,
		changeDataTopic: changeDataTopic,
		connErrTopic:    connErrTopic,
		writeDataTopic:  writeDataTopic,
		conf:            mqttConf,
		routineTrans:    &routineTrans{},
		dataChangeTrans: &dataChangeTrans{},
		connErrTrans:    &connectErrTrans{},
		dataWriteTrans:  &dataWriteTrans{},
	}, nil
}

type mqttSource struct {
	routineTopic    string
	changeDataTopic string
	connErrTopic    string
	writeDataTopic  string

	conf            *config.Config
	mqttServ        mqtt.MqttSubOnlyServer
	routineTrans    *routineTrans
	dataChangeTrans *dataChangeTrans
	connErrTrans    *connectErrTrans
	dataWriteTrans  *dataWriteTrans

	log log.Logger
}

func (m *mqttSource) SetLog(l log.Logger) {
	m.log = l
}

func (m *mqttSource) AddRoutineDataSubscriber(s SubscriberRoutineData) {
	m.routineTrans.addSubscriber(s)
}

func (m *mqttSource) RemoveRoutineDataSubscriber(s SubscriberRoutineData) {
	m.routineTrans.removeSubscriber(s)
}

func (m *mqttSource) AddConnectErrSubscriber(s SubscriberConnectErr) {
	m.connErrTrans.addSubscriber(s)
}

func (m *mqttSource) AddDataChangeSubscriber(s SubscriberDataChange) {
	m.dataChangeTrans.addSubscriber(s)
}

func (m *mqttSource) RemoveDataChangeSubscriber(s SubscriberDataChange) {
	m.dataChangeTrans.removeSubscriber(s)
	s.Close()
}

func (m *mqttSource) AddDataWriteSubscriber(s SubscriberDataWrite) {
	m.dataWriteTrans.addSubscriber(s)
}

func (m *mqttSource) RemoveDataWriteSubscriber(s SubscriberDataWrite) {
	m.dataWriteTrans.removeSubscriber(s)
	s.Close()
}

func (m *mqttSource) Run(ctx context.Context) {
	var err error
	m.mqttServ, err = mqtt.NewMqttSubOnlyServ(m.conf, map[string]trans.Trans{
		m.routineTopic:    m.routineTrans,
		m.changeDataTopic: m.dataChangeTrans,
		m.connErrTopic:    m.connErrTrans,
		m.writeDataTopic:  m.dataWriteTrans,
	})
	if err != nil {
		panic(err)
	}
	if m.conf.Debug {
		m.mqttServ.SetLog(m.log.GetLogging())
	}
	m.mqttServ.Run(ctx)
}

func (m *mqttSource) Close() {
	if m.mqttServ == nil {
		return
	}
	m.mqttServ.Close()
}

func (m *mqttSource) Statue() error {
	return m.mqttServ.Statue()
}

type routineTrans struct {
	subscribers []SubscriberRoutineData
}

func (s *routineTrans) addSubscriber(sub SubscriberRoutineData) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *routineTrans) removeSubscriber(sub SubscriberRoutineData) {
	for i, v := range s.subscribers {
		if v == sub {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			return
		}
	}
}

func (s *routineTrans) Close() {}

func (s *routineTrans) Send(topic string, b []byte) error {
	var data SubscribeData
	err := json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "json unmarshl fail")
	}

	for _, sub := range s.subscribers {
		sub.RecieveRoutineData(&data)
	}
	return nil
}

type dataChangeTrans struct {
	subscribers []SubscriberDataChange
}

func (s *dataChangeTrans) addSubscriber(sub SubscriberDataChange) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *dataChangeTrans) removeSubscriber(sub SubscriberDataChange) {
	for i, v := range s.subscribers {
		if v == sub {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			return
		}
	}
}

func (s *dataChangeTrans) Close() {
	for _, sub := range s.subscribers {
		sub.Close()
	}
}

func (s *dataChangeTrans) Send(topic string, b []byte) error {
	var data map[string]any
	err := json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "json unmarshl fail")
	}
	for _, sub := range s.subscribers {
		sub.RecieveChangeData(data)
	}
	return nil
}

type connectErrTrans struct {
	subscribers []SubscriberConnectErr
}

func (s *connectErrTrans) addSubscriber(sub SubscriberConnectErr) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *connectErrTrans) Close() {}

func (s *connectErrTrans) Send(topic string, b []byte) error {
	var data ConnectErrMsg
	err := json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "json unmarshl fail")
	}
	for _, sub := range s.subscribers {
		sub.RecieveConnectErr(&data)
	}
	return nil
}

type dataWriteTrans struct {
	subscribers []SubscriberDataWrite
}

func (s *dataWriteTrans) addSubscriber(sub SubscriberDataWrite) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *dataWriteTrans) removeSubscriber(sub SubscriberDataWrite) {
	for i, v := range s.subscribers {
		if v == sub {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			return
		}
	}
}

func (s *dataWriteTrans) Close() {
	for _, sub := range s.subscribers {
		sub.Close()
	}
}

func (s *dataWriteTrans) Send(topic string, b []byte) error {
	var data map[string]any
	err := json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "json unmarshl fail")
	}
	for _, sub := range s.subscribers {
		sub.RecieveDataWrite(data)
	}
	return nil
}
