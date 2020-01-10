package service

import (
	"context"
	"time"
	"valerian/app/service/topic/api"
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

// createTopic 创建话题
func (p *Service) createTopic(c context.Context, node sqalx.Node, aid int64, arg *api.ArgCreateTopic) (id int64, err error) {
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
		item.Avatar = arg.GetAvatarValue()
	}

	if arg.Bg != nil {
		item.Bg = arg.GetBgValue()
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

// updateTopic 更新话题
func (p *Service) updateTopic(c context.Context, node sqalx.Node, aid int64, arg *api.ArgUpdateTopic) (err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, node, arg.ID); err != nil {
		return
	}

	var isTopicMember bool
	if isTopicMember, err = p.isTopicMember(c, node, aid, arg.ID); err != nil {
		return
	}

	if isTopicMember {
		var s *model.AccountTopicSetting
		if s, err = p.d.GetAccountTopicSettingByCond(c, node, map[string]interface{}{"account_id": aid, "topic_id": t.ID}); err != nil {
			return
		} else if s == nil {
			err = ecode.AccountTopicSettingNotExist
			return
		}

		if arg.Important != nil {
			s.Important = types.BitBool(arg.GetImportantValue())
		}

		if arg.MuteNotification != nil {
			s.MuteNotification = types.BitBool(arg.GetMuteNotificationValue())
		}

		if err = p.d.UpdateAccountTopicSetting(c, node, s); err != nil {
			return
		}
	}

	// 如果不是管理员，则不更新话题内容
	var hasTopicManagePermission bool
	if hasTopicManagePermission, err = p.hasTopicManagePermission(c, node, aid, arg.ID); err != nil {
		return
	} else if !hasTopicManagePermission {
		return
	}

	if arg.Avatar != nil && arg.GetAvatarValue() != "" {
		t.Avatar = arg.GetAvatarValue()
	}

	if arg.Bg != nil && arg.GetBgValue() != "" {
		t.Bg = arg.GetBgValue()
	}

	if arg.Name != nil && arg.GetNameValue() != "" {
		t.Name = arg.GetNameValue()
	}

	if arg.Introduction != nil && arg.GetIntroductionValue() != "" {
		t.Introduction = arg.GetIntroductionValue()
	}

	if arg.JoinPermission != nil && arg.GetJoinPermissionValue() != "" {
		t.JoinPermission = arg.GetJoinPermissionValue()
	}

	if arg.EditPermission != nil && arg.GetEditPermissionValue() != "" {
		t.EditPermission = arg.GetEditPermissionValue()
	}

	if arg.ViewPermission != nil && arg.GetViewPermissionValue() != "" {
		t.ViewPermission = arg.GetViewPermissionValue()
	}

	if arg.CatalogViewType != nil && arg.GetCatalogViewTypeValue() != "" {
		t.CatalogViewType = arg.GetCatalogViewTypeValue()
	}

	if arg.IsPrivate != nil {
		t.IsPrivate = types.BitBool(arg.GetIsPrivateValue())
	}

	if arg.AllowChat != nil {
		t.AllowChat = types.BitBool(arg.GetAllowChatValue())
	}

	if arg.AllowDiscuss != nil {
		t.AllowDiscuss = types.BitBool(arg.GetAllowDiscussValue())
	}

	t.UpdatedAt = time.Now().Unix()

	if err = p.d.UpdateTopic(c, node, t); err != nil {
		return
	}

	return
}

// getAccountTopicSetting 获取用户话题设置
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
			p.d.SetAccountTopicSettingCache(context.Background(), item)
		})
	}

	return
}

// getTopic 获取话题信息
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
			p.d.SetTopicCache(context.Background(), item)
		})
	}

	return
}
