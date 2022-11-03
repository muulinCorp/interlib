package mid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/muulin/interlib"
	"bitbucket.org/muulin/interlib/auth/dao"
	"bitbucket.org/muulin/interlib/diutil"
	"bitbucket.org/muulin/interlib/types"
	interlibUtil "bitbucket.org/muulin/interlib/util"

	"github.com/94peter/sterna/api"
	sternaMid "github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/auth"
	"github.com/94peter/sterna/db"
	"github.com/94peter/sterna/model/cache"
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

func NewInterAuthMid() sternaMid.AuthMidInter {
	return &interAuthMiddle{

		authMap:  make(map[string]uint8),
		groupMap: make(map[string][]auth.UserPerm),
	}
}

func NewGinInterAuthMid(serviceName string) sternaMid.AuthGinMidInter {
	return &interAuthMiddle{
		service:  serviceName,
		authMap:  make(map[string]uint8),
		groupMap: make(map[string][]auth.UserPerm),
	}
}

func (lm *interAuthMiddle) GetName() string {
	return "auth"
}

type interAuthMiddle struct {
	service string
	// clt      interAuth.AuthClient
	authMap  map[string]uint8
	groupMap map[string][]auth.UserPerm
}

func (am *interAuthMiddle) outputErr(c *gin.Context, err error) {
	api.GinOutputErr(c, am.service, err)
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

func getReqUserByAuthKey(c *gin.Context) auth.TargetReqUser {
	authKey := c.GetHeader("X-AuthKey")
	if authKey == "" {
		return nil
	}
	result, err := diutil.RedisReqHandler(c.Request, types.Redis_DB_User, func(clt db.RedisClient) (any, error) {
		sToken := dao.NewSerializationTokenByKey(authKey)
		simpleCache := cache.NewCache(sToken, clt)
		data, err := simpleCache.Get()
		if err != nil {
			return nil, err
		}
		err = sToken.Decode(data)
		if err != nil {
			return nil, err
		}
		pr := sToken.GetParseResult()

		return auth.NewTargetReqUser(pr.Target(), auth.NewReqUser(pr.Host(), pr.Sub(), pr.Account(), pr.Name(), pr.Perms())), nil
	})
	if err != nil || result == nil {
		return nil
	}
	return result.(auth.TargetReqUser)
}

func (am *interAuthMiddle) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if am.IsAuth(path, c.Request.Method) {
			reqUser := getReqUserByAuthKey(c)
			host := util.GetHost(c.Request)
			if reqUser == nil {
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
				grpconf := interlib.GetGrpcConfByCtx(c.Request.Context())
				authClt, err := grpconf.NewAuthClient()
				if err != nil {
					am.outputErr(c, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, err.Error()))
					c.Abort()
					return
				}
				reqUser, err = authClt.ValidateToken(host, authToken)
				if err != nil {
					am.outputErr(c, types.NewErrorWaper(types.ErrInvalidToken, err.Error()))
					c.Abort()
					return
				}
			}

			if reqUser.Host() != host && reqUser.Target() != host {
				am.outputErr(c, types.NewErrorWaper(types.ErrHostNotMatch, "host not match"))
				c.Abort()
				return
			}

			if hasPerm := am.HasPerm(path, c.Request.Method, reqUser.GetPerm()); !hasPerm {
				am.outputErr(c, types.NewErrorWaper(types.ErrNoPermission, "perm not allow"))
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
				grpconf := interlib.GetGrpcConfByCtx(r.Context())
				authClt, err := grpconf.NewAuthClient()
				if err != nil {
					api.OutputErr(w, types.NewErrorWaper(types.ErrAuthGrpcConnectFail, err.Error()))
					return
				}
				reqUser, err := authClt.ValidateToken(host, authToken)
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

func (pr parseTokenResult) Target() string {
	return pr["target"].(string)
}
