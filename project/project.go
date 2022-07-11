package project

import (
	"net/http"
)

func NewLib(clt *http.Client, url string) ProejctLib {
	return &projectImpl{
		clt: clt,
		url: url,
	}
}

type projectImpl struct {
	clt *http.Client
	url string
}

type ProejctLib interface {
}
