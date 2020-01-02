package service

import (
	"context"
	"time"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic-feed/model"
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
	info := new(def.MsgCatalogArticleAdded)
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("topic-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("topic-feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("topic-feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()
	var history *model.ArticleHistory
	if history, err = p.getArticleHistory(c, tx, info.ArticleHistoryID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("topic-feed: GetArticleHistory", "GetArticleHistory(), id(%d),error(%+v)", info.ArticleHistoryID, err)
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, tx, info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("topic-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: def.ActionTypeCreateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateArticle,
		ActorID:    article.CreatedBy,
		ActorType:  def.ActorTypeUser,
		TargetID:   history.ID,
		TargetType: def.TargetTypeArticleHistory,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(context.Background(), tx, feed); err != nil {
		PromError("topic-feed: AddTopicFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("topic-feed: tx.Commit", "tx.Commit(), error(%+v)", err)
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
		PromError("topic-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("topic-feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("topic-feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var history *model.ArticleHistory
	if history, err = p.getArticleHistory(c, tx, info.ArticleHistoryID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("topic-feed: GetArticleHistory", "GetArticleHistory(), id(%d),error(%+v)", info.ArticleHistoryID, err)
		return
	}

	var article *model.Article
	if article, err = p.getArticle(c, tx, info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("topic-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCond(c, tx, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": article.ID,
	}); err != nil {
		PromError("topic-feed: GetTopicCatalogsByCond", "GetTopicCatalogsByCond(), type(%s),ref_id(%d),error(%+v)", model.TopicCatalogArticle, article.ID, err)
		return
	}

	for _, v := range catalogs {
		feed := &model.TopicFeed{
			ID:         gid.NewID(),
			TopicID:    v.TopicID,
			ActionType: def.ActionTypeUpdateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   history.ID,
			TargetType: def.TargetTypeArticleHistory,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddTopicFeed(context.Background(), tx, feed); err != nil {
			PromError("topic-feed: AddTopicFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("topic-feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}
	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgCatalogArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("topic-feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("topic-feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("topic-feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
			}
			return
		}
	}()

	var article *model.Article
	if article, err = p.getArticle(c, tx, info.ArticleID); err != nil {
		if ecode.IsNotExistEcode(err) {
			m.Ack()
			return
		}
		PromError("topic-feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: def.ActionTypeDeleteArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextDeleteArticle,
		ActorID:    info.ActorID,
		ActorType:  def.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(context.Background(), tx, feed); err != nil {
		PromError("topic-feed: AddTopicFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
		return
	}

	if err = tx.Commit(); err != nil {
		PromError("topic-feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}
