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
	_addArticleVersionSQL        = "INSERT INTO article_versions( id,article_id,locale,name,content,seq,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"
	_updateArticleVersionSQL     = "UPDATE article_versions SET article_id=?,locale=?,name=?,content=?,seq=?,updated_at=? WHERE id=?"
	_getArticleVersionSQL        = "SELECT a.* FROM  article_versions a  WHERE a.id=? AND a.deleted=0"
	_getArticleVersionsSQL       = "SELECT a.id, a.article_id, a.seq, a.name, a.content from article_versions a WHERE a.deleted=0 AND a.article_id=?"
	_delArticleVersionSQL        = "UPDATE article_versions SET deleted=1 WHERE id=?"
	_getArticleVersionByNameSQL  = "SELECT a.id, a.article_id, a.seq, a.name, a.content from article_versions a WHERE a.deleted=0 AND a.article_id=? AND a.name=?"
	_getArticleVersionsMaxSeqSQL = "SELECT a.seq FROM article_versions a WHERE a.deleted=0 AND a.article_id=? ORDER BY a.seq DESC LIMIT 1"
)

func (p *Dao) AddArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error) {
	if _, err = node.ExecContext(c, _addArticleVersionSQL,
		item.ID, item.ArticleID, item.Locale, item.Name, item.Content, item.Seq, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleVersion error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error) {
	if _, err = node.ExecContext(c, _updateArticleVersionSQL, item.ArticleID, item.Locale, item.Name, item.Content, item.Seq, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleVersion error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) GetArticleVersion(c context.Context, node sqalx.Node, id int64) (item *model.ArticleVersion, err error) {
	item = new(model.ArticleVersion)

	if err = node.GetContext(c, item, _getArticleVersionSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleVersion error(%+v) id(%d)", err, id))
	}

	return
}

func (p *Dao) DelArticleVersion(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delArticleVersionSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleVersion error(%+v), id(%d)", err, id))
	}
	return
}

func (p *Dao) GetArticleVersionByName(c context.Context, node sqalx.Node, articleID int64, versionName string) (item *model.ArticleVersionResp, err error) {
	item = new(model.ArticleVersionResp)

	if err = node.GetContext(c, item, _getArticleVersionByNameSQL, articleID, versionName); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleVersionByName error(%+v), article_id(%d) version_name(%s)", err, articleID, versionName))
	}

	return
}

func (p *Dao) GetArticleVersions(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleVersionResp, err error) {
	items = make([]*model.ArticleVersionResp, 0)
	if err = node.SelectContext(c, &items, _getArticleVersionsSQL, articleID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleVersions error(%+v) article_id(%d)", err, articleID))

	}
	return

}

func (p *Dao) GetArticleVersionsMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error) {
	if err = node.GetContext(c, &seq, _getArticleVersionsMaxSeqSQL, articleID); err != nil {
		if err == sql.ErrNoRows {
			seq = 0
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesMaxSeq error(%+v), article_id(%d)", err, articleID))
	}
	return
}
