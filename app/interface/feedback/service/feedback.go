package service

import (
	"context"
	"time"
	"valerian/app/interface/feedback/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"

	"github.com/jinzhu/copier"
)

func (p *Service) FeedbackTypeList(c context.Context, fType string) (items []*model.FeedbackTypeResp, err error) {
	items = make([]*model.FeedbackTypeResp, 0)

	var data []*model.FeedbackType
	if data, err = p.d.GetFeedbackTypesByCond(c, p.d.DB(), map[string]interface{}{"type": fType}); err != nil {
		return
	}

	if err = copier.Copy(&items, &data); err != nil {
		return
	}

	return
}

func (p *Service) AddFeedback(c context.Context, arg *model.ArgAddFeedback) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var fType *model.FeedbackType
	if fType, err = p.d.GetFeedbackTypeByID(c, p.d.DB(), arg.Type); err != nil {
		return
	}

	if fType == nil {
		err = ecode.InvalidFeedbackType
		return
	}

	item := &model.Feedback{
		ID:           gid.NewID(),
		TargetID:     arg.TargetID,
		TargetType:   arg.TargetType,
		FeedbackType: fType.ID,
		FeedbackDesc: arg.Desc,
		CreatedBy:    aid,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	if err = p.d.AddFeedback(c, p.d.DB(), item); err != nil {
		return
	}

	return
}
