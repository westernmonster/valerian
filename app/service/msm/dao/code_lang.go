package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/msm/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) CodesLang(c context.Context, node sqalx.Node) (codes map[int]map[string]string, lcode *model.CodeLangs, err error) {
	sqlSelect := "SELECT a.locale,a.code,a.message,a.updated_at FROM code_langs a WHERE a.deleted=0 ORDER BY a.id "

	var (
		tmp int64
	)

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CodesLang err(%+v)", err))
		return
	}
	defer rows.Close()
	lcode = &model.CodeLangs{}
	codes = make(map[int]map[string]string)

	for rows.Next() {
		var (
			code      int
			message   string
			locale    string
			updatedAt int64
		)
		t := make(map[string]string)
		if err = rows.Scan(&locale, &code, &message, &updatedAt); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.CodesLang err(%+v)", err))
			return
		}

		if len(locale) > 0 {
			t[locale] = message
		}

		codes[code] = t
		if updatedAt > tmp {
			lcode.Code = code
			lcode.Ver = updatedAt
			lcode.Msg = t
			tmp = updatedAt
		}
	}

	err = rows.Err()
	return
}

// GetAll get all records
func (p *Dao) GetCodeLangs(c context.Context, node sqalx.Node) (items []*model.CodeLang, err error) {
	items = make([]*model.CodeLang, 0)
	sqlSelect := "SELECT a.id,a.locale,a.code,a.message,a.deleted,a.created_at,a.updated_at FROM code_langs a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCodeLangs err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetCodeLangsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.CodeLang, err error) {
	items = make([]*model.CodeLang, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =?"
		condition = append(condition, val)
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =?"
		condition = append(condition, val)
	}
	if val, ok := cond["message"]; ok {
		clause += " AND a.message =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.locale,a.code,a.message,a.deleted,a.created_at,a.updated_at FROM code_langs a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCodeLangsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetCodeLangByID(c context.Context, node sqalx.Node, id int64) (item *model.CodeLang, err error) {
	item = new(model.CodeLang)
	sqlSelect := "SELECT a.id,a.locale,a.code,a.message,a.deleted,a.created_at,a.updated_at FROM code_langs a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCodeLangByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetCodeLangByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.CodeLang, err error) {
	item = new(model.CodeLang)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =?"
		condition = append(condition, val)
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =?"
		condition = append(condition, val)
	}
	if val, ok := cond["message"]; ok {
		clause += " AND a.message =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.locale,a.code,a.message,a.deleted,a.created_at,a.updated_at FROM code_langs a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCodeLangsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddCodeLang(c context.Context, node sqalx.Node, item *model.CodeLang) (err error) {
	sqlInsert := "INSERT INTO code_langs( id,locale,code,message,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Locale, item.Code, item.Message, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddCodeLangs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateCodeLang(c context.Context, node sqalx.Node, item *model.CodeLang) (err error) {
	sqlUpdate := "UPDATE code_langs SET locale=?,code=?,message=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Locale, item.Code, item.Message, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateCodeLangs err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelCodeLang(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE code_langs SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelCodeLangs err(%+v), item(%+v)", err, id))
		return
	}

	return
}
