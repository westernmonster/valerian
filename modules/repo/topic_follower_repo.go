package repo

import (
	"database/sql"
	"fmt"
	packr "github.com/gobuffalo/packr"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
	"time"
)

type TopicFollower struct {
	ID          int64 `db:"id" json:"id,string"`                     // ID ID
	TopicID     int64 `db:"topic_id" json:"topic_id,string"`         // TopicID 话题ID
	FollowersID int64 `db:"followers_id" json:"followers_id,string"` // FollowersID 关注者ID
	Deleted     int   `db:"deleted" json:"deleted"`                  // Deleted 是否删除
	CreatedAt   int64 `db:"created_at" json:"created_at"`            // CreatedAt 创建时间
	UpdatedAt   int64 `db:"updated_at" json:"updated_at"`            // UpdatedAt 更新时间
}

type TopicFollowerRepository struct{}

// QueryListPaged get paged records by condition
func (p *TopicFollowerRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*TopicFollower, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*TopicFollower, 0)

	box := packr.NewBox("./sql/topic_follower")
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
func (p *TopicFollowerRepository) GetAll(node sqalx.Node) (items []*TopicFollower, err error) {
	items = make([]*TopicFollower, 0)
	sqlSelect := packr.NewBox("./sql/topic_follower").String("GET_ALL.sql")

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
func (p *TopicFollowerRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*TopicFollower, err error) {
	items = make([]*TopicFollower, 0)
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
	if val, ok := cond["followers_id"]; ok {
		clause += " AND a.followers_id =:followers_id"
		condition["followers_id"] = val
	}

	box := packr.NewBox("./sql/topic_follower")
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
func (p *TopicFollowerRepository) GetByID(node sqalx.Node, id int64) (item *TopicFollower, exist bool, err error) {
	item = new(TopicFollower)
	sqlSelect := packr.NewBox("./sql/topic_follower").String("GET_BY_ID.sql")

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
func (p *TopicFollowerRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *TopicFollower, exist bool, err error) {
	item = new(TopicFollower)
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
	if val, ok := cond["followers_id"]; ok {
		clause += " AND a.followers_id =:followers_id"
		condition["followers_id"] = val
	}

	box := packr.NewBox("./sql/topic_follower")
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
func (p *TopicFollowerRepository) Insert(node sqalx.Node, item *TopicFollower) (err error) {
	sqlInsert := packr.NewBox("./sql/topic_follower").String("INSERT.sql")

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
func (p *TopicFollowerRepository) Update(node sqalx.Node, item *TopicFollower) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic_follower").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicFollowerRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/topic_follower").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicFollowerRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
