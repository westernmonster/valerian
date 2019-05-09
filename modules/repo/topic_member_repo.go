package repo

import (
	"database/sql"
	"fmt"
	"time"
	"valerian/models"

	"valerian/library/database/sqalx"

	packr "github.com/gobuffalo/packr"
	tracerr "github.com/ztrue/tracerr"
)

type TopicMember struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64  `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64  `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string `db:"role" json:"role"`                    // Role 成员角色
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicMemberRepository struct{}

// QueryListPaged get paged records by condition
func (p *TopicMemberRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*TopicMember, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*TopicMember, 0)

	box := packr.NewBox("./sql/topic_member")
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
func (p *TopicMemberRepository) GetAll(node sqalx.Node) (items []*TopicMember, err error) {
	items = make([]*TopicMember, 0)
	sqlSelect := packr.NewBox("./sql/topic_member").String("GET_ALL.sql")

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

// GetAllTopicMembers
func (p *TopicMemberRepository) GetAllTopicMembers(node sqalx.Node, topicID int64) (items []*models.TopicMember, err error) {
	items = make([]*models.TopicMember, 0)
	sqlSelect := "SELECT a.account_id, a.role, b.user_name, b.avatar FROM topic_members a LEFT JOIN accounts b ON a.account_id = b.id WHERE a.topic_id=? ORDER BY a.id DESC"

	err = node.Select(&items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *TopicMemberRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*TopicMember, err error) {
	items = make([]*TopicMember, 0)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =:role"
		condition["role"] = val
	}

	box := packr.NewBox("./sql/topic_member")
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
func (p *TopicMemberRepository) GetByID(node sqalx.Node, id int64) (item *TopicMember, exist bool, err error) {
	item = new(TopicMember)
	sqlSelect := packr.NewBox("./sql/topic_member").String("GET_BY_ID.sql")

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
func (p *TopicMemberRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *TopicMember, exist bool, err error) {
	item = new(TopicMember)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =:account_id"
		condition["account_id"] = val
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =:role"
		condition["role"] = val
	}

	box := packr.NewBox("./sql/topic_member")
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
func (p *TopicMemberRepository) Insert(node sqalx.Node, item *TopicMember) (err error) {
	sqlInsert := packr.NewBox("./sql/topic_member").String("INSERT.sql")

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
func (p *TopicMemberRepository) Update(node sqalx.Node, item *TopicMember) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic_member").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicMemberRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/topic_member").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicMemberRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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