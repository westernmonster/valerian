package middleware

import (
	"context"
	"strings"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/db"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
	"valerian/modules/repo"
	"valerian/modules/usecase"
)

type authFunc func(*mars.Context) (int64, error)

type Auth struct {
	OauthUsecase interface {
		GetByID(c context.Context, bizCtx *biz.BizContext, userID int64) (item *repo.Account, err error)
		GetTokenInfo(tokenStr string) (claims *infrastructure.TokenClaims, err error)
	}
}

func New() *Auth {
	node, err := db.InitDatabase()
	if err != nil {
		panic(err)
	}
	auth := &Auth{
		OauthUsecase: &usecase.OauthUsecase{
			Node:                       node,
			AccountRepository:          &repo.AccountRepository{},
			ValcodeRepository:          &repo.ValcodeRepository{},
			SessionRepository:          &repo.SessionRepository{},
			OauthClientRepository:      &repo.OauthClientRepository{},
			OauthAccessTokenRepository: &repo.OauthAccessTokenRepository{},
		},
	}
	return auth
}

// User is used to mark path as access required.
// If `access_key` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (p *Auth) User(ctx *mars.Context) {
	req := ctx.Request
	cookie, _ := req.Cookie("token")
	if cookie != nil {
		p.UserWeb(ctx)
	} else {
		p.UserMobile(ctx)
	}
}

// UserWeb is used to mark path as web access required.
func (p *Auth) UserWeb(ctx *mars.Context) {
	p.midAuth(ctx, p.AuthCookie)
}

// UserMobile is used to mark path as mobile access required.
func (p *Auth) UserMobile(ctx *mars.Context) {
	p.midAuth(ctx, p.AuthToken)
}

// AuthToken is used to authorize request by token
func (p *Auth) AuthToken(ctx *mars.Context) (accountID int64, err error) {
	req := ctx.Request

	tokenStr := ""
	if v, ok := getBearer(req.Header["Authorization"]); ok {
		tokenStr = v
	}

	if tokenStr == "" {
		return 0, ecode.NoLogin
	}

	token, err := p.OauthUsecase.GetTokenInfo(tokenStr)
	if err != nil {
		return 0, ecode.NoLogin
	}

	return token.AccountID, nil
}

// AuthCookie is used to authorize request by cookie
func (p *Auth) AuthCookie(ctx *mars.Context) (int64, error) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		return 0, ecode.NoLogin
	}

	tokenStr := cookie.Value
	token, err := p.OauthUsecase.GetTokenInfo(tokenStr)
	if err != nil {
		return 0, ecode.NoLogin
	}

	return token.AccountID, nil
}

func (p *Auth) midAuth(ctx *mars.Context, auth authFunc) {
	accountID, err := auth(ctx)
	if err != nil {
		ctx.JSON(nil, ecode.Unauthorized)
		return
	}
	setAccountID(ctx, accountID)
}

// set account id into context
func setAccountID(ctx *mars.Context, id int64) {
	ctx.Set("aid", id)
	if md, ok := metadata.FromContext(ctx); ok {
		md[metadata.Aid] = id
		return
	}
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
