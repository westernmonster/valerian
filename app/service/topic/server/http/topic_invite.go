package http

import (
	"strconv"

	"valerian/app/service/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func inviteFans(c *mars.Context) {
	arg := new(model.ArgTopicInvite)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Invite(c, arg))
}

func memberFansList(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
		query  string
	)

	params := c.Request.Form

	if offset, err = strconv.Atoi(params.Get("offset")); err != nil {
		offset = 0
	} else if offset < 0 {
		offset = 0
	}

	if limit, err = strconv.Atoi(params.Get("limit")); err != nil {
		limit = 10
	} else if limit < 0 {
		limit = 10
	}

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	query = params.Get("query")

	c.JSON(srv.GetMemberFansList(c, id, query, limit, offset))
}
