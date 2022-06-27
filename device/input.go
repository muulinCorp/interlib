package device

import (
	"errors"
	"time"

	"github.com/94peter/sterna/util"
)

type NewDevice struct {
	Mac         string
	Model       string
	Description string
	NewID       string
}

func (d *NewDevice) Valid() error {
	if !util.IsMAC(d.Mac) {
		return errors.New("invalid mac (400)")
	}

	return nil
}

type UpsertData struct {
	Mac  string
	DvID string
	Data map[uint16]Data
}

type Data struct {
	Value float64
	Time  time.Time
	DP    uint8
}
