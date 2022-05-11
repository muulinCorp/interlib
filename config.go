package interlib

import (
	"net/http"

	"bitbucket.org/muulin/interlib/rawdata"
)

type Conf struct {
	Url string
}

func (c *Conf) NewRawDataLib(clt *http.Client) rawdata.RawdataLib {
	util.NewRequest(aa)
	return rawdata.NewLib(clt, c.Url)
}
