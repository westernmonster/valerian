package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetFeedPaged(c context.Context, node sqalx.Node, accountID int64, limit, offset int) (items []*model.Feed, err error) {
	items = make([]*model.Feed, 0)

	sql := "SELECT a.* FROM feeds a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, accountID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedPaged error(%+v), account_id(%d) limit(%d) offset(%d)", err, accountID, limit, offset))
	}
	return
}

// GetAll get all records
func (p *Dao) GetFeeds(c context.Context, node sqalx.Node) (items []*model.Feed, err error) {
	items = make([]*model.Feed, 0)
	sqlSelect := "SELECT a.* FROM feeds a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeeds err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetFeedByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Feed, err error) {
	item = new(model.Feed)
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
	if val, ok := cond["actor_id"]; ok {
		clause += " AND a.actor_id =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM feeds a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

func (p *Dao) GetFeedsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Feed, err error) {
	items = make([]*model.Feed, 0)
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
	if val, ok := cond["actor_id"]; ok {
		clause += " AND a.actor_id =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM feeds a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFeedsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// Insert insert a new record
func (p *Dao) AddFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error) {
	sqlInsert := "INSERT INTO feeds( id,account_id,action_type,action_time,action_text,actor_id,actor_type,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFeed(c context.Context, node sqalx.Node, item *model.Feed) (err error) {
	sqlUpdate := "UPDATE feeds SET account_id=?,action_type=?,action_time=?,action_text=?,actor_id=?,actor_type=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeed(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeeds err(%+v), item(%+v)", err, id))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFeedByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE feeds SET deleted=1 WHERE  target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFeed err(%+v), target_type(%+v), target_id(%+v)", err, targetType, targetID))
		return
	}

	return
}
