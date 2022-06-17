package channel

import (
	"errors"
	"fmt"
	"net/http"

	"bitbucket.org/muulin/interlib/util"
)

func NewLib(clt *http.Client, url string) ChannelLib {
	return &channelImpl{
		clt: clt,
		url: url,
	}
}

type ChannelLib interface {
	GetConf(channelID string, env string) ([]byte, error)
}

type channelImpl struct {
	clt *http.Client
	url string
}

func (rd *channelImpl) GetConf(channelID string, env string) (conf []byte, err error) {
	const path = "/v1/channel/%s/conf/%s"
	rep, err := util.NewRequest(rd.clt).Url(fmt.Sprintf(rd.url+path, channelID, env)).Get()
	if err != nil {
		return
	}
	if rep.Status != http.StatusOK {
		// 錯誤處理
		repErr := util.ParserErrorResp(rep)
		err = errors.New(repErr.Title)
		return
	}
	conf = rep.Body
	return
}
