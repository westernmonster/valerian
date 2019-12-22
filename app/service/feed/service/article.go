package service

import (
	"context"
	"time"

	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
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

func (p *Service) onCatalogArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgCatalogArticleAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
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
		PromError("feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetTopicMemberIDs(c, tx, info.TopicID); err != nil {
		PromError("feed: GetTopicMemberIDs", "GetTopicMemberIDs(), topic_id(%d),error(%+v)", info.TopicID, err)
		return
	}

	for _, v := range ids {
		var data *model.Feed
		if data, err = p.d.GetFeedByCond(c, tx, map[string]interface{}{
			"account_id":  v,
			"action_type": def.ActionTypeCreateArticle,
			"target_id":   article.ID,
		}); err != nil {
			PromError("feed: GetFeedByCond", "GetFeedByCond(), aid(%d), article_id(%d), action_type(%s),error(%+v)", v, article.ID, def.ActionTypeCreateArticle, err)
			return
		} else if data != nil {
			continue
		}

		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCreateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateArticle,
			ActorID:    article.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
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
		PromError("feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCreateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateArticle,
			ActorID:    article.CreatedBy,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
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
		PromError("feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeUpdateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Commit", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	// 强制使用强制使用Master库
	c = sqalx.NewContext(c, true)
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		PromError("feed: Unmarshal data", "info.Umarshal() ,error(%+v)", err)
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		PromError("feed: tx.Beginx", "tx.Beginx(), error(%+v)", err)
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				PromError("feed: tx.Rollback", "tx.Rollback(), error(%+v)", err)
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
		PromError("feed: GetArticle", "GetArticle(), id(%d),error(%+v)", info.ArticleID, err)
		return
	}

	var setting *model.SettingResp
	if setting, err = p.getAccountSetting(c, tx, info.ActorID); err != nil {
		PromError("feed: getAccountSetting", "getAccountSetting(), id(%d),error(%+v)", info.ActorID, err)
		return
	}

	if !setting.ActivityLike {
		m.Ack()
		return
	}

	var ids []int64
	if ids, err = p.d.GetFansIDs(c, tx, info.ActorID); err != nil {
		PromError("feed: GetFansIDs", "GetFansIDs(), aid(%d),error(%+v)", info.ActorID, err)
		return
	}

	for _, v := range ids {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeLikeArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextLikeArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			PromError("feed: AddFeed", "AddFeed(), feed(%+v),error(%+v)", feed, err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		PromError("feed: tx.Rollback", "tx.Commit(), error(%+v)", err)
		return
	}

	m.Ack()
}
