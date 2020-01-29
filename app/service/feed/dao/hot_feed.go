package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetHotFeeds(c context.Context, node sqalx.Node) (items []*model.HotFeed, err error) {
	items = make([]*model.HotFeed, 0)
	sqlSelect := "SELECT a.id,a.target_id,a.target_type,a.score,a.deleted,a.created_at,a.updated_at FROM hot_feeds a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetHotFeeds err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetHotFeedsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.HotFeed, err error) {
	items = make([]*model.HotFeed, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
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
	if val, ok := cond["score"]; ok {
		clause += " AND a.score =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.score,a.deleted,a.created_at,a.updated_at FROM hot_feeds a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetHotFeedsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetHotFeedByID(c context.Context, node sqalx.Node, id int64) (item *model.HotFeed, err error) {
	item = new(model.HotFeed)
	sqlSelect := "SELECT a.id,a.target_id,a.target_type,a.score,a.deleted,a.created_at,a.updated_at FROM hot_feeds a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetHotFeedByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetHotFeedByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.HotFeed, err error) {
	item = new(model.HotFeed)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
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
	if val, ok := cond["score"]; ok {
		clause += " AND a.score =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.target_id,a.target_type,a.score,a.deleted,a.created_at,a.updated_at FROM hot_feeds a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetHotFeedsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddHotFeed(c context.Context, node sqalx.Node, item *model.HotFeed) (err error) {
	sqlInsert := "INSERT INTO hot_feeds( id,target_id,target_type,score,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TargetID, item.TargetType, item.Score, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddHotFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateHotFeed(c context.Context, node sqalx.Node, item *model.HotFeed) (err error) {
	sqlUpdate := "UPDATE hot_feeds SET target_id=?,target_type=?,score=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TargetID, item.TargetType, item.Score, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateHotFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelHotFeed(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE hot_feeds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelHotFeeds err(%+v), item(%+v)", err, id))
		return
	}

	return
}
