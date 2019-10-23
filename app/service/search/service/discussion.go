package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	"valerian/library/ecode"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded Unmarshal failed %#v", err))
		return
	}

	var v *model.Discussion
	if v, err = p.d.GetDiscussionByID(c, p.d.DB(), info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded GetDiscussionByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESDiscussion{
		ID:          v.ID,
		Title:       v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}

	var acc *model.Account
	if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
		return
	} else if acc == nil {
		err = ecode.UserNotExist
		return
	}

	item.Creator = &model.ESCreator{
		ID:           acc.ID,
		UserName:     &acc.UserName,
		Avatar:       &acc.Avatar,
		Introduction: &acc.Introduction,
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, p.d.DB(), v.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	item.Topic = &model.ESDiscussionTopic{
		ID:           t.ID,
		Name:         &t.Name,
		Avatar:       t.Avatar,
		Introduction: &t.Introduction,
	}

	if v.CategoryID != -1 {
		var cate *model.DiscussCategory
		if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
			return
		} else if cate == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		item.Category = &model.ESDiscussionCategory{
			ID:   cate.ID,
			Name: &cate.Name,
			Seq:  &cate.Seq,
		}
	}

	if err = p.d.PutDiscussion2ES(c, item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgDiscussionUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated Unmarshal failed %#v", err))
		return
	}

	var v *model.Discussion
	if v, err = p.d.GetDiscussionByID(c, p.d.DB(), info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated GetDiscussionByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESDiscussion{
		ID:          v.ID,
		Title:       v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}

	var acc *model.Account
	if acc, err = p.d.GetAccountByID(c, p.d.DB(), v.CreatedBy); err != nil {
		return
	} else if acc == nil {
		err = ecode.UserNotExist
		return
	}

	item.Creator = &model.ESCreator{
		ID:           acc.ID,
		UserName:     &acc.UserName,
		Avatar:       &acc.Avatar,
		Introduction: &acc.Introduction,
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, p.d.DB(), v.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	item.Topic = &model.ESDiscussionTopic{
		ID:           t.ID,
		Name:         &t.Name,
		Avatar:       t.Avatar,
		Introduction: &t.Introduction,
	}

	if v.CategoryID != -1 {
		var cate *model.DiscussCategory
		if cate, err = p.d.GetDiscussCategoryByID(c, p.d.DB(), v.CategoryID); err != nil {
			return
		} else if cate == nil {
			err = ecode.DiscussCategoryNotExist
			return
		}

		item.Category = &model.ESDiscussionCategory{
			ID:   cate.ID,
			Name: &cate.Name,
			Seq:  &cate.Seq,
		}
	}

	if err = p.d.PutDiscussion2ES(c, item); err != nil {
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
	m.Ack()

	if err = p.d.DelESDiscussion(c, info.DiscussionID); err != nil {
		return
	}

}
