package dao

import (
	"context"
	discuss "valerian/app/service/discuss/api"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	return p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id})
}

func (p *Dao) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *discuss.UserDiscussionsResp, err error) {
	return p.discussRPC.GetUserDiscussionsPaged(c, &discuss.UserDiscussionsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)})
}
