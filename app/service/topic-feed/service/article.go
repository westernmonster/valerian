package service

import (
	"context"
	"time"
	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic-feed/model"
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
