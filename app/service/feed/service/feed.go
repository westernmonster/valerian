package service

import (
	"context"
	"valerian/app/service/feed/model"
)

func (p *Service) GetFeedPaged(c context.Context, topicID int64, limit, offset int) (items []*model.Feed, err error) {
	return p.d.GetFeedPaged(c, p.d.DB(), topicID, limit, offset)
}
