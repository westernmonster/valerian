package dao

import (
	"context"
	"fmt"
	discuss "valerian/app/service/discuss/api"
	"valerian/library/log"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	if info, err = p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussion, error(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *discuss.UserDiscussionsResp, err error) {
	if resp, err = p.discussRPC.GetUserDiscussionsPaged(c, &discuss.UserDiscussionsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionsPaged, error(%+v), aid(%d), limit(%d), offset(%d)`", err, aid, limit, offset))
	}
	return
}

func (p *Dao) GetUserDiscussionIDsPaged(c context.Context, aid int64, limit, offset int) (resp *discuss.IDsResp, err error) {
	if resp, err = p.discussRPC.GetUserDiscussionIDsPaged(c, &discuss.UserDiscussionsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionIDsPaged, error(%+v), aid(%d), limit(%d), offset(%d)`", err, aid, limit, offset))
	}
	return
}
