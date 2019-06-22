package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addArticleSetSQL = "INSERT INTO article_sets( id,deleted,created_at,updated_at) VALUES ( ?,?,?,?)"
	_delArticleSetSQL = "UPDATE article_sets SET deleted=1 WHERE id=?"
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
