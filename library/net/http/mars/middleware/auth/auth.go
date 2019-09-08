package auth

import (
	"strings"
	idtv1 "valerian/app/service/identify/api/grpc"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
	"valerian/library/net/rpc/warden"

	"github.com/pkg/errors"
)

// Config is the identify config model.
type Config struct {
	Identify *warden.ClientConfig
	// csrf switch.
	DisableCSRF bool
}

// Auth is the authorization middleware
type Auth struct {
	idtv1.IdentifyClient

	conf *Config
}

var _defaultConf = &Config{
	Identify:    nil,
	DisableCSRF: false,
}

type authFunc func(*mars.Context) (int64, error)

// New is used to create an authorization middleware
func New(conf *Config) *Auth {
	if conf == nil {
		conf = _defaultConf
	}
	identify, err := idtv1.NewClient(conf.Identify)
	if err != nil {
		panic(errors.WithMessage(err, "Failed to dial identify service"))
	}
	auth := &Auth{
		IdentifyClient: identify,
		conf:           conf,
	}
	return auth
}

func (a *Auth) User(c *mars.Context) {
	req := c.Request
	cookie, _ := req.Cookie("token")
	if cookie != nil {
		a.UserWeb(c)
	} else {
		a.UserMobile(c)
	}
}

// UserWeb is used to mark path as web access required.
func (a *Auth) UserWeb(ctx *mars.Context) {
	a.midAuth(ctx, a.AuthCookie)
}

// UserMobile is used to mark path as mobile access required.
func (a *Auth) UserMobile(ctx *mars.Context) {
	a.midAuth(ctx, a.AuthToken)
}

// AuthToken is used to authorize request by token
func (a *Auth) AuthToken(ctx *mars.Context) (accountID int64, err error) {
	req := ctx.Request

	tokenStr := ""
	if v, ok := getBearer(req.Header["Authorization"]); ok {
		tokenStr = v
	}

	if tokenStr == "" {
		return 0, ecode.NoLogin
	}

	reply, err := a.GetTokenInfo(ctx, &idtv1.TokenReq{Token: tokenStr})
	if err != nil {
		return 0, err
	}

	if !reply.Login {
		return 0, ecode.NoLogin
	}

	return reply.Aid, nil
}

// AuthCookie is used to authorize request by cookie
func (a *Auth) AuthCookie(ctx *mars.Context) (int64, error) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		return 0, ecode.NoLogin
	}

	reply, err := a.GetTokenInfo(ctx, &idtv1.TokenReq{Token: cookie.Value})
	if err != nil {
		return 0, err
	}

	if !reply.Login {
		return 0, ecode.NoLogin
	}

	return reply.Aid, nil
}

func (a *Auth) midAuth(ctx *mars.Context, auth authFunc) {
	accountID, err := auth(ctx)
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
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
