package service

import (
	"context"
	"fmt"
	"time"
	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgCatalogArticleAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ArticleID); err != nil {
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: def.ActionTypeCreateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: def.ActionTextCreateArticle,
		ActorID:    article.Creator.ID,
		ActorType:  def.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: def.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.onArticleAdded() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	c := context.Background()

	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseUpdated Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
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

	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCond(c, tx, map[string]interface{}{
		"type":   model.TopicCatalogArticle,
		"ref_id": article.ID,
	}); err != nil {
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
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddTopicFeed(context.Background(), p.d.DB(), feed); err != nil {
			log.Errorf("service.onArticleUpdated() failed %#v", err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgCatalogArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelTopicFeedByCond(context.Background(), p.d.DB(), info.TopicID, def.TargetTypeArticle, info.ArticleID); err != nil {
		log.Errorf("service.DelTopicFeedByCond() failed %#v", err)
		return
	}

	m.Ack()
}
