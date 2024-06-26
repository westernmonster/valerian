package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetArticleRevisesPaged(c context.Context, node sqalx.Node, articleID int64, sort string, limit, offset int) (items []*model.Revise, err error) {
	items = make([]*model.Revise, 0)
	sqlSelect := "SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at  FROM revises a LEFT JOIN revise_stats b ON a.id = b.revise_id WHERE a.deleted=0 AND a.article_id=? %s limit ?,?"
	s := "ORDER BY a.id DESC"
	if sort == "hot" {
		s = "ORDER BY b.like_count DESC, b.dislike_count DESC, a.id DESC"
	}

	sqlSelect = fmt.Sprintf(sqlSelect, s)

	if err = node.SelectContext(c, &items, sqlSelect, articleID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleRevisesPaged err(%+v) article_id(%d) limit(%d) offset(%d)", err, articleID, limit, offset))
	}
	return
}

func (p *Dao) GetUserReviseIDsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.id FROM revises a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC limit ?,?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserReviseIDsPaged err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserReviseIDsPaged err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
	return
}

// GetAll get all records
func (p *Dao) GetRevises(c context.Context, node sqalx.Node) (items []*model.Revise, err error) {
	items = make([]*model.Revise, 0)
	sqlSelect := "SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at   FROM revises a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRevises err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetRevisesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Revise, err error) {
	items = make([]*model.Revise, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["article_id"]; ok {
		clause += " AND a.article_id =?"
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
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at FROM revises a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRevisesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetReviseByID(c context.Context, node sqalx.Node, id int64) (item *model.Revise, err error) {
	item = new(model.Revise)
	sqlSelect := "SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at  FROM revises a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetReviseByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetReviseByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Revise, err error) {
	item = new(model.Revise)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["article_id"]; ok {
		clause += " AND a.article_id =?"
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
	if val, ok := cond["created_by"]; ok {
		clause += " AND a.created_by =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at FROM revises a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetRevisesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error) {
	sqlInsert := "INSERT INTO revises( id,article_id,title,content,content_text,created_by,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.ArticleID, item.Title, item.Content, item.ContentText, item.CreatedBy, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRevises err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateRevise(c context.Context, node sqalx.Node, item *model.Revise) (err error) {
	sqlUpdate := "UPDATE revises SET article_id=?,title=?,content=?,content_text=?,created_by=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.ArticleID, item.Title, item.Content, item.ContentText, item.CreatedBy, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateRevises err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// GetAll get all records
func (p *Dao) GetRevisesByIDs(c context.Context, node sqalx.Node, ids []int64) (items []*model.Revise, err error) {
	items = make([]*model.Revise, 0)
	strIDs := Int64Array2StringArray(ids)
	sqlSelect := "SELECT a.id,a.article_id,a.title,a.content,a.content_text,a.created_by,a.deleted,a.created_at,a.updated_at   FROM revises a WHERE a.deleted=0 AND a.id IN (?) ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect, strings.Join(strIDs, ",")); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRevisesByIDs err(%+v)", err))
		return
	}
	return
}


