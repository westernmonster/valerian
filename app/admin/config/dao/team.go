package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetTeams(c context.Context, node sqalx.Node) (items []*model.Team, err error) {
	items = make([]*model.Team, 0)
	sqlSelect := "SELECT a.id,a.name,a.env,a.zone,a.deleted,a.created_at,a.updated_at FROM teams a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTeams err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTeamsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Team, err error) {
	items = make([]*model.Team, 0)
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
	if val, ok := cond["env"]; ok {
		clause += " AND a.env =?"
		condition = append(condition, val)
	}
	if val, ok := cond["zone"]; ok {
		clause += " AND a.zone =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.env,a.zone,a.deleted,a.created_at,a.updated_at FROM teams a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTeamsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTeamByID(c context.Context, node sqalx.Node, id int64) (item *model.Team, err error) {
	item = new(model.Team)
	sqlSelect := "SELECT a.id,a.name,a.env,a.zone,a.deleted,a.created_at,a.updated_at FROM teams a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTeamByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTeamByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Team, err error) {
	item = new(model.Team)
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
	if val, ok := cond["env"]; ok {
		clause += " AND a.env =?"
		condition = append(condition, val)
	}
	if val, ok := cond["zone"]; ok {
		clause += " AND a.zone =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.env,a.zone,a.deleted,a.created_at,a.updated_at FROM teams a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTeamsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTeam(c context.Context, node sqalx.Node, item *model.Team) (err error) {
	sqlInsert := "INSERT INTO teams( id,name,env,zone,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.Env, item.Zone, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTeams err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTeam(c context.Context, node sqalx.Node, item *model.Team) (err error) {
	sqlUpdate := "UPDATE teams SET name=?,env=?,zone=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.Env, item.Zone, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTeams err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTeam(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE teams SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTeams err(%+v), item(%+v)", err, id))
		return
	}

	return
}
