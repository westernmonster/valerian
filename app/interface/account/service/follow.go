package service

import (
	"context"
	"time"
	"valerian/app/interface/account/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *model.ArgFollow) (isFollow bool, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.AccountFollower
	if item, err = p.d.GetAccountFollowerByCond(c, p.d.DB(), map[string]interface{}{"account_id": arg.AccountID, "follower_id": aid}); err != nil {
		return
	} else if item == nil {
		isFollow = true
		af := &model.AccountFollower{
			ID:         gid.NewID(),
			AccountID:  arg.AccountID,
			FollowerID: aid,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}
		if err = p.d.AddAccountFollower(c, p.d.DB(), af); err != nil {
			return
		}
	} else {
		isFollow = false
		if err = p.d.DelAccountFollower(c, p.d.DB(), item.ID); err != nil {
			return
		}
	}

	return
}

func (p *Service) FansPaged(c context.Context, aid int64) (isFollow bool, err error) {
	return
}
