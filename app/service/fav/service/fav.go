package service

import (
	"context"
	"time"
	"valerian/app/service/fav/api"
	"valerian/app/service/fav/model"
	"valerian/library/gid"
)

func (p *Service) IsFav(c context.Context, aid int64, targetID int64, targetType string) (isFav bool, err error) {
	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		isFav = true
		return
	}
	return
}

func (p *Service) Fav(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav != nil {
		return
	}

	fav = &model.Fav{
		ID:         gid.NewID(),
		AccountID:  aid,
		TargetID:   targetID,
		TargetType: targetType,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddFav(c, p.d.DB(), fav); err != nil {
		return
	}

	return
}

func (p *Service) Unfav(c context.Context, aid, targetID int64, targetType string) (err error) {

	var fav *model.Fav
	if fav, err = p.d.GetFavByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":  aid,
		"target_type": targetType,
		"target_id":   targetID,
	}); err != nil {
		return
	} else if fav == nil {
		return
	} else {
		if err = p.d.DelFav(c, p.d.DB(), fav.ID); err != nil {
			return
		}
	}

	return
}

func (p *Service) GetFavIDsPaged(c context.Context, arg *api.UserFavsReq) (ids []int64, err error) {
	return p.d.GetFavIDsPaged(c, p.d.DB(), arg.AccountID, arg.TargetType, arg.Limit, arg.Offset)
}
