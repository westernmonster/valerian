package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded Unmarshal failed %#v", err))
		return
	}

	var v *topic.TopicInfo
	if v, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetTopicByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESTopic{
		ID:              v.ID,
		Name:            &v.Name,
		Avatar:          &v.Avatar,
		Bg:              &v.Bg,
		Introduction:    &v.Introduction,
		ViewPermission:  &v.ViewPermission,
		EditPermission:  &v.EditPermission,
		JoinPermission:  &v.JoinPermission,
		CatalogViewType: &v.CatalogViewType,
		CreatedAt:       &v.CreatedAt,
		UpdatedAt:       &v.UpdatedAt,
		AllowDiscuss:    &v.AllowDiscuss,
		AllowChat:       &v.AllowChat,
		IsPrivate:       &v.IsPrivate,
	}

	item.Creator = &model.ESCreator{
		ID:           v.Creator.ID,
		UserName:     &v.Creator.UserName,
		Avatar:       &v.Creator.Avatar,
		Introduction: &v.Creator.Introduction,
	}

	if err = p.d.PutTopic2ES(c, item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onTopicUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicUpdated Unmarshal failed %#v", err))
		return
	}

	var v *topic.TopicInfo
	if v, err = p.d.GetTopic(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetTopicByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESTopic{
		ID:              v.ID,
		Name:            &v.Name,
		Avatar:          &v.Avatar,
		Bg:              &v.Bg,
		Introduction:    &v.Introduction,
		ViewPermission:  &v.ViewPermission,
		EditPermission:  &v.EditPermission,
		JoinPermission:  &v.JoinPermission,
		CatalogViewType: &v.CatalogViewType,
		CreatedAt:       &v.CreatedAt,
		UpdatedAt:       &v.UpdatedAt,
		AllowDiscuss:    &v.AllowDiscuss,
		AllowChat:       &v.AllowChat,
		IsPrivate:       &v.IsPrivate,
	}

	item.Creator = &model.ESCreator{
		ID:           v.Creator.ID,
		UserName:     &v.Creator.UserName,
		Avatar:       &v.Creator.Avatar,
		Introduction: &v.Creator.Introduction,
	}

	if err = p.d.PutTopic2ES(c, item); err != nil {
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

	if err = p.d.DelESTopic(c, info.TopicID); err != nil {
		return
	}

	m.Ack()
}
