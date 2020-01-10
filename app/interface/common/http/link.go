package http

import (
	"valerian/app/interface/common/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取链接信息
// @Description 获取链接信息
// @Tags editor
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.common.model.ArgLinkInfo true "请求"
// @Success 200 {object}  app.interface.common.model.LinkInfoResp "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /editor/link_info [post]
func linkInfo(c *mars.Context) {
	arg := new(model.ArgLinkInfo)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.LinkInfo(c, arg))
}
