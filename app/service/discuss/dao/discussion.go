package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetTopicDiscussionsPaged(c context.Context, node sqalx.Node, topicID, categoryID int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)

	if categoryID == 0 {
		sqlSelect := "SELECT a.id,a.topic_id,a.category_id,a.created_by,a.title,a.content,a.content_text,a.deleted,a.created_at,a.updated_at FROM discussions a WHERE a.deleted=0 AND a.topic_id=?  ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sqlSelect, topicID, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged error(%+v), topic_id(%d) limit(%d) offset(%d)", err, topicID, limit, offset))
		}
	} else {
		sqlSelect := "SELECT a.id,a.topic_id,a.category_id,a.created_by,a.title,a.content,a.content_text,a.deleted,a.created_at,a.updated_at FROM discussions a WHERE a.deleted=0 AND a.topic_id=? and category_id=? ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sqlSelect, topicID, categoryID, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussionsPaged error(%+v), topic_id(%d) category_id(%d) limit(%d) offset(%d)", err, topicID, categoryID, limit, offset))
		}
	}

	return
}

func (p *Dao) GetUserDiscussionIDsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := `
    SELECT a.id
	FROM discussions a LEFT JOIN topics b ON a.topic_id=b.id
	WHERE b.deleted =0 AND a.deleted=0 AND a.created_by=?
	ORDER BY a.id DESC limit ?,?`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionIDsPaged err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionIDsPaged err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) GetUserDiscussionsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Discussion, err error) {
	items = make([]*model.Discussion, 0)
	sqlSelect := `
    SELECT a.id,a.topic_id,a.category_id,a.created_by,a.title,a.content,a.content_text,a.deleted,a.created_at,a.updated_at
	FROM discussions a LEFT JOIN topics b ON a.topic_id = b.id
	WHERE a.deleted=0 AND b.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?`

	if err = node.SelectContext(c, &items, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDiscussionsPaged err(%+v) aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetDiscussionByID(c context.Context, node sqalx.Node, id int64) (item *model.Discussion, err error) {
	item = new(model.Discussion)
	sqlSelect := "SELECT a.id,a.topic_id,a.category_id,a.created_by,a.title,a.content,a.content_text,a.deleted,a.created_at,a.updated_at FROM discussions a WHERE a.deleted=0 AND a.id=?"

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
