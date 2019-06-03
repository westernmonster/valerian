package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func topicCatalogs(c *mars.Context) {
	var (
		id  int64
		err error
	)
	params := c.Request.Form
	if id, err = strconv.ParseInt(params.Get("id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetCatalogHierarchy(c, id))
}
