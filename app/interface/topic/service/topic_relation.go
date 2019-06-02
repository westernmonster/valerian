package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) bulkSaveRelations(c context.Context, node sqalx.Node, topicID int64, relations []*model.ArgRelatedTopic) (err error) {
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

	return
}

func (p *Service) BulkSaveRelations(c context.Context, arg *model.ArgBatchSaveRelatedTopics) (err error) {
	return p.bulkSaveRelations(c, p.d.DB(), arg.TopicID, arg.RelatedTopics)
}
