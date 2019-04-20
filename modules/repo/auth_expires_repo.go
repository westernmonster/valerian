package repo

import (
	"database/sql"
	"fmt"
	"time"

	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
)

type AuthExpires struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	Token     string `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64  `db:"expires_at" json:"expires_at,string"` // ExpiresAt 过期时间
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AuthExpiresRepository struct{}

// QueryListPaged get paged records by condition
func (p *AuthExpiresRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*AuthExpires, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*AuthExpires, 0)

	box := packr.NewBox("./sql/auth_expires")
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
func (p *AuthExpiresRepository) GetAll(node sqalx.Node) (items []*AuthExpires, err error) {
	items = make([]*AuthExpires, 0)
	sqlSelect := packr.NewBox("./sql/auth_expires").String("GET_ALL.sql")

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
func (p *AuthExpiresRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*AuthExpires, err error) {
	items = make([]*AuthExpires, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =:token"
		condition["token"] = val
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =:expires_at"
		condition["expires_at"] = val
	}

	box := packr.NewBox("./sql/auth_expires")
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
func (p *AuthExpiresRepository) GetByID(node sqalx.Node, id int64) (item *AuthExpires, exist bool, err error) {
	item = new(AuthExpires)
	sqlSelect := packr.NewBox("./sql/auth_expires").String("GET_BY_ID.sql")

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
func (p *AuthExpiresRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *AuthExpires, exist bool, err error) {
	item = new(AuthExpires)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =:token"
		condition["token"] = val
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =:expires_at"
		condition["expires_at"] = val
	}

	box := packr.NewBox("./sql/auth_expires")
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
func (p *AuthExpiresRepository) Insert(node sqalx.Node, item *AuthExpires) (err error) {
	sqlInsert := packr.NewBox("./sql/auth_expires").String("INSERT.sql")

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
func (p *AuthExpiresRepository) Update(node sqalx.Node, item *AuthExpires) (err error) {
	sqlUpdate := packr.NewBox("./sql/auth_expires").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *AuthExpiresRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/auth_expires").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *AuthExpiresRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
