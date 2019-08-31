package dao

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/msm/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Codes get all codes.
func (d *Dao) Codes(c context.Context) (codes map[int]string, lcode *model.Code, err error) {
	return
}

// Diff get change codes.
func (d *Dao) Diff(c context.Context, node sqalx.Node, ver int64) (vers *list.List, err error) {
	items := make([]*model.Code, 0)
	sqlSelect := "SELECT a.* FROM codes a WHERE a.deleted=0 AND a.created_at > ? ORDER BY a.created_at LIMIT 100 "

	if err = node.SelectContext(c, &items, sqlSelect, ver); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DiffCodes err(%+v)", err))
		return
	}

	vers = list.New()
	for _, v := range items {
		vers.PushBack(v)
	}
	return
}

// CodesLang get all codes.
func (d *Dao) CodesLang(c context.Context, node sqalx.Node) (codes map[int]map[string]string, lcode *model.CodeLangs, err error) {
	items := make([]*model.Code, 0)
	sqlSelect := "SELECT a.* FROM codes a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCodes err(%+v)", err))
		return
	}

	codes = make(map[int]map[string]string)
	for _, v := range items {
		// codes[v.Code]
	}
	return
}

// GetAll get all records
func (p *Dao) GetCodes(c context.Context, node sqalx.Node) (items []*model.Code, err error) {
	items = make([]*model.Code, 0)
	sqlSelect := "SELECT a.* FROM codes a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCodes err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetCodesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Code, err error) {
	items = make([]*model.Code, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =?"
		condition = append(condition, val)
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =?"
		condition = append(condition, val)
	}
	if val, ok := cond["message"]; ok {
		clause += " AND a.message =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM codes a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCodesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetCodeByID(c context.Context, node sqalx.Node, id int64) (item *model.Code, err error) {
	item = new(model.Code)
	sqlSelect := "SELECT a.* FROM codes a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCodeByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetCodeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Code, err error) {
	item = new(model.Code)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =?"
		condition = append(condition, val)
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =?"
		condition = append(condition, val)
	}
	if val, ok := cond["message"]; ok {
		clause += " AND a.message =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM codes a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCodesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddCode(c context.Context, node sqalx.Node, item *model.Code) (err error) {
	sqlInsert := "INSERT INTO codes( id,code,locale,message,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Code, item.Locale, item.Message, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddCodes err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateCode(c context.Context, node sqalx.Node, item *model.Code) (err error) {
	sqlUpdate := "UPDATE codes SET code=?,locale=?,message=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Code, item.Locale, item.Message, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateCodes err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelCode(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE codes SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelCodes err(%+v), item(%+v)", err, id))
		return
	}

	return
}
