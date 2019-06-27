package http

import (
	"strconv"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取文章版本列表
// @Description 获取文章版本列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_set_id query string true "文章集合ID"
// @Success 200 {array} model.ArticleVersionResp "文章版本"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/versions [get]
func articleVersions(c *mars.Context) {
	var (
		articleSetID int64
		err          error
	)

	params := c.Request.Form
	if articleSetID, err = strconv.ParseInt(params.Get("article_set_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleVersions(c, articleSetID))
}

// @Summary 新建文章版本
// @Description 新建文章版本
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgNewArticleVersion true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/versions [post]
func addArticleVersion(c *mars.Context) {

	arg := new(model.ArgNewArticleVersion)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	id, err := srv.AddArticleVersion(c, arg)
	c.JSON(strconv.FormatInt(id, 10), err)
}

// @Summary 合并文章版本
// @Description 合并文章版本，需要当前用户为合并两个文章集中，是所有文章的成员并且是管理员或主理人
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body model.ArgMergeArticleVersion true "请求"
// @Failure 20 "获取用户ID失败，一般是因为未登录造成"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/versions/merge [post]
func mergeArticleVersion(c *mars.Context) {
}

// @Summary 保存文章版本（排序，重命名）
// @Description 保存文章版本（排序，重命名）
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body ArgSaveArticleVersions true "请求"
// @Failure 20 "获取用户ID失败，一般是因为未登录造成"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/versions/save [post]
func saveArticleVersion(c *mars.Context) {
	arg := new(model.ArgSaveArticleVersions)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveArticleVersions(c, arg))
}
