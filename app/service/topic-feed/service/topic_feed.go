package service

import (
	"context"
	"valerian/app/service/topic-feed/model"
)

func (p *Service) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (items []*model.TopicFeed, err error) {
	return p.d.GetTopicFeedPaged(c, p.d.DB(), topicID, limit, offset)
}
