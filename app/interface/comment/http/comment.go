package http

import (
	"valerian/library/net/http/mars"
)

// @Summary 评论列表
// @Description 评论列表
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param type query string true "类型"
// @Param resource_id query string true "目标ID"
// @Param limit query integer false "每页大小"
// @Param offset query integer false "offset"
// @Success 200 {object} model.CommentListResp "评论列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /comment/list/comments [get]
func commentList(c *mars.Context) {
}
