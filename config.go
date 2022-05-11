package interlib

import (
	"net/http"

	"bitbucket.org/muulin/interlib/rawdata"
)

type Conf struct {
	Url string
}

func (c *Conf) NewRawDataLib(clt *http.Client) rawdata.RawdataLib {
	return rawdata.NewLib(clt, c.Url)
}
