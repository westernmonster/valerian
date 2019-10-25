package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/app/service/recent/model"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) GetUserRecentPubsPaged(c context.Context, aid int64, targetType string, limit, offset int) (items []*model.RecentPub, err error) {
	if items, err = p.d.GetUserRecentPubsPaged(c, p.d.DB(), aid, targetType, limit, offset); err != nil {
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
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var data *model.RecentPub
	if data, err = p.d.GetRecentPubByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeArticle,
		"target_id":   info.ArticleID,
	}); err != nil {
		return
	} else if data != nil {
		m.Ack()
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

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelRecentPubByCond(c, p.d.DB(), model.TargetTypeArticle, info.ArticleID); err != nil {
		return
	}

	if err = p.d.DelRecentViewByCond(c, p.d.DB(), model.TargetTypeArticle, info.ArticleID); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded Unmarshal failed %#v", err))
		return
	}

	if _, err = p.d.GetRevise(c, info.ReviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseAdded GetRevise failed %#v", err))
		if ecode.Cause(err) == ecode.ReviseNotExist {
			m.Ack()
		}
		return
	}

	var data *model.RecentPub
	if data, err = p.d.GetRecentPubByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeRevise,
		"target_id":   info.ReviseID,
	}); err != nil {
		return
	} else if data != nil {
		m.Ack()
		return
	}

	if err = p.d.AddRecentPub(c, p.d.DB(), &model.RecentPub{
		ID:         gid.NewID(),
		AccountID:  info.ActorID,
		TargetID:   info.ReviseID,
		TargetType: model.TargetTypeRevise,
		CreatedAt:  info.ActionTime,
		UpdatedAt:  info.ActionTime,
	}); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onReviseDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgReviseDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onReviseDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelRecentPubByCond(c, p.d.DB(), model.TargetTypeRevise, info.ReviseID); err != nil {
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
		if ecode.Cause(err) == ecode.DiscussionNotExist {
			m.Ack()
		}
		return
	}

	var data *model.RecentPub
	if data, err = p.d.GetRecentPubByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeDiscussion,
		"target_id":   info.DiscussionID,
	}); err != nil {
		return
	} else if data != nil {
		m.Ack()
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

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelRecentPubByCond(c, p.d.DB(), model.TargetTypeDiscussion, info.DiscussionID); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onTopicDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelRecentPubByCond(c, p.d.DB(), model.TargetTypeTopic, info.TopicID); err != nil {
		return
	}

	m.Ack()
}
