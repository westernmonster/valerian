package service

import (
	"context"
	"fmt"
	"valerian/app/service/fav/model"
	"valerian/app/service/feed/def"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onArticleFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleFaved Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleFaved GetArticle failed %#v", err))
		return
	}

	item := &model.Fav{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.ArticleID,
		TargetType: model.TargetTypeArticle,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}

	if err = p.d.AddFav(c, p.d.DB(), item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionFaved GetDiscussion failed %#v", err))
		return
	}

	item := &model.Fav{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.DiscussionID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}

	if err = p.d.AddFav(c, p.d.DB(), item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onTopicFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFaved Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicFaved GetTopic failed %#v", err))
		return
	}

	item := &model.Fav{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.TopicID,
		TargetType: model.TargetTypeTopic,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}

	if err = p.d.AddFav(c, p.d.DB(), item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onReviseFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseFaved Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseFaved GetRevise failed %#v", err))
		return
	}

	item := &model.Fav{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.ReviseID,
		TargetType: model.TargetTypeRevise,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}

	if err = p.d.AddFav(c, p.d.DB(), item); err != nil {
		return
	}

	m.Ack()
}
