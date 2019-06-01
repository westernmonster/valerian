package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/ecode"
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

func (p *Service) GetTopic(c context.Context, topicID int64) (item *model.TopicResp, err error) {
	item, _ = p.d.TopicCache(c, topicID)

	if item == nil {
		if item, err = p.getTopic(c, topicID); err != nil {
			return
		}
	}

	if item.TopicMeta, err = p.GetTopicMeta(c, item); err != nil {
		return
	}

	return
}

func (p *Service) getTopic(c context.Context, topicID int64) (item *model.TopicResp, err error) {
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

	if item.MembersCount, item.Members, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
		return
	}

	if item.Versions, err = p.d.GetTopicVersions(c, p.d.DB(), t.TopicSetID); err != nil {
		return
	}

	if item.RelatedTopics, err = p.d.GetAllRelatedTopics(c, p.d.DB(), topicID); err != nil {
		return
	}

	if item.Catalogs, err = p.getCatalogHierarchyOfAll(c, p.d.DB(), topicID); err != nil {
		return
	}

	return
}
