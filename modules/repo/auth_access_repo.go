package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type AuthAccess struct {
	ID           int64  `db:"id" json:"id,string"`                 // ID ID
	ClientID     string `db:"client_id" json:"client_id"`          // ClientID Client ID
	Authorize    string `db:"authorize" json:"authorize"`          // Authorize Authorize
	Previous     string `db:"previous" json:"previous"`            // Previous Previous
	AccessToken  string `db:"access_token" json:"access_token"`    // AccessToken Access Token
	RefreshToken string `db:"refresh_token" json:"refresh_token"`  // RefreshToken Refresh Token
	ExpiredIn    int64  `db:"expired_in" json:"expired_in,string"` // ExpiredIn 过期时间
	Scope        string `db:"scope" json:"scope"`                  // Scope Scope
	RedirectURI  string `db:"redirect_uri" json:"redirect_uri"`    // RedirectURI Redirect URI
	Extra        string `db:"extra" json:"extra"`                  // Extra Extra
	Deleted      int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt    int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type AuthAccessRepository struct{}

// QueryListPaged get paged records by condition
func (p *AuthAccessRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*AuthAccess, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*AuthAccess, 0)

	box := packr.NewBox("./sql/auth_access")
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
func (p *AuthAccessRepository) GetAll(node sqalx.Node) (items []*AuthAccess, err error) {
	items = make([]*AuthAccess, 0)
	sqlSelect := packr.NewBox("./sql/auth_access").String("GET_ALL.sql")

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
func (p *AuthAccessRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*AuthAccess, err error) {
	items = make([]*AuthAccess, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id =:client_id"
		condition["client_id"] = val
	}
	if val, ok := cond["authorize"]; ok {
		clause += " AND a.authorize =:authorize"
		condition["authorize"] = val
	}
	if val, ok := cond["previous"]; ok {
		clause += " AND a.previous =:previous"
		condition["previous"] = val
	}
	if val, ok := cond["access_token"]; ok {
		clause += " AND a.access_token =:access_token"
		condition["access_token"] = val
	}
	if val, ok := cond["refresh_token"]; ok {
		clause += " AND a.refresh_token =:refresh_token"
		condition["refresh_token"] = val
	}
	if val, ok := cond["expired_in"]; ok {
		clause += " AND a.expired_in =:expired_in"
		condition["expired_in"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}
	if val, ok := cond["redirect_uri"]; ok {
		clause += " AND a.redirect_uri =:redirect_uri"
		condition["redirect_uri"] = val
	}
	if val, ok := cond["extra"]; ok {
		clause += " AND a.extra =:extra"
		condition["extra"] = val
	}

	box := packr.NewBox("./sql/auth_access")
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
func (p *AuthAccessRepository) GetByID(node sqalx.Node, id int64) (item *AuthAccess, exist bool, err error) {
	item = new(AuthAccess)
	sqlSelect := packr.NewBox("./sql/auth_access").String("GET_BY_ID.sql")

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
func (p *AuthAccessRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *AuthAccess, exist bool, err error) {
	item = new(AuthAccess)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id =:client_id"
		condition["client_id"] = val
	}
	if val, ok := cond["authorize"]; ok {
		clause += " AND a.authorize =:authorize"
		condition["authorize"] = val
	}
	if val, ok := cond["previous"]; ok {
		clause += " AND a.previous =:previous"
		condition["previous"] = val
	}
	if val, ok := cond["access_token"]; ok {
		clause += " AND a.access_token =:access_token"
		condition["access_token"] = val
	}
	if val, ok := cond["refresh_token"]; ok {
		clause += " AND a.refresh_token =:refresh_token"
		condition["refresh_token"] = val
	}
	if val, ok := cond["expired_in"]; ok {
		clause += " AND a.expired_in =:expired_in"
		condition["expired_in"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}
	if val, ok := cond["redirect_uri"]; ok {
		clause += " AND a.redirect_uri =:redirect_uri"
		condition["redirect_uri"] = val
	}
	if val, ok := cond["extra"]; ok {
		clause += " AND a.extra =:extra"
		condition["extra"] = val
	}

	box := packr.NewBox("./sql/auth_access")
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
func (p *AuthAccessRepository) Insert(node sqalx.Node, item *AuthAccess) (err error) {
	sqlInsert := packr.NewBox("./sql/auth_access").String("INSERT.sql")

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
func (p *AuthAccessRepository) Update(node sqalx.Node, item *AuthAccess) (err error) {
	sqlUpdate := packr.NewBox("./sql/auth_access").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *AuthAccessRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/auth_access").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *AuthAccessRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
