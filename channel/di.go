package channel

import (
	"github.com/94peter/di"
)

type DI interface {
	di.DI
	SetChannel(string)
	GetChannel() string
}
