package http

import "valerian/library/net/http/mars"

// @Summary 获取文章补充列表
// @Description 获取文章补充列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "article_id"
// @Success 200 {object} model.ReviseListResp "补充列表"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/revises [get]
func getRevises(c *mars.Context) {
}

// @Summary 新增文章补充
// @Description 新增文章补充
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgAddRevise true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/revise/add [post]
func addRevise(c *mars.Context) {
}

// @Summary 删除文章补充
// @Description 删除文章补充
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgDelRevise true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/revise/del [post]
func delRevise(c *mars.Context) {

}

// @Summary 更新文章补充
// @Description 更新文章补充
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgUpdateRevise true "请求"
// @Success 200 "成功,返回revise_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/revise/edit [post]
func updateRevise(c *mars.Context) {

}
