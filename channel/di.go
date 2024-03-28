package channel

import (
	"github.com/94peter/microservice/di"
)

type DI interface {
	di.DI
	SetChannel(string)
	GetChannel() string
}
