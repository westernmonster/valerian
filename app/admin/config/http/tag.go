package http

import (
	"valerian/app/admin/config/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func createTag(c *mars.Context) {
	arg := new(model.ArgCreateTag)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, svr.CreateTag(c, arg))
}
