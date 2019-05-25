package repo

import (
	"context"
	"database/sql"
	"time"
	"valerian/library/database/sqalx"
	types "valerian/library/database/sqlx/types"

	tracerr "github.com/ztrue/tracerr"
)

type Session struct {
	ID          int64         `db:"id" json:"id,string"`                 // ID ID
	SessionType int           `db:"session_type" json:"session_type"`    // SessionType 类型
	Used        int           `db:"used" json:"used"`                    // Used 类型, 0未使用，1使用
	AccountID   int64         `db:"account_id" json:"account_id,string"` // AccountID 账户ID
	Deleted     types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type SessionRepository struct{}

// GetByID get record by ID
func (p *SessionRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *Session, exist bool, err error) {
	item = new(Session)
	sqlSelect := "SELECT a.* FROM session a WHERE a.id=? AND a.deleted=0"

	if e := node.GetContext(ctx, item, sqlSelect); e != nil {
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
func (p *SessionRepository) Insert(ctx context.Context, node sqalx.Node, item *Session) (err error) {
	sqlInsert := "INSERT INTO session( id,session_type,used,account_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.SessionType, item.Used, item.AccountID, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
