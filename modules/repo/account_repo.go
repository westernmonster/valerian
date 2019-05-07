package repo

import (
	"database/sql"
	"fmt"
	"time"

	packr "github.com/gobuffalo/packr"
	types "github.com/jmoiron/sqlx/types"
	sqalx "github.com/westernmonster/sqalx"
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
	Deleted      types.BitBool `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64         `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64         `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type AccountRepository struct{}

// QueryListPaged get paged records by condition
func (p *AccountRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Account, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Account, 0)

	box := packr.NewBox("./sql/account")
	sqlCount := fmt.Sprintf(box.String("QUERY_LIST_PAGED_COUNT.sql"), clause)
	sqlSelect := fmt.Sprintf(box.String("QUERY_LIST_PAGED_DATA.sql"), clause)

	stmtCount, err := node.PrepareNamed(sqlCount)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtCount.Get(&total, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	condition["limit"] = pageSize
	condition["offset"] = offset

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAll get all records
func (p *AccountRepository) GetAll(node sqalx.Node) (items []*Account, err error) {
	items = make([]*Account, 0)
	sqlSelect := packr.NewBox("./sql/account").String("GET_ALL.sql")

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, map[string]interface{}{})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *AccountRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*Account, err error) {
	items = make([]*Account, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["mobile"]; ok {
		clause += " AND a.mobile =:mobile"
		condition["mobile"] = val
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =:email"
		condition["email"] = val
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =:password"
		condition["password"] = val
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =:role"
		condition["role"] = val
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =:salt"
		condition["salt"] = val
	}
	if val, ok := cond["gender"]; ok {
		clause += " AND a.gender =:gender"
		condition["gender"] = val
	}
	if val, ok := cond["birth_year"]; ok {
		clause += " AND a.birth_year =:birth_year"
		condition["birth_year"] = val
	}
	if val, ok := cond["birth_month"]; ok {
		clause += " AND a.birth_month =:birth_month"
		condition["birth_month"] = val
	}
	if val, ok := cond["birth_day"]; ok {
		clause += " AND a.birth_day =:birth_day"
		condition["birth_day"] = val
	}
	if val, ok := cond["location"]; ok {
		clause += " AND a.location =:location"
		condition["location"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =:avatar"
		condition["avatar"] = val
	}
	if val, ok := cond["source"]; ok {
		clause += " AND a.source =:source"
		condition["source"] = val
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =:ip"
		condition["ip"] = val
	}

	box := packr.NewBox("./sql/account")
	sqlSelect := fmt.Sprintf(box.String("GET_ALL_BY_CONDITION.sql"), clause)

	stmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtSelect.Select(&items, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get record by ID
func (p *AccountRepository) GetByID(node sqalx.Node, id int64) (item *Account, exist bool, err error) {
	item = new(Account)
	sqlSelect := packr.NewBox("./sql/account").String("GET_BY_ID.sql")

	tmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if e := tmtSelect.Get(item, map[string]interface{}{"id": id}); e != nil {
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
func (p *AccountRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *Account, exist bool, err error) {
	item = new(Account)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["mobile"]; ok {
		clause += " AND a.mobile =:mobile"
		condition["mobile"] = val
	}
	if val, ok := cond["email"]; ok {
		clause += " AND a.email =:email"
		condition["email"] = val
	}
	if val, ok := cond["password"]; ok {
		clause += " AND a.password =:password"
		condition["password"] = val
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =:role"
		condition["role"] = val
	}
	if val, ok := cond["salt"]; ok {
		clause += " AND a.salt =:salt"
		condition["salt"] = val
	}
	if val, ok := cond["gender"]; ok {
		clause += " AND a.gender =:gender"
		condition["gender"] = val
	}
	if val, ok := cond["birth_year"]; ok {
		clause += " AND a.birth_year =:birth_year"
		condition["birth_year"] = val
	}
	if val, ok := cond["birth_month"]; ok {
		clause += " AND a.birth_month =:birth_month"
		condition["birth_month"] = val
	}
	if val, ok := cond["birth_day"]; ok {
		clause += " AND a.birth_day =:birth_day"
		condition["birth_day"] = val
	}
	if val, ok := cond["location"]; ok {
		clause += " AND a.location =:location"
		condition["location"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["avatar"]; ok {
		clause += " AND a.avatar =:avatar"
		condition["avatar"] = val
	}
	if val, ok := cond["source"]; ok {
		clause += " AND a.source =:source"
		condition["source"] = val
	}
	if val, ok := cond["ip"]; ok {
		clause += " AND a.ip =:ip"
		condition["ip"] = val
	}

	box := packr.NewBox("./sql/account")
	sqlSelect := fmt.Sprintf(box.String("GET_BY_CONDITION.sql"), clause)

	tmtSelect, err := node.PrepareNamed(sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if e := tmtSelect.Get(item, condition); e != nil {
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
func (p *AccountRepository) Insert(node sqalx.Node, item *Account) (err error) {
	sqlInsert := packr.NewBox("./sql/account").String("INSERT.sql")

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlInsert, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *AccountRepository) Update(node sqalx.Node, item *Account) (err error) {
	sqlUpdate := packr.NewBox("./sql/account").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *AccountRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/account").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *AccountRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
	tx, err := node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	defer tx.Rollback()
	for _, id := range ids {
		errDelete := p.Delete(tx, id)
		if errDelete != nil {
			err = tracerr.Wrap(err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
