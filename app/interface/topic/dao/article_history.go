package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getArticleHistorySQL    = "SELECT a.* FROM article_histories a WHERE a.id=? AND a.deleted=0"
	_addArticleHistorySQL    = "INSERT INTO article_histories( id,article_id,updated_by,content,content_text,seq,diff,description,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"
	_updateArticleHistorySQL = "UPDATE article_histories SET article_id=?,updated_by=?,content=?,content_text=?,seq=?,diff=?,description=?,updated_at=? WHERE id=?"

	_getArticleHistoryMaxSeqSQL      = "SELECT a.seq FROM article_histories a WHERE a.deleted=0 AND a.article_id=? ORDER BY a.seq DESC LIMIT 1"
	_getOtherMemberArticleHistorySQL = "SELECT COUNT(1) as count FROM article_histories a WHERE a.deleted=0 AND a.article_id=? AND a.updated_by!=?"
)

func (p *Dao) GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error) {
	item = new(model.ArticleHistory)
	if err = node.GetContext(c, item, _getArticleHistorySQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetArticleByID error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	if _, err = node.ExecContext(c, _addArticleHistorySQL, item.ID, item.ArticleID, item.UpdatedBy, item.Content, item.ContentText, item.Seq, item.Diff, item.Description, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleHistory error(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) UpdateArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	if _, err = node.ExecContext(c, _updateArticleHistorySQL, item.ArticleID, item.UpdatedBy, item.Content, item.ContentText, item.Seq, item.Diff, item.Description, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleHistory error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) GetArticleHistories(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleHistory, err error) {
	return
}

func (p *Dao) GetArticleHistoryMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error) {
	if err = node.GetContext(c, &seq, _getArticleHistoryMaxSeqSQL, articleID); err != nil {
		if err == sql.ErrNoRows {
			seq = 0
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoryMaxSeq error(%+v), article id(%d)", err, articleID))
		return
	}
	return
}

func (p *Dao) GetOrderMemberArticleHistoriesCount(c context.Context, node sqalx.Node, articleID int64, aid int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getOtherMemberArticleHistorySQL, articleID, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetOrderMemberHistoriesCount error(%+v), article id(%d)", err, articleID))
		return
	}
	return
}
