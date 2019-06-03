package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

func (p *Service) getTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	var addCache = true

	if items, err = p.d.TopicVersionCache(c, topicSetID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetTopicVersions(c, node, topicSetID); err != nil {
		return
	}

	if addCache {
		p.d.SetTopicVersionCache(context.TODO(), topicSetID, items)
	}

	return
}
