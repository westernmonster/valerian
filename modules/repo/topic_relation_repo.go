package repo

import (
	"database/sql"
	"fmt"
	"time"
	"valerian/models"

	packr "github.com/gobuffalo/packr"
	types "github.com/jmoiron/sqlx/types"
	sqalx "github.com/westernmonster/sqalx"
	tracerr "github.com/ztrue/tracerr"
)

type TopicRelation struct {
	ID          int64         `db:"id" json:"id,string"`                       // ID ID
	FromTopicID int64         `db:"from_topic_id" json:"from_topic_id,string"` // FromTopicID 话题ID
	ToTopicID   int64         `db:"to_topic_id" json:"to_topic_id,string"`     // ToTopicID 关联话题ID
	Relation    string        `db:"relation" json:"relation"`                  // Relation 关系
	Deleted     types.BitBool `db:"deleted" json:"deleted"`                    // Deleted 是否删除
	CreatedAt   int64         `db:"created_at" json:"created_at"`              // CreatedAt 创建时间
	UpdatedAt   int64         `db:"updated_at" json:"updated_at"`              // UpdatedAt 更新时间
}

type TopicRelationRepository struct{}

// QueryListPaged get paged records by condition
func (p *TopicRelationRepository) QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*TopicRelation, err error) {
	offset := (page - 1) * pageSize
	condition := make(map[string]interface{})
	clause := ""

	items = make([]*TopicRelation, 0)

	box := packr.NewBox("./sql/topic_relation")
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
func (p *TopicRelationRepository) GetAll(node sqalx.Node) (items []*TopicRelation, err error) {
	items = make([]*TopicRelation, 0)
	sqlSelect := packr.NewBox("./sql/topic_relation").String("GET_ALL.sql")

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

func (p *TopicRelationRepository) GetAllRelatedTopics(node sqalx.Node, topicID int64) (items []*models.RelatedTopic, err error) {
	items = make([]*models.RelatedTopic, 0)
	sqlSelect := "SELECT a.to_topic_id AS topic_id,b.name,b.version_name,b.version_lang,a.relation AS type FROM topic_relations a LEFT JOIN topics b ON a.to_topic_id=b.id WHERE a.from_topic_id=?"

	err = node.Select(&items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *TopicRelationRepository) GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*TopicRelation, err error) {
	items = make([]*TopicRelation, 0)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["from_topic_id"]; ok {
		clause += " AND a.from_topic_id =:from_topic_id"
		condition["from_topic_id"] = val
	}
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =:to_topic_id"
		condition["to_topic_id"] = val
	}
	if val, ok := cond["relation"]; ok {
		clause += " AND a.relation =:relation"
		condition["relation"] = val
	}

	box := packr.NewBox("./sql/topic_relation")
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
func (p *TopicRelationRepository) GetByID(node sqalx.Node, id int64) (item *TopicRelation, exist bool, err error) {
	item = new(TopicRelation)
	sqlSelect := packr.NewBox("./sql/topic_relation").String("GET_BY_ID.sql")

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
func (p *TopicRelationRepository) GetByCondition(node sqalx.Node, cond map[string]string) (item *TopicRelation, exist bool, err error) {
	item = new(TopicRelation)
	condition := make(map[string]interface{})
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =:id"
		condition["id"] = val
	}
	if val, ok := cond["from_topic_id"]; ok {
		clause += " AND a.from_topic_id =:from_topic_id"
		condition["from_topic_id"] = val
	}
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =:to_topic_id"
		condition["to_topic_id"] = val
	}
	if val, ok := cond["relation"]; ok {
		clause += " AND a.relation =:relation"
		condition["relation"] = val
	}

	box := packr.NewBox("./sql/topic_relation")
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
func (p *TopicRelationRepository) Insert(node sqalx.Node, item *TopicRelation) (err error) {
	sqlInsert := packr.NewBox("./sql/topic_relation").String("INSERT.sql")

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
func (p *TopicRelationRepository) Update(node sqalx.Node, item *TopicRelation) (err error) {
	sqlUpdate := packr.NewBox("./sql/topic_relation").String("UPDATE.sql")

	item.UpdatedAt = time.Now().Unix()

	_, err = node.NamedExec(sqlUpdate, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicRelationRepository) Delete(node sqalx.Node, id int64) (err error) {
	sqlDelete := packr.NewBox("./sql/topic_relation").String("DELETE.sql")

	_, err = node.NamedExec(sqlDelete, map[string]interface{}{"id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// BatchDelete logic batch delete records
func (p *TopicRelationRepository) BatchDelete(node sqalx.Node, ids []int64) (err error) {
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
