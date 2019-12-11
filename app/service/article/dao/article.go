package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetUserArticlesPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Article, err error) {
	items = make([]*model.Article, 0)
	sqlSelect := "SELECT a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at FROM articles a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserArticlesPaged err(%+v) aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
		return
	}
	return
}

func (p *Dao) GetUserArticleIDsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.id FROM articles a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserArticleIDsPaged err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserArticleIDsPaged err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

// GetAll get all records
func (p *Dao) GetArticles(c context.Context, node sqalx.Node) (items []*model.Article, err error) {
	items = make([]*model.Article, 0)
	sqlSelect := "SELECT  a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at  FROM articles a WHERE a.deleted=0 ORDER BY a.id DESC "

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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at  FROM articles a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticlesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error) {
	item = new(model.Article)
	sqlSelect := "SELECT a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at  FROM articles a WHERE a.id=? AND a.deleted=0"

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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at  FROM articles a WHERE a.deleted=0 %s", clause)

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

// GetArticlesByIDs get all records by ids
func (p *Dao) GetArticlesByIDs(c context.Context, node sqalx.Node, ids []int64) (items []*model.Article, err error) {
	items = make([]*model.Article, 0)
	sqlSelect := "SELECT  a.id,a.title,a.content,a.content_text,a.disable_revise,a.disable_comment,a.created_by,a.deleted,a.created_at,a.updated_at  FROM articles a WHERE a.deleted=0 AND a.di in ? ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect, ids); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticlesByIDs err(%+v)", err))
		return
	}
	return
}
