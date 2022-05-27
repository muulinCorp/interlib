package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
)

func NewLib(clt *http.Client, url string) ControllerLib {
	return &controllerImpl{
		clt: clt,
		url: url,
	}
}

type ControllerLib interface {
	Upsert(channelID string, inputData *UpsertData) error
}

type controllerImpl struct {
	clt *http.Client
	url string
}

func (ct *controllerImpl) Upsert(channelID string, inputData *UpsertData) error {
	const path = "/internal/v1/controller"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputData)
	if err != nil {
		return err
	}

	resp, err := util.NewRequest(ct.clt).
		AddHeader("X-Service", channelID).
		Body(&buf).Url(ct.url + path).Post()
	if err != nil {
		return err
	}
	if resp.Status != http.StatusOK {
		repErr := util.ParserErrorResp(resp)
		return errors.New(repErr.Title+"("+repErr.Status+")")
	}

	return nil
}