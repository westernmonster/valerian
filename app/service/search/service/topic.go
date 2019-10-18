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

func (p *Service) onTopicAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgTopicAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded Unmarshal failed %#v", err))
		return
	}

	var v *model.Topic
	if v, err = p.d.GetTopicByID(c, p.d.DB(), info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicAdded GetTopicByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESTopic{
		ID:              v.ID,
		Name:            &v.Name,
		Avatar:          v.Avatar,
		Bg:              v.Bg,
		Introduction:    &v.Introduction,
		ViewPermission:  &v.ViewPermission,
		EditPermission:  &v.EditPermission,
		JoinPermission:  &v.JoinPermission,
		CatalogViewType: &v.CatalogViewType,
		CreatedAt:       &v.CreatedAt,
		UpdatedAt:       &v.UpdatedAt,
	}

	allowDiscuss := bool(v.AllowDiscuss)
	allowChat := bool(v.AllowChat)
	isPrivate := bool(v.IsPrivate)
	item.AllowDiscuss = &allowDiscuss
	item.AllowChat = &allowChat
	item.IsPrivate = &isPrivate

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
		Introduction: acc.Introduction,
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

	var v *model.Topic
	if v, err = p.d.GetTopicByID(c, p.d.DB(), info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onTopicUpdated GetTopicByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESTopic{
		ID:              v.ID,
		Name:            &v.Name,
		Avatar:          v.Avatar,
		Bg:              v.Bg,
		Introduction:    &v.Introduction,
		ViewPermission:  &v.ViewPermission,
		EditPermission:  &v.EditPermission,
		JoinPermission:  &v.JoinPermission,
		CatalogViewType: &v.CatalogViewType,
		CreatedAt:       &v.CreatedAt,
		UpdatedAt:       &v.UpdatedAt,
	}

	allowDiscuss := bool(v.AllowDiscuss)
	allowChat := bool(v.AllowChat)
	isPrivate := bool(v.IsPrivate)
	item.AllowDiscuss = &allowDiscuss
	item.AllowChat = &allowChat
	item.IsPrivate = &isPrivate

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
		Introduction: acc.Introduction,
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
