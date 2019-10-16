package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/app/service/recent/model"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) GetUserRecentViewsPaged(c context.Context, aid int64, limit, offset int) (items []*model.RecentView, err error) {
	if items, err = p.d.GetUserRecentViewsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) onArticleViewed(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleViewed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleViewed Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleViewed GetArticle failed %#v", err))
		return
	}

	if err = p.d.AddRecentView(c, p.d.DB(), &model.RecentView{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.ArticleID,
		TargetType: model.TargetTypeArticle,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onTopicViewed(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicViewed)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicViewed Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicViewed GetTopic failed %#v", err))
		return
	}

	if err = p.d.AddRecentView(c, p.d.DB(), &model.RecentView{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.TopicID,
		TargetType: model.TargetTypeTopic,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}); err != nil {
		return
	}

	m.Ack()
}
