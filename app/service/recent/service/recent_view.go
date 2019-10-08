package service

import (
	"context"
	"valerian/app/service/recent/model"
)

func (p *Service) GetUserRecentViewsPaged(c context.Context, aid int64, limit, offset int) (items []*model.RecentView, err error) {
	if items, err = p.d.GetUserRecentViewsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}
