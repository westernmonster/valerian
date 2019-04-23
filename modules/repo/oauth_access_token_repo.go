package repo

import (
	"database/sql"
	"fmt"
	"time"

	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
)

type OauthAccessToken struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	ClientID  string `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64  `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64  `db:"expires_at" json:"expires_at,string"` // ExpiresAt 过期时间
	Scope     string `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type OauthAccessTokenRepository struct{}

// QueryListPaged get paged records by condition
func (p *OauthAccessTokenRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*OauthAccessToken, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*OauthAccessToken, 0)

	box := packr.NewBox("./sql/oauth_access_token")
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
func (p *OauthAccessTokenRepository) GetAll(node sqalx.Node) (items []*OauthAccessToken, err error) {
	items = make([]*OauthAccessToken, 0)
	sqlSelect := packr.NewBox("./sql/oauth_access_token").String("GET_ALL.sql")

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
func (p *OauthAccessTokenRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*OauthAccessToken, err error) {
	items = make([]*OauthAccessToken, 0)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =:token"
		condition["token"] = val
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =:expires_at"
		condition["expires_at"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}

	box := packr.NewBox("./sql/oauth_access_token")
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
func (p *OauthAccessTokenRepository) GetByID(node sqalx.Node, id int64) (item *OauthAccessToken, exist bool, err error) {
	item = new(OauthAccessToken)
	sqlSelect := packr.NewBox("./sql/oauth_access_token").String("GET_BY_ID.sql")

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
func (p *OauthAccessTokenRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *OauthAccessToken, exist bool, err error) {
	item = new(OauthAccessToken)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =:token"
		condition["token"] = val
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =:expires_at"
		condition["expires_at"] = val
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =:scope"
		condition["scope"] = val
	}

	box := packr.NewBox("./sql/oauth_access_token")
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
func (p *OauthAccessTokenRepository) Insert(node sqalx.Node, item *OauthAccessToken) (err error) {
	sqlInsert := packr.NewBox("./sql/oauth_access_token").String("INSERT.sql")

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
func (p *OauthAccessTokenRepository) Update(node sqalx.Node, item *OauthAccessToken) (err error) {
	sqlUpdate := packr.NewBox("./sql/oauth_access_token").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *OauthAccessTokenRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/oauth_access_token").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *OauthAccessTokenRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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

// Delete delete records  by condition
func (p *OauthAccessTokenRepository) DeleteByCondition(node sqalx.Node, cond map[string]string) (err error) {
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id =:client_id"
		condition["client_id"] = val
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}

	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at <=:expires_at"
		condition["expires_at"] = val
	}

	box := packr.NewBox("./sql/oauth_access_token")
	sqlDelete := fmt.Sprintf(box.String("DELETE_BY_CONDITION.sql"), clause)

	_, err = node.NamedExec(sqlDelete, condition)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
