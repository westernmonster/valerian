package http

import (
	"fmt"
	"strconv"

	"valerian/app/interface/discuss/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取讨论分类列表
// @Tags discussion
// @Description 获取讨论分类列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Success 200 {array}  app.interface.discuss.model.DiscussCategoryResp "讨论分类"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/list/categories [get]
func discussCategories(c *mars.Context) {
	var (
		id  int64
		err error
	)

	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetDiscussCategories(c, id))
}

// @Summary 批量更新讨论分类
// @Description 批量更新讨论分类
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.discuss.model.ArgSaveDiscussCategories true "请求"
// @Success 200 "成功"
// @Failure 18 "话题不存在"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/categories [post]
func editDiscussCategories(c *mars.Context) {
	arg := new(model.ArgSaveDiscussCategories)
	if e := c.Bind(arg); e != nil {
		return
	}

	fmt.Println(arg)
	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveDiscussCategories(c, arg))

}
