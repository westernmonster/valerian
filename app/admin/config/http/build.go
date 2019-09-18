package http

import (
	"valerian/app/admin/config/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func createBuild(c *mars.Context) {
	arg := new(model.ArgCreateBuild)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(svr.CreateBuild(c, arg))
}
