package http

import (
	"strconv"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 获取话题类目
// @Description 获取话题类目
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Success 200 {array} model.TopicLevel1Catalog "话题类目"
// @Failure 400 "请求验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic/list/catalogs [get]
func topicCatalogs(c *mars.Context) {
	var (
		id  int64
		err error
	)
	params := c.Request.Form
	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetCatalogHierarchy(c, id))
}
