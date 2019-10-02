package service

import (
	"context"
	"valerian/app/service/feed/model"
)

func (p *Service) GetAccountFeedPaged(c context.Context, accountID int64, limit, offset int) (items []*model.AccountFeed, err error) {
	return p.d.GetAccountFeedPaged(c, p.d.DB(), accountID, limit, offset)
}
