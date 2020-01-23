package http

import (
	"strconv"
	"valerian/app/service/msm/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func codes(c *mars.Context) {
	var (
		err  error
		code *model.ErrCodes
		ver  int64
	)

	verStr := c.Request.Form.Get("ver")
	if ver, err = strconv.ParseInt(verStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if code, err = svr.Codes(c, ver); err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(code, nil)
}

func codesLangs(c *mars.Context) {
	var (
		err  error
		code *model.CodesLangs
		ver  int64
	)

	verStr := c.Request.Form.Get("ver")
	if ver, err = strconv.ParseInt(verStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if code, err = svr.CodesLangs(c, ver); err != nil {
		c.JSON(nil, err)
		return
	}

	c.JSON(code, nil)
}
