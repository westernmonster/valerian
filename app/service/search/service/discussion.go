package service

import (
	"context"
	"fmt"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	topic "valerian/app/service/topic/api"
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

	var v *discuss.DiscussionInfo
	if v, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionAdded GetDiscussionByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESDiscussion{
		ID:          v.ID,
		Title:       &v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}

	item.Creator = &model.ESCreator{
		ID:           v.Creator.ID,
		UserName:     &v.Creator.UserName,
		Avatar:       &v.Creator.Avatar,
		Introduction: &v.Creator.Introduction,
	}

	var t *topic.TopicInfo
	if t, err = p.d.GetTopic(c, v.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	item.Topic = &model.ESDiscussionTopic{
		ID:           t.ID,
		Name:         &t.Name,
		Avatar:       &t.Avatar,
		Introduction: &t.Introduction,
	}

	if v.CategoryID != -1 {
		item.Category = &model.ESDiscussionCategory{
			ID:   v.CategoryInfo.ID,
			Name: &v.CategoryInfo.Name,
			Seq:  &v.CategoryInfo.Seq,
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

	var v *discuss.DiscussionInfo
	if v, err = p.d.GetDiscussion(c, info.DiscussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onDiscussionUpdated GetDiscussionByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESDiscussion{
		ID:          v.ID,
		Title:       &v.Title,
		Content:     &v.Content,
		ContentText: &v.ContentText,
		CreatedAt:   &v.CreatedAt,
		UpdatedAt:   &v.UpdatedAt,
	}

	item.Creator = &model.ESCreator{
		ID:           v.Creator.ID,
		UserName:     &v.Creator.UserName,
		Avatar:       &v.Creator.Avatar,
		Introduction: &v.Creator.Introduction,
	}

	var t *topic.TopicInfo
	if t, err = p.d.GetTopic(c, v.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	item.Topic = &model.ESDiscussionTopic{
		ID:           t.ID,
		Name:         &t.Name,
		Avatar:       &t.Avatar,
		Introduction: &t.Introduction,
	}

	if v.CategoryID != -1 {
		item.Category = &model.ESDiscussionCategory{
			ID:   v.CategoryInfo.ID,
			Name: &v.CategoryInfo.Name,
			Seq:  &v.CategoryInfo.Seq,
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
