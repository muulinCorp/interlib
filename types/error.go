package types

import (
	"fmt"
	"net/http"

	apiErr "github.com/94peter/sterna/api/err"
)

var (
	ErrJsonEncodeFail      = apiErr.NewWithKey(http.StatusInternalServerError, "json encdoe fail: %s", "500001")
	ErrRequestGetFail      = apiErr.NewWithKey(http.StatusBadGateway, "request fail: %s", "502001")
	ErrAuthGrpcConnectFail = apiErr.NewWithKey(http.StatusServiceUnavailable, "json encdoe fail: %s", "503001")
	ErrPathNotFound        = apiErr.NewWithKey(http.StatusNotFound, "path not found: %s", "404001")
)

var (
	ErrTokenTimeout    = apiErr.NewWithKey(http.StatusUnauthorized, "token timeout", "401001")
	ErrMissingToken    = apiErr.NewWithKey(http.StatusUnauthorized, "missing", "401002")
	ErrInvalidToken    = apiErr.NewWithKey(http.StatusUnauthorized, "invalid token", "401003")
	ErrTokenParserFail = apiErr.NewWithKey(http.StatusUnauthorized, "token parser fail: %s", "401004")
	ErrHostNotMatch    = apiErr.NewWithKey(http.StatusUnauthorized, "host is not match", "401005")
	ErrNoPermission    = apiErr.NewWithKey(http.StatusUnauthorized, "you do not have permission", "401006")
)

// error key for device: XXX100~XXX159
// error key for controller: XXX160~179
// error key for sensor: XXX180~199

func NewErrorWaper(err apiErr.ApiError, detail string) apiErr.ApiError {
	return apiErr.NewWithKey(err.GetStatus(), fmt.Sprintf(err.GetErrorMsg(), detail), err.GetErrorKey())
}
