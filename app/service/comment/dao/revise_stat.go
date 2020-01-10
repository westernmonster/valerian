package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

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

func (p *Dao) IncrReviseStat(c context.Context, node sqalx.Node, item *model.ReviseStat) (err error) {
	sqlUpdate := "UPDATE revise_stats SET like_count=like_count+?,dislike_count=dislike_count+?,comment_count=comment_count+?,updated_at=? WHERE revise_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, time.Now().Unix(), item.ReviseID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrReviseStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
