package service

import (
	"context"
	"time"

	"valerian/app/service/like/model"
	"valerian/library/gid"
)

func (p *Service) IsLike(c context.Context, aid int64, targetID int64, targetType string) (isLike bool, err error) {
	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isLike = true
		return
	}
	return
}

func (p *Service) Like(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Like{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddLike(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

func (p *Service) CancelLike(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Like
	if fav, err = p.d.GetLikeByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	} else {
		if err = p.d.DelLike(c, p.d.DB(), fav.ID); err != nil {
			return
		}
	}

	return
}
