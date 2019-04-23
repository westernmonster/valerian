package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type OauthScope struct {
	ID          int64  `db:"id" json:"id,string"`            // ID ID
	Scope       string `db:"scope" json:"scope"`             // Scope Scope
	Description string `db:"description" json:"description"` // Description Description
	IsDefault   int    `db:"is_default" json:"is_default"`   // IsDefault 是否默认, 0 否，1 是
	Deleted     int    `db:"deleted" json:"deleted"`         // Deleted 是否删除
	CreatedAt   int64  `db:"created_at" json:"created_at"`   // CreatedAt 创建时间
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`   // UpdatedAt 更新时间
}

type OauthScopeRepository struct{}

// QueryListPaged get paged records by condition
func (p *OauthScopeRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*OauthScope, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*OauthScope, 0)

	box := packr.NewBox("./sql/oauth_scope")
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
func (p *OauthScopeRepository) GetAll(node sqalx.Node) (items []*OauthScope, err error) {
	items = make([]*OauthScope, 0)
	sqlSelect := packr.NewBox("./sql/oauth_scope").String("GET_ALL.sql")

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
func (p *OauthScopeRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*OauthScope, err error) {
	items = make([]*OauthScope, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}
	if val, ok := cond["description"]; ok {
		clause += " AND a.description =:description"
		condition["description"] = val
	}
	if val, ok := cond["is_default"]; ok {
		clause += " AND a.is_default =:is_default"
		condition["is_default"] = val
	}

	box := packr.NewBox("./sql/oauth_scope")
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
func (p *OauthScopeRepository) GetByID(node sqalx.Node, id int64) (item *OauthScope, exist bool, err error) {
	item = new(OauthScope)
	sqlSelect := packr.NewBox("./sql/oauth_scope").String("GET_BY_ID.sql")

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
func (p *OauthScopeRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *OauthScope, exist bool, err error) {
	item = new(OauthScope)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}
	if val, ok := cond["description"]; ok {
		clause += " AND a.description =:description"
		condition["description"] = val
	}
	if val, ok := cond["is_default"]; ok {
		clause += " AND a.is_default =:is_default"
		condition["is_default"] = val
	}

	box := packr.NewBox("./sql/oauth_scope")
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
func (p *OauthScopeRepository) Insert(node sqalx.Node, item *OauthScope) (err error) {
	sqlInsert := packr.NewBox("./sql/oauth_scope").String("INSERT.sql")

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
func (p *OauthScopeRepository) Update(node sqalx.Node, item *OauthScope) (err error) {
	sqlUpdate := packr.NewBox("./sql/oauth_scope").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *OauthScopeRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/oauth_scope").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *OauthScopeRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
