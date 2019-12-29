package http

import (
	"fmt"
	"strconv"
	"valerian/app/interface/discuss/model"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

// @Summary 批量更新文件
// @Description 批量更新文件
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.discuss.model.ArgSaveDiscussionFiles true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/files [post]
func editDiscussionFiles(c *mars.Context) {
	arg := new(model.ArgSaveDiscussionFiles)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		log.For(c).Error(fmt.Sprintf("validation error(%+v)", e))
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveDiscussionFiles(c, arg))
}

// @Summary 获取文件列表
// @Description 获取文件列表
// @Tags discussion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param discussion_id query string true "文章ID"
// @Success 200 {array}  app.interface.discuss.model.DiscussionFileResp "文件列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /discussion/list/files [get]
func discussionFiles(c *mars.Context) {
	var (
		id  int64
		err error
	)

	params := c.Request.Form

	if id, err = strconv.ParseInt(params.Get("discussion_id"), 10, 64); err != nil {
		log.For(c).Error(fmt.Sprintf("req error(%+v)", err))
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetDiscussionFiles(c, id))
}
