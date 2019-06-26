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
	_getArticleFileSQL    = "SELECT a.* FROM article_files a WHERE a.id=? AND a.deleted=0"
	_addArticleFileSQL    = "INSERT INTO article_files( id,file_name,file_url,seq,article_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"
	_updateArticleFileSQL = "UPDATE article_files SET file_name=?,file_url=?,seq=?,article_id=?,updated_at=? WHERE id=?"
	_delArticleFileSQL    = "UPDATE article_files SET deleted=1 WHERE id=? "
	_getArticleFilesSQL   = "SELECT a.* FROM article_files a WHERE a.deleted=0 AND a.article_id=? ORDER BY a.seq "
)

func (p *Dao) GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error) {
	item = new(model.ArticleFile)
	if err = node.GetContext(c, item, _getArticleFilesSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetArticleFileByID error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error) {
	items = make([]*model.ArticleFile, 0)
	if err = node.SelectContext(c, &items, _getArticleFileSQL, articleID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFiles error(%+v), article_id(%d)", err, articleID))
	}
	return
}

func (p *Dao) AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	if _, err = node.ExecContext(c, _addArticleFileSQL,
		item.ID,
		item.FileName,
		item.FileURL,
		item.Seq,
		item.ArticleID,
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleFile error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	if _, err = node.ExecContext(c, _updateArticleFileSQL,
		item.FileName,
		item.FileURL,
		item.Seq,
		item.ArticleID,
		item.UpdatedAt,
		item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleFile error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delArticleFileSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleFile error(%+v), id(%d)", err, id))
	}
	return
}
