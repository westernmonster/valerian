package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/fav/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetFavsPaged(c context.Context, node sqalx.Node, aid int64, targetType string, limit, offset int) (items []*model.Fav, err error) {
	items = make([]*model.Fav, 0)

	if targetType == "all" {
		sql := "SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sql, aid, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFavsPaged error(%+v), aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
		}
	} else {
		sql := "SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.deleted=0 AND a.account_id=? AND target_type=? ORDER BY a.id DESC limit ?,?"
		if err = node.SelectContext(c, &items, sql, aid, targetType, offset, limit); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFavsPaged error(%+v), aid(%d) target_type(%s) limit(%d) offset(%d)", err, aid, targetType, limit, offset))
		}
	}
	return
}

// GetAll get all records
func (p *Dao) GetFavs(c context.Context, node sqalx.Node) (items []*model.Fav, err error) {
	items = make([]*model.Fav, 0)
	sqlSelect := "SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFavs err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetFavsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Fav, err error) {
	items = make([]*model.Fav, 0)
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFavsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetFavByID(c context.Context, node sqalx.Node, id int64) (item *model.Fav, err error) {
	item = new(model.Fav)
	sqlSelect := "SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFavByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetFavByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Fav, err error) {
	item = new(model.Fav)
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM favs a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFavsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddFav(c context.Context, node sqalx.Node, item *model.Fav) (err error) {
	sqlInsert := "INSERT INTO favs( id,account_id,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.AccountID, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddFavs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateFav(c context.Context, node sqalx.Node, item *model.Fav) (err error) {
	sqlUpdate := "UPDATE favs SET account_id=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.AccountID, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFavs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelFav(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE favs SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelFavs err(%+v), item(%+v)", err, id))
		return
	}

	return
}
