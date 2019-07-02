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
	_getArticleHistorySQL         = "SELECT a.* FROM article_histories a WHERE a.id=? AND a.deleted=0"
	_getArticleHistoriesSQL       = "SELECT a.* FROM article_histories a WHERE a.article_version_id=? AND a.deleted=0"
	_addArticleHistorySQL         = "INSERT INTO article_histories( id,article_version_id,updated_by,content,content_text,seq,diff,change_id,description,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"
	_updateArticleHistorySQL      = "UPDATE article_histories SET article_version_id=?,updated_by=?,content=?,content_text=?,seq=?,diff=?,change_id=?,description=?,updated_at=? WHERE id=?"
	_getArticleHistoriesMaxSeqSQL = "SELECT a.seq FROM article_histories a WHERE a.deleted=0 AND a.article_version_id=? ORDER BY a.seq DESC LIMIT 1"
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
	if _, err = node.ExecContext(c, _addArticleHistorySQL, item.ID, item.ArticleVersionID, item.UpdatedBy, item.Content, item.ContentText, item.Seq, item.Diff, item.ChangeID, item.Description, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleHistory error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	if _, err = node.ExecContext(c, _updateArticleHistorySQL, item.ArticleVersionID, item.UpdatedBy, item.Content, item.ContentText, item.Seq, item.Diff, item.ChangeID, item.Description, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleHistory error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) GetArticleHistories(c context.Context, node sqalx.Node, articleVersionID int64) (items []*model.ArticleHistory, err error) {
	items = make([]*model.ArticleHistory, 0)
	if err = node.SelectContext(c, &items, _getArticleHistoriesSQL, articleVersionID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistories error(%+v) article_version_id(%d)", err, articleVersionID))

	}
	return
}

func (p *Dao) GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleVersionID int64) (seq int, err error) {
	if err = node.GetContext(c, &seq, _getArticleHistoriesMaxSeqSQL, articleVersionID); err != nil {
		if err == sql.ErrNoRows {
			seq = 0
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoryMaxSeq error(%+v), article_version_id(%d)", err, articleVersionID))
	}
	return
}
