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

type CountryCode struct {
	ID        int64  `db:"id" json:"id,string"`          // ID ID
	Name      string `db:"name" json:"name"`             // Name 国家名
	Emoji     string `db:"emoji" json:"emoji"`           // Emoji 国旗
	CnName    string `db:"cn_name" json:"cn_name"`       // CnName 国家中文名
	Code      string `db:"code" json:"code"`             // Code 编码
	Prefix    string `db:"prefix" json:"prefix"`         // Prefix 前缀
	Deleted   int    `db:"deleted" json:"deleted"`       // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"` // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"` // UpdatedAt 更新时间
}

type CountryCodeRepository struct{}

// QueryListPaged get paged records by condition
func (p *CountryCodeRepository) QueryListPaged(ctx context.Context, node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*CountryCode, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*CountryCode, 0)

	box := packr.NewBox("./sql/country_code")
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
func (p *CountryCodeRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*CountryCode, err error) {
	items = make([]*CountryCode, 0)
	sqlSelect := "SELECT a.* FROM country_codes a WHERE a.deleted=0 ORDER BY a.id "

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *CountryCodeRepository) GetAllByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*CountryCode, err error) {
	items = make([]*CountryCode, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["cn_name"]; ok {
		clause += " AND a.cn_name =:cn_name"
		condition["cn_name"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["prefix"]; ok {
		clause += " AND a.prefix =:prefix"
		condition["prefix"] = val
	}

	box := packr.NewBox("./sql/country_code")
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
func (p *CountryCodeRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *CountryCode, exist bool, err error) {
	item = new(CountryCode)
	sqlSelect := packr.NewBox("./sql/country_code").String("GET_BY_ID.sql")

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
func (p *CountryCodeRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *CountryCode, exist bool, err error) {
	item = new(CountryCode)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["cn_name"]; ok {
		clause += " AND a.cn_name =:cn_name"
		condition["cn_name"] = val
	}
	if val, ok := cond["code"]; ok {
		clause += " AND a.code =:code"
		condition["code"] = val
	}
	if val, ok := cond["prefix"]; ok {
		clause += " AND a.prefix =:prefix"
		condition["prefix"] = val
	}

	box := packr.NewBox("./sql/country_code")
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
func (p *CountryCodeRepository) Insert(ctx context.Context, node sqalx.Node, item *CountryCode) (err error) {
	sqlInsert := packr.NewBox("./sql/country_code").String("INSERT.sql")

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
func (p *CountryCodeRepository) Update(ctx context.Context, node sqalx.Node, item *CountryCode) (err error) {
	sqlUpdate := packr.NewBox("./sql/country_code").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExecContext(ctx, sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *CountryCodeRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/country_code").String("DELETE.sql")

	_, err = node.NamedExecContext(ctx, sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *CountryCodeRepository) BatchDelete(ctx context.Context, node sqalx.Node, ids []int64) (err error) {
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
