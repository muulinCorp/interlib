package client

import (
	"errors"

	"github.com/muulinCorp/interlib/communicate/pb"
)

var (
	ErrNotFount    = errors.New("not found")
	ErrConnectFail = errors.New("connect fail")
)

type SensorReader interface {
	GetValue(name string) (float64, error)
}

func NewSensorReader(sensors []*pb.GetSensorsResponse_Sensor) SensorReader {
	return &sensorReaderImpl{
		sensors: sensors,
	}
}

type sensorReaderImpl struct {
	sensors []*pb.GetSensorsResponse_Sensor
}

func (impl *sensorReaderImpl) GetValue(name string) (float64, error) {
	l := len(impl.sensors)
	for i := 0; i < l; i++ {
		if impl.sensors[i].Name == name {
			if impl.sensors[i].GetIsConnErr() {
				return 0, ErrConnectFail
			}
			return impl.sensors[i].GetValue(), nil
		}
	}
	return 0, ErrNotFount
}
