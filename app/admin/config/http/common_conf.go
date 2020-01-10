package http

import (
	"fmt"
	"strconv"
	"strings"
	"valerian/app/admin/config/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func updateCommonConfig(c *mars.Context) {
	arg := new(model.ArgUpdateCommonConfig)
	if e := c.Bind(arg); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if e := arg.Validate(); e != nil {
		fmt.Println(e)
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, svr.UpdatCommonConf(c, arg))
}

func createCommonConfig(c *mars.Context) {
	arg := new(model.ArgCreateCommonConfig)
	if e := c.Bind(arg); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if e := arg.Validate(); e != nil {
		fmt.Println(e)
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, svr.CreateCommonConf(c, arg))
}

func commonConfigList(c *mars.Context) {
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

	operator := strings.TrimSpace(params.Get("operator"))
	if operator != "" {
		cond["operator"] = operator
	}

	state := strings.TrimSpace(params.Get("state"))
	if state != "" {
		if val, err := strconv.Atoi(state); err != nil {
			c.JSON(nil, ecode.RequestErr)
			return
		} else {
			cond["state"] = val
		}
	}

	c.JSON(svr.GetCommonConfsPaged(c, cond, int32(page), int32(pageSize)))
}
