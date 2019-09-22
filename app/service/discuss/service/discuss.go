package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

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

func (p *Service) onDelDiscussion(c context.Context, id int64) {
	if err := p.d.NotifyDiscussionDeleted(c, id); err != nil {
		log.For(c).Error(fmt.Sprintf("NotifyDiscussionDeleted(%d) : %+v", id, err))
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

func (p *Service) GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	if items, err = p.d.GetUserDiscussionsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) GetTopicDiscussionsPaged(c context.Context, topicID int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	if items, err = p.d.GetTopicDiscussionsPaged(c, p.d.DB(), topicID, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) AddDiscussion(c context.Context, aid int64, arg *model.ArgAddDiscuss) (id int64, err error) {
	// TODO: 权限检测
	// 有什么权限的可以添加

	if err = p.checkCategory(c, p.d.DB(), arg.CategoryID); err != nil {
		return
	}

	if err = p.initStat(c, aid, arg.TopicID); err != nil {
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

	// Update Stat
	if err = p.d.IncrAccountStat(c, tx, &model.AccountResStat{AccountID: aid, DiscussionCount: 1}); err != nil {
		return
	}
	if err = p.d.IncrTopicStat(c, tx, &model.TopicResStat{TopicID: aid, DiscussionCount: 1}); err != nil {
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

func (p *Service) UpdateDiscussion(c context.Context, aid int64, arg *model.ArgUpdateDiscuss) (err error) {
	// TODO: 编辑权限检测
	// 有什么权限的可以对其进行编辑

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
		err = ecode.ModifyDiscussionNotAllowed
		return
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

func (p *Service) DelDiscussion(c context.Context, aid int64, arg *model.ArgDelDiscuss) (err error) {
	var item *model.Discussion
	if item, err = p.d.GetDiscussionByID(c, p.d.DB(), arg.ID); err != nil {
		return
	} else if item == nil {
		err = ecode.DiscussionNotExist
		return
	}

	// TODO: 编辑权限检测
	// 有什么权限的可以对其进行编辑
	if err = p.initStat(c, aid, item.TopicID); err != nil {
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

	if item.CreatedBy != aid {
		err = ecode.ModifyDiscussionNotAllowed
		return
	}

	if err = p.d.DelDiscussion(c, tx, arg.ID); err != nil {
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
		p.onDelDiscussion(context.Background(), arg.ID)
	})

	return
}
