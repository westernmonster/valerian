package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAllTopicRelationsSQL      = "SELECT a.* FROM topic_relations a WHERE a.from_topic_id=?"
	_getAllRelatedTopicsSQL       = "SELECT a.to_topic_id AS topic_id,b.name,b.version_name,a.relation AS type FROM topic_relations a LEFT JOIN topics b ON a.to_topic_id=b.id WHERE a.from_topic_id=?"
	_getAllRelatedTopicsDetailSQL = "SELECT a.to_topic_id AS topic_id,b.name,b.version_name,a.relation AS type,b.cover, b.introduction FROM topic_relations a LEFT JOIN topics b ON a.to_topic_id=b.id WHERE a.from_topic_id=?"
	_addTopicRelationSQL          = "INSERT INTO topic_relations( id,from_topic_id,to_topic_id,relation,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"
	_updateTopicRelationSQL       = "UPDATE topic_relations SET from_topic_id=?,to_topic_id=?,relation=?,updated_at=? WHERE id=?"
	_deleteTopicRelationSQL       = "UPDATE topic_relations SET deleted=1 WHERE id=? "
)

func (p *Dao) GetAllTopicRelations(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicRelation, err error) {
	items = make([]*model.TopicRelation, 0)

	if err = node.SelectContext(c, &items, _getAllTopicRelationsSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllTopicRelations error(%+v), topic id(%d)", err, topicID))
	}
	return
}

func (p *Dao) GetAllRelatedTopics(c context.Context, node sqalx.Node, topicID int64) (items []*model.RelatedTopicShort, err error) {
	items = make([]*model.RelatedTopicShort, 0)

	if err = node.SelectContext(c, &items, _getAllRelatedTopicsSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllTopicRelations error(%+v), topic id(%d)", err, topicID))
	}
	return
}

func (p *Dao) GetAllRelatedTopicsDetail(c context.Context, node sqalx.Node, topicID int64) (items []*model.RelatedTopicResp, err error) {
	items = make([]*model.RelatedTopicResp, 0)

	if err = node.SelectContext(c, &items, _getAllRelatedTopicsDetailSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllRelatedTopicsDetail error(%+v), topic id(%d)", err, topicID))
	}
	return
}

// Insert insert a new record
func (p *Dao) AddTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error) {
	if _, err = node.ExecContext(c, _addTopicRelationSQL, item.ID, item.FromTopicID, item.ToTopicID, item.Relation, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicRelation error(%+v), item(%+v)", err, item))
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error) {
	if _, err = node.ExecContext(c, _updateTopicRelationSQL, item.FromTopicID, item.ToTopicID, item.Relation, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicRelation error(%+v), item(%+v)", err, item))
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DeleteTopicRelation(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _deleteTopicRelationSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicRelationsPaged error(%+v), topic member id(%d)", err, id))
	}

	return
}
