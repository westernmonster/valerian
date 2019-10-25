package http

import (
	"net/http"
	"valerian/app/interface/passport-register/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 用户邮件注册
// @Description  用户邮件注册
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.passport-register.model.ArgEmail true "请求"
// @Success 200 {object}  app.interface.passport-register.model.LoginResp "成功"
// @Failure 626 "用户不存在"
// @Failure 116 "Client 不存在"
// @Failure 2 "验证码已失效"
// @Failure 3 "验证码错误"
// @Failure 4 "用户已经存在"
// @Failure 5 "未找到地址信息"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/registration/email [post]
func emailRegister(c *mars.Context) {
	arg := new(model.ArgEmail)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	resp, err := srv.EmailRegister(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	if arg.Source == model.SourceWeb {
		cookie, err := createCookie(resp.AccessToken)
		if err != nil {
			c.JSON(nil, err)
			return
		}
		http.SetCookie(c.Writer, cookie)
	}

	c.JSON(resp, nil)
}

// @Summary 用户手机注册
// @Description  用户手机注册
// @Tags auth
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.passport-register.model.ArgMobile true "请求"
// @Success 200 {object}  app.interface.passport-register.model.LoginResp "成功"
// @Failure 626 "用户不存在"
// @Failure 116 "Client 不存在"
// @Failure 2 "验证码已失效"
// @Failure 3 "验证码错误"
// @Failure 4 "用户已经存在"
// @Failure 5 "未找到地址信息"
// @Failure 400 "验证失败"
// @Failure 500 "服务器端错误"
// @Router /oauth/registration/mobile [post]
func mobileRegister(c *mars.Context) {
	arg := new(model.ArgMobile)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	resp, err := srv.MobileRegister(c, arg)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	if arg.Source == model.SourceWeb {
		cookie, err := createCookie(resp.AccessToken)
		if err != nil {
			c.JSON(nil, err)
			return
		}
		http.SetCookie(c.Writer, cookie)
	}

	c.JSON(resp, nil)
}

func createCookie(token string) (cookie *http.Cookie, err error) {
	age := model.WebCookieExpires
	cookie = new(http.Cookie)
	cookie.Name = "token"
	cookie.MaxAge = age
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = token
	return
}
