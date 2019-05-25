package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/models"

	"valerian/library/database/sqalx"
	types "valerian/library/database/sqlx/types"

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

func (p *TopicRelationRepository) GetAllTopicRelations(ctx context.Context, node sqalx.Node, topicID int64) (items []*TopicRelation, err error) {
	items = make([]*TopicRelation, 0)
	sqlSelect := "SELECT a.* FROM topic_relations a WHERE a.from_topic_id=?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicRelationRepository) GetAllRelatedTopics(ctx context.Context, node sqalx.Node, topicID int64) (items []*models.RelatedTopicShort, err error) {
	items = make([]*models.RelatedTopicShort, 0)
	sqlSelect := "SELECT a.to_topic_id AS topic_id,b.name,b.version_name,a.relation AS type FROM topic_relations a LEFT JOIN topics b ON a.to_topic_id=b.id WHERE a.from_topic_id=?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *TopicRelationRepository) GetAllRelatedTopicsDetail(ctx context.Context, node sqalx.Node, topicID int64) (items []*models.RelatedTopic, err error) {
	items = make([]*models.RelatedTopic, 0)
	sqlSelect := "SELECT a.to_topic_id AS topic_id,b.name,b.version_name,a.relation AS type,b.cover, b.introduction FROM topic_relations a LEFT JOIN topics b ON a.to_topic_id=b.id WHERE a.from_topic_id=?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByCondition get a record by condition
func (p *TopicRelationRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *TopicRelation, exist bool, err error) {
	item = new(TopicRelation)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["from_topic_id"]; ok {
		clause += " AND a.from_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["relation"]; ok {
		clause += " AND a.relation =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_relations a WHERE a.deleted=0 %s", clause)

	if e := node.GetContext(ctx, item, sqlSelect, condition...); e != nil {
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
func (p *TopicRelationRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicRelation) (err error) {
	sqlInsert := "INSERT INTO topic_relations( id,from_topic_id,to_topic_id,relation,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.FromTopicID, item.ToTopicID, item.Relation, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicRelationRepository) Update(ctx context.Context, node sqalx.Node, item *TopicRelation) (err error) {
	sqlUpdate := "UPDATE topic_relations SET from_topic_id=?,to_topic_id=?,relation=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.FromTopicID, item.ToTopicID, item.Relation, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
