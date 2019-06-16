package service

import (
	"context"
	"time"
	"valerian/app/interface/draft/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) GetUserDrafts(c context.Context, categoryID *int64) (items []*model.DraftResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	cond := make(map[string]interface{})
	if categoryID != nil {
		cond["category_id"] = *categoryID
	}

	data, err := p.d.GetUserDrafts(c, p.d.DB(), aid, cond)
	if err != nil {
		return
	}

	for _, v := range data {
		item := &model.DraftResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Text:      v.Text,
			AccountID: v.AccountID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		if v.CategoryID != nil {
			if item.Category, err = p.getDraftCategory(c, aid, *v.CategoryID); err != nil {
				return
			}
		}

		items = append(items, item)
	}

	return
}

func (p *Service) AddDraft(c context.Context, arg *model.ArgAddDraft) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if arg.CategoryID != nil {
		var category *model.DraftCategory
		if category, err = p.d.GetDraftCategory(c, p.d.DB(), *arg.CategoryID); err != nil {
			return
		} else if category == nil {
			err = ecode.DraftCategoryNotExist
			return
		}
	}

	item := &model.Draft{
		ID:         gid.NewID(),
		Title:      arg.Title,
		Content:    arg.Content,
		AccountID:  aid,
		CategoryID: arg.CategoryID,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddDraft(c, p.d.DB(), item); err != nil {
		return
	}

	id = item.ID

	return
}

func (p *Service) UpdateDraft(c context.Context, arg *model.ArgUpdateDraft) (err error) {
	_, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if arg.CategoryID != nil {
		var category *model.DraftCategory
		if category, err = p.d.GetDraftCategory(c, p.d.DB(), *arg.CategoryID); err != nil {
			return
		} else if category == nil {
			err = ecode.DraftCategoryNotExist
			return
		}
	}

	var item *model.Draft
	if item, err = p.d.GetDraft(c, p.d.DB(), arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DraftNotExist
		return
	}

	item.Title = arg.Title
	item.Content = arg.Content
	item.CategoryID = arg.CategoryID
	item.UpdatedAt = time.Now().Unix()

	if err = p.d.UpdateDraft(c, p.d.DB(), item); err != nil {
		return
	}

	return
}

func (p *Service) DelDraft(c context.Context, id int64) (err error) {
	_, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.DelDraft(c, p.d.DB(), id); err != nil {
		return
	}

	return
}
