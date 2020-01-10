package service

import (
	"context"
	"valerian/app/interface/locale/model"
)

func (p *Service) GetAllLocales(ctx context.Context) (items []*model.Locale, err error) {
	var addCache = true

	items = make([]*model.Locale, 0)
	if items, err = p.d.LocalesCache(ctx); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetAllLocales(ctx, p.d.DB()); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetLocalesCache(context.Background(), items)
		})
	}
	return
}
