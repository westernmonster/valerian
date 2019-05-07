package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type TopicCategory struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64  `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	Name      string `db:"name" json:"name"`                    // Name 分类名
	ParentID  int64  `db:"parent_id" json:"parent_id,string"`   // ParentID 父级ID, 一级分类的父ID为 0
	CreatedBy int64  `db:"created_by" json:"created_by,string"` // CreatedBy 创建人n
	Seq       int    `db:"seq" json:"seq"`                      // Seq 顺序
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicCategoryRepository struct{}

// QueryListPaged get paged records by condition
func (p *TopicCategoryRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*TopicCategory, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*TopicCategory, 0)

	box := packr.NewBox("./sql/topic_category")
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
func (p *TopicCategoryRepository) GetAll(node sqalx.Node) (items []*TopicCategory, err error) {
	items = make([]*TopicCategory, 0)
	sqlSelect := packr.NewBox("./sql/topic_category").String("GET_ALL.sql")

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
func (p *TopicCategoryRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*TopicCategory, err error) {
	items = make([]*TopicCategory, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =:topic_id"
		condition["topic_id"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =:parent_id"
		condition["parent_id"] = val
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =:seq"
		condition["seq"] = val
	}

	box := packr.NewBox("./sql/topic_category")
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

// GetByID get a record by ID
func (p *TopicCategoryRepository) GetByID(node sqalx.Node, id int64) (item *TopicCategory, exist bool, err error) {
	item = new(TopicCategory)
	sqlSelect := packr.NewBox("./sql/topic_category").String("GET_BY_ID.sql")

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

// GetByCondition get a record by condition
func (p *TopicCategoryRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *TopicCategory, exist bool, err error) {
	item = new(TopicCategory)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =:topic_id"
		condition["topic_id"] = val
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =:name"
		condition["name"] = val
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =:parent_id"
		condition["parent_id"] = val
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =:created_by"
		condition["created_by"] = val
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =:seq"
		condition["seq"] = val
	}

	box := packr.NewBox("./sql/topic_category")
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
func (p *TopicCategoryRepository) Insert(node sqalx.Node, item *TopicCategory) (err error) {
	sqlInsert := packr.NewBox("./sql/topic_category").String("INSERT.sql")

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
func (p *TopicCategoryRepository) Update(node sqalx.Node, item *TopicCategory) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic_category").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicCategoryRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/topic_category").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicCategoryRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
