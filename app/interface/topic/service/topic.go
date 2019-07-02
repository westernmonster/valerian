package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) GetTopic(c context.Context, topicID int64, include string) (item *model.TopicResp, err error) {
	if item, err = p.getTopic(c, p.d.DB(), topicID); err != nil {
		return
	}
	inc := includeParam(include)
	if inc["members"] {
		if item.MembersCount, item.Members, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
			return
		}
	}

	if inc["versions"] {
		if item.Versions, err = p.getTopicVersionsResp(c, p.d.DB(), item.ID); err != nil {
			return
		}
	}

	if inc["related_topics"] {
		if item.RelatedTopics, err = p.getAllRelatedTopics(c, p.d.DB(), topicID); err != nil {
			return
		}
	}

	if inc["versions[*].catalogs"] {
		for _, v := range item.Versions {
			if v.Catalogs, err = p.getCatalogsHierarchy(c, p.d.DB(), v.ID); err != nil {
				return
			}
		}
	}

	if inc["related_topics[*].meta"] {
		for _, v := range item.RelatedTopics {
			var t *model.TopicResp
			if t, err = p.getTopic(c, p.d.DB(), v.TopicID); err != nil {
				return
			}
			if v.TopicMeta, err = p.GetTopicMeta(c, t); err != nil {
				return
			}
		}
	}

	if inc["meta"] {
		if item.TopicMeta, err = p.GetTopicMeta(c, item); err != nil {
			return
		}
	}

	return
}

func (p *Service) getTopic(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, node, topicID); err != nil {
		return
	} else if t == nil {
		return nil, ecode.TopicNotExist
	}

	item = &model.TopicResp{
		ID:              t.ID,
		Cover:           t.Cover,
		Bg:              t.Bg,
		Name:            t.Name,
		Introduction:    t.Introduction,
		CatalogViewType: t.CatalogViewType,
		TopicType:       t.TopicType,
		TopicHome:       t.TopicHome,
		IsPrivate:       bool(t.IsPrivate),
		AllowChat:       bool(t.AllowChat),
		AllowDiscuss:    bool(t.AllowDiscuss),
		EditPermission:  t.EditPermission,
		ViewPermission:  t.ViewPermission,
		JoinPermission:  t.JoinPermission,
		CreatedAt:       t.CreatedAt,
	}

	item.Members = make([]*model.TopicMemberResp, 0)
	item.RelatedTopics = make([]*model.RelatedTopicShort, 0)
	item.Versions = make([]*model.TopicVersionResp, 0)

	var tType *model.TopicType
	if tType, err = p.d.GetTopicType(c, node, t.TopicType); err != nil {
		return
	} else if tType != nil {
		item.TopicTypeName = tType.Name
	}
	var s *model.AccountTopicSetting
	if s, err = p.getAccountTopicSetting(c, node, aid, topicID); err != nil {
		return
	}

	item.Important = bool(s.Important)
	item.MuteNotification = bool(s.MuteNotification)

	if addCache {
		p.addCache(func() {
			p.d.SetTopicCache(context.TODO(), item)
		})
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
		Cover:           arg.Cover,
		Bg:              arg.Bg,
		Introduction:    arg.Introduction,
		IsPrivate:       types.BitBool(arg.IsPrivate),
		AllowChat:       types.BitBool(arg.AllowChat),
		AllowDiscuss:    types.BitBool(arg.AllowDiscuss),
		EditPermission:  arg.EditPermission,
		ViewPermission:  arg.ViewPermission,
		JoinPermission:  arg.JoinPermission,
		CatalogViewType: arg.CatalogViewType,
		TopicHome:       arg.TopicHome,
		CreatedBy:       aid,
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
	}

	if v, e := p.d.GetTopicType(c, tx, arg.TopicType); e != nil {
		err = e
		return
	} else if v == nil {
		err = ecode.TopicTypeNotExist
		return
	}

	item.TopicType = arg.TopicType

	for _, v := range arg.Versions {
		set := &model.TopicVersion{
			ID:        gid.NewID(),
			Name:      v.Name,
			Seq:       v.Seq,
			TopicID:   item.ID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddTopicVersion(c, tx, set); err != nil {
			return
		}

		if err = p.bulkCreateCatalogs(c, tx, item.ID, set.ID, v.Catalogs); err != nil {
			return
		}

	}

	if err = p.d.AddTopic(c, tx, item); err != nil {
		return
	}

	setting := &model.AccountTopicSetting{
		ID:               gid.NewID(),
		AccountID:        aid,
		TopicID:          item.ID,
		Important:        types.BitBool(arg.Important),
		MuteNotification: types.BitBool(arg.MuteNotification),
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}

	if err = p.d.AddAccountTopicSetting(c, tx, setting); err != nil {
		return
	}

	if err = p.bulkCreateMembers(c, tx, aid, item.ID, arg); err != nil {
		return
	}

	if err = p.bulkSaveRelations(c, tx, item.ID, arg.RelatedTopics); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicVersionCache(context.TODO(), item.ID)
	})

	topicID = item.ID
	return
}

func (p *Service) UpdateTopic(c context.Context, arg *model.ArgUpdateTopic) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	fmt.Println(aid)
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCondition(c, p.d.DB(), arg.ID, aid); err != nil {
		return
	} else if member == nil {
		return ecode.NotBelongToTopic
	} else if member.Role != model.MemberRoleAdmin && member.Role != model.MemberRoleOwner {
		return ecode.NotTopicAdmin
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

	if arg.TopicType != nil {
		var topicType *model.TopicType
		if topicType, err = p.d.GetTopicType(c, tx, *arg.TopicType); err != nil {
			return
		} else if topicType == nil {
			return ecode.TopicTypeNotExist
		}
		t.TopicType = *arg.TopicType

	}

	if arg.Cover != nil && *arg.Cover != "" {
		t.Cover = arg.Cover
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

	if arg.TopicHome != nil && *arg.TopicHome != "" {
		t.TopicHome = *arg.TopicHome
	}

	if arg.IsPrivate != nil {
		t.IsPrivate = types.BitBool(*arg.IsPrivate)
	}

	if arg.AllowChat != nil {
		t.AllowChat = types.BitBool(*arg.AllowChat)
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
	if s, err = p.d.GetAccountTopicSetting(c, tx, aid, t.ID); err != nil {
		return
	} else if s == nil {
		setting := &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          t.ID,
			Important:        types.BitBool(important),
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

func (p *Service) DelTopic(c context.Context, topicID int64) (err error) {
	return
}

func (p *Service) getAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error) {
	var addCache = true
	if item, err = p.d.AccountTopicSettingCache(c, aid, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetAccountTopicSetting(c, node, aid, topicID); err != nil {
		return
	} else if item == nil {
		item = &model.AccountTopicSetting{
			ID:               gid.NewID(),
			AccountID:        aid,
			TopicID:          topicID,
			Important:        false,
			MuteNotification: false,
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
