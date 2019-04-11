package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/helper"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type AccountCtrl struct {
	infrastructure.BaseCtrl

	AccountUsecase interface {
		GetByID(bizCtx *biz.BizContext, userID int64) (item *repo.Account, err error)
		EmailLogin(bizCtx *biz.BizContext, req *models.EmailLoginReq, ip string) (item *repo.Account, err error)
		MobileLogin(bizCtx *biz.BizContext, req *models.MobileLoginReq, ip string) (item *repo.Account, err error)
		EmailRegister(bizCtx *biz.BizContext, req *models.EmailRegisterReq, ip string) (err error)
		MobileRegister(bizCtx *biz.BizContext, req *models.MobileRegisterReq, ip string) (err error)
		ForgetPassword(bizCtx *biz.BizContext, req *models.ForgetPasswordReq) (sessionID int64, err error)
		ResetPassword(bizCtx *biz.BizContext, req *models.ResetPasswordReq) (err error)
		ChangePassword(bizCtx *biz.BizContext, req *models.ChangePasswordReq) (err error)
		UpdateProfile(bizCtx *biz.BizContext, req *models.UpdateProfileReq) (err error)
		GetProfile(bizCtx *biz.BizContext) (profile *models.ProfileResp, err error)
	}
}

func GetBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) == 2 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

// Auth godoc
func (p *AccountCtrl) Auth(ctx *gin.Context) {
	tokenStr := ""

	if v, ok := GetBearer(ctx.Request.Header["Authorization"]); ok {
		tokenStr = v
	} else {
		cookieValue, err := ctx.Cookie("token")
		if err == nil {
			tokenStr = cookieValue
		}
	}

	if tokenStr == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "jwt token is not set",
		})
		return
	}

	claims := new(infrastructure.TokenClaims)

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(infrastructure.TokenSigningKey), nil
	})

	if err != nil {
		logrus.Error(fmt.Sprintf("Error parsing token: %v", err))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "服务器端发生错误",
		})
		return
	}

	if claims, ok := token.Claims.(*infrastructure.TokenClaims); ok && token.Valid {
		account, err := p.AccountUsecase.GetByID(nil, claims.AccountID)
		if err != nil {
			logrus.Error(fmt.Sprintf("Error on fetch user info: %v", err))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"success": false,
				"message": "服务器端发生错误",
			})
			return
		}

		ctx.Set("AccountID", account.ID)
		ctx.Set("Role", claims.Role)
		ctx.Next()
	} else {
		logrus.Error(fmt.Sprintf("Error parsing token: %v", err))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"success": false,
			"message": "服务器端发生错误",
		})
		return
	}
}

