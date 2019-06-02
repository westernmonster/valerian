package http

import (
	"valerian/app/interface/valcode/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func emailValcode(c *mars.Context) {
	arg := new(model.ArgEmailValcode)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.EmailValcode(c, arg))
}

func mobileValcode(c *mars.Context) {
	arg := new(model.ArgMobileValcode)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.MobileValcode(c, arg))
}
