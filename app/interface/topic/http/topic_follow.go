package http

import (
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 关注话题
// @Description 关注话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgTopicFollow true "请求"
// @Success 200  "返回关注状态"
// @Failure 18 "话题不存在"
// @Failure 626 "用户不存在"
// @Failure 37 "需要身份认证"
// @Failure 38 "需要工作认证"
// @Failure 39 "必须管理员添加"
// @Failure 40 "需要购买"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/follow [post]
func followTopic(c *mars.Context) {
	arg := new(model.ArgTopicFollow)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.Follow(c, arg))
}
