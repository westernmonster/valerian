package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
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

// GetByID get record by ID
func (p *TopicFollowerRepository) GetByID(ctx context.Context, node sqalx.Node, id int64) (item *TopicFollower, exist bool, err error) {
	item = new(TopicFollower)
	sqlSelect := "SELECT a.* FROM topic_followers a WHERE a.id=? AND a.deleted=0"

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

// GetByCondition get record by condition
func (p *TopicFollowerRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *TopicFollower, exist bool, err error) {
	item = new(TopicFollower)
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
	if val, ok := cond["followers_id"]; ok {
		clause += " AND a.followers_id =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_followers a WHERE a.deleted=0 %s", clause)

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
func (p *TopicFollowerRepository) Insert(ctx context.Context, node sqalx.Node, item *TopicFollower) (err error) {
	sqlInsert := "INSERT INTO topic_followers( id,topic_id,followers_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert, item.ID, item.TopicID, item.FollowersID, item.Deleted, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete logic delete a exist record
func (p *TopicFollowerRepository) Delete(ctx context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_followers SET deleted=1 WHERE id=? "

	_, err = node.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
