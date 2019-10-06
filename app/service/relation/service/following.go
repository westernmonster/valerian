package service

import (
	"context"
	"fmt"
	"valerian/app/service/relation/model"
	"valerian/library/log"
)

// Following 分页获取关注列表
func (p *Service) FollowingPaged(c context.Context, aid int64, limit, offset int) (resp []*model.FollowingResp, err error) {
	var (
		addCache = true
		items    []*model.AccountFollowing
	)

	resp = make([]*model.FollowingResp, 0)

	log.For(c).Info(fmt.Sprintf("service.FollowingPaged aid(%d)", aid))

	if items, err = p.d.FollowingsCache(c, aid, limit, offset); err != nil {
		addCache = false
	}

	if items == nil {
		if items, err = p.d.GetFollowingsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetFollowingsCache(context.TODO(), aid, limit, offset, items)
		})
	}

	for _, v := range items {
		resp = append(resp, &model.FollowingResp{
			AccountID: v.TargetAccountID,
			Attribute: v.Attribute,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})

	}
	return
}
