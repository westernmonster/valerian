package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/recent/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetRecentViews(c context.Context, node sqalx.Node) (items []*model.RecentView, err error) {
	items = make([]*model.RecentView, 0)
	sqlSelect := "SELECT a.* FROM recent_views a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViews err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetRecentViewsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.RecentView, err error) {
	items = make([]*model.RecentView, 0)
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
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM recent_views a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetRecentViewByID(c context.Context, node sqalx.Node, id int64) (item *model.RecentView, err error) {
	item = new(model.RecentView)
	sqlSelect := "SELECT a.* FROM recent_views a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetRecentViewByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.RecentView, err error) {
	item = new(model.RecentView)
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
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM recent_views a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetRecentViewsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddRecentView(c context.Context, node sqalx.Node, item *model.RecentView) (err error) {
	sqlInsert := "INSERT INTO recent_views( id,account_id,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRecentViews err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateRecentView(c context.Context, node sqalx.Node, item *model.RecentView) (err error) {
	sqlUpdate := "UPDATE recent_views SET account_id=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateRecentViews err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelRecentView(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE recent_views SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecentViews err(%+v), item(%+v)", err, id))
		return
	}

	return
}
