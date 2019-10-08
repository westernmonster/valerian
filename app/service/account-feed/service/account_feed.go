package service

import (
	"context"
	"valerian/app/service/account-feed/model"
)

func (p *Service) GetAccountFeedPaged(c context.Context, topicID int64, limit, offset int) (items []*model.AccountFeed, err error) {
	return p.d.GetAccountFeedPaged(c, p.d.DB(), topicID, limit, offset)
}
