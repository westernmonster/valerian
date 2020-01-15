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

func (p *Dao) GetArticleStatByID(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleStat, err error) {
	item = new(model.ArticleStat)
	sqlSelect := "SELECT a.article_id,a.like_count,a.dislike_count,a.revise_count,a.comment_count,a.created_at,a.updated_at FROM article_stats a WHERE a.article_id=?"

	if err = node.GetContext(c, item, sqlSelect, articleID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleStatByID err(%+v), article_id(%+v)", err, articleID))
	}

	return
}

func (p *Dao) IncrArticleStat(c context.Context, node sqalx.Node, item *model.ArticleStat) (err error) {
	sqlUpdate := "UPDATE article_stats SET like_count=like_count+?,dislike_count=dislike_count+?, revise_count=revise_count+?,comment_count=comment_count+?,updated_at=? WHERE article_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.ReviseCount, item.CommentCount, time.Now().Unix(), item.ArticleID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrArticleStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
