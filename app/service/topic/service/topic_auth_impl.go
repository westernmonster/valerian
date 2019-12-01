package service

import (
	"context"
	"time"

	"valerian/app/service/topic/api"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
)

func (p *Service) isAuthTopic(c context.Context, node sqalx.Node, toTopicID, fromTopicID int64) (isAuth bool, err error) {
	var t *model.AuthTopic
	if t, err = p.d.GetAuthTopicByCond(c, node, map[string]interface{}{"to_topic_id": toTopicID, "topic_id": fromTopicID}); err != nil {
		return
	} else if t != nil {
		isAuth = true
	}

	return
}

func (p *Service) checkAuthTopics(c context.Context, node sqalx.Node, topicID int64, items []*api.ArgAuthTopic) (err error) {
	// must unique and not equal to current topic
	dic := make(map[int64]bool)
	for _, v := range items {
		if v.TopicID == topicID {
			err = ecode.MustNotUseCurrentTopic
			return
		}

		if dic[v.TopicID] {
			err = ecode.AuthTopicDuplicate
			return
		}

		if err = p.checkTopic(c, node, v.TopicID); err != nil {
			return
		}

		dic[v.TopicID] = true
	}

	return
}

func (p *Service) loadAuthTopicsMap(c context.Context, node sqalx.Node, topicID int64) (dic map[int64]bool, err error) {
	dic = make(map[int64]bool)
	var dbItems []*model.AuthTopic
	if dbItems, err = p.d.GetAuthTopicsByCond(c, node, map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	for _, v := range dbItems {
		dic[v.ID] = false
	}

	return
}

func (p *Service) bulkSaveAuthTopics(c context.Context, node sqalx.Node, topicID int64, items []*api.ArgAuthTopic) (err error) {
	if err = p.checkAuthTopics(c, node, topicID, items); err != nil {
		return
	}

	if err = p.checkTopic(c, node, topicID); err != nil {
		return
	}

	var dic map[int64]bool
	if dic, err = p.loadAuthTopicsMap(c, node, topicID); err != nil {
		return
	}

	for _, v := range items {
		var dItem *model.AuthTopic
		if dItem, err = p.d.GetAuthTopicByCond(c, node, map[string]interface{}{"topic_id": topicID, "to_topic_id": v.TopicID}); err != nil {
			return
		}

		if dItem == nil {
			item := &model.AuthTopic{
				ID:         gid.NewID(),
				TopicID:    topicID,
				ToTopicID:  v.TopicID,
				Permission: v.Permission,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddAuthTopic(c, node, item); err != nil {
				return
			}

			continue
		}

		dic[dItem.ID] = true
		dItem.Permission = v.Permission
		if err = p.d.UpdateAuthTopic(c, node, dItem); err != nil {
			return
		}
	}

	for k, used := range dic {
		if used {
			continue
		}
		if err = p.d.DelAuthTopic(c, node, k); err != nil {
			return
		}
	}

	return
}

func (p *Service) getAuthTopicsResp(c context.Context, node sqalx.Node, topicID int64) (items []*api.AuthTopicInfo, err error) {
	var data []*model.AuthTopic
	if data, err = p.getAuthTopics(c, node, topicID); err != nil {
		return
	}

	items = make([]*api.AuthTopicInfo, len(data))

	for i, v := range data {
		var t *model.Topic
		if t, err = p.getTopic(c, node, v.ToTopicID); err != nil {
			return
		}

		// TODO: need cache ?
		var stat *model.TopicStat
		if stat, err = p.d.GetTopicStatByID(c, node, topicID); err != nil {
			return
		}

		item := &api.AuthTopicInfo{
			ToTopicID:      v.ToTopicID,
			Permission:     v.Permission,
			Avatar:         t.Avatar,
			Name:           t.Name,
			MemberCount:    stat.MemberCount,
			EditPermission: t.EditPermission,
		}

		items[i] = item
	}

	return
}

func (p *Service) getAuthTopics(c context.Context, node sqalx.Node, topicID int64) (data []*model.AuthTopic, err error) {
	var addCache = true
	if data, err = p.d.AuthTopicsCache(c, topicID); err != nil {
		addCache = false
	} else if data != nil {
		return
	}

	if data, err = p.d.GetAuthTopicsByCond(c, node, map[string]interface{}{"topic_id": topicID}); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAuthTopicsCache(context.TODO(), topicID, data)
		})
	}
	return
}

// 获取授权了当前话题的话题ID
func (p *Service) GetAuthed2CurrentTopicIDsPaged(c context.Context, topicID int64, limit, offset int32) (ids []int64, err error) {
	if ids, err = p.d.GetAuthed2CurrentTopicIDsPaged(c, p.d.DB(), topicID, limit, offset); err != nil {
		return
	}

	return
}
