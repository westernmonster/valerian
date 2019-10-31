package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/search/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetLastArticleHistory(c context.Context, node sqalx.Node, articleID int64) (item *model.ArticleHistory, err error) {
	item = new(model.ArticleHistory)
	sqlSelect := "SELECT a.* FROM article_histories a WHERE a.article_id=? AND a.deleted=0 ORDER BY a.id DESC LIMIT 1"

	if err = node.GetContext(c, item, sqlSelect, articleID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetLastArticleHistory err(%+v), article_id(%+v)", err, articleID))
	}

	return
}

// GetAll get all records
func (p *Dao) GetArticles(c context.Context, node sqalx.Node) (items []*model.Article, err error) {
	items = make([]*model.Article, 0)
	sqlSelect := "SELECT a.* FROM articles a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticles err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetArticlesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Article, err error) {
	items = make([]*model.Article, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["disable_revise"]; ok {
		clause += " AND a.disable_revise =?"
		condition = append(condition, val)
	}
	if val, ok := cond["disable_comment"]; ok {
		clause += " AND a.disable_comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM articles a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticlesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error) {
	item = new(model.Article)
	sqlSelect := "SELECT a.* FROM articles a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetArticleByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Article, err error) {
	item = new(model.Article)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["title"]; ok {
		clause += " AND a.title =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["disable_revise"]; ok {
		clause += " AND a.disable_revise =?"
		condition = append(condition, val)
	}
	if val, ok := cond["disable_comment"]; ok {
		clause += " AND a.disable_comment =?"
		condition = append(condition, val)
	}
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM articles a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticlesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	sqlInsert := "INSERT INTO articles( id,title,content,content_text,disable_revise,disable_comment,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Title, item.Content, item.ContentText, item.DisableRevise, item.DisableComment, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	sqlUpdate := "UPDATE articles SET title=?,content=?,content_text=?,disable_revise=?,disable_comment=?,created_by=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Title, item.Content, item.ContentText, item.DisableRevise, item.DisableComment, item.CreatedBy, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelArticle(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE articles SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticles err(%+v), item(%+v)", err, id))
		return
	}

	return
}
