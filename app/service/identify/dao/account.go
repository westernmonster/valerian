package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/log"
)

func (p *Dao) GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error) {
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.email=? AND a.deleted=0 AND a.deactive=0"
	item = &model.Account{}
	if err = node.GetContext(c, item, sqlSelect, email); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail error(%+v), email(%s)", err, email))
	}

	return
}

func (p *Dao) GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error) {
	item = &model.Account{}
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.mobile=? AND a.deleted=0 AND a.deactive=0"
	if err = node.GetContext(c, item, sqlSelect, mobile); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByMobile error(%+v), email(%s)", err, mobile))
	}

	return
}

func (p *Dao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	item = &model.Account{}
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.id=? AND a.deleted=0"
	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByID error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) SetPassword(c context.Context, node sqalx.Node, password, salt string, id int64) (err error) {
	sqlSelect := "UPDATE accounts SET password=?, salt =? WHERE id=? AND deleted=0"
	if _, err = node.ExecContext(c, sqlSelect, password, salt, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SetPassword error(%+v), id(%d)", err, id))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	sqlInsert := "INSERT INTO accounts( id,mobile,user_name,email,password,role,salt,gender,birth_year,birth_month,birth_day,location,introduction,avatar,source,ip,id_cert,work_cert,is_org,is_vip,deleted,created_at,updated_at,prefix,is_lock,deactive) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVip, item.Deleted, item.CreatedAt, item.UpdatedAt, item.Prefix, item.IsLock, item.Deactive); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccounts err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error) {
	sqlUpdate := "UPDATE accounts SET mobile=?,user_name=?,email=?,password=?,role=?,salt=?,gender=?,birth_year=?,birth_month=?,birth_day=?,location=?,introduction=?,avatar=?,source=?,ip=?,id_cert=?,work_cert=?,is_org=?,is_vip=?,updated_at=?,prefix=?,is_lock=?,deactive=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVip, item.UpdatedAt, item.Prefix, item.IsLock, item.Deactive, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccounts err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) GetAccounts(c context.Context, node sqalx.Node) (items []*model.Account, err error) {
	items = make([]*model.Account, 0)
	sqlSelect := "SELECT a.id,a.mobile,a.user_name,a.email,a.password,a.role,a.salt,a.gender,a.birth_year,a.birth_month,a.birth_day,a.location,a.introduction,a.avatar,a.source,a.ip,a.id_cert,a.work_cert,a.is_org,a.is_vip,a.deleted,a.created_at,a.updated_at,a.prefix,a.is_lock,a.deactive FROM accounts a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccounts err(%+v)", err))
		return
	}
	return
}

func (p *Dao) DeactiveAccount(c context.Context, node sqalx.Node, aid int64) (err error) {
	sqlSelect := "UPDATE accounts SET user_name=?, deactive=?, updated_at=? WHERE id=? AND deleted=0"
	if _, err = node.ExecContext(c, sqlSelect, "已注销", types.BitBool(true), aid, time.Now().Unix()); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DeactiveAccount error(%+v), id(%d)", err, aid))
		return
	}
	return
}

func (p *Dao) AccountSetLock(c context.Context, node sqalx.Node, accountID int64, isLock bool) (err error) {
	sqlUpdate := "UPDATE accounts SET is_lock=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, types.BitBool(isLock), accountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AccountSetLock err(%+v)", err))
		return
	}

	return
}
