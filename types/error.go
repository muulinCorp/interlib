package types

import (
	"fmt"
	"net/http"

	apiErr "github.com/94peter/api-toolkit/errors"
)

var (
	ErrJsonEncodeFail      = apiErr.New(http.StatusInternalServerError, "json encdoe fail")
	ErrRequestGetFail      = apiErr.New(http.StatusBadGateway, "request fail")
	ErrAuthGrpcConnectFail = apiErr.New(http.StatusServiceUnavailable, "json encdoe fail")
	ErrPathNotFound        = apiErr.New(http.StatusNotFound, "path not found")
)

var (
	ErrTokenTimeout    = apiErr.New(http.StatusUnauthorized, "token timeout")
	ErrMissingToken    = apiErr.New(http.StatusUnauthorized, "missing")
	ErrInvalidToken    = apiErr.New(http.StatusUnauthorized, "invalid token")
	ErrTokenParserFail = apiErr.New(http.StatusUnauthorized, "token parser fail")
	ErrHostNotMatch    = apiErr.New(http.StatusUnauthorized, "host is not match")
	ErrNoPermission    = apiErr.New(http.StatusUnauthorized, "you do not have permission")
)

// error key for device: XXX100~XXX159
// error key for controller: XXX160~179
// error key for sensor: XXX180~199

func NewErrorWaper(err apiErr.ApiError, detail string) apiErr.ApiError {
	return apiErr.PkgError(err.GetStatus(), fmt.Errorf("%s:%s", err.Error(), detail))
}
