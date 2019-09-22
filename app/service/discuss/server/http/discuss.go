package http

import (
	"strconv"

	"valerian/app/service/discuss/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func getDiscusstionsByTopic(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetTopicDiscussionsPaged(c, id, limit, offset))
}

func getDiscusstionsByAccount(c *mars.Context) {
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

	if id, err = strconv.ParseInt(params.Get("aid"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetUserDiscussionsPaged(c, id, limit, offset))
}

func addDiscussion(c *mars.Context) {
	var aid int64
	var err error
	params := c.Request.Form
	if aid, err = strconv.ParseInt(params.Get("aid"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	arg := new(model.ArgAddDiscuss)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	id, err := srv.AddDiscussion(c, aid, arg)
	c.JSON(strconv.FormatInt(id, 10), err)
}

func updateDiscussion(c *mars.Context) {
	var aid int64
	var err error
	params := c.Request.Form
	if aid, err = strconv.ParseInt(params.Get("aid"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	arg := new(model.ArgUpdateDiscuss)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateDiscussion(c, aid, arg))
}

func delDiscussion(c *mars.Context) {
	var aid int64
	var err error
	params := c.Request.Form
	if aid, err = strconv.ParseInt(params.Get("aid"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	arg := new(model.ArgDelDiscuss)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.DelDiscussion(c, aid, arg))
}
