package service

import (
	"context"

	"valerian/app/service/relation/model"
)

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
