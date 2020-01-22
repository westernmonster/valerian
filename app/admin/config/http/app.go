package http

import (
	"strconv"
	"strings"
	"valerian/app/admin/config/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func getAppByName(c *mars.Context) {
	params := c.Request.Form

	env := strings.TrimSpace(params.Get("env"))
	name := strings.TrimSpace(params.Get("name"))

	c.JSON(svr.GetAppByName(c, env, name))
}

func createApp(c *mars.Context) {
	arg := new(model.ArgCreateApp)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, svr.CreateApp(c, arg))
}

func appList(c *mars.Context) {
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

	env := strings.TrimSpace(params.Get("env"))
	if env != "" {
		cond["env"] = env
	}

	zone := strings.TrimSpace(params.Get("zone"))
	if zone != "" {
		cond["zone"] = zone
	}

	name := strings.TrimSpace(params.Get("name"))
	if name != "" {
		cond["name"] = name
	}

	token := strings.TrimSpace(params.Get("token"))
	if token != "" {
		cond["token"] = token
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

	c.JSON(svr.GetAppsPaged(c, cond, int32(page), int32(pageSize)))
}
