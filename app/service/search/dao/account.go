package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetAccounts(c context.Context, node sqalx.Node) (items []*model.Account, err error) {
	items = make([]*model.Account, 0)
	sqlSelect := "SELECT a.* FROM accounts a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccounts err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAccountsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Account, err error) {
	items = make([]*model.Account, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mobile"]; ok {
		clause += " AND a.mobile =?"
		condition = append(condition, val)
	}
	if val, ok := cond["user_name"]; ok {
		clause += " AND a.user_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =?"
		condition = append(condition, val)
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =?"
		condition = append(condition, val)
	}
	if val, ok := cond["gender"]; ok {
		clause += " AND a.gender =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_year"]; ok {
		clause += " AND a.birth_year =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_month"]; ok {
		clause += " AND a.birth_month =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_day"]; ok {
		clause += " AND a.birth_day =?"
		condition = append(condition, val)
	}
	if val, ok := cond["location"]; ok {
		clause += " AND a.location =?"
		condition = append(condition, val)
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =?"
		condition = append(condition, val)
	}
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =?"
		condition = append(condition, val)
	}
	if val, ok := cond["source"]; ok {
		clause += " AND a.source =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_cert"]; ok {
		clause += " AND a.id_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["work_cert"]; ok {
		clause += " AND a.work_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_org"]; ok {
		clause += " AND a.is_org =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_vip"]; ok {
		clause += " AND a.is_vip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["prefix"]; ok {
		clause += " AND a.prefix =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM accounts a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	item = new(model.Account)
	sqlSelect := "SELECT a.* FROM accounts a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAccountByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Account, err error) {
	item = new(model.Account)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mobile"]; ok {
		clause += " AND a.mobile =?"
		condition = append(condition, val)
	}
	if val, ok := cond["user_name"]; ok {
		clause += " AND a.user_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =?"
		condition = append(condition, val)
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =?"
		condition = append(condition, val)
	}
	if val, ok := cond["gender"]; ok {
		clause += " AND a.gender =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_year"]; ok {
		clause += " AND a.birth_year =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_month"]; ok {
		clause += " AND a.birth_month =?"
		condition = append(condition, val)
	}
	if val, ok := cond["birth_day"]; ok {
		clause += " AND a.birth_day =?"
		condition = append(condition, val)
	}
	if val, ok := cond["location"]; ok {
		clause += " AND a.location =?"
		condition = append(condition, val)
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =?"
		condition = append(condition, val)
	}
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =?"
		condition = append(condition, val)
	}
	if val, ok := cond["source"]; ok {
		clause += " AND a.source =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["id_cert"]; ok {
		clause += " AND a.id_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["work_cert"]; ok {
		clause += " AND a.work_cert =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_org"]; ok {
		clause += " AND a.is_org =?"
		condition = append(condition, val)
	}
	if val, ok := cond["is_vip"]; ok {
		clause += " AND a.is_vip =?"
		condition = append(condition, val)
	}
	if val, ok := cond["prefix"]; ok {
		clause += " AND a.prefix =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM accounts a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	sqlInsert := "INSERT INTO accounts( id,mobile,user_name,email,password,role,salt,gender,birth_year,birth_month,birth_day,location,introduction,avatar,source,ip,id_cert,work_cert,is_org,is_vip,deleted,created_at,updated_at,prefix) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVip, item.Deleted, item.CreatedAt, item.UpdatedAt, item.Prefix); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccounts err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	sqlUpdate := "UPDATE accounts SET mobile=?,user_name=?,email=?,password=?,role=?,salt=?,gender=?,birth_year=?,birth_month=?,birth_day=?,location=?,introduction=?,avatar=?,source=?,ip=?,id_cert=?,work_cert=?,is_org=?,is_vip=?,updated_at=?,prefix=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVip, item.UpdatedAt, item.Prefix, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccounts err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccount(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE accounts SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccounts err(%+v), item(%+v)", err, id))
		return
	}

	return
}
