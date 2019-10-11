package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
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

func (p *Service) checkAuthTopics(c context.Context, node sqalx.Node, topicID int64, items []*model.ArgAuthTopic) (err error) {
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
		dic[v.ToTopicID] = false
	}

	return
}

func (p *Service) bulkSaveAuthTopics(c context.Context, node sqalx.Node, topicID int64, items []*model.ArgAuthTopic) (err error) {
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

	if err = p.checkAuthTopics(c, tx, topicID, items); err != nil {
		return
	}

	if err = p.checkTopic(c, tx, topicID); err != nil {
		return
	}

	var dic map[int64]bool
	if dic, err = p.loadAuthTopicsMap(c, tx, topicID); err != nil {
		return
	}

	for _, v := range items {
		var dItem *model.AuthTopic
		if dItem, err = p.d.GetAuthTopicByCond(c, tx, map[string]interface{}{"topic_id": topicID, "to_topic_id": v.TopicID}); err != nil {
			return
		}
		dic[v.TopicID] = true

		if dItem == nil {
			item := &model.AuthTopic{
				ID:         gid.NewID(),
				TopicID:    topicID,
				ToTopicID:  v.TopicID,
				Permission: v.Permission,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			}
			if err = p.d.AddAuthTopic(c, tx, item); err != nil {
				return
			}

			continue
		}
		dItem.Permission = v.Permission
		if err = p.d.UpdateAuthTopic(c, tx, dItem); err != nil {
			return
		}
	}

	for k, used := range dic {
		if used {
			continue
		}
		if err = p.d.DelAuthTopic(c, tx, k); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelAuthTopicsCache(context.TODO(), topicID)
	})
	return
}

func (p *Service) SaveAuthTopics(c context.Context, arg *model.ArgSaveAuthTopics) (err error) {
	return p.bulkSaveAuthTopics(c, p.d.DB(), arg.TopicID, arg.AuthTopics)
}

func (p *Service) GetAuthTopics(c context.Context, topicID int64) (items []*model.AuthTopicResp, err error) {
	return p.getAuthTopicsResp(c, p.d.DB(), topicID)
}

func (p *Service) getAuthTopicsResp(c context.Context, node sqalx.Node, topicID int64) (items []*model.AuthTopicResp, err error) {
	var data []*model.AuthTopic
	if data, err = p.getAuthTopics(c, node, topicID); err != nil {
		return
	}

	items = make([]*model.AuthTopicResp, len(data))

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

		item := &model.AuthTopicResp{
			ToTopicID:   v.ToTopicID,
			Permission:  v.Permission,
			Avatar:      t.Avatar,
			Name:        t.Name,
			MemberCount: stat.MemberCount,
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

func (p *Service) GetUserCanEditTopics(c context.Context) (resp *model.CanEditTopicsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var data []*model.TopicIDItem
	if data, err = p.d.GetUserCanEditTopicIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	resp = &model.CanEditTopicsResp{
		Items:  make([]*model.CanEditTopicItem, len(data)),
		Paging: &model.Paging{},
	}

	for i, v := range data {
		var t *model.Topic
		if t, err = p.getTopic(c, p.d.DB(), v.TopicID); err != nil {
			return
		}
		item := &model.CanEditTopicItem{
			ID:             v.TopicID,
			Name:           t.Name,
			Introduction:   t.Introduction,
			EditPermission: t.EditPermission,
			Avatar:         t.Avatar,
		}

		var stat *model.TopicStat
		if stat, err = p.GetTopicStat(c, v.TopicID); err != nil {
			return
		}

		item.MemberCount = stat.MemberCount
		item.ArticleCount = stat.ArticleCount
		item.DiscussionCount = stat.DiscussionCount

		if item.HasCatalogTaxonomy, err = p.d.HasTaxonomy(c, p.d.DB(), v.TopicID); err != nil {
			return
		}

		resp.Items[i] = item
	}

	resp.Paging.IsEnd = true

	return
}
