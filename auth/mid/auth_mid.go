package mid

import (
	"fmt"
	"net/http"
	"strings"

	authClient "bitbucket.org/muulin/interlib/auth/client"
	"bitbucket.org/muulin/interlib/auth/pb"
	"bitbucket.org/muulin/interlib/types"

	apiErr "github.com/94peter/sterna/api/err"
	sternaMid "github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/auth"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"
)

const (
	authValue = uint8(1 << iota)
)

func getPathKey(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

func NewGinInterAuthMid(address, serviceName string) (sternaMid.AuthGinMidInter, error) {
	authSDK := authClient.New(address)
	return &interAuthMiddle{
		service:  serviceName,
		authSDK:  authSDK,
		authMap:  make(map[string]uint8),
		groupMap: make(map[string][]auth.UserPerm),
	}, nil
}

func (lm *interAuthMiddle) GetName() string {
	return "auth"
}

type interAuthMiddle struct {
	service string
	authSDK authClient.AuthClient
	// clt      interAuth.AuthClient
	authMap  map[string]uint8
	groupMap map[string][]auth.UserPerm
}

func (am *interAuthMiddle) outputErr(c *gin.Context, err error) {
	apiErr.GinOutputErr(c, am.service, err)
}

func (am *interAuthMiddle) AddAuthPath(path string, method string, isAuth bool, group []auth.UserPerm) {
	value := uint8(0)
	if isAuth {
		value = value | authValue
	}
	key := getPathKey(path, method)
	am.authMap[key] = uint8(value)
	am.groupMap[key] = group
}

func (am *interAuthMiddle) IsAuth(path string, method string) bool {
	key := getPathKey(path, method)
	value, ok := am.authMap[key]
	if ok {
		return (value & authValue) > 0
	}
	return false
}

func (am *interAuthMiddle) HasPerm(path, method string, perm []string) bool {
	key := fmt.Sprintf("%s:%s", path, method)
	groupAry, ok := am.groupMap[key]
	if !ok || groupAry == nil || len(groupAry) == 0 {
		return true
	}
	for _, g := range groupAry {
		if util.IsStrInList(string(g), perm...) {
			return true
		}
	}
	return false
}

func (am *interAuthMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if am.IsAuth(path, c.Request.Method) {

			host := util.GetHost(c.Request)

			authToken := c.GetHeader(sternaMid.BearerAuthTokenKey)
			if authToken == "" {
				am.outputErr(c, types.NewErrorWaper(types.ErrMissingToken, path))
				c.Abort()
				return
			}
			if !strings.HasPrefix(authToken, "Bearer ") {
				am.outputErr(c, types.NewErrorWaper(types.ErrInvalidToken, "not bearer token"))
				c.Abort()
				return
			}

			user, err := am.authSDK.GetUserInfo(c, &pb.GetUserInfoRequest{
				Host: host, Token: authToken[7:],
			})
			if err != nil {
				am.outputErr(c, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, err.Error()))
				c.Abort()
				return
			}
			if user.StatusCode != http.StatusOK {
				am.outputErr(c, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, user.Message))
				c.Abort()
				return
			}

			if hasPerm := am.HasPerm(path, c.Request.Method, user.Roles); !hasPerm {
				am.outputErr(c, types.NewErrorWaper(types.ErrNoPermission, "perm not allow"))
				c.Abort()
				return
			}

			if channel := c.GetHeader("X-Channel"); channel != "" && !util.IsStrInList(channel, user.Channels...) {
				am.outputErr(c, types.NewErrorWaper(types.ErrNoPermission, "channel not allow"))
				c.Abort()
				return
			}

			c.Set(string(auth.CtxUserInfoKey), auth.NewReqUser(host, user.Sub, user.Account, user.Name, user.Roles))
		}
		c.Next()
	}
}
