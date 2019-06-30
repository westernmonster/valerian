package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

const (
	_addArticleVersionSQL       = "INSERT INTO article_versions( id,name,content,seq,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"
	_updateArticleVersionSQL    = "UPDATE article_versions SET name=?,content=?,seq=?,updated_at=? WHERE id=?"
	_getArticleVersionsSQL      = "SELECT a.id FROM  articles a  WHERE a.article_set_id=? AND a.deleted=0"
	_delArticleVersionSQL       = "UPDATE article_sets SET deleted=1 WHERE id=?"
	_getArticleVersionByNameSQL = "SELECT a.article_set_id,a.title AS article_title,a.id AS article_id,a.version_name,a.seq FROM articles a WHERE a.article_set_id=? AND a.version_name=? AND a.deleted=0 LIMIT 1"
)

func (p *Dao) AddArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error) {
	if _, err = node.ExecContext(c, _addArticleVersionSQL,
		item.ID, item.Name, item.Content, item.Seq, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleVersion error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error) {
	if _, err = node.ExecContext(c, _updateArticleVersionSQL, item.Name, item.Content, item.Seq, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleVersion error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) DelArticleVersion(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delArticleVersionSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleVersion error(%+v), id(%d)", err, id))
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

func (p *Dao) GetArticleVersions(c context.Context, node sqalx.Node, articleSetID int64) (items []int64, err error) {
	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, _getArticleVersionsSQL, articleSetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleVersions error(%+v), article set id(%d)", err, articleSetID))
		return
	}

	defer rows.Close()

	items = make([]int64, 0)

	for rows.Next() {
		var tid int64
		err = rows.Scan(&tid)
		if err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetArticleVersions scan error(%+v), article set id(%d)", err, articleSetID))
			return nil, err
		}

		items = append(items, tid)
	}

	return
}
