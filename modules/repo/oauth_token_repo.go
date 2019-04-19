package repo

import (
	"database/sql"
	"fmt"
	"time"

	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
)

type OAUTHToken struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	ExpiredAt int64  `db:"expired_at" json:"expired_at,string"` // ExpiredAt 过期时间
	Code      string `db:"code" json:"code"`                    // Code Authorization code
	Access    string `db:"access" json:"access"`                // Access Authorization code
	Refresh   string `db:"refresh" json:"refresh"`              // Refresh Authorization code
	Data      string `db:"data" json:"data"`                    // Data Authorization code
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type OAUTHTokenRepository struct{}

// QueryListPaged get paged records by condition
func (p *OAUTHTokenRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*OAUTHToken, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*OAUTHToken, 0)

	box := packr.NewBox("./sql/oauth_token")
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
func (p *OAUTHTokenRepository) GetAll(node sqalx.Node) (items []*OAUTHToken, err error) {
	items = make([]*OAUTHToken, 0)
	sqlSelect := packr.NewBox("./sql/oauth_token").String("GET_ALL.sql")

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
func (p *OAUTHTokenRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*OAUTHToken, err error) {
	items = make([]*OAUTHToken, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["expired_at"]; ok {
		clause += " AND a.expired_at =:expired_at"
		condition["expired_at"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["access"]; ok {
		clause += " AND a.access =:access"
		condition["access"] = val
	}
	if val, ok := cond["refresh"]; ok {
		clause += " AND a.refresh =:refresh"
		condition["refresh"] = val
	}
	if val, ok := cond["data"]; ok {
		clause += " AND a.data =:data"
		condition["data"] = val
	}

	box := packr.NewBox("./sql/oauth_token")
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
func (p *OAUTHTokenRepository) GetByID(node sqalx.Node, id int64) (item *OAUTHToken, exist bool, err error) {
	item = new(OAUTHToken)
	sqlSelect := packr.NewBox("./sql/oauth_token").String("GET_BY_ID.sql")

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
func (p *OAUTHTokenRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *OAUTHToken, exist bool, err error) {
	item = new(OAUTHToken)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["expired_at"]; ok {
		clause += " AND a.expired_at =:expired_at"
		condition["expired_at"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["access"]; ok {
		clause += " AND a.access =:access"
		condition["access"] = val
	}
	if val, ok := cond["refresh"]; ok {
		clause += " AND a.refresh =:refresh"
		condition["refresh"] = val
	}
	if val, ok := cond["data"]; ok {
		clause += " AND a.data =:data"
		condition["data"] = val
	}

	box := packr.NewBox("./sql/oauth_token")
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
func (p *OAUTHTokenRepository) Insert(node sqalx.Node, item *OAUTHToken) (err error) {
	sqlInsert := packr.NewBox("./sql/oauth_token").String("INSERT.sql")

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
func (p *OAUTHTokenRepository) Update(node sqalx.Node, item *OAUTHToken) (err error) {
	sqlUpdate := packr.NewBox("./sql/oauth_token").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *OAUTHTokenRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/oauth_token").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *OAUTHTokenRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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

func (p *OAUTHTokenRepository) DeleteExpired(node sqalx.Node, expiredAt int64) (err error) {
	sqlDelete := packr.NewBox("./sql/oauth_token").String("DELETE_EXPIRED.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"expired_at": expiredAt})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *OAUTHTokenRepository) GetExpiredCount(node sqalx.Node, expiredAt int64) (total int, err error) {

	sqlCount := packr.NewBox("./sql/oauth_token").String("GET_EXPIRED.sql")

	stmtCount, err := node.PrepareNamed(sqlCount)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	err = stmtCount.Get(&total, map[string]interface{}{
		"expired_at": expiredAt,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
