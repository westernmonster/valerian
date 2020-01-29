package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

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
