package http

import (
	"valerian/app/service/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func followTopic(c *mars.Context) {
	arg := new(model.ArgTopicFollow)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.Follow(c, arg))
}
