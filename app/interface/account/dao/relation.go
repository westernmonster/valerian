package dao

import (
	"context"
	relation "valerian/app/service/relation/api"
)

func (p *Dao) GetFollowings(c context.Context, accountID int64, limit, offset int) (resp *relation.FollowingResp, err error) {
	return p.relationRPC.GetFollowingPaged(c, &relation.RelationReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)})
}

func (p *Dao) GetFans(c context.Context, accountID int64, limit, offset int) (resp *relation.FansResp, err error) {
	return p.relationRPC.GetFansPaged(c, &relation.RelationReq{AccountID: accountID, Limit: int32(limit), Offset: int32(offset)})
}

func (p *Dao) Follow(c context.Context, accountID, targetAccountID int64) (err error) {
	_, err = p.relationRPC.Follow(c, &relation.FollowReq{AccountID: accountID, TargetAccountID: targetAccountID})
	return
}

func (p *Dao) Unfollow(c context.Context, accountID, targetAccountID int64) (err error) {
	_, err = p.relationRPC.Unfollow(c, &relation.FollowReq{AccountID: accountID, TargetAccountID: targetAccountID})
	return
}

func (p *Dao) Stat(c context.Context, aid, targetID int64) (resp *relation.StatInfo, err error) {
	return p.relationRPC.Stat(c, &relation.FollowReq{AccountID: aid, TargetAccountID: targetID})
}
