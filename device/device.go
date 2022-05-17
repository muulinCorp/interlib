package device

import (
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
	GetService(mac, gwid string) (string, error)
}

type deviceImpl struct {
	clt *http.Client
	url string
}

func (dv *deviceImpl) GetService(mac, gwid string) (string, error) {
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