package http

import "valerian/library/net/http/mars"

func topicTypeList(c *mars.Context) {
	c.JSON(srv.GetTopicTypes(c))
}
