package http

import (
	"context"
	"net/http"
	"strconv"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/infrastructure/helper"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
	"valerian/models"
	"valerian/modules/repo"
)

type AuthCtrl struct {
	infrastructure.BaseCtrl

	OauthUsecase interface {
		GetByID(c context.Context, ctx *biz.BizContext, userID int64) (item *repo.Account, err error)
		EmailLogin(c context.Context, ctx *biz.BizContext, req *models.EmailLoginReq, ip string) (loginResult *models.LoginResult, err error)
		MobileLogin(c context.Context, ctx *biz.BizContext, req *models.MobileLoginReq, ip string) (loginResult *models.LoginResult, err error)
		DigitLogin(c context.Context, ctx *biz.BizContext, req *models.DigitLoginReq, ip string) (loginResult *models.LoginResult, err error)
		EmailRegister(c context.Context, ctx *biz.BizContext, req *models.EmailRegisterReq, ip string) (accountID int64, err error)
		MobileRegister(c context.Context, ctx *biz.BizContext, req *models.MobileRegisterReq, ip string) (accountID int64, err error)
		ForgetPassword(c context.Context, ctx *biz.BizContext, req *models.ForgetPasswordReq) (sessionID int64, err error)
		ResetPassword(c context.Context, ctx *biz.BizContext, req *models.ResetPasswordReq) (err error)
	}

	AccountUsecase interface {
		GetProfileByID(c context.Context, accountID int64) (profile *models.ProfileResp, err error)
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
// @Router /oauth/login/email [post]
func (p *AuthCtrl) EmailLogin(ctx *mars.Context) {
	req := new(models.EmailLoginReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ip := "0.0.0.0"
	result, err := p.OauthUsecase.EmailLogin(ctx.Context, p.GetBizContext(ctx), req, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(ctx.Context, result.AccountID)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(result, err)
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
// @Router /oauth/login/mobile [post]
func (p *AuthCtrl) MobileLogin(ctx *mars.Context) {
	req := new(models.MobileLoginReq)
	ctx.Bind(req)
	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ip := metadata.String(ctx, metadata.RemoteIP)
	bizCtx := p.GetBizContext(ctx)
	result, err := p.OauthUsecase.MobileLogin(ctx.Context, bizCtx, req, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(ctx.Context, result.AccountID)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(result, err)

	return
}

// DigitLogin 验证码登录
// @Summary 验证码登录
// @Description 验证码登录
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.DigitLoginReq true "手机登录"
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/login/digit [post]
func (p *AuthCtrl) DigitLogin(ctx *mars.Context) {
	req := new(models.DigitLoginReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ip := metadata.String(ctx, metadata.RemoteIP)
	bizCtx := p.GetBizContext(ctx)
	result, err := p.OauthUsecase.DigitLogin(ctx.Context, bizCtx, req, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(ctx.Context, result.AccountID)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(result, err)

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
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/registration/email [post]
func (p *AuthCtrl) EmailRegister(ctx *mars.Context) {
	req := new(models.EmailRegisterReq)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ip := metadata.String(ctx, metadata.RemoteIP)
	_, err := p.OauthUsecase.EmailRegister(ctx.Context, p.GetBizContext(ctx), req, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result, err := p.OauthUsecase.EmailLogin(ctx.Context, p.GetBizContext(ctx), &models.EmailLoginReq{
		Source:   req.Source,
		Email:    req.Email,
		Password: req.Password,
		ClientID: "532c28d5412dd75bf975fb951c740a30",
	}, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(ctx.Context, result.AccountID)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(result, err)

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
// @Success 200 {object} models.LoginResult "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/registration/mobile [post]
func (p *AuthCtrl) MobileRegister(ctx *mars.Context) {
	req := new(models.MobileRegisterReq)

	ctx.Bind(req)
	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	ip := metadata.String(ctx, metadata.RemoteIP)
	_, err := p.OauthUsecase.MobileRegister(ctx.Context, p.GetBizContext(ctx), req, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result, err := p.OauthUsecase.MobileLogin(ctx.Context, p.GetBizContext(ctx), &models.MobileLoginReq{
		Source:   req.Source,
		Mobile:   req.Mobile,
		Password: req.Password,
		Prefix:   req.Prefix,
		ClientID: "532c28d5412dd75bf975fb951c740a30",
	}, ip)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(ctx.Context, result.AccountID)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(result, err)

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
// @Router /oauth/password/reset [put]
func (p *AuthCtrl) ForgetPassword(ctx *mars.Context) {
	req := new(models.ForgetPasswordReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	sessionID, err := p.OauthUsecase.ForgetPassword(ctx.Context, p.GetBizContext(ctx), req)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	strID := strconv.FormatInt(sessionID, 10)
	b64Str := helper.Base64Encode(strID)

	ctx.JSON(b64Str, err)

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
// @Router /oauth/password/reset/confirm [put]
func (p *AuthCtrl) ResetPassword(ctx *mars.Context) {
	req := new(models.ResetPasswordReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	err := p.OauthUsecase.ResetPassword(ctx.Context, p.GetBizContext(ctx), req)
	ctx.JSON(nil, err)

	return
}

func (p *AuthCtrl) createCookie(token string) (cookie *http.Cookie, err error) {
	age := models.ExpiresIn
	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = age
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = token
	return
}

// Logout 退出登录
// @Summary 退出登录
// @Description 退出登录
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "来源" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/logout [post]
func (p *AuthCtrl) Logout(ctx *mars.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = 1
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = ""
	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(nil, nil)
	return
}
