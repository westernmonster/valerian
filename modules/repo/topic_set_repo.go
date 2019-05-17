package repo

import (
	"context"
	"time"

	"valerian/library/database/sqalx"

	"github.com/jmoiron/sqlx/types"
	tracerr "github.com/ztrue/tracerr"
)

type TopicSet struct {
	ID        int64         `db:"id" json:"id,string"`          // ID ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type TopicSetRepository struct{}

// Insert insert a new record
func (p *TopicSetRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicSet) (err error) {
	sqlInsert := "INSERT INTO topic_sets( id,deleted,created_at,updated_at) VALUES ( ?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
