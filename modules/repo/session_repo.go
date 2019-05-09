package repo

import (
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type Session struct {
	ID          int64 `db:"id" json:"id,string"`                 // ID ID
	SessionType int   `db:"session_type" json:"session_type"`    // SessionType 类型
	Used        int   `db:"used" json:"used"`                    // Used 类型, 0未使用，1使用
	AccountID   int64 `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	Deleted     int   `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64 `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type SessionRepository struct{}

// QueryListPaged get paged records by condition
func (p *SessionRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Session, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Session, 0)

	box := packr.NewBox("./sql/session")
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
func (p *SessionRepository) GetAll(node sqalx.Node) (items []*Session, err error) {
	items = make([]*Session, 0)
	sqlSelect := packr.NewBox("./sql/session").String("GET_ALL.sql")

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
func (p *SessionRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*Session, err error) {
	items = make([]*Session, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["session_type"]; ok {
		clause += " AND a.session_type =:session_type"
		condition["session_type"] = val
	}
	if val, ok := cond["used"]; ok {
		clause += " AND a.used =:used"
		condition["used"] = val
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}

	box := packr.NewBox("./sql/session")
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
func (p *SessionRepository) GetByID(node sqalx.Node, id int64) (item *Session, exist bool, err error) {
	item = new(Session)
	sqlSelect := packr.NewBox("./sql/session").String("GET_BY_ID.sql")

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
func (p *SessionRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *Session, exist bool, err error) {
	item = new(Session)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["session_type"]; ok {
		clause += " AND a.session_type =:session_type"
		condition["session_type"] = val
	}
	if val, ok := cond["used"]; ok {
		clause += " AND a.used =:used"
		condition["used"] = val
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}

	box := packr.NewBox("./sql/session")
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
func (p *SessionRepository) Insert(node sqalx.Node, item *Session) (err error) {
	sqlInsert := packr.NewBox("./sql/session").String("INSERT.sql")

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
func (p *SessionRepository) Update(node sqalx.Node, item *Session) (err error) {
	sqlUpdate := packr.NewBox("./sql/session").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *SessionRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/session").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *SessionRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
