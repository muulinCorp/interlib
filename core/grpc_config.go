package core

import (
	"github.com/muulinCorp/interlib/core/interceptor"
	"github.com/muulinCorp/interlib/util"
	"google.golang.org/grpc"
)

type GrpcConfig struct {
	Port           int  `env:"GRPC_PORT"`
	ReflectService bool `env:"GRPC_REFLECT"`

	Logger              Log
	registerServiceFunc func(grpcServer *grpc.Server)
	interceptors        []interceptor.Interceptor
}

func (c *GrpcConfig) SetRegisterServiceFunc(f func(grpcServer *grpc.Server)) {
	c.registerServiceFunc = f
}

func (c *GrpcConfig) SetInterceptors(i ...interceptor.Interceptor) {
	c.interceptors = i
}

func GetConfigFromEnv() (*GrpcConfig, error) {
	var cfg GrpcConfig
	err := util.GetFromEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Log interface {
	Infof(format string, a ...any)
	Fatalf(format string, a ...any)
}
