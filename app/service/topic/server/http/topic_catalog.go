package http

import (
	"strconv"
	"valerian/app/service/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func topicCatalogTaxonomies(c *mars.Context) {
	var (
		id  int64
		err error
	)
	params := c.Request.Form
	if id, err = strconv.ParseInt(params.Get("topic_id"), 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetCatalogTaxonomiesHierarchy(c, id))
}

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

	c.JSON(srv.GetCatalogsHierarchy(c, id))
}

func editTopicCatalogs(c *mars.Context) {
	arg := new(model.ArgSaveTopicCatalog)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.SaveCatalogs(c, arg))
}
