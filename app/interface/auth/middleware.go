package auth

import (
	"strings"
	"valerian/app/conf"
	"valerian/app/interface/passport-auth/service"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

type authFunc func(*mars.Context) (int64, error)

var (
	srv *service.Service
)

func Init(c *conf.Config) {
	srv = service.New(c)
}

func User(c *mars.Context) {
	req := c.Request
	cookie, _ := req.Cookie("token")
	if cookie != nil {
		UserWeb(c)
	} else {
		UserMobile(c)
	}
}

// UserWeb is used to mark path as web access required.
func UserWeb(ctx *mars.Context) {
	midAuth(ctx, AuthCookie)
}

// UserMobile is used to mark path as mobile access required.
func UserMobile(ctx *mars.Context) {
	midAuth(ctx, AuthToken)
}

// AuthToken is used to authorize request by token
func AuthToken(ctx *mars.Context) (accountID int64, err error) {
	req := ctx.Request

	tokenStr := ""
	if v, ok := getBearer(req.Header["Authorization"]); ok {
		tokenStr = v
	}

	if tokenStr == "" {
		return 0, ecode.NoLogin
	}

	reply, err := srv.Auth(ctx, tokenStr)
	if err != nil {
		return 0, ecode.NoLogin
	}

	return reply.Aid, nil
}

// AuthCookie is used to authorize request by cookie
func AuthCookie(ctx *mars.Context) (int64, error) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		return 0, ecode.NoLogin
	}

	reply, err := srv.Auth(ctx, cookie.Value)
	if err != nil {
		return 0, ecode.NoLogin
	}

	return reply.Aid, nil
}

func midAuth(ctx *mars.Context, auth authFunc) {
	accountID, err := auth(ctx)
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		return
	}
	setAccountID(ctx, accountID)
}

func getBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) == 2 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

// set account id into context
func setAccountID(ctx *mars.Context, id int64) {
	ctx.Set("aid", id)
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Aid] = id
		return
	}
}
