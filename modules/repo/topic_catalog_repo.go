package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"
	types "valerian/library/database/sqlx/types"

	tracerr "github.com/ztrue/tracerr"
)

type TopicCatalog struct {
	ID        int64         `db:"id" json:"id,string"`                   // ID ID
	Name      string        `db:"name" json:"name"`                      // Name 名称
	Seq       int           `db:"seq" json:"seq"`                        // Seq 顺序
	Type      string        `db:"type" json:"type"`                      // Type 类型
	ParentID  int64         `db:"parent_id" json:"parent_id,string"`     // ParentID 父ID
	RefID     *int64        `db:"ref_id" json:"ref_id,omitempty,string"` // RefID 引用ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`       // TopicID 话题ID
	Deleted   types.BitBool `db:"deleted" json:"deleted"`                // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`          // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`          // UpdatedAt 更新时间
}

type TopicCatalogRepository struct{}

// GetAll get all records
func (p *TopicCatalogRepository) GetAll(ctx context.Context, node sqalx.Node) (items []*TopicCatalog, err error) {
	items = make([]*TopicCatalog, 0)
	sqlSelect := "SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 ORDER BY a.seq"

	err = node.SelectContext(ctx, &items, sqlSelect)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *TopicCatalogRepository) GetAllByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*TopicCatalog, err error) {
	items = make([]*TopicCatalog, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["type"]; ok {
		clause += " AND a.type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ref_id"]; ok {
		clause += " AND a.ref_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 %s ORDER BY a.seq ", clause)

	err = node.SelectContext(ctx, &items, sqlSelect, condition...)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get a record by ID
func (p *TopicCatalogRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *TopicCatalog, exist bool, err error) {
	item = new(TopicCatalog)
	sqlSelect := "SELECT a.* FROM topic_catalogs a WHERE a.id=? AND a.deleted=0"

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

// GetByCondition get a record by condition
func (p *TopicCatalogRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *TopicCatalog, exist bool, err error) {
	item = new(TopicCatalog)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["type"]; ok {
		clause += " AND a.type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["ref_id"]; ok {
		clause += " AND a.ref_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_catalogs a WHERE a.deleted=0 %s", clause)

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
func (p *TopicCatalogRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicCatalog) (err error) {
	sqlInsert := "INSERT INTO topic_catalogs( id,name,seq,type,parent_id,ref_id,topic_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicCatalogRepository) Update(ctx context.Context, node sqalx.Node, item *TopicCatalog) (err error) {
	sqlUpdate := "UPDATE topic_catalogs SET name=?,seq=?,type=?,parent_id=?,ref_id=?,topic_id=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.Name, item.Seq, item.Type, item.ParentID, item.RefID, item.TopicID, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicCatalogRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_catalogs SET deleted=1 WHERE id=? "

	_, err = node.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *TopicCatalogRepository) GetChildrenCount(ctx context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error) {
	sqlCount := "SELECT COUNT(1) as count FROM topic_catalogs a WHERE a.topic_id=? AND a.parent_id = ?"
	err = node.GetContext(ctx, &count, sqlCount, topicID, parentID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
