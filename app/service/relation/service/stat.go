package service

import (
	"context"
	"time"

	"valerian/app/service/relation/model"
)

func (p *Service) Stat(c context.Context, aid, targetAid int64) (stat *model.AccountRelationStat, isFollowing bool, err error) {
	if stat, err = p.d.GetStatByID(c, p.d.DB(), targetAid); err != nil {
		return
	} else if stat == nil {
		stat = &model.AccountRelationStat{
			AccountID: targetAid,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
	}

	if isFollowing, err = p.IsFollowing(c, aid, targetAid); err != nil {
		return
	}

	return
}

func (p *Service) IsFollowing(c context.Context, aid, targetID int64) (isFollowing bool, err error) {

	if aid == targetID {
		return
	}

	var following *model.AccountFollowing
	if following, err = p.d.GetFollowingByCond(c, p.d.DB(), map[string]interface{}{
		"account_id":        aid,
		"target_account_id": targetID,
	}); err != nil {
		return
	} else if following == nil {
		return
	}

	isFollowing = following.Following()
	return
}
