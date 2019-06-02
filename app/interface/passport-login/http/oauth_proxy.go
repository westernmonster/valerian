package http

import (
	"net/http"
	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

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

	if arg.Source == model.SourceWeb {
		cookie := createCookie(resp.AccessToken)
		http.SetCookie(c.Writer, cookie)
	}

	c.JSON(resp, nil)

}

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

	if arg.Source == model.SourceWeb {
		cookie := createCookie(resp.AccessToken)
		http.SetCookie(c.Writer, cookie)
	}

	c.JSON(resp, nil)

}

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

	if arg.Source == model.SourceWeb {
		cookie := createCookie(resp.AccessToken)
		http.SetCookie(c.Writer, cookie)
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
