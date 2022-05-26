package device

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
)

func NewLib(clt *http.Client, url string) DeviceLib {
	return &deviceImpl{
		clt: clt,
		url: url,
	}
}

type DeviceLib interface {
	CreateDevice(channel string, inputDevice *NewDevice) error
	GetChannel(mac, gwid string) (string, error)
}

type deviceImpl struct {
	clt *http.Client
	url string
}

func (dv *deviceImpl) CreateDevice(channel string, inputDevice *NewDevice) error {
	const path = "/internal/v1/device"

	err := inputDevice.Valid()
	if err != nil {
		return err
	} 

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(inputDevice)
	if err != nil {
		return err
	}

	resp, err := util.NewRequest(dv.clt).
		AddHeader("X-Service", channel).
		Body(&buf).Url(dv.url + path).Post()
	if err != nil {
		return err
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		return errors.New(repErr.Title+"("+repErr.Status+")")
	}

	return nil	
}

func (dv *deviceImpl) GetChannel(mac, gwid string) (string, error) {
	const path = "/internal/v1/device/service"

	req := util.NewRequest(dv.clt).Url(dv.url+path).
			AddQuery("macAddress", mac).
			AddQuery("gwID", gwid)

	resp, err := req.Get()
	if err != nil {
		return "", err
	}

	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		return "", errors.New(repErr.Title)
	}

	return string(resp.Body), nil
}