package http

import (
	"valerian/app/interface/passport-auth/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func renewToken(c *mars.Context) {
	arg := new(model.ArgRenewToken)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
	}

	c.JSON(srv.RenewToken(c, arg))
}
