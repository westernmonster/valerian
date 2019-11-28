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
func (p *Dao) GetReviseFiles(c context.Context, node sqalx.Node) (items []*model.ReviseFile, err error) {
	items = make([]*model.ReviseFile, 0)
	sqlSelect := "SELECT a.id,a.file_name,a.file_url,a.seq,a.revise_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM revise_files a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFiles err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetReviseFilesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ReviseFile, err error) {
	items = make([]*model.ReviseFile, 0)
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
	if val, ok := cond["revise_id"]; ok {
		clause += " AND a.revise_id =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.file_name,a.file_url,a.seq,a.revise_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM revise_files a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFilesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetReviseFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ReviseFile, err error) {
	item = new(model.ReviseFile)
	sqlSelect := "SELECT a.id,a.file_name,a.file_url,a.seq,a.revise_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM revise_files a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFileByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetReviseFileByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ReviseFile, err error) {
	item = new(model.ReviseFile)
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
	if val, ok := cond["revise_id"]; ok {
		clause += " AND a.revise_id =?"
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

	sqlSelect := fmt.Sprintf("SELECT a.id,a.file_name,a.file_url,a.seq,a.revise_id,a.deleted,a.created_at,a.updated_at,a.file_type,a.pdf_url FROM revise_files a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetReviseFilesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddReviseFile(c context.Context, node sqalx.Node, item *model.ReviseFile) (err error) {
	sqlInsert := "INSERT INTO revise_files( id,file_name,file_url,seq,revise_id,deleted,created_at,updated_at,file_type,pdf_url) VALUES ( ?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.FileName, item.FileURL, item.Seq, item.ReviseID, item.Deleted, item.CreatedAt, item.UpdatedAt, item.FileType, item.PdfURL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddReviseFiles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateReviseFile(c context.Context, node sqalx.Node, item *model.ReviseFile) (err error) {
	sqlUpdate := "UPDATE revise_files SET file_name=?,file_url=?,seq=?,revise_id=?,updated_at=?,file_type=?,pdf_url=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.FileName, item.FileURL, item.Seq, item.ReviseID, item.UpdatedAt, item.FileType, item.PdfURL, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateReviseFiles err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelReviseFile(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE revise_files SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelReviseFiles err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) DelReviseFileByCond(c context.Context, node sqalx.Node, reviseID int64) (err error) {
	sqlDelete := "UPDATE revise_files SET deleted=1 WHERE revise_id=? "

	if _, err = node.ExecContext(c, sqlDelete, reviseID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelReviseFileByCond err(%+v), revise_id(%+v)", err, reviseID))
		return
	}

	return
}
