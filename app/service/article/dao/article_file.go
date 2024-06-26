package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetArticleFiles(c context.Context, node sqalx.Node) (items []*model.ArticleFile, err error) {
	items = make([]*model.ArticleFile, 0)
	sqlSelect := "SELECT a.id,a.file_name,a.file_url,a.seq,a.article_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM article_files a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFiles err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetArticleFilesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ArticleFile, err error) {
	items = make([]*model.ArticleFile, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_name"]; ok {
		clause += " AND a.file_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_url"]; ok {
		clause += " AND a.file_url =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["article_id"]; ok {
		clause += " AND a.article_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_type"]; ok {
		clause += " AND a.file_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["pdf_url"]; ok {
		clause += " AND a.pdf_url =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.file_name,a.file_url,a.seq,a.article_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM article_files a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFilesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error) {
	item = new(model.ArticleFile)
	sqlSelect := "SELECT a.id,a.file_name,a.file_url,a.seq,a.article_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM article_files a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFileByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetArticleFileByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ArticleFile, err error) {
	item = new(model.ArticleFile)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_name"]; ok {
		clause += " AND a.file_name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_url"]; ok {
		clause += " AND a.file_url =?"
		condition = append(condition, val)
	}
	if val, ok := cond["seq"]; ok {
		clause += " AND a.seq =?"
		condition = append(condition, val)
	}
	if val, ok := cond["article_id"]; ok {
		clause += " AND a.article_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["file_type"]; ok {
		clause += " AND a.file_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["pdf_url"]; ok {
		clause += " AND a.pdf_url =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.file_name,a.file_url,a.seq,a.article_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM article_files a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetArticleFilesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	sqlInsert := "INSERT INTO article_files( id,file_name,file_url,seq,article_id,deleted,created_at,updated_at,file_type,pdf_url) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.FileName, item.FileURL, item.Seq, item.ArticleID, item.Deleted, item.CreatedAt, item.UpdatedAt, item.FileType, item.PdfURL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddArticleFiles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	sqlUpdate := "UPDATE article_files SET file_name=?,file_url=?,seq=?,article_id=?,updated_at=?,file_type=?,pdf_url=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.FileName, item.FileURL, item.Seq, item.ArticleID, item.UpdatedAt, item.FileType, item.PdfURL, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateArticleFiles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE article_files SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelArticleFiles err(%+v), item(%+v)", err, id))
		return
	}

	return
}
