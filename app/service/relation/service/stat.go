package service

import (
	"context"
	"time"

	"valerian/app/service/relation/model"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Stat(c context.Context, aid int64) (stat *model.AccountRelationStat, isFollowing bool, err error) {
	if stat, err = p.d.GetStatByID(c, p.d.DB(), aid); err != nil {
		return
	} else if stat == nil {
		stat = &model.AccountRelationStat{
			AccountID: aid,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
	}

	if isFollowing, err = p.IsFollowing(c, aid); err != nil {
		return
	}

	return
}

func (p *Service) IsFollowing(c context.Context, targetID int64) (isFollowing bool, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

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
