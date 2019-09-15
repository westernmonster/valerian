package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) checkTopic(c context.Context, node sqalx.Node, topicID int64) (err error) {
	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	return
}

func (p *Service) CreateTopic(c context.Context, arg *model.ArgCreateTopic) (topicID int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	item := &model.Topic{
		ID:              gid.NewID(),
		Name:            arg.Name,
		Avatar:          arg.Avatar,
		Bg:              arg.Bg,
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

	if err = p.d.AddTopic(c, tx, item); err != nil {
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

	if err = p.d.AddAccountTopicSetting(c, tx, setting); err != nil {
		return
	}

	if err = p.createTopicMemberOwner(c, tx, aid, item.ID); err != nil {
		return
	}

	if err = p.d.AddTopicMemberStat(c, tx, &model.TopicMemberStat{
		TopicID:     item.ID,
		MemberCount: 1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	topicID = item.ID
	return
}

func (p *Service) UpdateTopic(c context.Context, arg *model.ArgUpdateTopic) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.checkTopic(c, p.d.DB(), arg.ID); err != nil {
		return
	}

	if err = p.checkTopicMemberAdmin(c, p.d.DB(), arg.ID, aid); err != nil {
		return
	}

	return p.updateTopic(c, p.d.DB(), aid, arg)
}

func (p *Service) updateTopic(c context.Context, node sqalx.Node, aid int64, arg *model.ArgUpdateTopic) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, arg.ID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	if arg.Avatar != nil && *arg.Avatar != "" {
		t.Avatar = arg.Avatar
	}

	if arg.Bg != nil && *arg.Bg != "" {
		t.Bg = arg.Bg
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

	important := false
	muteNotification := false
	if arg.Important != nil {
		important = *arg.Important
	}

	if arg.MuteNotification != nil {
		muteNotification = *arg.MuteNotification
	}

	var s *model.AccountTopicSetting
	if s, err = p.d.GetAccountTopicSettingByCond(c, tx, map[string]interface{}{"account_id": aid, "topic_id": t.ID}); err != nil {
		return
	} else if s == nil {
		setting := &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          t.ID,
			Important:        types.BitBool(important),
			Fav:              types.BitBool(false),
			MuteNotification: types.BitBool(muteNotification),
			CreatedAt:        time.Now().Unix(),
			UpdatedAt:        time.Now().Unix(),
		}

		if err = p.d.AddAccountTopicSetting(c, tx, setting); err != nil {
			return
		}
	} else {
		s.Important = types.BitBool(important)
		s.MuteNotification = types.BitBool(muteNotification)
		if err = p.d.UpdateAccountTopicSetting(c, tx, s); err != nil {
			return
		}
	}

	if err = p.d.UpdateTopic(c, tx, t); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelAccountTopicSettingCache(context.TODO(), aid, t.ID)
		p.d.DelTopicCache(context.TODO(), arg.ID)
	})
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

func (p *Service) getTopicResp(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicResp, err error) {
	var t *model.Topic
	if t, err = p.getTopic(c, node, topicID); err != nil {
		return
	}

	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
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

func (p *Service) GetTopic(c context.Context, topicID int64, include string) (item *model.TopicResp, err error) {
	if item, err = p.getTopicResp(c, p.d.DB(), topicID); err != nil {
		return
	}
	inc := includeParam(include)
	if inc["members"] {
		if item.MembersCount, item.Members, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
			return
		}
	}

	if inc["catalogs"] {
		if item.Catalogs, err = p.getCatalogsHierarchy(c, p.d.DB(), topicID); err != nil {
			return
		}
	}

	if inc["discuss_categories"] {
		if item.DiscussCategories, err = p.getDiscussCategories(c, p.d.DB(), topicID); err != nil {
			return
		}
	}

	if inc["auth_topics"] {
		// if item.AuthTopics, err = p.getAuthTopics(c, p.d.DB(), topicID); err != nil {
		// 	return
		// }
	}

	if inc["meta"] {
		if item.TopicMeta, err = p.GetTopicMeta(c, item); err != nil {
			return
		}
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

func (p *Service) DelTopic(c context.Context, topicID int64) (err error) {
	return
}

func (p *Service) FavTopic(c context.Context, topicID int64) (faved bool, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var item *model.AccountTopicSetting
	if item, err = p.d.GetAccountTopicSettingByCond(c, tx, map[string]interface{}{"account_id": aid, "topic_id": topicID}); err != nil {
		return
	} else if item == nil {
		item = &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          topicID,
			Important:        false,
			MuteNotification: false,
			Fav:              true,
			CreatedAt:        time.Now().Unix(),
			UpdatedAt:        time.Now().Unix(),
		}

		faved = true
		if err = p.d.AddAccountTopicSetting(c, tx, item); err != nil {
			return
		}
	} else {
		item.Fav = !item.Fav
		faved = bool(item.Fav)

		if err = p.d.UpdateAccountTopicSetting(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}
