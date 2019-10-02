package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetUserTopicsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Topic, err error) {
	items = make([]*model.Topic, 0)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserTopics err(%+v)", err))
		return
	}
	return
}

// GetAll get all records
func (p *Dao) GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error) {
	items = make([]*model.Topic, 0)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopics err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Topic, err error) {
	items = make([]*model.Topic, 0)
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
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =?"
		condition = append(condition, val)
	}
	if val, ok := cond["bg"]; ok {
		clause += " AND a.bg =?"
		condition = append(condition, val)
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_discuss"]; ok {
		clause += " AND a.allow_discuss =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_chat"]; ok {
		clause += " AND a.allow_chat =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =?"
		condition = append(condition, val)
	}
	if val, ok := cond["view_permission"]; ok {
		clause += " AND a.view_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["edit_permission"]; ok {
		clause += " AND a.edit_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["join_permission"]; ok {
		clause += " AND a.join_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["catalog_view_type"]; ok {
		clause += " AND a.catalog_view_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_home"]; ok {
		clause += " AND a.topic_home =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topics a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)
	sqlSelect := "SELECT a.* FROM topics a WHERE a.id=? AND a.deleted=0"

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

// GetByCondition get a record by condition
func (p *Dao) GetTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Topic, err error) {
	item = new(model.Topic)
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
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =?"
		condition = append(condition, val)
	}
	if val, ok := cond["bg"]; ok {
		clause += " AND a.bg =?"
		condition = append(condition, val)
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_discuss"]; ok {
		clause += " AND a.allow_discuss =?"
		condition = append(condition, val)
	}
	if val, ok := cond["allow_chat"]; ok {
		clause += " AND a.allow_chat =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_private"]; ok {
		clause += " AND a.is_private =?"
		condition = append(condition, val)
	}
	if val, ok := cond["view_permission"]; ok {
		clause += " AND a.view_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["edit_permission"]; ok {
		clause += " AND a.edit_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["join_permission"]; ok {
		clause += " AND a.join_permission =?"
		condition = append(condition, val)
	}
	if val, ok := cond["catalog_view_type"]; ok {
		clause += " AND a.catalog_view_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_home"]; ok {
		clause += " AND a.topic_home =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topics a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	sqlInsert := "INSERT INTO topics( id,name,avatar,bg,introduction,allow_discuss,allow_chat,is_private,view_permission,edit_permission,join_permission,catalog_view_type,topic_home,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Avatar, item.Bg, item.Introduction, item.AllowDiscuss, item.AllowChat, item.IsPrivate, item.ViewPermission, item.EditPermission, item.JoinPermission, item.CatalogViewType, item.TopicHome, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	sqlUpdate := "UPDATE topics SET name=?,avatar=?,bg=?,introduction=?,allow_discuss=?,allow_chat=?,is_private=?,view_permission=?,edit_permission=?,join_permission=?,catalog_view_type=?,topic_home=?,created_by=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Avatar, item.Bg, item.Introduction, item.AllowDiscuss, item.AllowChat, item.IsPrivate, item.ViewPermission, item.EditPermission, item.JoinPermission, item.CatalogViewType, item.TopicHome, item.CreatedBy, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topics SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopics err(%+v), item(%+v)", err, id))
		return
	}

	return
}
