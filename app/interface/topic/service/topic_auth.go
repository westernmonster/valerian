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
)

func (p *Service) checkAuthTopics(c context.Context, topicID int64, items []*model.ArgAuthTopic) (err error) {
	// must unique  and could equal to current topic
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

	if err = p.checkAuthTopics(c, topicID, items); err != nil {
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
		// p.d.DelAuthTopicsCache(context.TODO(), topicID)
	})
	return
}

func (p *Service) BulkSaveRelations(c context.Context, arg *model.ArgSaveAuthTopics) (err error) {
	return p.bulkSaveAuthTopics(c, p.d.DB(), arg.TopicID, arg.AuthTopics)
}

// func (p *Service) GetAllRelatedTopicsWithMeta(c context.Context, topicID int64, include string) (items []*model.RelatedTopicResp, err error) {
// 	inc := includeParam(include)
// 	var data []*model.TopicRelation
// 	if data, err = p.getTopicRelations(c, p.d.DB(), topicID); err != nil {
// 		return
// 	}

// 	items = make([]*model.RelatedTopicResp, 0)

// 	for _, v := range data {
// 		item := &model.RelatedTopicResp{
// 			TopicVersionID: v.ToTopicVersionID,
// 			TopicID:        v.ToTopicID,
// 			Seq:            v.Seq,
// 			Type:           v.Relation,
// 		}

// 		var t *model.TopicResp
// 		if t, err = p.getTopic(c, p.d.DB(), item.TopicID); err != nil {
// 			return
// 		}

// 		var ver *model.TopicVersion
// 		if ver, err = p.d.GetTopicVersion(c, p.d.DB(), v.ToTopicVersionID); err != nil {
// 			return
// 		} else if ver == nil {
// 			err = ecode.TopicVersionNotExist
// 			return
// 		}

// 		item.TopicName = t.Name
// 		item.VersionName = ver.Name
// 		item.Cover = t.Cover
// 		item.Introduction = t.Introduction
// 		if item.MembersCount, _, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
// 			return
// 		}
// 		if inc["meta"] {
// 			if item.TopicMeta, err = p.GetTopicMeta(c, t); err != nil {
// 				return
// 			}
// 		}

// 		items = append(items, item)

// 	}

// 	return
// }

// func (p *Service) getAllRelatedTopics(c context.Context, node sqalx.Node, topicID int64) (items []*model.RelatedTopicShort, err error) {
// 	var data []*model.TopicRelation
// 	if data, err = p.getTopicRelations(c, node, topicID); err != nil {
// 		return
// 	}

// 	items = make([]*model.RelatedTopicShort, 0)

// 	for _, v := range data {
// 		item := &model.RelatedTopicShort{
// 			TopicID: v.ToTopicID,
// 			Seq:     v.Seq,
// 			Type:    v.Relation,
// 		}

// 		var t *model.TopicResp
// 		if t, err = p.getTopic(c, node, item.TopicID); err != nil {
// 			return
// 		}

// 		var ver *model.TopicVersion
// 		if ver, err = p.d.GetTopicVersion(c, p.d.DB(), v.ToTopicVersionID); err != nil {
// 			return
// 		} else if ver == nil {
// 			err = ecode.TopicVersionNotExist
// 			return
// 		}
// 		item.TopicName = t.Name
// 		item.VersionName = ver.Name

// 		items = append(items, item)
// 	}
// 	return
// }

// func (p *Service) getTopicRelations(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicRelation, err error) {
// 	var addCache = true

// 	if items, err = p.d.TopicRelationCache(c, topicID); err != nil {
// 		addCache = false
// 	} else if items != nil {
// 		return
// 	}

// 	if items, err = p.d.GetAllTopicRelations(c, node, topicID); err != nil {
// 		return
// 	}

// 	if addCache {
// 		p.addCache(func() {
// 			p.d.SetTopicRelationCache(context.TODO(), topicID, items)
// 		})
// 	}

// 	return
// }
