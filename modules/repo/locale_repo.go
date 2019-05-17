package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type Locale struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	Locale    string `db:"locale" json:"locale"`         // Locale 语言编码
	Name      string `db:"name" json:"name"`             // Name 语言名称
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type LocaleRepository struct{}

// QueryListPaged get paged records by condition
func (p *LocaleRepository) QueryListPaged(ctx context.Context, node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Locale, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Locale, 0)

	box := packr.NewBox("./sql/locale")
	sqlCount := fmt.Sprintf(box.String("QUERY_LIST_PAGED_COUNT.sql"), clause)
	sqlSelect := fmt.Sprintf(box.String("QUERY_LIST_PAGED_DATA.sql"), clause)

	stmtCount, err := node.PrepareNamedContext(ctx, sqlCount)
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

	stmtSelect, err := node.PrepareNamedContext(ctx, sqlSelect)
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
func (p *LocaleRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*Locale, err error) {
	items = make([]*Locale, 0)
	sqlSelect := packr.NewBox("./sql/locale").String("GET_ALL.sql")

	stmtSelect, err := node.PrepareNamedContext(ctx, sqlSelect)
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
func (p *LocaleRepository) GetAllByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*Locale, err error) {
	items = make([]*Locale, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =:locale"
		condition["locale"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}

	box := packr.NewBox("./sql/locale")
	sqlSelect := fmt.Sprintf(box.String("GET_ALL_BY_CONDITION.sql"), clause)

	stmtSelect, err := node.PrepareNamedContext(ctx, sqlSelect)
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
func (p *LocaleRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *Locale, exist bool, err error) {
	item = new(Locale)
	sqlSelect := packr.NewBox("./sql/locale").String("GET_BY_ID.sql")

	tmtSelect, err := node.PrepareNamedContext(ctx, sqlSelect)
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
func (p *LocaleRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *Locale, exist bool, err error) {
	item = new(Locale)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["locale"]; ok {
		clause += " AND a.locale =:locale"
		condition["locale"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}

	box := packr.NewBox("./sql/locale")
	sqlSelect := fmt.Sprintf(box.String("GET_BY_CONDITION.sql"), clause)

	tmtSelect, err := node.PrepareNamedContext(ctx, sqlSelect)
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
func (p *LocaleRepository) Insert(ctx context.Context, node sqalx.Node, item *Locale) (err error) {
	sqlInsert := packr.NewBox("./sql/locale").String("INSERT.sql")

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExecContext(ctx, sqlInsert, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *LocaleRepository) Update(ctx context.Context, node sqalx.Node, item *Locale) (err error) {
	sqlUpdate := packr.NewBox("./sql/locale").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExecContext(ctx, sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *LocaleRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/locale").String("DELETE.sql")

	_, err = node.NamedExecContext(ctx, sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *LocaleRepository) BatchDelete(ctx context.Context, node sqalx.Node, ids []int64) (err error) {
	tx, err := node.Beginx(ctx)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	defer tx.Rollback()
	for _, id := range ids {
		errDelete := p.Delete(ctx, tx, id)
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
