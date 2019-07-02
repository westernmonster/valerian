package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAllTopicRelationsSQL = "SELECT a.* FROM topic_relations a WHERE a.from_topic_id=? AND a.deleted=0 ORDER BY a.seq"
	_addTopicRelationSQL     = "INSERT INTO topic_relations( id,from_topic_id,to_topic_version_id,to_topic_id,relation,seq,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"
	_updateTopicRelationSQL  = "UPDATE topic_relations SET from_topic_id=?,to_topic_version_id=?,to_topic_id=?,relation=?,seq=?,updated_at=? WHERE id=?"
	_deleteTopicRelationSQL  = "UPDATE topic_relations SET deleted=1 WHERE id=? "
	_isRelationExistSQL      = "SELECT a.* FROM topic_relations a WHERE a.deleted=0 AND a.from_topic_id=? AND a.to_topic_version_id=?"
)

func (p *Dao) GetAllTopicRelations(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicRelation, err error) {
	items = make([]*model.TopicRelation, 0)

	if err = node.SelectContext(c, &items, _getAllTopicRelationsSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllTopicRelations error(%+v), topic id(%d)", err, topicID))
	}
	return
}

// Insert insert a new record
func (p *Dao) AddTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error) {
	if _, err = node.ExecContext(c, _addTopicRelationSQL, item.ID, item.FromTopicID, item.ToTopicVersionID, item.ToTopicID, item.Relation, item.Seq, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicRelation error(%+v), item(%+v)", err, item))
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error) {
	if _, err = node.ExecContext(c, _updateTopicRelationSQL, item.FromTopicID, item.ToTopicVersionID, item.ToTopicID, item.Relation, item.Seq, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicRelation error(%+v), item(%+v)", err, item))
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicRelation(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _deleteTopicRelationSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicRelationsPaged error(%+v), topic member id(%d)", err, id))
	}

	return
}

func (p *Dao) GetTopicRelationByCondition(c context.Context, node sqalx.Node, fromTopicID, toTopicVersionID int64) (item *model.TopicRelation, err error) {
	item = new(model.TopicRelation)

	if err = node.GetContext(c, item, _isRelationExistSQL, fromTopicID, toTopicVersionID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.IsTopicRelationExist error(%+v), from_topic_id(%d), to_topic_version_id(%d)", err, fromTopicID, toTopicVersionID))
	}

	return
}
