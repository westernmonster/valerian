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
