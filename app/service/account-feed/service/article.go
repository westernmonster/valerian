package service

import (
	"context"
	"time"
	"valerian/app/service/account-feed/model"
	"valerian/app/service/feed/def"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"

	"github.com/nats-io/stan.go"
)

func (p *Service) getArticleHistory(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleHistory, err error) {
	if item, err = p.d.GetArticleHistoryByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleHistoryNotExist
		return
	}

	return
}

func (p *Service) getArticle(c context.Context, node sqalx.Node, articleID int64) (item *model.Article, err error) {
	var addCache = true
	if item, err = p.d.ArticleCache(c, articleID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetArticleByID(c, p.d.DB(), articleID); err != nil {
		return
	} else if item == nil {
		err = ecode.ArticleNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetArticleCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		PromError("account-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeCreateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var history *model.ArticleHistory
	if history, err = p.getArticleHistory(c, p.d.DB(), info.ArticleHistoryID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		PromError("account-feed: getArticleHistory", "getArticleHistory(), id(%d),error(%+v)", info.ArticleHistoryID, err)
		return
	}

	if _, err = p.getArticle(c, p.d.DB(), info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		PromError("account-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeUpdateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextUpdateArticle,
		TargetID:   history.ID,
		TargetType: def.TargetTypeArticleHistory,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("account-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, p.d.DB(), info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
		}
		PromError("account-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, p.d.DB(), info.ActorID); err != nil {
		PromError("account-feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityLike {
		m.Ack()
		return
	}

	feed := &model.AccountFeed{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		ActionType: def.ActionTypeLikeArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextLikeArticle,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddAccountFeed(c, p.d.DB(), feed); err != nil {
		PromError("account-feed: AddAccountFeed", "AddAccountFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	m.Ack()
}
