package service

import (
	"context"
	"time"

	"valerian/app/service/like/model"
	"valerian/library/gid"
)

func (p *Service) IsDislike(c context.Context, aid int64, targetID int64, targetType string) (isDislike bool, err error) {
	var fav *model.Dislike
	if fav, err = p.d.GetDislikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isDislike = true
		return
	}
	return
}

func (p *Service) Dislike(c context.Context, aid, targetID int64, targetType string) (err error) {
	var fav *model.Dislike
	if fav, err = p.d.GetDislikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Dislike{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddDislike(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

func (p *Service) CancelDislike(c context.Context, aid, targetID int64, targetType string) (err error) {
	var fav *model.Dislike
	if fav, err = p.d.GetDislikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	} else {
		if err = p.d.DelDislike(c, p.d.DB(), fav.ID); err != nil {
			return
		}
	}

	return
}
