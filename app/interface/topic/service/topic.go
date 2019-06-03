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

func (p *Service) FollowTopic(c context.Context, mid int64, topicID int64) (err error) {
	return
}

func (p *Service) SearchTopics(c context.Context, query string) (err error) {
	return
}
func (p *Service) GetAllRelatedTopics(c context.Context, topicID int64) (items []*model.RelatedTopicResp, err error) {
	return
}

func (p *Service) DeleteTopic(c context.Context, topicID int64) (err error) {
	return
}

func (p *Service) GetTopic(c context.Context, topicID int64) (item *model.TopicResp, err error) {

	fmt.Println(topicID)
	if item, err = p.getTopic(c, topicID); err != nil {
		return
	}

	if item.MembersCount, item.Members, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
		return
	}

	if item.Versions, err = p.getTopicVersions(c, p.d.DB(), item.TopicSetID); err != nil {
		return
	}

	if item.RelatedTopics, err = p.getAllRelatedTopics(c, p.d.DB(), topicID); err != nil {
		return
	}

	// if item.Catalogs, err = p.GetCatalogHierarchyOfAll(c, p.d.DB(), topicID); err != nil {
	// 	return
	// }

	if item.TopicMeta, err = p.GetTopicMeta(c, item); err != nil {
		return
	}

	return
}

func (p *Service) getTopic(c context.Context, topicID int64) (item *model.TopicResp, err error) {
	var addCache = true
	if item, err = p.d.TopicCache(c, topicID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	fmt.Println(topicID)
	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, p.d.DB(), topicID); err != nil {
		return
	} else if t == nil {
		return nil, ecode.TopicNotExist
	}

	item = &model.TopicResp{
		ID:               t.ID,
		TopicSetID:       t.TopicSetID,
		Cover:            t.Cover,
		Bg:               t.Bg,
		Name:             t.Name,
		Introduction:     t.Introduction,
		CatalogViewType:  t.CatalogViewType,
		TopicType:        t.TopicType,
		TopicHome:        t.TopicHome,
		VersionName:      t.VersionName,
		IsPrivate:        bool(t.IsPrivate),
		AllowChat:        bool(t.AllowChat),
		AllowDiscuss:     bool(t.AllowDiscuss),
		EditPermission:   t.EditPermission,
		ViewPermission:   t.ViewPermission,
		JoinPermission:   t.JoinPermission,
		Important:        bool(t.Important),
		MuteNotification: bool(t.MuteNotification),
		CreatedAt:        t.CreatedAt,
	}

	item.Members = make([]*model.TopicMemberResp, 0)
	item.RelatedTopics = make([]*model.RelatedTopicShort, 0)
	item.Catalogs = make([]*model.TopicLevel1Catalog, 0)
	item.Versions = make([]*model.TopicVersionResp, 0)

	var tType *model.TopicType
	if tType, err = p.d.GetTopicType(c, p.d.DB(), t.TopicType); err != nil {
		return
	} else if tType != nil {
		item.TopicTypeName = tType.Name
	}

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
		ID:               gid.NewID(),
		Name:             arg.Name,
		Cover:            arg.Cover,
		Bg:               arg.Bg,
		Introduction:     arg.Introduction,
		IsPrivate:        types.BitBool(arg.IsPrivate),
		AllowChat:        types.BitBool(arg.AllowChat),
		AllowDiscuss:     types.BitBool(arg.AllowDiscuss),
		EditPermission:   arg.EditPermission,
		ViewPermission:   arg.ViewPermission,
		JoinPermission:   arg.JoinPermission,
		Important:        types.BitBool(arg.Important),
		MuteNotification: types.BitBool(arg.MuteNotification),
		CatalogViewType:  arg.CatalogViewType,
		TopicHome:        arg.TopicHome,
		VersionName:      arg.VersionName,
		CreatedBy:        aid,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}

	if v, e := p.d.GetTopicType(c, tx, arg.TopicType); e != nil {
		err = e
		return
	} else if v == nil {
		err = ecode.TopicTypeNotExist
		return
	}

	item.TopicType = arg.TopicType

	if arg.TopicSetID != nil {
		if v, e := p.d.GetTopicVersionByName(c, tx, *arg.TopicSetID, arg.VersionName); e != nil {
			err = e
			return
		} else if v != nil {
			err = ecode.TopicVersionNameExist
			return
		}

		item.TopicSetID = *arg.TopicSetID
	} else {
		set := &model.TopicSet{
			ID:        gid.NewID(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		item.TopicSetID = set.ID

		if err = p.d.AddTopicSet(c, tx, set); err != nil {
			return
		}
	}

	if err = p.d.AddTopic(c, tx, item); err != nil {
		return
	}

	if err = p.bulkCreateCatalogs(c, tx, item.ID, arg.Catalogs); err != nil {
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
		p.d.DelTopicVersionCache(context.TODO(), item.TopicSetID)
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

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCondition(c, p.d.DB(), arg.ID, aid); err != nil {
		return
	} else if member == nil {
		return ecode.NotBelongToTopic
	} else if member.Role != model.MemberRoleAdmin && member.Role != model.MemberRoleAdmin {
		return ecode.NotTopicAdmin
	}

	return p.updateTopic(c, p.d.DB(), arg)
}

func (p *Service) updateTopic(c context.Context, node sqalx.Node, arg *model.ArgUpdateTopic) (err error) {
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

	if arg.VersionName != nil && *arg.VersionName != "" {
		var ver *model.TopicVersionResp
		if ver, err = p.d.GetTopicVersionByName(c, tx, t.TopicSetID, *arg.VersionName); err != nil {
			return
		} else if ver != nil && ver.TopicID != t.ID {
			return ecode.TopicVersionNameExist
		}

		t.VersionName = *arg.VersionName
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

	if arg.Important != nil {
		t.Important = types.BitBool(*arg.Important)
	}

	if arg.MuteNotification != nil {
		t.MuteNotification = types.BitBool(*arg.MuteNotification)
	}

	if err = p.d.UpdateTopic(c, tx, t); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelTopicCache(context.TODO(), arg.ID)
	})
	return
}
