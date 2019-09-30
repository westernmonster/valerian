package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func isFav(c *mars.Context) {
	var (
		accountID, targetID int64
		targetType          string
	)

	if id, err := strconv.ParseInt(c.Request.Form.Get("account_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		accountID = id
	}

	if id, err := strconv.ParseInt(c.Request.Form.Get("target_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		targetID = id
	}

	targetType = c.Request.Form.Get("target_id")

	c.JSON(srv.IsFav(c, accountID, targetID, targetType))
}
