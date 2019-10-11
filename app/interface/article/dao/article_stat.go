package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetArticleStatByID(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleStat, err error) {
	item = new(model.ArticleStat)
	sqlSelect := "SELECT a.* FROM article_stats a WHERE a.article_id=?"

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

// Insert insert a new record
func (p *Dao) AddArticleStat(c context.Context, node sqalx.Node, item *model.ArticleStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO article_stats( article_id,like_count,dislike_count,revise_count,comment_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ArticleID, item.LikeCount, item.DislikeCount, item.ReviseCount, item.CommentCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateArticleStat(c context.Context, node sqalx.Node, item *model.ArticleStat) (err error) {
	sqlUpdate := "UPDATE article_stats SET like_count=?,dislike_count=?,revise_count=?,comment_count=?,updated_at=? WHERE article_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.LikeCount, item.DislikeCount, item.ReviseCount, item.CommentCount, item.UpdatedAt, item.ArticleID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleStats err(%+v), item(%+v)", err, item))
		return
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
