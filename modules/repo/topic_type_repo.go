package repo

import (
	"context"
	"database/sql"

	"valerian/library/database/sqalx"

	"github.com/jmoiron/sqlx/types"
	tracerr "github.com/ztrue/tracerr"
)

type TopicType struct {
	ID        int           `db:"id" json:"id"`                 // ID ID
	Name      string        `db:"name" json:"name"`             // Name 话题类型
	Deleted   types.BitBool `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type TopicTypeRepository struct{}

// GetAll get all records
func (p *TopicTypeRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*TopicType, err error) {
	items = make([]*TopicType, 0)
	sqlSelect := "SELECT a.* FROM topic_types a WHERE a.deleted=0 ORDER BY a.id DESC "

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get a record by ID
func (p *TopicTypeRepository) GetByID(ctx context.Context, node sqalx.Node, id int) (item *TopicType, exist bool, err error) {
	item = new(TopicType)
	sqlSelect := "SELECT a.* FROM topic_types a WHERE a.id=? AND a.deleted=0"

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
