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
	_getArticleSQL    = "SELECT a.* FROM articles a WHERE a.id=? AND a.deleted=0"
	_addArticleSQL    = "INSERT INTO articles( id,title,cover,locale,introduction,private,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"
	_updateArticleSQL = "UPDATE articles SET title=?,cover=?,locale=?,introduction=?,private=?,created_by=?,updated_at=? WHERE id=?"
	_delArticleSQL    = "UPDATE articles SET deleted=1 WHERE id=?"
)

func (p *Dao) GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error) {
	item = new(model.Article)

	if err = node.GetContext(c, item, _getArticleSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetArticleByID error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	if _, err = node.ExecContext(c, _addArticleSQL, item.ID, item.Title, item.Cover, item.Locale, item.Introduction, item.Private, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticle error(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	if _, err = node.ExecContext(c, _updateArticleSQL,
		item.Title, item.Cover, item.Locale, item.Introduction, item.Private, item.CreatedBy, item.UpdatedAt, item.ID,
	); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticle error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) DelArticle(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delArticleSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticle error(%+v), id(%d)", err, id))
	}
	return
}
