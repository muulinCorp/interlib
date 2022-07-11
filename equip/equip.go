package rawdata

import (
	"errors"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
)

func NewLib(clt *http.Client, url string) EquipLib {
	return &equipImpl{
		clt: clt,
		url: url,
	}
}

type EquipLib interface {
	Test() error
	TestMux() error
}

type equipImpl struct {
	clt *http.Client
	url string
}

func (rd *equipImpl) Test() error {
	const path = "/v1/equip"

	rep, err := util.NewRequest(rd.clt).Url(rd.url + path).Get()
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

func (rd *equipImpl) TestMux() error {
	const path = "/mux/v1/equip"

	rep, err := util.NewRequest(rd.clt).Url(rd.url + path).Get()
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
