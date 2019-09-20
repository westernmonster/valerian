package service

import (
	"context"
	"valerian/app/service/relation/model"
)

// Following 分页获取关注列表
func (p *Service) Following(c context.Context, aid int64, limit, offset int) (resp []*model.FollowingResp, err error) {
	var (
		addCache = true
		items    []*model.AccountRelation
	)

	resp = make([]*model.FollowingResp, 0)

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
			AccountID: v.AccountID,
			Attribute: v.Attribute,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})

	}
	return
}
