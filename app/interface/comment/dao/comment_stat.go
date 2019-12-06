package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetCommentStatByID(c context.Context, node sqalx.Node, id int64) (item *model.CommentStat, err error) {
	item = new(model.CommentStat)
	sqlSelect := "SELECT a.comment_id,a.like_count,a.dislike_count,a.children_count,a.created_at,a.updated_at FROM comment_stats a WHERE a.comment_id=?"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCommentStatByID err(%+v), id(%+v)", err, id))
	}

	return
}

func (p *Dao) AddCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO comment_stats( comment_id,like_count,dislike_count,children_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.CommentID, item.LikeCount, item.DislikeCount, item.ChildrenCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddCommentStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) UpdateCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error) {
	sqlUpdate := "UPDATE comment_stats SET like_count=?,dislike_count=?,children_count=?,updated_at=? WHERE comment_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.ChildrenCount, item.UpdatedAt, item.CommentID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateCommentStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrCommentStat(c context.Context, node sqalx.Node, item *model.CommentStat) (err error) {
	sqlUpdate := "UPDATE comment_stats SET like_count=like_count+?,dislike_count=dislike_count+?,children_count=children_count+?,updated_at=? WHERE comment_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.ChildrenCount, item.UpdatedAt, item.CommentID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrCommentStat err(%+v), item(%+v)", err, item))
		return
	}

	return
}
