package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/app/interface/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetReviseStatByID(c context.Context, node sqalx.Node, reviseID int64) (item *model.ReviseStat, err error) {
	item = new(model.ReviseStat)
	sqlSelect := "SELECT a.revise_id,a.like_count,a.dislike_count,a.comment_count,a.created_at,a.updated_at FROM revise_stats a WHERE a.revise_id=?"

	if err = node.GetContext(c, item, sqlSelect, reviseID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetReviseStatByID err(%+v), revise_id(%+v)", err, reviseID))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddReviseStat(c context.Context, node sqalx.Node, item *model.ReviseStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO revise_stats( revise_id,like_count,dislike_count,comment_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ReviseID, item.LikeCount, item.DislikeCount, item.CommentCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddReviseStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateReviseStat(c context.Context, node sqalx.Node, item *model.ReviseStat) (err error) {
	sqlUpdate := "UPDATE revise_stats SET revise_id=?,like_count=?,dislike_count=?,comment_count=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, item.UpdatedAt, item.ReviseID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateReviseStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrReviseStat(c context.Context, node sqalx.Node, item *model.ReviseStat) (err error) {
	sqlUpdate := "UPDATE revise_stats SET like_count=like_count+?,dislike_count=dislike_count+?,comment_count=comment_count+?,updated_at=? WHERE revise_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, time.Now().Unix(), item.ReviseID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrReviseStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
