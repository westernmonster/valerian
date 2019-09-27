package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error) {
	sqlSelect := "SELECT COUNT(1) as count FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id = ? AND a.deleted=0"
	if err = node.GetContext(c, &count, sqlSelect, topicID, parentID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogChildrenCount error(%+v), topic_id(%d) ,parent id (%d)", err, topicID, parentID))
	}
	return
}

// GetAll get all records
func (p *Dao) GetTopicCatalogs(c context.Context, node sqalx.Node) (items []*model.TopicCatalog, err error) {
	items = make([]*model.TopicCatalog, 0)
	sqlSelect := "SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 ORDER BY a.seq "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogs err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicCatalogsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error) {
	items = make([]*model.TopicCatalog, 0)
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
	if val, ok := cond["is_primary"]; ok {
		clause += " AND a.is_primary =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 %s ORDER BY a.seq", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error) {
	item = new(model.TopicCatalog)
	sqlSelect := "SELECT a.* FROM topic_catalogs a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTopicCatalogByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error) {
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
	if val, ok := cond["is_primary"]; ok {
		clause += " AND a.is_primary =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	sqlInsert := "INSERT INTO topic_catalogs( id,name,seq,type,parent_id,ref_id,topic_id,is_primary,permission,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.IsPrimary, item.Permission, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicCatalogs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	sqlUpdate := "UPDATE topic_catalogs SET name=?,seq=?,type=?,parent_id=?,ref_id=?,topic_id=?,is_primary=?,permission=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.IsPrimary, item.Permission, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicCatalogs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_catalogs SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicCatalogs err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error) {
	sqlSelect := "SELECT a.seq FROM topic_catalogs a WHERE a.deleted=0 AND a.topic_id=? AND a.parent_id=? ORDER BY a.seq  LIMIT 1"
	if err = node.GetContext(c, &seq, sqlSelect, topicID, parentID); err != nil {
		if err == sql.ErrNoRows {
			seq = 0
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogMaxChildrenSeq error(%+v), topic_id(%d) parent_id(%d)", err, topicID, parentID))
		return
	}
	return
}
