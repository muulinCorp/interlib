package interlib

import (
	"context"
	"fmt"
	"io/ioutil"

	"bitbucket.org/muulin/interlib/auth"
	"bitbucket.org/muulin/interlib/channel"
	"bitbucket.org/muulin/interlib/comm"
	appDevice "bitbucket.org/muulin/interlib/device/app"
	coreDevice "bitbucket.org/muulin/interlib/device/core"
	"bitbucket.org/muulin/interlib/message"
	"github.com/94peter/sterna/log"
	"github.com/94peter/sterna/util"
	"gopkg.in/yaml.v2"
)

const (
	CtxGrpcConfKey = util.CtxKey("gRPConf")
)

func GetGrpcConfByCtx(ctx context.Context) GrpcRouterConf {
	val := ctx.Value(CtxGrpcConfKey)
	if conf, ok := val.(GrpcRouterConf); ok {
		return conf
	}
	return nil
}

type GrpcRouterConf map[string]string

func (conf *GrpcRouterConf) InitConfByByte(b []byte) {
	err := yaml.Unmarshal(b, conf)
	if err != nil {
		panic(fmt.Errorf("yaml unmarshal %s fail: %s", string(b), err.Error()))
	}
}

func (conf *GrpcRouterConf) InitConfByFile(f string) {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		panic(fmt.Errorf("load conf %s fail: %s", f, err.Error()))
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(fmt.Errorf("yaml unmarshal %s fail: %s", f, err.Error()))
	}
}

func (conf GrpcRouterConf) getAddress(key string) (string, error) {
	address, ok := conf[key]
	if !ok {
		return "", fmt.Errorf("config not set router key [%s]", key)
	}
	return address, nil
}

func (conf GrpcRouterConf) NewChannelClient() (channel.ChannelClient, error) {
	address, err := conf.getAddress(channel.RouterKey)
	if err != nil {
		return nil, err
	}
	return channel.NewGrpcClient(address)
}

func (conf GrpcRouterConf) NewCoreDeviceClient() (coreDevice.CoreDeviceClient, error) {
	address, err := conf.getAddress(coreDevice.RouterKey)
	if err != nil {
		return nil, err
	}
	return coreDevice.NewGrpcClient(address)
}

func (conf GrpcRouterConf) NewAppDeviceClient() (appDevice.AppDeviceClient, error) {
	address, err := conf.getAddress(appDevice.RouterKey)
	if err != nil {
		return nil, err
	}
	return appDevice.NewGrpcClient(address)
}

func (conf GrpcRouterConf) NewMessageClient() (message.MessageClient, error) {
	address, err := conf.getAddress(message.RouterKey)
	if err != nil {
		return nil, err
	}
	return message.NewGrpcClient(address)
}

func (conf GrpcRouterConf) NewAuthClient() (auth.AuthClient, error) {
	address, err := conf.getAddress(auth.RouterKey)
	if err != nil {
		return nil, err
	}
	return auth.NewGrpcClient(address)
}

func (conf GrpcRouterConf) NewCommClient(l log.Logger) (comm.CommClient, error) {
	address, err := conf.getAddress(comm.RouterKey)
	if err != nil {
		return nil, err
	}
	return comm.NewGrpcClient(address, l)
}
