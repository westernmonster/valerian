package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)
	sqlSelect := "SELECT a.id,a.name,a.avatar,a.bg,a.introduction,a.allow_discuss,a.allow_chat,a.is_private,a.view_permission,a.edit_permission,a.join_permission,a.catalog_view_type,a.topic_home,a.created_by,a.deleted,a.created_at,a.updated_at FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}

func (p *Dao) GetTopicEditPermissionByID(c context.Context, node sqalx.Node, id int64) (editPermission string, err error) {
	item := new(model.Topic)
	sqlSelect := "SELECT a.id,a.name,a.avatar,a.bg,a.introduction,a.allow_discuss,a.allow_chat,a.is_private,a.view_permission,a.edit_permission,a.join_permission,a.catalog_view_type,a.topic_home,a.created_by,a.deleted,a.created_at,a.updated_at FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			editPermission = ""
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicEditPermissionByID err(%+v), id(%+v)", err, id))
		return
	}

	editPermission = item.EditPermission

	return
}

func (p *Dao) GetTopicViewPermissionByID(c context.Context, node sqalx.Node, id int64) (viewPermission string, err error) {
	item := new(model.Topic)
	sqlSelect := "SELECT a.id,a.name,a.avatar,a.bg,a.introduction,a.allow_discuss,a.allow_chat,a.is_private,a.view_permission,a.edit_permission,a.join_permission,a.catalog_view_type,a.topic_home,a.created_by,a.deleted,a.created_at,a.updated_at FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			viewPermission = ""
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicViewPermissionByID err(%+v), id(%+v)", err, id))
		return
	}

	viewPermission = item.ViewPermission

	return
}
