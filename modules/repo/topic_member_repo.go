package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/models"

	"valerian/library/database/sqalx"

	types "github.com/jmoiron/sqlx/types"
	tracerr "github.com/ztrue/tracerr"
)

type TopicMember struct {
	ID        int64         `db:"id" json:"id,string"`                 // ID ID
	TopicID   int64         `db:"topic_id" json:"topic_id,string"`     // TopicID 分类ID
	AccountID int64         `db:"account_id" json:"account_id,string"` // AccountID 成员ID
	Role      string        `db:"role" json:"role"`                    // Role 成员角色
	Deleted   types.BitBool `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64         `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64         `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type TopicMemberRepository struct{}

// GetAllTopicMembers
func (p *TopicMemberRepository) GetTopicMembers(ctx context.Context, node sqalx.Node, topicID int64, limit int) (items []*models.TopicMember, err error) {
	items = make([]*models.TopicMember, 0)
	sqlSelect := "SELECT a.account_id, a.role, b.user_name, b.avatar FROM topic_members a LEFT JOIN accounts b ON a.account_id = b.id WHERE a.topic_id=? ORDER BY a.id DESC limit ?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID, limit)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetAllTopicMembers
func (p *TopicMemberRepository) GetAllTopicMembers(ctx context.Context, node sqalx.Node, topicID int64) (items []*TopicMember, err error) {
	items = make([]*TopicMember, 0)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.topic_id=? ORDER BY a.id DESC"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetTopicMembersCount
func (p *TopicMemberRepository) GetTopicMembersCount(ctx context.Context, node sqalx.Node, topicID int64) (count int, err error) {
	sqlCount := "SELECT COUNT(1) as count FROM topic_members a LEFT JOIN accounts b ON a.account_id = b.id WHERE a.topic_id=?"
	err = node.GetContext(ctx, &count, sqlCount, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetTopicMembersPaged
func (p *TopicMemberRepository) GetTopicMembersPaged(ctx context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*models.TopicMember, err error) {
	items = make([]*models.TopicMember, 0)
	offset := (page - 1) * pageSize

	sqlCount := "SELECT COUNT(1) as count FROM topic_members a LEFT JOIN accounts b ON a.account_id = b.id WHERE a.topic_id=?"
	err = node.GetContext(ctx, &count, sqlCount, topicID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	sqlSelect := "SELECT a.account_id, a.role, b.user_name, b.avatar FROM topic_members a LEFT JOIN accounts b ON a.account_id = b.id WHERE a.topic_id=? ORDER BY a.id DESC limit ?,?"

	err = node.SelectContext(ctx, &items, sqlSelect, topicID, offset, pageSize)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// GetByCondition get a record by condition
func (p *TopicMemberRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *TopicMember, exist bool, err error) {
	item = new(TopicMember)
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
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_members a WHERE a.deleted=0 %s", clause)

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
func (p *TopicMemberRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicMember) (err error) {
	sqlInsert := "INSERT INTO topic_members( id,topic_id,account_id,role,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.TopicID, item.AccountID, item.Role, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Update update a exist record
func (p *TopicMemberRepository) Update(ctx context.Context, node sqalx.Node, item *TopicMember) (err error) {
	sqlUpdate := "UPDATE topic_members SET topic_id=?,account_id=?,role=?,updated_at=? WHERE id=?"

	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlUpdate, item.TopicID, item.AccountID, item.Role, item.UpdatedAt, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicMemberRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_members SET deleted=1 WHERE id=? "

	_, err = node.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
