package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicDiscussionsPaged(c context.Context, node sqalx.Node, topicID, categoryID int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)

	if categoryID == 0 {
		sql := "SELECT a.* FROM discussions a WHERE a.deleted=0 AND a.topic_id=?  ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sql, topicID, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged error(%+v), topic_id(%d) limit(%d) offset(%d)", err, topicID, limit, offset))
		}
	} else {
		sql := "SELECT a.* FROM discussions a WHERE a.deleted=0 AND a.topic_id=? and category_id=? ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sql, topicID, categoryID, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged error(%+v), topic_id(%d) category_id(%d) limit(%d) offset(%d)", err, topicID, categoryID, limit, offset))
		}
	}

	return
}

func (p *Dao) GetUserDiscussionsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)

	sql := "SELECT a.* FROM discussions a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
	}
	return
}

// GetAll get all records
func (p *Dao) GetDiscussions(c context.Context, node sqalx.Node) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	sqlSelect := "SELECT a.* FROM discussions a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussions err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetDiscussionsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["category_id"]; ok {
		clause += " AND a.category_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM discussions a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetDiscussionByID(c context.Context, node sqalx.Node, id int64) (item *model.Discussion, err error) {
	item = new(model.Discussion)
	sqlSelect := "SELECT a.* FROM discussions a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetDiscussionByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Discussion, err error) {
	item = new(model.Discussion)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["category_id"]; ok {
		clause += " AND a.category_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM discussions a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error) {
	sqlInsert := "INSERT INTO discussions( id,topic_id,category_id,created_by,title,content,content_text,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.CategoryID, item.CreatedBy, item.Title, item.Content, item.ContentText, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDiscussions err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error) {
	sqlUpdate := "UPDATE discussions SET topic_id=?,category_id=?,created_by=?,title=?,content=?,content_text=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.CategoryID, item.CreatedBy, item.Title, item.Content, item.ContentText, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDiscussions err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelDiscussion(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE discussions SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDiscussions err(%+v), item(%+v)", err, id))
		return
	}

	return
}