// EmailLogin 用户邮件登录
// @Summary 用户邮件登录
// @Description 用户邮件登录
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "来源" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.EmailLoginReq true "用户登录"
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /auth/login/email [post]
func (p *AccountCtrl) EmailLogin(ctx *gin.Context) {
	req := new(models.EmailLoginReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	ip := ctx.ClientIP()
	user, err := p.AccountUsecase.EmailLogin(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	cookie, err := p.createCookie(user)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, models.LoginResult{
		Token: cookie.Value,
		Role:  "user",
	})

	return
}

// EmailLogin 用户手机登录
// @Summary 用户手机登录
// @Description 用户手机登录
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.MobileLoginReq true "手机登录"
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /auth/login/mobile [post]
func (p *AccountCtrl) MobileLogin(ctx *gin.Context) {
	req := new(models.MobileLoginReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	ip := ctx.ClientIP()
	bizCtx := p.GetBizContext(ctx)
	user, err := p.AccountUsecase.MobileLogin(bizCtx, req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	cookie, err := p.createCookie(user)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, models.LoginResult{
		Token: cookie.Value,
		Role:  "user",
	})

	return
}

// @Summary 用户邮件注册
// @Description  用户邮件注册
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.EmailRegisterReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /auth/registration/email [post]
func (p *AccountCtrl) EmailRegister(ctx *gin.Context) {
	req := new(models.EmailRegisterReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		fmt.Println(e)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	ip := ctx.ClientIP()
	err := p.AccountUsecase.EmailRegister(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// @Summary 用户手机注册
// @Description  用户手机注册
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.MobileRegisterReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /auth/registration/mobile [post]
func (p *AccountCtrl) MobileRegister(ctx *gin.Context) {
	req := new(models.MobileRegisterReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	ip := ctx.ClientIP()
	err := p.AccountUsecase.MobileRegister(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// @Summary 忘记密码
// @Description  忘记密码，此为重设密码第一步，提交用户标识（手机号、邮箱），和用户输入的验证码进行验证，并返回一个 Session ID
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.ForgetPasswordReq true "请求"
// @Success 200 {object} infrastructure.RespCommon "成功"
// @Failure 400 {object} infrastructure.RespCommon "验证失败"
// @Failure 500 {object} infrastructure.RespCommon "服务器端错误"
// @Router /auth/password/reset [put]
func (p *AccountCtrl) ForgetPassword(ctx *gin.Context) {
	req := new(models.ForgetPasswordReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	sessionID, err := p.AccountUsecase.ForgetPassword(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	strID := strconv.FormatInt(sessionID, 10)
	b64Str := helper.Base64Encode(strID)

	p.SuccessResp(ctx, b64Str)

	return
}

// @Summary 重设密码
// @Description 重设密码第二步，传入新密码和Session ID，如果返回的Code值为307，则表示Session已经失效，前端可以根据这个值做对应的处理
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.ResetPasswordReq true "请求"
// @Success 200 {object} infrastructure.RespCommon "成功"
// @Failure 400 {object} infrastructure.RespCommon "验证失败"
// @Failure 500 {object} infrastructure.RespCommon "服务器端错误"
// @Router /auth/password/reset/confirm [put]
func (p *AccountCtrl) ResetPassword(ctx *gin.Context) {
	req := new(models.ResetPasswordReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	err := p.AccountUsecase.ResetPassword(p.GetBizContext(ctx), req)
	if err != nil {
		switch err.(type) {
		case *berr.BizError:
			ctx.AbortWithStatusJSON(http.StatusOK, infrastructure.RespCommon{
				Success: false,
				Code:    http.StatusTemporaryRedirect,
				Message: "Session 错误或失效",
			})
			return
		default:
			p.HandleError(ctx, err)
			return

		}
	}

	p.SuccessResp(ctx, nil)

	return
}

func (p *AccountCtrl) createCookie(user *repo.Account) (cookie *http.Cookie, err error) {
	age := 90 * 24 * 60 * 60 // 90 days, present by second
	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = age
	cookie.HttpOnly = true

	claims := infrastructure.TokenClaims{
		user.ID,
		"role",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			Issuer:    "web",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(infrastructure.TokenSigningKey))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	cookie.Value = tokenString

	return

}

// @Summary 更改密码
// @Description 更改密码，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateProfileReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/password [put]
func (p *AccountCtrl) ChangePassword(ctx *gin.Context) {

	req := new(models.ChangePasswordReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	err := p.AccountUsecase.ChangePassword(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// UpdateProfile 更改用户资料
// @Summary 更改用户资料
// @Description 更改用户资料，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateProfileReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me [put]
func (p *AccountCtrl) UpdateProfile(ctx *gin.Context) {

	req := new(models.UpdateProfileReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	err := p.AccountUsecase.UpdateProfile(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// GetProfile 获取当前用户资料
// @Summary 获取当前用户资料
// @Description 获取当前用户资料，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object} models.ProfileResp "用户资料"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me [get]
func (p *AccountCtrl) GetProfile(ctx *gin.Context) {

	item, err := p.AccountUsecase.GetProfile(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, item)

	return
}
