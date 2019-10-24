package http

import (
	"net/http"
	"valerian/app/admin/login/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 用户邮件登录
// @Description 用户邮件登录
// @Tags admin
// @Accept json
// @Produce json
// @Param Source header int true "来源" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgEmailLogin true "用户登录"
// @Success 200 {object} model.LoginResp "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/login/email [post]
func emailLogin(c *mars.Context) {
	arg := new(model.ArgEmailLogin)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	resp, err := srv.EmailLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(resp, nil)

}

// @Summary 用户手机登录
// @Description 用户手机登录
// @Tags admin
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgMobileLogin true "手机登录"
// @Success 200 {object} model.LoginResp "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/login/mobile [post]
func mobileLogin(c *mars.Context) {
	arg := new(model.ArgMobileLogin)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	resp, err := srv.MobileLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(resp, nil)

}

// @Summary 验证码登录
// @Description 验证码登录
// @Tags admin
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDigitLogin true "手机登录"
// @Success 200 {object} model.LoginResp "成功"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /admin/login/digit [post]
func digitLogin(c *mars.Context) {
	arg := new(model.ArgDigitLogin)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	resp, err := srv.DigitLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(resp, nil)
}

func createCookie(token string) (cookie *http.Cookie) {
	age := model.WebCookieExpires
	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = age
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = token
	return
}
