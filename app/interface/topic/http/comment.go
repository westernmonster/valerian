package http

import "valerian/library/net/http/mars"

// @Summary 获取评论列表
// @Description 获取评论列表
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param resource_id query string true "resource_id，话题讨论、文章、补充的ID"
// @Param type query string true "type，目前为 discuss, article, revise"
// @Success 200 {object} model.CommentListResp "评论列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/list [get]
func getComments(c *mars.Context) {
}

// @Summary 新增评论
// @Description 新增评论
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddComment true "请求"
// @Success 200 "成功,返回comment_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/add [post]
func addComment(c *mars.Context) {

}

// @Summary 删除评论
// @Description 删除评论
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDelComment true "请求"
// @Success 200 "成功,返回comment_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/del [post]
func delComment(c *mars.Context) {

}
