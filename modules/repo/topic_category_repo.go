package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
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

// GetAllByCondition get records by condition
func (p *TopicCategoryRepository) GetAllByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (items []*TopicCategory, err error) {
	items = make([]*TopicCategory, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["parent_id"]; ok {
		clause += " AND a.parent_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_categories a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	err = node.SelectContext(ctx, &items, sqlSelect, condition...)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByID get a record by ID
func (p *TopicCategoryRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *TopicCategory, exist bool, err error) {
	item = new(TopicCategory)
	sqlSelect := "SELECT a.* FROM topic_categories a WHERE a.id=? AND a.deleted=0"

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

// Insert insert a new record
func (p *TopicCategoryRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicCategory) (err error) {
	sqlInsert := "INSERT INTO topic_categories( id,topic_id,name,parent_id,created_by,seq,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.TopicID, item.Name, item.ParentID, item.CreatedBy, item.Seq, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicCategoryRepository) Update(ctx context.Context, node sqalx.Node, item *TopicCategory) (err error) {
	sqlUpdate := "UPDATE topic_categories SET topic_id=?,name=?,parent_id=?,created_by=?,seq=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.TopicID, item.Name, item.ParentID, item.CreatedBy, item.Seq, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicCategoryRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_categories SET deleted=1 WHERE id=? "

	_, err = node.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
