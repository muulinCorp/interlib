package types

import (
	"fmt"
	"time"
)

func NewClockTime(h, m, s int) *ClockTime {
	c := ClockTime(time.Date(1911, 1, 1, h, m, s, 0, time.Now().Location()))
	return &c
}

type ClockTime time.Time

const clockTimeFormat = "15:04:05Z07:00"

func (ct ClockTime) String() string {
	return time.Time(ct).Format(clockTimeFormat)
}

func (ct ClockTime) Equal(tt ClockTime) bool {
	t1 := time.Time(ct)
	t2 := time.Time(tt)
	return (t1.Hour() == t2.Hour() && t1.Second() == t2.Second())
}

func (ct ClockTime) applyTime(tt time.Time) time.Time {
	t := time.Time(ct)
	return time.Date(tt.Year(), tt.Month(), tt.Day(), t.Hour(), t.Minute(), t.Second(), 0, tt.Location())
}

func (ct ClockTime) Sub(tt time.Time) time.Duration {
	t := ct.applyTime(tt)
	totalMins := (t.Hour() - tt.Hour()) * 60
	totalMins += t.Minute() - tt.Minute()
	return time.Duration(t.Second()-tt.Second()+totalMins*60) * time.Second
}

func (ct ClockTime) MarshalYAML() (interface{}, error) {
	tt := time.Time(ct)
	return tt.Format(clockTimeFormat), nil
}

func (ct *ClockTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err != nil {
		return err
	}
	t, err := time.Parse(clockTimeFormat, str)
	if err != nil {
		return err
	}
	*ct = ClockTime(t)
	return nil
}

func (ct *ClockTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	result := str[1 : len(str)-1]
	if len(result) == 0 {
		return nil
	}
	t, err := time.Parse(clockTimeFormat, result)
	if err != nil {
		return err
	}
	*ct = ClockTime(t)
	return nil
}

func (ct *ClockTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(*ct)
	return []byte(fmt.Sprintf("\"%s\"", tt.Format(clockTimeFormat))), nil
}
