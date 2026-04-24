package types

import (
	"errors"
	"time"
)

type ServiceSetting struct {
	Mock    bool          `mapstructure:"mock, omitempty"`
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func (s ServiceSetting) Validate() error {
	if !s.Mock {
		if s.Address == "" {
			return errors.New("address is required")
		}
		if s.Timeout == 0 {
			return errors.New("timeout is required")
		}
	}

	return nil
}