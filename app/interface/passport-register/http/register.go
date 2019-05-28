package http

import (
	"net/http"
	"valerian/app/interface/passport-register/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func emailRegister(c *mars.Context) {
	arg := new(model.ArgEmail)
	c.Bind(arg)

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

func mobileRegister(c *mars.Context) {
	arg := new(model.ArgMobile)
	c.Bind(arg)

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
