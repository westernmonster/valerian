package service

import (
	"context"
	"time"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
)

// checkTopic 检测Topic是否存在
func (p *Service) checkTopic(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	return
}

func (p *Service) createTopic(c context.Context, node sqalx.Node, aid int64, arg *model.ArgCreateTopic) (id int64, err error) {
	item := &model.Topic{
		ID:              gid.NewID(),
		Name:            arg.Name,
		Introduction:    arg.Introduction,
		TopicHome:       model.TopicHomeFeed,
		IsPrivate:       false,
		AllowChat:       types.BitBool(arg.AllowChat),
		AllowDiscuss:    types.BitBool(arg.AllowDiscuss),
		ViewPermission:  model.ViewPermissionJoin,
		EditPermission:  model.EditPermissionAdmin,
		JoinPermission:  model.JoinPermissionMemberApprove,
		CatalogViewType: arg.CatalogViewType,
		CreatedBy:       aid,
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
	}

	if arg.Avatar != nil {
		item.Avatar = *arg.Avatar
	}

	if arg.Bg != nil {
		item.Bg = *arg.Bg
	}

	if err = p.d.AddTopic(c, node, item); err != nil {
		return
	}

	setting := &model.AccountTopicSetting{
		ID:               gid.NewID(),
		AccountID:        aid,
		TopicID:          item.ID,
		Important:        types.BitBool(false),
		MuteNotification: types.BitBool(false),
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}

	if err = p.d.AddAccountTopicSetting(c, node, setting); err != nil {
		return
	}

	if err = p.createOwner(c, node, aid, item.ID); err != nil {
		return
	}

	if err = p.d.AddTopicStat(c, node, &model.TopicStat{
		TopicID:     item.ID,
		MemberCount: 1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.d.IncrAccountStat(c, node, &model.AccountStat{AccountID: item.CreatedBy, TopicCount: 1}); err != nil {
		return
	}

	return item.ID, nil

}

func (p *Service) updateTopic(c context.Context, node sqalx.Node, aid int64, arg *model.ArgUpdateTopic) (err error) {
	var isAdmin bool
	if isAdmin, err = p.isTopicMemberAdmin(c, node, arg.ID, aid); err != nil {
		return
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, node, arg.ID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	var s *model.AccountTopicSetting
	if s, err = p.d.GetAccountTopicSettingByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": t.ID}); err != nil {
		return
	} else if s == nil {
		err = ecode.AccountTopicSettingNotExist
		return
	}

	if arg.Important != nil {
		s.Important = types.BitBool(*arg.Important)
	}

	if arg.MuteNotification != nil {
		s.MuteNotification = types.BitBool(*arg.MuteNotification)
	}

	if err = p.d.UpdateAccountTopicSetting(c, node, s); err != nil {
		return
	}

	// 如果不是管理员，则不更新话题内容
	if !isAdmin {
		return
	}

	if arg.Avatar != nil && *arg.Avatar != "" {
		t.Avatar = *arg.Avatar
	}

	if arg.Bg != nil && *arg.Bg != "" {
		t.Bg = *arg.Bg
	}

	if arg.Name != nil && *arg.Name != "" {
		t.Name = *arg.Name
	}

	if arg.Introduction != nil && *arg.Introduction != "" {
		t.Introduction = *arg.Introduction
	}

	if arg.JoinPermission != nil && *arg.JoinPermission != "" {
		t.JoinPermission = *arg.JoinPermission
	}

	if arg.EditPermission != nil && *arg.EditPermission != "" {
		t.EditPermission = *arg.EditPermission
	}

	if arg.ViewPermission != nil && *arg.ViewPermission != "" {
		t.ViewPermission = *arg.ViewPermission
	}

	if arg.CatalogViewType != nil && *arg.CatalogViewType != "" {
		t.CatalogViewType = *arg.CatalogViewType
	}

	if arg.IsPrivate != nil {
		t.IsPrivate = types.BitBool(*arg.IsPrivate)
	}

	if arg.AllowChat != nil {
		t.AllowChat = types.BitBool(*arg.AllowChat)
	}

	if arg.AllowDiscuss != nil {
		t.AllowDiscuss = types.BitBool(*arg.AllowDiscuss)
	}

	if err = p.d.UpdateTopic(c, node, t); err != nil {
		return
	}

	return
}

func (p *Service) getAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error) {
	var addCache = true
	if item, err = p.d.AccountTopicSettingCache(c, aid, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetAccountTopicSettingByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if item == nil {
		item = &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          topicID,
			Important:        false,
			MuteNotification: false,
			Fav:              false,
			CreatedAt:        time.Now().Unix(),
			UpdatedAt:        time.Now().Unix(),
		}
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAccountTopicSettingCache(context.TODO(), item)
		})
	}

	return
}

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.Topic, err error) {
	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if item == nil {
		return nil, ecode.TopicNotExist
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicCache(context.TODO(), item)
		})
	}

	return
}

func (p *Service) getTopicResp(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.TopicResp, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, node, topicID); err != nil {
		return
	}
	item = &model.TopicResp{
		ID:              t.ID,
		Avatar:          t.Avatar,
		Bg:              t.Bg,
		Name:            t.Name,
		Introduction:    t.Introduction,
		CatalogViewType: t.CatalogViewType,
		TopicHome:       t.TopicHome,
		IsPrivate:       bool(t.IsPrivate),
		AllowChat:       bool(t.AllowChat),
		AllowDiscuss:    bool(t.AllowDiscuss),
		ViewPermission:  t.ViewPermission,
		EditPermission:  t.EditPermission,
		JoinPermission:  t.JoinPermission,
		CreatedAt:       t.CreatedAt,
	}

	var s *model.AccountTopicSetting
	if s, err = p.getAccountTopicSetting(c, node, aid, topicID); err != nil {
		return
	}

	item.Important = bool(s.Important)
	item.MuteNotification = bool(s.MuteNotification)
	return
}
