package dao

import (
	"context"
	"fmt"
	discuss "valerian/app/service/discuss/api"
	"valerian/library/log"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	if info, err = p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussion err(%+v)", err))
	}
	return
}

func (p *Dao) GetDiscussionCategories(c context.Context, topicID int64) (resp *discuss.CategoriesResp, err error) {
	if resp, err = p.discussRPC.GetDiscussionCategories(c, &discuss.CategoriesReq{TopicID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionCategories err(%+v)", err))
	}
	return
}
