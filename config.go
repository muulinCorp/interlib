package interlib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/muulin/interlib/channel"
	appDevice "bitbucket.org/muulin/interlib/device/app"
	coreDevice "bitbucket.org/muulin/interlib/device/core"
	"bitbucket.org/muulin/interlib/message"
	"bitbucket.org/muulin/interlib/rawdata"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Url string
}

func (c *Conf) NewRawDataLib(clt *http.Client) rawdata.RawdataLib {
	return rawdata.NewLib(clt, c.Url)
}

type GrpcRouterConf map[string]string

func (conf *GrpcRouterConf) InitConfByFile(f string) {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("load conf fail: " + f)
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
}

func (conf GrpcRouterConf) getAddress(key string) (string, error) {
	address, ok := conf[key]
	if !ok {
		return "", fmt.Errorf("config not set router key [%s]", channel.RouterKey)
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
