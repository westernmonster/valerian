package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/db"
	"git.flywk.com/flywiki/api/infrastructure/ecode"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
	"github.com/gin-gonic/gin"
)

type authFunc func(*gin.Context) (int64, error)

type Auth struct {
	OauthUsecase interface {
		GetByID(bizCtx *biz.BizContext, userID int64) (item *repo.Account, err error)
		GetTokenInfo(tokenStr string) (claims *infrastructure.TokenClaims, err error)
	}
}

func New() *Auth {
	db, node, err := db.InitDatabase()
	if err != nil {
		panic(err)
	}
	auth := &Auth{
		OauthUsecase: &usecase.OauthUsecase{
			Node:                       node,
			DB:                         db,
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
func (p *Auth) User(ctx *gin.Context) {
	req := ctx.Request
	cookie, _ := req.Cookie("token")
	if cookie != nil {
		p.UserWeb(ctx)
	} else {
		p.UserMobile(ctx)
	}
}

// UserWeb is used to mark path as web access required.
func (p *Auth) UserWeb(ctx *gin.Context) {
	p.midAuth(ctx, p.AuthCookie)
}

// UserMobile is used to mark path as mobile access required.
func (p *Auth) UserMobile(ctx *gin.Context) {
	p.midAuth(ctx, p.AuthToken)
}

// AuthToken is used to authorize request by token
func (p *Auth) AuthToken(ctx *gin.Context) (accountID int64, err error) {
	req := ctx.Request

	tokenStr := ""
	if v, ok := getBearer(req.Header["Authorization"]); ok {
		tokenStr = v
	}

	if tokenStr == "" {
		return 0, ecode.NoLogin
	}

	token, err := p.OauthUsecase.GetTokenInfo(tokenStr)
	fmt.Println(token)
	if err != nil {
		return 0, ecode.NoLogin
	}

	return token.AccountID, nil
}

// AuthCookie is used to authorize request by cookie
func (p *Auth) AuthCookie(ctx *gin.Context) (int64, error) {

	tokenStr := ""
	cookieValue, err := ctx.Cookie("token")
	if err == nil {
		tokenStr = cookieValue
	} else {
		return 0, ecode.NoLogin
	}

	fmt.Println(tokenStr)

	token, err := p.OauthUsecase.GetTokenInfo(tokenStr)
	if err != nil {
		fmt.Println(err)
		return 0, ecode.NoLogin
	}

	return token.AccountID, nil
}

func (p *Auth) midAuth(ctx *gin.Context, auth authFunc) {
	accountID, err := auth(ctx)
	if err != nil {
		bcode := ecode.Cause(err)
		ctx.Header(models.HeaderStatusCode, bcode.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: bcode.Error(),
		})

		return
	}
	setAccountID(ctx, accountID)
}

// set account id into context
func setAccountID(ctx *gin.Context, id int64) {
	ctx.Set("AccountID", id)
	// if md, ok := metadata.FromContext(ctx); ok {
	// 	md[metadata.Mid] = mid
	// 	return
	// }
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
