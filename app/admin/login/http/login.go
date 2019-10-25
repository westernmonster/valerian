package http

import (
	"net/http"
	"net/url"
	"time"
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
// @Param req body  app.admin.login.model.ArgEmailLogin true "用户登录"
// @Success 200 {object}  app.admin.login.model.LoginResp "成功"
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

	resp, sid, err := srv.EmailLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	setHTTPCookie(c, cfg.Session.CookieName, sid)
	c.JSON(resp, nil)

}

// @Summary 用户手机登录
// @Description 用户手机登录
// @Tags admin
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.login.model.ArgMobileLogin true "手机登录"
// @Success 200 {object}  app.admin.login.model.LoginResp "成功"
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

	resp, sid, err := srv.MobileLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	setHTTPCookie(c, cfg.Session.CookieName, sid)
	c.JSON(resp, nil)

}

// @Summary 验证码登录
// @Description 验证码登录
// @Tags admin
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.admin.login.model.ArgDigitLogin true "手机登录"
// @Success 200 {object}  app.admin.login.model.LoginResp "成功"
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

	resp, sid, err := srv.DigitLogin(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	setHTTPCookie(c, cfg.Session.CookieName, sid)
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

func setHTTPCookie(ctx *mars.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     "/",
		HttpOnly: true,
		Domain:   cfg.Session.Domain,
	}
	cookie.MaxAge = cfg.Session.CookieLifeTime
	cookie.Expires = time.Now().Add(time.Duration(cfg.Session.CookieLifeTime) * time.Second)
	http.SetCookie(ctx.Writer, cookie)
}
