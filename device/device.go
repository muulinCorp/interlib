package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
	apiErr "github.com/94peter/sterna/api/err"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewLib(clt *http.Client, url string) DeviceLib {
	return &deviceImpl{
		clt: clt,
		url: url,
	}
}

type DeviceLib interface {
	CheckDvExist(oID primitive.ObjectID, channel string) (bool, error)
	CreateDevice(channel string, inputDevice *NewDevice) error
	GetChannel(mac, gwid string) (string, error)
	UpsertSenValue(channelID string, inputData *UpsertData) apiErr.ApiError
	UpsertConValue(channelID string, inputData *UpsertData) apiErr.ApiError
}

type deviceImpl struct {
	clt *http.Client
	url string
}

func (dv *deviceImpl) CheckDvExist(oID primitive.ObjectID, channel string) (bool, error) {
	const (
		errKey = "%v102"
		path   = "internal/v1/device/exist"
	)

	resp, err := util.NewRequest(dv.clt).
		AddHeader("X-Channel", channel).Url(dv.url + path).Get()
	if err != nil {
		return false, err
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf(errKey, repErr.Status)
		return false, apiErr.NewWithKey(repErr.Status, repErr.Title, key)
	}

	var res bool
	err = json.Unmarshal(resp.Body, &res)
	if err != nil {
		return false, err
	}

	return res, nil
}

func (dv *deviceImpl) CreateDevice(channel string, inputDevice *NewDevice) error {
	const (
		errKey = "%v100"
		path   = "/internal/v1/device"
	)

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
		key := fmt.Sprintf(errKey, repErr.Status)
		return apiErr.NewWithKey(repErr.Status, repErr.Title, key)
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
		return "", apiErr.NewWithKey(repErr.Status, repErr.Title, key)
	}

	return string(resp.Body), nil
}

func (ct *deviceImpl) UpsertSenValue(channelID string, inputData *UpsertData) apiErr.ApiError {
	const (
		path   = "/internal/v1/sensor"
		errKey = "%v180"
	)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputData)
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusBadRequest)
		return apiErr.NewWithKey(http.StatusBadRequest, err.Error(), key)
	}

	resp, err := util.NewRequest(ct.clt).
		AddHeader("X-Service", channelID).
		Body(&buf).Url(ct.url + path).Post()
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusInternalServerError)
		return apiErr.NewWithKey(http.StatusInternalServerError, err.Error(), key)
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf(errKey, repErr.Status)
		return apiErr.NewWithKey(repErr.Status, repErr.Title, key)
	}

	return nil
}

func (ct *deviceImpl) UpsertConValue(channelID string, inputData *UpsertData) apiErr.ApiError {
	const (
		path   = "/internal/v1/controller"
		errKey = "%v160"
	)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputData)
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusBadRequest)
		return apiErr.NewWithKey(http.StatusBadRequest, err.Error(), key)
	}

	resp, err := util.NewRequest(ct.clt).
		AddHeader("X-Service", channelID).
		Body(&buf).Url(ct.url + path).Post()
	if err != nil {
		key := fmt.Sprintf(errKey, http.StatusInternalServerError)
		return apiErr.NewWithKey(http.StatusInternalServerError, err.Error(), key)
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		key := fmt.Sprintf(errKey, repErr.Status)
		return apiErr.NewWithKey(repErr.Status, repErr.Title, key)
	}

	return nil
}
