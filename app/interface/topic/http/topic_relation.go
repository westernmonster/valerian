package http

import (
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func editTopicRelations(c *mars.Context) {
	arg := new(model.ArgBatchSaveRelatedTopics)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.BulkSaveRelations(c, arg))

}
