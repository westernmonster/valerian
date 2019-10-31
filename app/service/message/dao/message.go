package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetMessages(c context.Context, node sqalx.Node) (items []*model.Message, err error) {
	items = make([]*model.Message, 0)
	sqlSelect := "SELECT a.* FROM messages a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetMessages err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetMessagesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Message, err error) {
	items = make([]*model.Message, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_type"]; ok {
		clause += " AND a.action_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_time"]; ok {
		clause += " AND a.action_time =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_text"]; ok {
		clause += " AND a.action_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["actors"]; ok {
		clause += " AND a.actors =?"
		condition = append(condition, val)
	}
	if val, ok := cond["extend"]; ok {
		clause += " AND a.extend =?"
		condition = append(condition, val)
	}
	if val, ok := cond["merge_count"]; ok {
		clause += " AND a.merge_count =?"
		condition = append(condition, val)
	}
	if val, ok := cond["actor_type"]; ok {
		clause += " AND a.actor_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_read"]; ok {
		clause += " AND a.is_read =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM messages a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetMessagesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetMessageByID(c context.Context, node sqalx.Node, id int64) (item *model.Message, err error) {
	item = new(model.Message)
	sqlSelect := "SELECT a.* FROM messages a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetMessageByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetMessageByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Message, err error) {
	item = new(model.Message)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_type"]; ok {
		clause += " AND a.action_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_time"]; ok {
		clause += " AND a.action_time =?"
		condition = append(condition, val)
	}
	if val, ok := cond["action_text"]; ok {
		clause += " AND a.action_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["actors"]; ok {
		clause += " AND a.actors =?"
		condition = append(condition, val)
	}
	if val, ok := cond["extend"]; ok {
		clause += " AND a.extend =?"
		condition = append(condition, val)
	}
	if val, ok := cond["merge_count"]; ok {
		clause += " AND a.merge_count =?"
		condition = append(condition, val)
	}
	if val, ok := cond["actor_type"]; ok {
		clause += " AND a.actor_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_read"]; ok {
		clause += " AND a.is_read =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM messages a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetMessagesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddMessage(c context.Context, node sqalx.Node, item *model.Message) (err error) {
	sqlInsert := "INSERT INTO messages( id,account_id,action_type,action_time,action_text,actors,extend,merge_count,actor_type,target_id,target_type,is_read,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.Actors, item.Extend, item.MergeCount, item.ActorType, item.TargetID, item.TargetType, item.IsRead, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddMessage err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateMessage(c context.Context, node sqalx.Node, item *model.Message) (err error) {
	sqlUpdate := "UPDATE messages SET account_id=?,action_type=?,action_time=?,action_text=?,actors=?,extend=?,merge_count=?,actor_type=?,target_id=?,target_type=?,is_read=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.Actors, item.Extend, item.MergeCount, item.ActorType, item.TargetID, item.TargetType, item.IsRead, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateMessages err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelMessage(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE messages SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelMessages err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) DelMessageByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE messages SET deleted=1 WHERE target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelMessages err(%+v), target_type(%s) traget_id(%d)", err, targetType, targetID))
		return
	}

	return
}
