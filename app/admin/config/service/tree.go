package service

import (
	"context"
	"valerian/app/admin/config/model"
)

func (p *Service) GetTreesPaged(c context.Context, cond map[string]interface{}, page, pageSize int32) (resp *model.TreeListResp, err error) {
	var count int32
	var data []*model.Tree
	if count, data, err = p.d.GetTreesByCondPaged(c, p.d.ConfigDB(), cond, page, pageSize); err != nil {
		return
	}

	resp = &model.TreeListResp{
		Total:    count,
		Page:     page,
		PageSize: pageSize,
		Items:    make([]*model.TreeItem, len(data)),
	}

	for i, v := range data {
		item := &model.TreeItem{
			TreeID:     v.TreeID,
			Name:       v.Name,
			Mark:       v.Mark,
			PlatformID: v.PlatformID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		switch v.PlatformID {
		case model.PlatformAdmin:
			item.PlatformName = "admin"
			break
		case model.PlatformInfra:
			item.PlatformName = "infra"
			break
		case model.PlatformService:
			item.PlatformName = "service"
			break
		case model.PlatformInterface:
			item.PlatformName = "interface"
			break
		}

		resp.Items[i] = item
	}

	return
}
