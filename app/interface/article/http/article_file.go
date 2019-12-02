package http

import (
	"fmt"
	"strconv"
	"valerian/app/interface/article/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 批量更新文件
// @Description 批量更新文件
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.article.model.ArgSaveArticleFiles true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/files [post]
func editArticleFiles(c *mars.Context) {
	arg := new(model.ArgSaveArticleFiles)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		fmt.Println(e)
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveArticleFiles(c, arg))
}

// @Summary 获取文件列表
// @Description 获取文件列表
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param article_id query string true "文章ID"
// @Success 200 {array}  app.interface.article.model.ArticleFileResp "文件列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/list/files [get]
func articleFiles(c *mars.Context) {
	var (
		articleID int64
		err       error
	)

	params := c.Request.Form
	if articleID, err = strconv.ParseInt(params.Get("article_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleFiles(c, articleID))
}

// @Summary 获取文件
// @Description 获取文件
// @Tags article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param file_id query string true "文件ID"
// @Success 200 {object}  app.interface.article.model.ArticleFileResp "文件列表"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /article/file/get [get]
func getArticleFile(c *mars.Context) {
	var (
		fileID int64
		err    error
	)

	params := c.Request.Form
	if fileID, err = strconv.ParseInt(params.Get("file_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetArticleFile(c, fileID))
}
