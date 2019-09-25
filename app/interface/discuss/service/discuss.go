package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/interface/discuss/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"

	"github.com/PuerkitoBio/goquery"
)

func (p *Service) onAddDiscussion(c context.Context, id int64) {
	if err := p.d.NotifyDiscussionAdded(c, id); err != nil {
		log.For(c).Error(fmt.Sprintf("NotifyDiscussionAdded(%d) : %+v", id, err))
		return
	}
}

func (p *Service) onUpdateDiscussion(c context.Context, id int64) {
	if err := p.d.NotifyDiscussionUpdated(c, id); err != nil {
		log.For(c).Error(fmt.Sprintf("NotifyDiscussionUpdated(%d) : %+v", id, err))
		return
	}
}

func (p *Service) onDelDiscussion(c context.Context, id int64, topicID int64) {
	if err := p.d.NotifyDiscussionDeleted(c, id, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("NotifyDiscussionDeleted, id(%d) topic_id(%d): %+v", id, topicID, err))
		return
	}
}

func (p *Service) initStat(c context.Context, aid, topicID int64) (err error) {
	if err = p.d.AddAccountStat(c, p.d.DB(), &model.AccountResStat{
		AccountID: aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.d.AddTopicStat(c, p.d.DB(), &model.TopicResStat{
		TopicID:   aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	return
}

func (p *Service) initDiscussionStat(c context.Context, discussionID int64) (err error) {
	if err = p.d.AddDiscussionStat(c, p.d.DB(), &model.DiscussionStat{
		DiscussionID: discussionID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	return
}

func (p *Service) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	if items, err = p.d.GetUserDiscussionsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) GetTopicDiscussionsPaged(c context.Context, topicID, categoryID int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	if items, err = p.d.GetTopicDiscussionsPaged(c, p.d.DB(), topicID, categoryID, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) AddDiscussion(c context.Context, arg *model.ArgAddDiscuss) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.checkCategory(c, p.d.DB(), arg.CategoryID); err != nil {
		return
	}

	if err = p.initStat(c, aid, arg.TopicID); err != nil {
		return
	}

	// 检测话题
	if _, err = p.d.GetTopic(c, arg.TopicID); err != nil {
		return
	}

	// 检测是否成员
	if err = p.checkTopicMember(c, arg.TopicID, aid); err != nil {
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	item := &model.Discussion{
		ID:         gid.NewID(),
		TopicID:    arg.TopicID,
		CategoryID: arg.CategoryID,
		CreatedBy:  aid,
		Title:      arg.Title,
		Content:    arg.Content,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}

	item.ContentText = doc.Text()

	if err = p.d.AddDiscussion(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateFiles(c, tx, item.ID, arg.Files); err != nil {
		return
	}

	// Update Stat
	if err = p.d.IncrAccountStat(c, tx, &model.AccountResStat{AccountID: aid, DiscussionCount: 1}); err != nil {
		return
	}
	if err = p.d.IncrTopicStat(c, tx, &model.TopicResStat{TopicID: aid, DiscussionCount: 1}); err != nil {
		return
	}

	if err = p.d.AddDiscussionStat(c, tx, &model.DiscussionStat{
		DiscussionID: item.ID,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	id = item.ID

	p.addCache(func() {
		p.onAddDiscussion(context.Background(), id)
	})

	return
}

func (p *Service) checkCategory(c context.Context, node sqalx.Node, categoryID int64) (err error) {
	if categoryID == -1 {
		return
	}

	var t *model.DiscussCategory
	if t, err = p.d.GetDiscussCategoryByID(c, node, categoryID); err != nil {
		return
	} else if t == nil {
		return ecode.DiscussCategoryNotExist
	}

	return
}

func (p *Service) UpdateDiscussion(c context.Context, arg *model.ArgUpdateDiscuss) (err error) {

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, tx, arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if item.CreatedBy != aid {
		if err = p.checkTopicManager(c, item.TopicID, aid); err != nil {
			return
		}
	}

	item.Title = arg.Title
	item.Content = arg.Content
	item.UpdatedAt = time.Now().Unix()

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		err = ecode.ParseHTMLFailed
		return
	}

	item.ContentText = doc.Text()

	if err = p.d.UpdateDiscussion(c, tx, item); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onUpdateDiscussion(context.Background(), arg.ID)
	})

	return
}

func (p *Service) DelDiscussion(c context.Context, id int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), id); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	if err = p.initStat(c, aid, item.TopicID); err != nil {
		return
	}

	if item.CreatedBy != aid {
		if err = p.checkTopicManager(c, item.TopicID, aid); err != nil {
			return
		}
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	if err = p.d.DelDiscussion(c, tx, id); err != nil {
		return
	}

	if err = p.d.DelDiscussionFiles(c, tx, id); err != nil {
		return
	}

	if err = p.d.IncrAccountStat(c, tx, &model.AccountResStat{AccountID: aid, DiscussionCount: -1}); err != nil {
		return
	}
	if err = p.d.IncrTopicStat(c, tx, &model.TopicResStat{TopicID: aid, DiscussionCount: -1}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onDelDiscussion(context.Background(), id, item.TopicID)
		p.d.DelDiscussionFilesCache(context.TODO(), id)
	})

	return
}

func (p *Service) GetDiscussion(c context.Context, discussionID int64) (resp *model.DiscussDetailResp, err error) {

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data *model.Discussion
	if data, err = p.getDiscussion(c, p.d.DB(), discussionID); err != nil {
		return
	}

	resp = &model.DiscussDetailResp{
		ID:      data.ID,
		TopicID: data.TopicID,
		Title:   data.Title,
		Content: data.Content,
		Files:   make([]*model.DiscussFileResp, 0),
	}

	if data.CategoryID == -1 {
		resp.Category = &model.DiscussItemCategory{
			ID:   data.CategoryID,
			Name: "问答",
		}
	} else {
		var category *model.DiscussCategory
		if category, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), data.CategoryID); err != nil {
			return
		} else if category == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		resp.Category = &model.DiscussItemCategory{
			ID:   data.CategoryID,
			Name: category.Name,
		}

	}
	if account, e := p.d.GetAccountBaseInfo(c, data.CreatedBy); e != nil {
		err = e
		return
	} else {
		resp.Creator = &model.Creator{
			ID:       account.ID,
			UserName: account.UserName,
			Avatar:   account.Avatar,
		}

		intro := account.GetIntroductionValue()
		resp.Creator.Introduction = &intro
	}

	var stat *model.DiscussionStat
	if stat, err = p.d.GetDiscussionStatByID(c, p.d.DB(), discussionID); err != nil {
		return
	}

	resp.LikeCount = stat.LikeCount
	resp.CommentCount = stat.CommentCount

	if aid == data.CreatedBy {
		return
	}

	var info *topic.MemberRoleReply
	if info, err = p.d.GetTopicMemberRole(c, data.TopicID, aid); err != nil {
		return
	}

	if info.IsMember && info.Role != model.MemberRoleUser {
		resp.CanEdit = true
	}

	return
}

func (p *Service) getDiscussion(c context.Context, node sqalx.Node, discussionID int64) (item *model.Discussion, err error) {
	if item, err = p.d.GetDiscussionByID(c, node, discussionID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	return
}
