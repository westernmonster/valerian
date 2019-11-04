package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/davecgh/go-spew/spew"
)

func (p *Dao) GetArticleHistoriesPaged(c context.Context, node sqalx.Node, articleID int64, limit, offset int) (items []*model.ArticleHistory, err error) {
	items = make([]*model.ArticleHistory, 0)
	sqlSelect := "SELECT a.* FROM article_histories a WHERE a.deleted=0 AND a.article_id=? ORDER BY a.id DESC limit ?,?"

	fmt.Println(articleID, limit, offset)

	if err = node.SelectContext(c, &items, sqlSelect, articleID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesPaged err(%+v) article_id(%d) limit(%d) offset(%d)", err, articleID, limit, offset))
		return
	}

	spew.Dump(items)
	return
}

func (p *Dao) GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error) {
	sqlSelect := "SELECT a.seq FROM article_histories a WHERE a.deleted=0 AND a.article_id=? ORDER BY a.ID desc LIMIT 1"
	if err = node.GetContext(c, &seq, sqlSelect, articleID); err != nil {
		if err == sql.ErrNoRows {
			seq = 0
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesMaxSeq error(%+v), article_id(%d)", err, articleID))
		return
	}
	return
}

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
func (p *Dao) GetArticleHistories(c context.Context, node sqalx.Node) (items []*model.ArticleHistory, err error) {
	items = make([]*model.ArticleHistory, 0)
	sqlSelect := "SELECT a.* FROM article_histories a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistories err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetArticleHistoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ArticleHistory, err error) {
	items = make([]*model.ArticleHistory, 0)
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
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["diff"]; ok {
		clause += " AND a.diff =?"
		condition = append(condition, val)
	}
	if val, ok := cond["updated_by"]; ok {
		clause += " AND a.updated_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["change_desc"]; ok {
		clause += " AND a.change_desc =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM article_histories a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error) {
	item = new(model.ArticleHistory)
	sqlSelect := "SELECT a.* FROM article_histories a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoryByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetArticleHistoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ArticleHistory, err error) {
	item = new(model.ArticleHistory)
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
	if val, ok := cond["content"]; ok {
		clause += " AND a.content =?"
		condition = append(condition, val)
	}
	if val, ok := cond["content_text"]; ok {
		clause += " AND a.content_text =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["diff"]; ok {
		clause += " AND a.diff =?"
		condition = append(condition, val)
	}
	if val, ok := cond["updated_by"]; ok {
		clause += " AND a.updated_by =?"
		condition = append(condition, val)
	}
	if val, ok := cond["change_desc"]; ok {
		clause += " AND a.change_desc =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM article_histories a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleHistoriesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	sqlInsert := "INSERT INTO article_histories( id,article_id,content,content_text,seq,diff,updated_by,change_desc,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.ArticleID, item.Content, item.ContentText, item.Seq, item.Diff, item.UpdatedBy, item.ChangeDesc, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleHistories err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	sqlUpdate := "UPDATE article_histories SET article_id=?,content=?,content_text=?,seq=?,diff=?,updated_by=?,change_desc=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.ArticleID, item.Content, item.ContentText, item.Seq, item.Diff, item.UpdatedBy, item.ChangeDesc, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleHistories err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelArticleHistory(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE article_histories SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleHistories err(%+v), item(%+v)", err, id))
		return
	}

	return
}
