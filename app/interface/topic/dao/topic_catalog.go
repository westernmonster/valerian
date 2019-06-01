package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addTopicCatalogSQL              = "INSERT INTO topic_catalogs( id,name,seq,type,parent_id,ref_id,topic_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"
	_updateTopicCatalogSQL           = "UPDATE topic_catalogs SET name=?,seq=?,type=?,parent_id=?,ref_id=?,topic_id=?,updated_at=? WHERE id=? AND deleted=0"
	_delTopicCatalogSQL              = "UPDATE topic_catalogs SET deleted=1 WHERE id=? "
	_getTopicCatalogChildrenCountSQL = "SELECT COUNT(1) as count FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id = ? AND a.deleted=0"
	_getTopicCatalogsByCondition     = "SELECT a.* FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id=? AND a.deleted=0 ORDER BY a.seq"
)

// Insert insert a new record
func (p *Dao) AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	if _, err = node.ExecContext(c, _addTopicCatalogSQL, item.ID, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicCatalog error(%+v), item(%+v)", err, item))
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	if _, err = node.ExecContext(c, _updateTopicCatalogSQL, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicCatalog error(%+v), item(%+v)", err, item))
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delTopicCatalogSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicCatalog error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getTopicCatalogChildrenCountSQL, topicID, parentID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogChildrenCount error(%+v), topic id(%d) ,parent id (%d)", err, topicID, parentID))
	}
	return
}

func (p *Dao) GetTopicCatalogsByCondition(c context.Context, node sqalx.Node, topicID, parentID int64) (items []*model.TopicCatalog, err error) {
	items = make([]*model.TopicCatalog, 0)
	if err = node.SelectContext(c, &items, _getTopicCatalogsByCondition, topicID, parentID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogsByCondition error(%+v), topic id(%d) ,parent id (%d)", err, topicID, parentID))
	}
	return
}
