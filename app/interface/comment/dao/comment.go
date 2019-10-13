package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error) {
	item = new(model.Comment)
	sqlSelect := "SELECT a.* FROM comments a WHERE a.id=? "

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

func (p *Dao) AddComment(c context.Context, node sqalx.Node, item *model.Comment) (err error) {
	sqlInsert := "INSERT INTO comments( id,content,owner_id,resource_id,featured,deleted,created_by,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Content, item.OwnerID, item.ResourceID, item.Featured, item.Deleted, item.CreatedBy, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddComments err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateComment(c context.Context, node sqalx.Node, item *model.Comment) (err error) {
	sqlUpdate := "UPDATE comments SET content=?,owner_id=?,resource_id=?,featured=?,created_by=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Content, item.OwnerID, item.ResourceID, item.Featured, item.CreatedBy, item.UpdatedAt, item.ID)
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
