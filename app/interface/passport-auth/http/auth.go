package http

import (
	"valerian/app/interface/passport-auth/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func renewToken(c *mars.Context) {
	arg := new(model.ArgRenewToken)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.RenewToken(c, arg))
}
