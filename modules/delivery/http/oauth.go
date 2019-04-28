package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/helper"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type AuthCtrl struct {
	infrastructure.BaseCtrl

	OauthUsecase interface {
		GetByID(bizCtx *biz.BizContext, userID int64) (item *repo.Account, err error)
		EmailLogin(ctx *biz.BizContext, req *models.EmailLoginReq, ip string) (loginResult *models.LoginResult, err error)
		MobileLogin(ctx *biz.BizContext, req *models.MobileLoginReq, ip string) (loginResult *models.LoginResult, err error)
		DigitLogin(ctx *biz.BizContext, req *models.DigitLoginReq, ip string) (loginResult *models.LoginResult, err error)
		EmailRegister(bizCtx *biz.BizContext, req *models.EmailRegisterReq, ip string) (accountID int64, err error)
		MobileRegister(bizCtx *biz.BizContext, req *models.MobileRegisterReq, ip string) (accountID int64, err error)
		ForgetPassword(bizCtx *biz.BizContext, req *models.ForgetPasswordReq) (sessionID int64, err error)
		ResetPassword(bizCtx *biz.BizContext, req *models.ResetPasswordReq) (err error)
	}

	AccountUsecase interface {
		GetProfileByID(accountID int64) (profile *models.ProfileResp, err error)
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
func (p *AuthCtrl) EmailLogin(ctx *gin.Context) {
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
	result, err := p.OauthUsecase.EmailLogin(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(result.AccountID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, result)

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
func (p *AuthCtrl) MobileLogin(ctx *gin.Context) {
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
	result, err := p.OauthUsecase.MobileLogin(bizCtx, req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(result.AccountID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, result)

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
func (p *AuthCtrl) DigitLogin(ctx *gin.Context) {
	req := new(models.DigitLoginReq)

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
	result, err := p.OauthUsecase.DigitLogin(bizCtx, req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(result.AccountID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result.Profile = profile

	cookie, err := p.createCookie(result.AccessToken)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, result)

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
func (p *AuthCtrl) EmailRegister(ctx *gin.Context) {
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
	_, err := p.OauthUsecase.EmailRegister(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result, err := p.OauthUsecase.EmailLogin(p.GetBizContext(ctx), &models.EmailLoginReq{
		Source:   req.Source,
		Email:    req.Email,
		Password: req.Password,
		ClientID: "532c28d5412dd75bf975fb951c740a30",
	}, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(result.AccountID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result.Profile = profile

	p.SuccessResp(ctx, result)

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
func (p *AuthCtrl) MobileRegister(ctx *gin.Context) {
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
	_, err := p.OauthUsecase.MobileRegister(p.GetBizContext(ctx), req, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result, err := p.OauthUsecase.MobileLogin(p.GetBizContext(ctx), &models.MobileLoginReq{
		Source:   req.Source,
		Mobile:   req.Mobile,
		Password: req.Password,
		ClientID: "532c28d5412dd75bf975fb951c740a30",
	}, ip)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	profile, err := p.AccountUsecase.GetProfileByID(result.AccountID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	result.Profile = profile

	p.SuccessResp(ctx, result)

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
func (p *AuthCtrl) ForgetPassword(ctx *gin.Context) {
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

	sessionID, err := p.OauthUsecase.ForgetPassword(p.GetBizContext(ctx), req)
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
// @Router /oauth/password/reset/confirm [put]
func (p *AuthCtrl) ResetPassword(ctx *gin.Context) {
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

	err := p.OauthUsecase.ResetPassword(p.GetBizContext(ctx), req)
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
func (p *AuthCtrl) Logout(ctx *gin.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = 1
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = ""
	http.SetCookie(ctx.Writer, cookie)
	p.SuccessResp(ctx, nil)
	return
}
