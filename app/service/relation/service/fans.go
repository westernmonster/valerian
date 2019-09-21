package service

import (
	"context"
	"valerian/app/service/relation/model"
)

// Fans 分页获取关注列表
func (p *Service) Fans(c context.Context, aid int64, limit, offset int) (resp []*model.FansResp, err error) {
	var (
		addCache = true
		items    []*model.AccountFans
	)

	resp = make([]*model.FansResp, 0)

	if items, err = p.d.FansCache(c, aid, limit, offset); err != nil {
		addCache = false
	}

	if items == nil {
		if items, err = p.d.GetFansPaged(c, p.d.DB(), aid, limit, offset); err != nil {
			return
		}
	}

	if items != nil && addCache {
		p.addCache(func() {
			p.d.SetFansCache(context.TODO(), aid, limit, offset, items)
		})
	}

	for _, v := range items {
		resp = append(resp, &model.FansResp{
			AccountID: v.AccountID,
			Attribute: v.Attribute,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})

	}
	return
}
