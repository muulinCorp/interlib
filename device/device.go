package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
	"github.com/94peter/sterna/api"
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
	UpsertSenValue(channelID string, inputData *UpsertData) api.ApiError
	UpsertConValue(channelID string, inputData *UpsertData) api.ApiError
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
		AddHeader("X-Channel", channel).
		Body(&buf).Url(dv.url + path).Post()
	if err != nil {
		return err
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf("%v100", repErr.Status)
		return api.NewApiErrorWithKey(repErr.Status, repErr.Title, key)
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
		key := fmt.Sprintf("%v101", repErr.Status)
		return "", api.NewApiErrorWithKey(repErr.Status, repErr.Title, key)
	}

	return string(resp.Body), nil
}

func (ct *deviceImpl) UpsertSenValue(channelID string, inputData *UpsertData) api.ApiError {
	const (
		path   = "/internal/v1/sensor"
		errKey = "%v180"
	)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputData)
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusBadRequest)
		return api.NewApiErrorWithKey(http.StatusBadRequest, err.Error(), key)
	}

	resp, err := util.NewRequest(ct.clt).
		AddHeader("X-Service", channelID).
		Body(&buf).Url(ct.url + path).Post()
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusInternalServerError)
		return api.NewApiErrorWithKey(http.StatusInternalServerError, err.Error(), key)
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf(errKey, repErr.Status)
		return api.NewApiErrorWithKey(repErr.Status, repErr.Title, key)
	}

	return nil
}

func (ct *deviceImpl) UpsertConValue(channelID string, inputData *UpsertData) api.ApiError {
	const (
		path   = "/internal/v1/controller"
		errKey = "%v160"
	)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputData)
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusBadRequest)
		return api.NewApiErrorWithKey(http.StatusBadRequest, err.Error(), key)
	}

	resp, err := util.NewRequest(ct.clt).
		AddHeader("X-Service", channelID).
		Body(&buf).Url(ct.url + path).Post()
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusInternalServerError)
		return api.NewApiErrorWithKey(http.StatusInternalServerError, err.Error(), key)
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf(errKey, repErr.Status)
		return api.NewApiErrorWithKey(repErr.Status, repErr.Title, key)
	}

	return nil
}
