package repo

import (
	"database/sql"
	"fmt"
	"time"

	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
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

// QueryListPaged get paged records by condition
func (p *TopicTypeRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*TopicType, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*TopicType, 0)

	box := packr.NewBox("./sql/topic_type")
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
func (p *TopicTypeRepository) GetAll(node sqalx.Node) (items []*TopicType, err error) {
	items = make([]*TopicType, 0)
	sqlSelect := packr.NewBox("./sql/topic_type").String("GET_ALL.sql")

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
func (p *TopicTypeRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*TopicType, err error) {
	items = make([]*TopicType, 0)
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

	box := packr.NewBox("./sql/topic_type")
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
func (p *TopicTypeRepository) GetByID(node sqalx.Node, id int) (item *TopicType, exist bool, err error) {
	item = new(TopicType)
	sqlSelect := packr.NewBox("./sql/topic_type").String("GET_BY_ID.sql")

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
func (p *TopicTypeRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *TopicType, exist bool, err error) {
	item = new(TopicType)
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

	box := packr.NewBox("./sql/topic_type")
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
func (p *TopicTypeRepository) Insert(node sqalx.Node, item *TopicType) (err error) {
	sqlInsert := packr.NewBox("./sql/topic_type").String("INSERT.sql")

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
func (p *TopicTypeRepository) Update(node sqalx.Node, item *TopicType) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic_type").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicTypeRepository) Delete(node sqalx.Node, id int) (err error) {
	sqlDelete := packr.NewBox("./sql/topic_type").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicTypeRepository) BatchDelete(node sqalx.Node, ids []int) (err error) {
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
