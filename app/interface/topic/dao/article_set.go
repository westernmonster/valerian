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
	_addArticleSetSQL           = "INSERT INTO article_sets( id,deleted,created_at,updated_at) VALUES ( ?,?,?,?)"
	_delArticleSetSQL           = "UPDATE article_sets SET deleted=1 WHERE id=?"
	_getArticleVersionByNameSQL = "SELECT a.article_set_id,a.title AS article_title,a.id AS article_id,a.version_name,a.seq FROM articles a WHERE a.article_set_id=? AND a.version_name=? AND a.deleted=0 LIMIT 1"
)

func (p *Dao) AddArticleSet(c context.Context, node sqalx.Node, item *model.ArticleSet) (err error) {
	if _, err = node.ExecContext(c, _addArticleSetSQL,
		item.ID,
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleSet error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) DelArticleSet(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delArticleSetSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleSet error(%+v), id(%d)", err, id))
	}
	return
}

func (p *Dao) GetArticleVersionByName(c context.Context, node sqalx.Node, articleSetID int64, versionName string) (item *model.ArticleVersionResp, err error) {
	item = new(model.ArticleVersionResp)

	if err = node.GetContext(c, item, _getArticleVersionByNameSQL, articleSetID, versionName); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleVersionByName error(%+v), article_set_id(%d) version_name(%s)", err, articleSetID, versionName))
	}

	return
}
