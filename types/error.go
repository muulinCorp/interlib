package types

import (
	"fmt"
	"net/http"

	"github.com/94peter/sterna/api"
)

var (
	ErrJsonEncodeFail = api.NewApiErrorWithKey(http.StatusInternalServerError, "json encdoe fail: %s", "500001")
	ErrRequestGetFail = api.NewApiErrorWithKey(http.StatusBadGateway, "request fail: %s", "502001")
	ErrPathNotFound   = api.NewApiErrorWithKey(http.StatusNotFound, "path not found: %s", "404001")
)

var (
	ErrTokenTimeout    = api.NewApiErrorWithKey(http.StatusUnauthorized, "token timeout", "401001")
	ErrMissingToken    = api.NewApiErrorWithKey(http.StatusUnauthorized, "missing", "401002")
	ErrInvalidToken    = api.NewApiErrorWithKey(http.StatusUnauthorized, "invalid token", "401003")
	ErrTokenParserFail = api.NewApiErrorWithKey(http.StatusUnauthorized, "token parser fail: %s", "401004")
	ErrHostNotMatch    = api.NewApiErrorWithKey(http.StatusUnauthorized, "host is not match", "401005")
	ErrNoPermission    = api.NewApiErrorWithKey(http.StatusUnauthorized, "you do not have permission", "401006")
)

func NewErrorWaper(err api.ApiError, detail string) api.ApiError {
	return api.NewApiErrorWithKey(err.GetStatus(), fmt.Sprintf(err.GetErrorMsg(), detail), err.GetErrorKey())
}
