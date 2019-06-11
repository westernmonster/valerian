package service

import (
	"context"
	"valerian/app/interface/draft/model"
)

func (p *Service) GetAllColors(c context.Context) (items []*model.Color, err error) {
	var addCache = true
	if items, err = p.d.ColorsCache(c); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetAllColors(c, p.d.DB()); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetColorsCache(context.TODO(), items)
		})
	}
	return
}
