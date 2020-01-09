package service

import (
	"context"
	"valerian/app/admin/config/model"
)

func (p *Service) GetPlatformList(c context.Context) (items []*model.PlatformListItem, err error) {
	items = make([]*model.PlatformListItem, 0)
	items = append(items, &model.PlatformListItem{
		ID:   1,
		Name: "admin",
	})

	items = append(items, &model.PlatformListItem{
		ID:   2,
		Name: "infra",
	})

	items = append(items, &model.PlatformListItem{
		ID:   3,
		Name: "service",
	})

	items = append(items, &model.PlatformListItem{
		ID:   4,
		Name: "interface",
	})

	return
}
