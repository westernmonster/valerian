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
	_addTopicCatalogSQL              = "INSERT INTO topic_catalogs( id,name,seq,type,parent_id,ref_id,topic_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"
	_updateTopicCatalogSQL           = "UPDATE topic_catalogs SET name=?,seq=?,type=?,parent_id=?,ref_id=?,topic_id=?,updated_at=? WHERE id=? AND deleted=0"
	_delTopicCatalogSQL              = "UPDATE topic_catalogs SET deleted=1 WHERE id=? "
	_getTopicCatalogChildrenCountSQL = "SELECT COUNT(1) as count FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id = ? AND a.deleted=0"
	_getTopicCatalogsByCondition     = "SELECT a.* FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id=? AND a.deleted=0 ORDER BY a.seq"
	_getTopicCatalogByCondition      = "SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 %s"
	_getTopicCatalogByID             = "SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 AND a.id=?"
)

func (p *Dao) GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error) {
	item = new(model.TopicCatalog)

	if err = node.GetContext(c, item, _getTopicCatalogByID, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogByID error(%+v), id(%d)", err, id))
	}

	return
}

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

func (p *Dao) GetTopicCatalogByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *model.TopicCatalog, err error) {
	item = new(model.TopicCatalog)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["type"]; ok {
		clause += " AND a.type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ref_id"]; ok {
		clause += " AND a.ref_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf(_getTopicCatalogByCondition, clause)

	if err = node.GetContext(ctx, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			return
		}
	}

	return
}
