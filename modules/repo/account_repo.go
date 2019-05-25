package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/library/database/sqalx"

	types "valerian/library/database/sqlx/types"

	tracerr "github.com/ztrue/tracerr"
)

type Account struct {
	ID           int64         `db:"id" json:"id,string"`                        // ID ID
	Mobile       string        `db:"mobile" json:"mobile"`                       // Mobile 手机
	Email        string        `db:"email" json:"email"`                         // Email 邮件地址
	UserName     string        `db:"user_name" json:"user_name"`                 // UserName 用户名
	Password     string        `db:"password" json:"password"`                   // Password 密码hash
	Role         string        `db:"role" json:"role"`                           // Role 角色
	Salt         string        `db:"salt" json:"salt"`                           // Salt 盐
	Gender       *int          `db:"gender" json:"gender,omitempty"`             // Gender 性别
	BirthYear    *int          `db:"birth_year" json:"birth_year,omitempty"`     // BirthYear 出生年
	BirthMonth   *int          `db:"birth_month" json:"birth_month,omitempty"`   // BirthMonth 出生月
	BirthDay     *int          `db:"birth_day" json:"birth_day,omitempty"`       // BirthDay 出生日
	Location     *int64        `db:"location" json:"location,omitempty,string"`  // Location 地区
	Introduction *string       `db:"introduction" json:"introduction,omitempty"` // Introduction 自我介绍
	Avatar       string        `db:"avatar" json:"avatar"`                       // Avatar 头像
	Source       int           `db:"source" json:"source"`                       // Source 注册来源
	IP           int64         `db:"ip" json:"ip,string"`                        // IP 注册IP
	IDCert       types.BitBool `db:"id_cert" json:"id_cert"`                     // IDCert 是否身份认证
	WorkCert     types.BitBool `db:"work_cert" json:"work_cert"`                 // WorkCert 是否工作认证
	IsOrg        types.BitBool `db:"is_org" json:"is_org"`                       // IsOrg 是否机构用户
	IsVIP        types.BitBool `db:"is_vip" json:"is_vip"`                       // IsVIP 是否VIP用户
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type AccountRepository struct{}

// GetByID get record by ID
func (p *AccountRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *Account, exist bool, err error) {
	item = new(Account)
	sqlSelect := "SELECT a.* FROM accounts a WHERE a.id=? AND a.deleted=0"

	if e := node.GetContext(ctx, item, sqlSelect, id); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return
}

// GetByCondition get record by condition
func (p *AccountRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *Account, exist bool, err error) {
	item = new(Account)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM accounts a WHERE a.deleted=0 %s", clause)

	if e := node.GetContext(ctx, item, sqlSelect, condition...); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return
}

// Insert insert a new record
func (p *AccountRepository) Insert(ctx context.Context, node sqalx.Node, item *Account) (err error) {
	sqlInsert := "INSERT INTO accounts( id,mobile,user_name,email,password,role,salt,gender,birth_year,birth_month,birth_day,location,introduction,avatar,source,ip,id_cert,work_cert,is_org,is_vip,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVIP, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *AccountRepository) Update(ctx context.Context, node sqalx.Node, item *Account) (err error) {
	sqlUpdate := "UPDATE accounts SET mobile=?,user_name=?,email=?,password=?,role=?,salt=?,gender=?,birth_year=?,birth_month=?,birth_day=?,location=?,introduction=?,avatar=?,source=?,ip=?,id_cert=?,work_cert=?,is_org=?,is_vip=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.Mobile, item.UserName, item.Email, item.Password, item.Role, item.Salt, item.Gender, item.BirthYear, item.BirthMonth, item.BirthDay, item.Location, item.Introduction, item.Avatar, item.Source, item.IP, item.IDCert, item.WorkCert, item.IsOrg, item.IsVIP, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
