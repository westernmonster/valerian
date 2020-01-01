package dao

import (
	"context"
	"fmt"
	discussion "valerian/app/service/discuss/api"
	"valerian/library/log"
)

func (p *Dao) GetDiscussion(c context.Context, req *discussion.IDReq) (info *discussion.DiscussionInfo, err error) {
	if info, err = p.discussionRPC.GetDiscussionInfo(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussion, error(%+v) req(%+v)", err, req))
	}
	return
}

// GetUserDiscussionsPaged
// aid 当前登录用户
// accountID 要用户讨论的用户
func (p *Dao) GetUserDiscussionsPaged(c context.Context, accountID, aid int64, limit, offset int) (resp *discussion.DiscussionsResp, err error) {
	if resp, err = p.discussionRPC.GetUserDiscussionsPaged(c, &discussion.UserDiscussionsReq{AccountID: accountID, Aid: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionsPaged, error(%+v), account_id(%d), aid(%d), limit(%d), offset(%d)`", err, accountID, aid, limit, offset))
	}
	return
}

func (p *Dao) GetTopicDiscussionsPaged(c context.Context, topicID, categoryID, aid int64, limit, offset int) (resp *discussion.DiscussionsResp, err error) {
	if resp, err = p.discussionRPC.GetTopicDiscussionsPaged(c, &discussion.TopicDiscussionsReq{TopicID: topicID, CategoryID: categoryID, Aid: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged, error(%+v), aid(%d), topic_id(%d), category_id(%d), limit(%d), offset(%d)", err, aid, topicID, categoryID, limit, offset))
	}
	return
}

func (p *Dao) GetDiscussionCategories(c context.Context, aid, topicID int64) (resp *discussion.CategoriesResp, err error) {
	if resp, err = p.discussionRPC.GetDiscussionCategories(c, &discussion.CategoriesReq{Aid: aid, TopicID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionCategories, error(%+v), aid(%d), tpic_id(%d)", err, aid, topicID))
	}
	return
}

func (p *Dao) GetUserDiscussionIDsPaged(c context.Context, aid int64, limit, offset int) (resp *discussion.IDsResp, err error) {
	if resp, err = p.discussionRPC.GetUserDiscussionIDsPaged(c, &discussion.UserDiscussionsReq{AccountID: aid, Limit: int32(limit), Offset: int32(offset)}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionIDsPaged, error(%+v), aid(%d), limit(%d), offset(%d)`", err, aid, limit, offset))
	}
	return
}

func (p *Dao) SaveDiscussionCategories(c context.Context, arg *discussion.ArgSaveDiscussCategories) (err error) {
	if _, err = p.discussionRPC.SaveDiscussionCategories(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SaveDiscussionCategories, error(%+v), arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) AddDiscussion(c context.Context, arg *discussion.ArgAddDiscussion) (id int64, err error) {
	var resp *discussion.IDResp
	if resp, err = p.discussionRPC.AddDiscussion(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDiscussion, error(%+v), arg(%+v)", err, arg))
	}

	id = resp.ID
	return
}

func (p *Dao) UpdateDiscussion(c context.Context, arg *discussion.ArgUpdateDiscussion) (err error) {
	if _, err = p.discussionRPC.UpdateDiscussion(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDiscussion, error(%+v), arg(%+v)", err, arg))
	}

	return
}

func (p *Dao) DelDiscussion(c context.Context, aid, discussionID int64) (err error) {
	if _, err = p.discussionRPC.DelDiscussion(c, &discussion.IDReq{Aid: aid, ID: discussionID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDiscussion, error(%+v), aid(%d), discussion_id(%d)", err, aid, discussionID))
	}
	return
}

func (p *Dao) GetDiscussionFiles(c context.Context, aid, discussionID int64) (resp *discussion.DiscussionFilesResp, err error) {
	if resp, err = p.discussionRPC.GetDiscussionFiles(c, &discussion.IDReq{Aid: aid, ID: discussionID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionFiles, error(%+v), aid(%d), discussion_id(%d)", err, aid, discussionID))
	}
	return
}

func (p *Dao) SaveDiscussionFiles(c context.Context, arg *discussion.ArgSaveDiscussionFiles) (err error) {
	if _, err = p.discussionRPC.SaveDiscussionFiles(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SaveDiscussionFiles, error(%+v), arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) CanEdit(c context.Context, req *discussion.IDReq) (canEdit bool, err error) {
	var resp *discussion.BoolResp
	if resp, err = p.discussionRPC.CanEdit(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CanEdit() error(%+v), req(%+v)", err, req))
		return
	}
	canEdit = resp.Result
	return
}

func (p *Dao) CanView(c context.Context, req *discussion.IDReq) (canView bool, err error) {
	var resp *discussion.BoolResp
	if resp, err = p.discussionRPC.CanView(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CanView() error(%+v), req(%+v)", err, req))
		return
	}
	canView = resp.Result
	return
}
