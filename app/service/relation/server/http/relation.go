package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func following(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
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

	if id, err = strconv.ParseInt(params.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.FollowingPaged(c, id, limit, offset))
}

func fans(c *mars.Context) {
	var (
		id     int64
		err    error
		offset int
		limit  int
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

	if id, err = strconv.ParseInt(params.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.FansPaged(c, id, limit, offset))
}
