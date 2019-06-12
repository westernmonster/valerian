package service

import (
	"context"
	"valerian/app/interface/draft/model"
	"valerian/library/ecode"
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

		if item.Category, err = p.getDraftCategory(c, aid, v.CategoryID); err != nil {
			return
		}

		items = append(items, item)
	}

	return
}
