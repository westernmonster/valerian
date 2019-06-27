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
	_getAccountArticleAttrSQL    = "SELECT a.* FROM account_article_attrs a WHERE a.account_id=? AND a.article_id=? AND a.deleted=0"
	_addAccountArticleAttrSQL    = "INSERT INTO account_article_attrs( id,account_id,article_id,`read`,`like`,fav,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"
	_updateAccountArticleAttrSQL = "UPDATE account_article_attrs SET account_id=?,article_id=?,`read`=?,`like`=?,fav=?,updated_at=? WHERE id=?"
	_delAccountArticleAttrSQL    = "UPDATE account_article_attrs SET deleted=1 WHERE id=? "
	_getArticleLikeCountSQL      = "SELECT COUNT(1) as count FROM account_article_attrs a WHERE a.deleted=0 AND a.`like`=1 AND a.article_id=?"
	_getArticleFavCountSQL       = "SELECT COUNT(1) as count FROM account_article_attrs a WHERE a.deleted=0 AND a.fav=1 AND a.article_id=?"
)

func (p *Dao) AddAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error) {
	if _, err = node.ExecContext(c, _addAccountArticleAttrSQL, item.ID, item.AccountID, item.ArticleID, item.Read, item.Like, item.Fav, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccountArticleAttr error(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) UpdateAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error) {
	if _, err = node.ExecContext(c, _updateAccountArticleAttrSQL, item.AccountID, item.ArticleID, item.Read, item.Like, item.Fav, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAccountArticleAttr error(%+v), item(%+v)", err, item))
	}
	return
}

// Delete logic delete a exist record
func (p *Dao) DelAccountArticleAttr(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delAccountArticleAttrSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAccountArticleAttr error(%+v), id(%d)", err, id))
	}

	return
}

func (p *Dao) GetAccountArticleAttr(c context.Context, node sqalx.Node, aid, articleID int64) (item *model.AccountArticleAttr, err error) {
	item = new(model.AccountArticleAttr)

	if err = node.GetContext(c, item, _getAccountArticleAttrSQL, aid, articleID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAccountArticleAttr error(%+v) account_id(%d) article_id(%d)", err, aid, articleID))
	}

	return
}

func (p *Dao) GetArticleFavCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getArticleFavCountSQL, articleID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFavCount error(%+v), article id(%d)", err, articleID))
		return
	}
	return
}

func (p *Dao) GetArticleLikeCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getArticleLikeCountSQL, articleID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleLikeCount error(%+v), article id(%d)", err, articleID))
		return
	}
	return
}
