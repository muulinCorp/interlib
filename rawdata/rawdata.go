package rawdata

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
)

func NewLib(clt *http.Client, url string) RawdataLib {
	return &rawdataImpl{
		clt: clt,
		url: url,
	}
}

type RawdataLib interface {
	CreateRawData(serviceID string, data map[string]interface{}) error
}

type rawdataImpl struct {
	clt *http.Client
	url string
}

func (rd *rawdataImpl) CreateRawData(serviceID string, data map[string]interface{}) error {
	const path = "/internal/v1/rawdata"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return err
	}
	rep, err := util.NewRequest(rd.clt).
		AddHeader("X-Service", serviceID).
		Body(&buf).Url(rd.url + path).Post()
	if err != nil {
		return err
	}
	if rep.Status != http.StatusOK {
		// 錯誤處理
		repErr := util.ParserErrorResp(rep)
		return errors.New(repErr.Title)
	}
	return nil
}
