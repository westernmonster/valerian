package http

import (
	"strconv"

	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func getDiscussion(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetDiscussion(c, id))
	}
}
