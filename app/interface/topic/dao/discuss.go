package dao

import (
	"context"
	"valerian/app/service/discuss/api"
	discuss "valerian/app/service/discuss/api"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	return p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id})
}

func (p *Dao) GetDiscussionCategories(c context.Context, topicID int64) (resp *api.CategoriesResp, err error) {
	return p.discussRPC.GetDiscussionCategories(c, &discuss.CategoriesReq{TopicID: topicID})
}
