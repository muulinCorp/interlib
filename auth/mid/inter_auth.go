package mid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	interAuth "bitbucket.org/muulin/interlib/auth"
	"bitbucket.org/muulin/interlib/types"
	interlibUtil "bitbucket.org/muulin/interlib/util"

	"github.com/94peter/sterna/api"
	sternaMid "github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/auth"
	"github.com/94peter/sterna/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

const (
	authValue = uint8(1 << iota)
)

func getPathKey(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

func NewInterAuthMid(authClt interAuth.AuthClient) sternaMid.AuthMidInter {
	return &interAuthMiddle{

		authMap:  make(map[string]uint8),
		groupMap: make(map[string][]auth.UserPerm),
	}
}

func NewGinInterAuthMid(authClt interAuth.AuthClient) sternaMid.AuthGinMidInter {
	return &interAuthMiddle{
		clt:      authClt,
		authMap:  make(map[string]uint8),
		groupMap: make(map[string][]auth.UserPerm),
	}
}

func (lm *interAuthMiddle) GetName() string {
	return "auth"
}

type interAuthMiddle struct {
	clt      interAuth.AuthClient
	authMap  map[string]uint8
	groupMap map[string][]auth.UserPerm
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
		path := c.Request.URL.Path
		if am.IsAuth(path, c.Request.Method) {
			authToken := c.GetHeader(sternaMid.BearerAuthTokenKey)
			if authToken == "" {
				api.GinOutputErr(c, types.NewErrorWaper(types.ErrMissingToken, path))
				c.Abort()
				return
			}
			if !strings.HasPrefix(authToken, "Bearer ") {
				api.GinOutputErr(c, types.NewErrorWaper(types.ErrInvalidToken, "not bearer token"))
				c.Abort()
				return
			}
			host := util.GetHost(c.Request)
			reqUser, err := am.clt.ValidateToken(host, authToken)
			if err != nil {
				fmt.Println("validate token fail")
				api.GinOutputErr(c, types.NewErrorWaper(types.ErrInvalidToken, err.Error()))
				c.Abort()
				return
			}

			if reqUser.Host() != host {
				api.GinOutputErr(c, types.NewErrorWaper(types.ErrHostNotMatch, "host not match"))
				c.Abort()
				return
			}

			if hasPerm := am.HasPerm(path, c.Request.Method, reqUser.GetPerm()); !hasPerm {
				api.GinOutputErr(c, types.NewErrorWaper(types.ErrNoPermission, "perm not allow"))
				c.Abort()
				return
			}
			c.Request = util.SetCtxKeyVal(c.Request, auth.CtxUserInfoKey, reqUser)
		}
		c.Next()
	}
}

func (am *interAuthMiddle) GetMiddleWare() func(f http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// one time scope setup area for middleware
		return func(w http.ResponseWriter, r *http.Request) {
			path, err := mux.CurrentRoute(r).GetPathTemplate()
			if err != nil {
				api.OutputErr(w, types.NewErrorWaper(types.ErrPathNotFound, path))
				return
			}
			if am.IsAuth(path, r.Method) {
				authToken := r.Header.Get(sternaMid.BearerAuthTokenKey)
				if authToken == "" {
					api.OutputErr(w, types.ErrMissingToken)
					return
				}
				if !strings.HasPrefix(authToken, "Bearer ") {
					api.OutputErr(w, types.ErrInvalidToken)
					return
				}
				// 打api取得token內容
				host := util.GetHost(r)
				reqUser, err := am.clt.ValidateToken(host, authToken)
				if err != nil {
					api.OutputErr(w, err)
					return
				}

				if reqUser.Host() != host {
					api.OutputErr(w, types.ErrHostNotMatch)
					return
				}

				if hasPerm := am.HasPerm(path, r.Method, reqUser.GetPerm()); !hasPerm {
					api.OutputErr(w, types.ErrNoPermission)
					return
				}
				r = util.SetCtxKeyVal(r, auth.CtxUserInfoKey, reqUser)
			}
			f(w, r)
		}
	}
}

func getParserToken(url, token string) (sternaMid.TokenParserResult, error) {

	req := interlibUtil.NewRequest(&http.Client{})

	res, err := req.AddHeader("Authorization", token).Url(url).Get()
	if err != nil {
		return nil, types.NewErrorWaper(types.ErrRequestGetFail, err.Error())
	}
	if res.Status != http.StatusOK {
		errRes := interlibUtil.ParserErrorResp(res)
		return nil, api.NewApiErrorWithKey(errRes.Status, errRes.Title, errRes.ErrorKey)
	}

	pr := parseTokenResult{}
	err = json.Unmarshal(res.Body, &pr)
	if err != nil {
		return nil, types.NewErrorWaper(types.ErrJsonEncodeFail, err.Error())
	}
	return pr, nil
}

type parseTokenResult map[string]interface{}

func (pr parseTokenResult) Account() string {
	return pr["account"].(string)
}

func (pr parseTokenResult) Host() string {
	return pr["host"].(string)
}

func (pr parseTokenResult) Name() string {
	return pr["name"].(string)
}

func (pr parseTokenResult) Perms() []string {
	plist := pr["perms"].([]interface{})
	l := len(plist)
	perms := make([]string, l)
	for i := 0; i < l; i++ {
		perms[i] = plist[i].(string)
	}
	return perms
}

func (pr parseTokenResult) Sub() string {
	return pr["sub"].(string)
}
