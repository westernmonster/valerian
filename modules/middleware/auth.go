package middleware

import (
	"go-common/library/ecode"
	"net/http"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func New() *Auth {
	auth := &Auth{}
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
	key := req.Form.Get("access_key")
	if key == "" {
		return 0, ecode.NoLogin
	}
	buvid := req.Header.Get("buvid")

	// reply, err := a.GetTokenInfo(ctx, &idtv1.GetTokenInfoReq{Token: key, Buvid: buvid})
	// if err != nil {
	// 	return 0, err
	// }
	// if !reply.IsLogin {
	// 	return 0, ecode.NoLogin
	// }

	// return reply.Mid, nil
	return
}

// AuthCookie is used to authorize request by cookie
func (p *Auth) AuthCookie(ctx *gin.Context) (int64, error) {
	req := ctx.Request
	ssDaCk, _ := req.Cookie("SESSDATA")
	if ssDaCk == nil {
		return 0, ecode.NoLogin
	}

	cookie := req.Header.Get("Cookie")
	reply, err := a.GetCookieInfo(ctx, &idtv1.GetCookieInfoReq{Cookie: cookie})
	if err != nil {
		return 0, err
	}
	if !reply.IsLogin {
		return 0, ecode.NoLogin
	}

	// check csrf
	clientCsrf := req.FormValue("csrf")
	if a.conf != nil && !a.conf.DisableCSRF && req.Method == "POST" {
		if clientCsrf != reply.Csrf {
			return 0, ecode.CsrfNotMatchErr
		}
	}

	return reply.Mid, nil
}

func (p *Auth) midAuth(ctx *gin.Context, auth authFunc) {
	accountID, err := auth(ctx)
	if err != nil {
		code := http.StatusOK
		c.Error = err
		bcode := ecode.Cause(err)
		ctx.Header(models.HeaderStatusCode, bcode.Code())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
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
