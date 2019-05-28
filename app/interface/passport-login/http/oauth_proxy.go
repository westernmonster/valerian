package http

import (
	"valerian/app/interface/passport-login/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func emailLogin(c *mars.Context) {
	arg := new(model.ArgEmailLogin)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.EmailLogin(c, arg))
}

func mobileLogin(c *mars.Context) {
	arg := new(model.ArgMobileLogin)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.MobileLogin(c, arg))
}

func digitLogin(c *mars.Context) {
	arg := new(model.ArgDigitLogin)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.DigitLogin(c, arg))
}
