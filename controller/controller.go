package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
	"github.com/94peter/sterna/api"
)

func NewLib(clt *http.Client, url string) ControllerLib {
	return &controllerImpl{
		clt: clt,
		url: url,
	}
}

type ControllerLib interface {
	Upsert(channelID string, inputData *UpsertData) api.ApiError
}

type controllerImpl struct {
	clt *http.Client
	url string
}

func (ct *controllerImpl) Upsert(channelID string, inputData *UpsertData) api.ApiError {
	const (
		path = "/internal/v1/controller"
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