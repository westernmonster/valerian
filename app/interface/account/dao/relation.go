package dao

import (
	"context"
	"fmt"
	relation "valerian/app/service/relation/api"
	"valerian/library/log"
)

func (p *Dao) GetFollowings(c context.Context, accountID int64, limit, offset int) (resp *relation.FollowingResp, err error) {
	if resp, err = p.relationRPC.GetFollowingPaged(c, &relation.RelationReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowings, error(%+v), aid(%d), limit(%d), offset(%d)", err, accountID, limit, offset))
	}
	return
}

func (p *Dao) GetFans(c context.Context, accountID int64, limit, offset int) (resp *relation.FansResp, err error) {
	if resp, err = p.relationRPC.GetFansPaged(c, &relation.RelationReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFans, error(%+v), aid(%d), limit(%d), offset(%d)", err, accountID, limit, offset))
	}

	return
}

func (p *Dao) Follow(c context.Context, accountID, targetAccountID int64) (err error) {
	log.For(c).Error(fmt.Sprintf("dao.Unfollow, error(%+v), aid(%d), target_aid(%d)", err, accountID, targetAccountID))
	if _, err = p.relationRPC.Follow(c, &relation.FollowReq{AccountID: accountID, TargetAccountID: targetAccountID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Follow, error(%+v), aid(%d), target_aid(%d)", err, accountID, targetAccountID))
	}
	return
}

func (p *Dao) Unfollow(c context.Context, accountID, targetAccountID int64) (err error) {
	log.For(c).Error(fmt.Sprintf("dao.Unfollow, error(%+v), aid(%d), target_aid(%d)", err, accountID, targetAccountID))
	if _, err = p.relationRPC.Unfollow(c, &relation.FollowReq{AccountID: accountID, TargetAccountID: targetAccountID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Unfollow, error(%+v), aid(%d), target_aid(%d)", err, accountID, targetAccountID))
	}
	return
}

func (p *Dao) IsFollowing(c context.Context, aid, targetAccountID int64) (IsFollowing bool, err error) {
	var resp *relation.IsFollowingResp
	if resp, err = p.relationRPC.IsFollowing(c, &relation.FollowReq{AccountID: aid, TargetAccountID: targetAccountID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsFollowing, error(%+v), aid(%d), target_aid(%d)", err, aid, targetAccountID))
		return
	}
	return resp.IsFollowing, nil
}
