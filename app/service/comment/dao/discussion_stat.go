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

func (p *Dao) GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error) {
	item = new(model.DiscussionStat)
	sqlSelect := "SELECT a.discussion_id,a.like_count,a.dislike_count,a.comment_count,a.created_at,a.updated_at FROM discussion_stats a WHERE a.discussion_id=?"

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

func (p *Dao) IncrDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error) {
	sqlUpdate := "UPDATE discussion_stats SET like_count=like_count+?,dislike_count=dislike_count+?,comment_count=comment_count+?,updated_at=? WHERE discussion_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.CommentCount, time.Now().Unix(), item.DiscussionID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrDiscussionStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
