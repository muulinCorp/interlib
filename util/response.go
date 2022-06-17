package util

import (
	"encoding/json"
)

type ErrorResp struct {
	Status   int
	Title    string
	ErrorKey string `json:"errorKey"`
}

func ParserErrorResp(res *response) *ErrorResp {
	// 錯誤處理
	errRes := &ErrorResp{
		Status: res.Status,
	}
	switch res.Header.Get("Content-Type") {
	case "application/json":
		json.Unmarshal(res.Body, errRes)
	case "text/plain; charset=utf-8":
		errRes.Title = string(res.Body)
	}
	return errRes
}
