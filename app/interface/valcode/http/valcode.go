package http

import (
	"valerian/app/interface/valcode/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 请求邮件验证码
// @Description 请求邮件验证码
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgEmailValcode true "请求"
// @Success 200 "成功"
// @Failure 1  "60秒下发一次验证码"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /valcode/email [post]
func emailValcode(c *mars.Context) {
	arg := new(model.ArgEmailValcode)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.EmailValcode(c, arg))
}

// @Summary 请求短信验证码
// @Description 请求短信验证码
// @Tags common
// @Accept json
// @Produce json
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgMobileValcode true "请求"
// @Success 200 "成功"
// @Failure 1  "60秒下发一次验证码"
// @Failure 400 "验证请求失败"
// @Failure 500 "服务器端错误"
// @Router /valcode/mobile [post]
func mobileValcode(c *mars.Context) {
	arg := new(model.ArgMobileValcode)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.MobileValcode(c, arg))
}
