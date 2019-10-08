package service

import (
	"context"
	"time"
	article "valerian/app/service/article/api"
	"valerian/app/service/topic-feed/model"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	info := new(model.MsgArticleAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(context.Background(), info.ID); err != nil {
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    info.TopicID,
		ActionType: model.ActionTypeCreateArticle,
		ActionTime: time.Now().Unix(),
		ActionText: model.ActionTextCreateArticle,
		ActorID:    article.Creator.ID,
		ActorType:  model.ActorTypeUser,
		TargetID:   article.ID,
		TargetType: model.TargetTypeArticle,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(context.Background(), p.d.DB(), feed); err != nil {
		log.Errorf("service.AddTopicFeed() failed %#v", err)
		return
	}

	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	info := new(model.MsgArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelTopicFeedByCond(context.Background(), p.d.DB(), info.TopicID, model.TargetTypeArticle, info.ID); err != nil {
		log.Errorf("service.DelTopicFeedByCond() failed %#v", err)
		return
	}

	m.Ack()
}
