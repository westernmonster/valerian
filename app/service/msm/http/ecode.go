package http

import (
	"valerian/app/service/msm/model"
	"valerian/library/net/http/mars"
)

func codes(c *mars.Context) {
	var (
		err   error
		code  *model.Codes
		param = new(struct {
			Ver int64 `form:"ver"`
		})
	)
	if err = c.Bind(param); err != nil {
		return
	}
	if code, err = svr.Codes(c, param.Ver); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(code, nil)
}

func codesLangs(c *mars.Context) {
	var (
		err   error
		code  *model.CodesLangs
		param = new(struct {
			Ver int64 `form:"ver"`
		})
	)
	if err = c.Bind(param); err != nil {
		return
	}
	if code, err = svr.CodesLangs(c, param.Ver); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(code, nil)
}
