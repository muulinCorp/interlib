package util

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Request interface {
	Url(u string) Request
	AddHeader(key, value string) Request
	Body(b io.Reader) Request

	Get() (*response, error)
	Post() (*response, error)
}

func NewRequest(clt *http.Client) Request {
	return &myRequest{
		clt: clt,
	}
}

type myRequest struct {
	clt    *http.Client
	url    string
	header http.Header
	body   io.Reader
}

func (r *myRequest) Url(u string) Request {
	r.url = u
	return r
}

func (r *myRequest) Body(b io.Reader) Request {
	r.body = b
	return r
}

func (r *myRequest) AddHeader(key, value string) Request {
	if r.header == nil {
		r.header = make(http.Header)
	}
	if v, ok := r.header[key]; ok {
		r.header[key] = append(v, value)
	} else {
		r.header[key] = []string{value}
	}
	return r
}

type response struct {
	Status int
	Header http.Header
	Body   []byte
}

func (r *myRequest) Get() (*response, error) {
	if r.clt == nil {
		return nil, errors.New("clt is nil")
	}
	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return nil, err
	}
	if r.header != nil {
		req.Header = r.header
	}
	res, err := r.clt.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &response{
		Status: res.StatusCode,
		Body:   data,
		Header: res.Header,
	}, nil
}

func (r *myRequest) Post() (*response, error) {
	if r.clt == nil {
		return nil, errors.New("clt is nil")
	}
	req, err := http.NewRequest("POST", r.url, r.body)
	if err != nil {
		return nil, err
	}
	if r.header != nil {
		req.Header = r.header
	}
	res, err := r.clt.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &response{
		Status: res.StatusCode,
		Body:   data,
		Header: res.Header,
	}, nil
}
