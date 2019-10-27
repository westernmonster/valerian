package dao

import (
	"context"
	"fmt"
	discuss "valerian/app/service/discuss/api"
	"valerian/library/log"
)

func (p *Dao) GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error) {
	if info, err = p.discussRPC.GetDiscussionInfo(c, &discuss.IDReq{ID: id, Include: "content,content_text"}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussion, error(%+v) id(%d)", err, id))
	}
	return
}

func (p *Dao) GetAllDiscussions(c context.Context) (items []*discuss.DiscussionInfo, err error) {
	var resp *discuss.AllDiscussionsResp
	if resp, err = p.discussRPC.GetAllDiscussions(c, &discuss.EmptyStruct{}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllDiscussions, error(%+v) ", err))
		return
	}

	items = make([]*discuss.DiscussionInfo, 0)
	if resp.Items != nil {
		items = resp.Items
	}
	return
}
