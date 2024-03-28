package channel

import (
	"github.com/94peter/micro-service/di"
)

type DI interface {
	di.DI
	SetChannel(string)
	GetChannel() string
}
