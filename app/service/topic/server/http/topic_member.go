package http

import (
	"strconv"
	"valerian/app/service/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func topicMembers(c *mars.Context) {
	var (
		id       int64
		err      error
		page     int
		pageSize int
	)

	params := c.Request.Form

	if page, err = strconv.Atoi(params.Get("page")); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if page <= 0 {
		page = 1
	}

	if pageSize, err = strconv.Atoi(params.Get("page_size")); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if pageSize <= 0 {
		pageSize = 10
	}

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetTopicMembersPaged(c, id, page, pageSize))
}

func editTopicMembers(c *mars.Context) {
	arg := new(model.ArgBatchSavedTopicMember)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.BulkSaveMembers(c, arg))

}

func leave(c *mars.Context) {
	params := c.Request.Form
	id, err := strconv.ParseInt(params.Get("topic_id"), 10, 64)
	if err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.Leave(c, id))

}
