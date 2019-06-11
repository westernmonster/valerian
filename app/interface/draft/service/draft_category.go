package service

import (
	"context"
	"time"
	"valerian/app/interface/draft/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) GetUserDraftCategories(c context.Context) (items []*model.DraftCategoryResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var addCache = true
	if items, err = p.d.DraftCategoriesCache(c, aid); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetUserDraftCategories(c, p.d.DB(), aid); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetDraftCategoriesCache(context.TODO(), aid, items)
		})
	}
	return
}

func (p *Service) AddDraftCategory(c context.Context, arg *model.ArgAddDraftCategory) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if _, err = p.getColor(c, aid, arg.ColorID); err != nil {
		return
	}

	item := &model.DraftCategory{
		ID:        gid.NewID(),
		Name:      arg.Name,
		ColorID:   arg.ColorID,
		AccountID: aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddDraftCategory(c, p.d.DB(), item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
	})

	return
}

func (p *Service) UpdateDraftCategory(c context.Context, arg *model.ArgUpdateDraftCategory) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if _, err = p.getColor(c, aid, arg.ColorID); err != nil {
		return
	}

	var item *model.DraftCategory
	if item, err = p.d.GetDraftCategory(c, p.d.DB(), arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DraftCategoryNotExist
		return
	} else if item.AccountID != aid {
		err = ecode.NotBelongToYou
		return
	}

	item.ColorID = arg.ColorID
	item.Name = arg.Name

	if err = p.d.UpdateDraftCategory(c, p.d.DB(), item); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
	})

	return
}

func (p *Service) DelDraftCategory(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.DraftCategory
	if item, err = p.d.GetDraftCategory(c, p.d.DB(), id); err != nil {
		return
	} else if item == nil {
		return
	} else if item.AccountID != aid {
		err = ecode.NotBelongToYou
		return
	}

	if err = p.d.DelDraftCategory(c, p.d.DB(), id); err != nil {
		return
	}

	p.addCache(func() {
		p.d.DelDraftCategoriesCache(context.TODO(), aid)
	})

	return
}
