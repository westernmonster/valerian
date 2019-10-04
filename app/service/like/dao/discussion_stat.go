package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/app/service/like/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error) {
	item = new(model.DiscussionStat)
	sqlSelect := "SELECT a.* FROM discussion_stats a WHERE a.discussion_id=?"

	if err = node.GetContext(c, item, sqlSelect, discussionID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionStatByID err(%+v), discussion_id(%+v)", err, discussionID))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO discussion_stats( discussion_id,like_count,dislike_count,comment_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.DiscussionID, item.LikeCount, item.DislikeCount, item.CommentCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDiscussionStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error) {
	sqlUpdate := "UPDATE discussion_stats SET like_count=?,dislike_count=?,comment_count=?,updated_at=?  WHERE discussion_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, item.UpdatedAt, item.DiscussionID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDiscussionStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error) {
	sqlUpdate := "UPDATE discussion_stats SET like_count=like_count+?,dislike_count=dislike_count+?,comment_count=discussion_count+?,updated_at=? WHERE discussion_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, time.Now().Unix(), item.DiscussionID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrDiscussionStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
