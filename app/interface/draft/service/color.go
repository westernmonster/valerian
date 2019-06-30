package service

import (
	"context"
	"time"
	"valerian/app/interface/draft/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) GetUserColors(c context.Context) (items []*model.Color, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var addCache = true
	if items, err = p.d.ColorsCache(c, aid); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetUserColors(c, p.d.DB(), aid); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetColorsCache(context.TODO(), aid, items)
		})
	}
	return
}

func (p *Service) AddColor(c context.Context, arg *model.ArgAddColor) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &model.Color{
		ID:        gid.NewID(),
		Name:      arg.Name,
		Color:     arg.Color,
		AccountID: aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddColor(c, p.d.DB(), item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
		p.d.DelColorsCache(context.TODO(), aid)
	})

	return
}

func (p *Service) getColor(c context.Context, aid, id int64) (item *model.Color, err error) {
	if item, err = p.d.GetColor(c, p.d.DB(), id); err != nil {
		return
	} else if item == nil {
		err = ecode.ColorNotExist
		return
	} else if item.AccountID != aid {
		err = ecode.NotBelongToYou
		return
	}

	return
}

func (p *Service) UpdateColor(c context.Context, arg *model.ArgUpdateColor) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.Color
	if item, err = p.d.GetColor(c, p.d.DB(), arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.ColorNotExist
		return
	} else if item.AccountID != aid {
		err = ecode.NotBelongToYou
		return
	}

	item.Color = arg.Color
	item.Name = arg.Name

	if err = p.d.UpdateColor(c, p.d.DB(), item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
		p.d.DelColorsCache(context.TODO(), aid)
	})

	return
}

func (p *Service) DelColor(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.Color
	if item, err = p.d.GetColor(c, p.d.DB(), id); err != nil {
		return
	} else if item == nil {
		return
	} else if item.AccountID != aid {
		err = ecode.NotBelongToYou
		return
	}

	if err = p.d.DelColor(c, p.d.DB(), id); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
		p.d.DelColorsCache(context.TODO(), aid)
	})

	return
}