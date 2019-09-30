package dao

import (
	"context"
	discuss "valerian/app/service/discuss/api"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	return p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id})
}
