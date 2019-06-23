package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addTopicCatalogSQL    = "INSERT INTO topic_catalogs( id,name,seq,type,parent_id,ref_id,topic_id,is_primary,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"
	_updateTopicCatalogSQL = "UPDATE topic_catalogs SET name=?,seq=?,type=?,parent_id=?,ref_id=?,topic_id=?,is_primary=?,updated_at=? WHERE id=? AND deleted=0"
	_delTopicCatalogSQL    = "UPDATE topic_catalogs SET deleted=1 WHERE id=? "
	_getTopicCatalogByID   = "SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 AND a.id=?"
	_getMaxChildrenSeqSQL  = "SELECT a.seq FROM topic_catalogs a WHERE a.deleted=0 AND a.topic_id=? AND a.parent_id=? ORDER BY a.seq DESC LIMIT 1"
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
	if _, err = node.ExecContext(c, _addTopicCatalogSQL, item.ID, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.IsPrimary, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicCatalog error(%+v), item(%+v)", err, item))
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	if _, err = node.ExecContext(c, _updateTopicCatalogSQL, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.IsPrimary, item.UpdatedAt, item.ID); err != nil {
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

func (p *Dao) GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error) {
	if err = node.GetContext(c, &seq, _getMaxChildrenSeqSQL, topicID, parentID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicCatalogMaxChildrenSeq error(%+v), topic id(%d) parent_id(%d)", err, topicID, parentID))
		return
	}
	return
}
