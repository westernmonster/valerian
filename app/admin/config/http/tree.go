package http

import (
	"strconv"
	"strings"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func treeList(c *mars.Context) {
	var (
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

	cond := make(map[string]interface{})

	name := strings.TrimSpace(params.Get("name"))
	if name != "" {
		cond["name"] = name
	}

	platformID := strings.TrimSpace(params.Get("platform_id"))
	if platformID != "" {
		if val, err := strconv.Atoi(platformID); err != nil {
			c.JSON(nil, ecode.RequestErr)
			return
		} else {
			cond["platform_id"] = val
		}
	}

	treeId := strings.TrimSpace(params.Get("tree_id"))
	if treeId != "" {
		if val, err := strconv.Atoi(treeId); err != nil {
			c.JSON(nil, ecode.RequestErr)
			return
		} else {
			cond["tree_id"] = val
		}
	}

	c.JSON(svr.GetTreesPaged(c, cond, int32(page), int32(pageSize)))
}
