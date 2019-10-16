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

func (p *Service) GetUserRecentPubsPaged(c context.Context, aid int64, limit, offset int) (items []*model.RecentPub, err error) {
	if items, err = p.d.GetUserRecentPubsPaged(c, p.d.DB(), aid, limit, offset); err != nil {
		return
	}

	return
}

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticle failed %#v", err))
		return
	}

	if err = p.d.AddRecentPub(c, p.d.DB(), &model.RecentPub{
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

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetTopic failed %#v", err))
		return
	}

	if err = p.d.AddRecentPub(c, p.d.DB(), &model.RecentPub{
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

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded GetDiscussion failed %#v", err))
		return
	}

	if err = p.d.AddRecentPub(c, p.d.DB(), &model.RecentPub{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.DiscussionID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}); err != nil {
		return
	}

	m.Ack()
}
