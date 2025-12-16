package types

import "time"

type ServiceSetting struct {
	Mock    bool          `mapstructure:"mock"`
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}
