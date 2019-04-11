package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type Article struct {
	ID           int64  `db:"id" json:"id,string"`                 // ID ID
	Title        string `db:"title" json:"title"`                  // Title 标题
	Cover        string `db:"cover" json:"cover"`                  // Cover 文章封面
	Introduction string `db:"introduction" json:"introduction"`    // Introduction 文章简介
	Important    int    `db:"important" json:"important"`          // Important 重要标记
	CreatedBy    int64  `db:"created_by" json:"created_by,string"` // CreatedBy 创建人
	Deleted      int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt    int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type ArticleRepository struct{}

// QueryListPaged get paged records by condition
func (p *ArticleRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*Article, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*Article, 0)

	box := packr.NewBox("./sql/article")
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
func (p *ArticleRepository) GetAll(node sqalx.Node) (items []*Article, err error) {
	items = make([]*Article, 0)
	sqlSelect := packr.NewBox("./sql/article").String("GET_ALL.sql")

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
func (p *ArticleRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*Article, err error) {
	items = make([]*Article, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =:title"
		condition["title"] = val
	}
	if val, ok := cond["cover"]; ok {
		clause += " AND a.cover =:cover"
		condition["cover"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =:important"
		condition["important"] = val
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}

	box := packr.NewBox("./sql/article")
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
func (p *ArticleRepository) GetByID(node sqalx.Node, id int64) (item *Article, exist bool, err error) {
	item = new(Article)
	sqlSelect := packr.NewBox("./sql/article").String("GET_BY_ID.sql")

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
func (p *ArticleRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *Article, exist bool, err error) {
	item = new(Article)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =:title"
		condition["title"] = val
	}
	if val, ok := cond["cover"]; ok {
		clause += " AND a.cover =:cover"
		condition["cover"] = val
	}
	if val, ok := cond["introduction"]; ok {
		clause += " AND a.introduction =:introduction"
		condition["introduction"] = val
	}
	if val, ok := cond["important"]; ok {
		clause += " AND a.important =:important"
		condition["important"] = val
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}

	box := packr.NewBox("./sql/article")
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
func (p *ArticleRepository) Insert(node sqalx.Node, item *Article) (err error) {
	sqlInsert := packr.NewBox("./sql/article").String("INSERT.sql")

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
func (p *ArticleRepository) Update(node sqalx.Node, item *Article) (err error) {
	sqlUpdate := packr.NewBox("./sql/article").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *ArticleRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/article").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *ArticleRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
