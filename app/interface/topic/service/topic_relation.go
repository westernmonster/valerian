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

func (p *Service) bulkSaveRelations(c context.Context, node sqalx.Node, topicID int64, relations []*model.ArgRelatedTopic) (err error) {
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
	if t, err = p.d.GetTopicByID(c, tx, topicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	for _, v := range relations {
		var relation *model.TopicRelation
		if relation, err = p.d.GetTopicRelationByCondition(c, tx, topicID, v.TopicID); err != nil {
			return
		}

		if relation != nil {
			relation.Relation = v.Type
			if err = p.d.UpdateTopicRelation(c, tx, relation); err != nil {
				return
			}

			continue
		}

		item := &model.TopicRelation{
			ID:          gid.NewID(),
			FromTopicID: topicID,
			ToTopicID:   v.TopicID,
			Seq:         v.Seq,
			Relation:    v.Type,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}
		if err = p.d.AddTopicRelation(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	p.addCache(func() {
		p.d.DelTopicRelationCache(context.TODO(), topicID)
	})
	return
}

func (p *Service) BulkSaveRelations(c context.Context, arg *model.ArgBatchSaveRelatedTopics) (err error) {
	return p.bulkSaveRelations(c, p.d.DB(), arg.TopicID, arg.RelatedTopics)
}

func (p *Service) GetAllRelatedTopicsWithMeta(c context.Context, topicID int64) (items []*model.RelatedTopicResp, err error) {
	var data []*model.TopicRelation
	if data, err = p.getTopicRelations(c, p.d.DB(), topicID); err != nil {
		return
	}

	items = make([]*model.RelatedTopicResp, 0)

	for _, v := range data {
		item := &model.RelatedTopicResp{
			TopicID: v.ToTopicID,
			Seq:     v.Seq,
			Type:    v.Relation,
		}

		var t *model.TopicResp
		if t, err = p.getTopic(c, item.TopicID); err != nil {
			return
		}
		item.TopicName = t.Name
		item.VersionName = t.VersionName
		item.Cover = t.Cover
		item.Introduction = t.Introduction
		if item.MembersCount, _, err = p.getTopicMembers(c, p.d.DB(), topicID, 10); err != nil {
			return
		}
		if item.TopicMeta, err = p.GetTopicMeta(c, t); err != nil {
			return
		}

		items = append(items, item)

	}

	return
}

func (p *Service) getAllRelatedTopics(c context.Context, node sqalx.Node, topicID int64) (items []*model.RelatedTopicShort, err error) {
	var data []*model.TopicRelation
	if data, err = p.getTopicRelations(c, node, topicID); err != nil {
		return
	}

	items = make([]*model.RelatedTopicShort, 0)

	for _, v := range data {
		item := &model.RelatedTopicShort{
			TopicID: v.ToTopicID,
			Seq:     v.Seq,
			Type:    v.Relation,
		}

		var t *model.TopicResp
		if t, err = p.getTopic(c, item.TopicID); err != nil {
			return
		}
		item.TopicName = t.Name
		item.VersionName = t.VersionName

		items = append(items, item)
	}
	return
}

func (p *Service) getTopicRelations(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicRelation, err error) {
	var addCache = true

	if items, err = p.d.TopicRelationCache(c, topicID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetAllTopicRelations(c, node, topicID); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicRelationCache(context.TODO(), topicID, items)
		})
	}

	return
}
