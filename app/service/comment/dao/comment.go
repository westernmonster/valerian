package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetCommentsPaged(c context.Context, node sqalx.Node, resourceID int64, targetType string, limit, offset int) (items []*model.Comment, err error) {
	items = make([]*model.Comment, 0)
	sqlSelect := "SELECT a.id,a.content,a.target_type,a.owner_id,a.resource_id,a.featured,a.deleted,a.reply_to,a.created_by,a.created_at,a.updated_at,a.owner_type FROM comments a WHERE a.resource_id=? AND target_type=? ORDER BY a.featured DESC,a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sqlSelect, resourceID, targetType, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCommentsPaged err(%+v) resource_id(%d) target_type(%s) limit(%d) offset(%d)", err, resourceID, targetType, limit, offset))
	}
	return
}

func (p *Dao) GetAllChildrenComments(c context.Context, node sqalx.Node, resourceID int64) (items []*model.Comment, err error) {
	sqlSelect := "SELECT a.id,a.content,a.target_type,a.owner_id,a.resource_id,a.featured,a.deleted,a.reply_to,a.created_by,a.created_at,a.updated_at,a.owner_type FROM comments a WHERE a.resource_id=? AND target_type='comment' ORDER BY a.id DESC"

	if err = node.SelectContext(c, &items, sqlSelect, resourceID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllChildrenComments err(%+v), resource_id(%+v)", err, resourceID))
		return
	}
	return
}

func (p *Dao) GetChildrenComments(c context.Context, node sqalx.Node, resourceID int64, limit int) (items []*model.Comment, err error) {
	sqlSelect := "SELECT a.id,a.content,a.target_type,a.owner_id,a.resource_id,a.featured,a.deleted,a.reply_to,a.created_by,a.created_at,a.updated_at,a.owner_type  FROM comments a WHERE a.resource_id=? AND target_type='comment' ORDER BY a.id DESC LIMIT ?"

	if err = node.SelectContext(c, &items, sqlSelect, resourceID, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllChildrenComments err(%+v), resource_id(%+v)", err, resourceID))
		return
	}
	return
}

func (p *Dao) GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error) {
	item = new(model.Comment)
	sqlSelect := "SELECT a.id,a.content,a.target_type,a.owner_id,a.resource_id,a.featured,a.deleted,a.reply_to,a.created_by,a.created_at,a.updated_at,a.owner_type FROM comments a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCommentByID err(%+v), id(%+v)", err, id))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddComment(c context.Context, node sqalx.Node, item *model.Comment) (err error) {
	sqlInsert := "INSERT INTO comments( id,content,target_type,owner_id,resource_id,featured,deleted,reply_to,created_by,created_at,updated_at,owner_type) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Content, item.TargetType, item.OwnerID, item.ResourceID, item.Featured, item.Deleted, item.ReplyTo, item.CreatedBy, item.CreatedAt, item.UpdatedAt, item.OwnerType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddComments err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateComment(c context.Context, node sqalx.Node, item *model.Comment) (err error) {
	sqlUpdate := "UPDATE comments SET content=?,target_type=?,owner_id=?,resource_id=?,featured=?,reply_to=?,created_by=?,updated_at=?,owner_type=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Content, item.TargetType, item.OwnerID, item.ResourceID, item.Featured, item.ReplyTo, item.CreatedBy, item.UpdatedAt, item.OwnerType, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateComments err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelComment(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE comments SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelComments err(%+v), item(%+v)", err, id))
		return
	}

	return
}
