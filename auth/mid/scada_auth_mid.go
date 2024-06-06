package mid

import (
	"context"
	"net/http"
	"strings"
	"time"

	authClient "github.com/muulinCorp/interlib/auth/client"
	"github.com/muulinCorp/interlib/auth/pb"
	"github.com/muulinCorp/interlib/types"

	apitool "github.com/94peter/api-toolkit"
	"github.com/94peter/api-toolkit/auth"
	"github.com/94peter/api-toolkit/errors"
	"github.com/gin-gonic/gin"
)

func NewGinScadaAuthMid(address string) auth.GinAuthMidInter {
	authSDK := authClient.New(address)
	return &scadaAuthMiddle{
		auth:    auth.NewGinBearAuthMid(true),
		authSDK: authSDK,
	}
}

func (lm *scadaAuthMiddle) GetName() string {
	return "auth"
}

type scadaAuthMiddle struct {
	errors.CommonApiErrorHandler
	auth    auth.GinAuthMidInter
	authSDK authClient.AuthClient
}

func (am *scadaAuthMiddle) AddAuthPath(path string, method string, isAuth bool, group []auth.ApiPerm) {
	am.auth.AddAuthPath(path, method, isAuth, group)
}
func (am *scadaAuthMiddle) IsAuth(path string, method string) bool {
	return am.auth.IsAuth(path, method)
}
func (am *scadaAuthMiddle) HasPerm(path, method string, perm []string) bool {
	return am.auth.HasPerm(path, method, perm)
}

func (am *scadaAuthMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if am.IsAuth(path, c.Request.Method) {

			host := apitool.GetHost(c.Request)

			authToken := c.GetHeader(auth.BearerAuthTokenKey)
			if authToken == "" {

				am.GinApiErrorHandler(c, types.NewErrorWaper(types.ErrMissingToken, path))
				c.Abort()
				return
			}
			if !strings.HasPrefix(authToken, "Bearer ") {
				am.GinApiErrorHandler(c, types.NewErrorWaper(types.ErrInvalidToken, "not bearer token"))
				c.Abort()
				return
			}
			timeout, cancel := context.WithTimeout(c, 3*time.Second)
			defer cancel()
			user, err := am.authSDK.GetTokenInfo(timeout, &pb.GetTokenInfoRequest{
				Host: host, Token: authToken[7:],
			})
			if err != nil {
				am.GinApiErrorHandler(c, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, err.Error()))
				c.Abort()
				return
			}
			if user.StatusCode != http.StatusOK {
				am.GinApiErrorHandler(c, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, user.Message))
				c.Abort()
				return
			}

			if hasPerm := am.HasPerm(path, c.Request.Method, user.Roles); !hasPerm {
				am.GinApiErrorHandler(c, types.NewErrorWaper(types.ErrNoPermission, "perm not allow"))
				c.Abort()
				return
			}

			auth.SetReqUserToGin(c, auth.NewReqUser(host, user.Sub, user.Account, user.Name, user.Roles, user.Useage))
		}
		c.Next()
	}
}
