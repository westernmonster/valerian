package http

import (
	"valerian/app/interface/valcode/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func emailValcode(c *mars.Context) {
	arg := new(model.ArgEmailValcode)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.EmailValcode(c, arg))
}

func mobileValcode(c *mars.Context) {
	arg := new(model.ArgMobileValcode)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.MobileValcode(c, arg))
}
