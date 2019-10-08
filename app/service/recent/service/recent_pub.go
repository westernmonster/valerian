package service

import (
	"context"
	"valerian/app/service/recent/model"
)

func (p *Service) GetUserRecentPubsPaged(c context.Context, aid int64, limit, offset int) (items []*model.RecentPub, err error) {
	if items, err = p.d.GetUserRecentPubsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}
