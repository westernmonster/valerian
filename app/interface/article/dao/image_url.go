package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetImageUrls(c context.Context, node sqalx.Node) (items []*model.ImageURL, err error) {
	items = make([]*model.ImageURL, 0)
	sqlSelect := "SELECT a.* FROM image_urls a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetImageUrls err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetImageUrlsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ImageURL, err error) {
	items = make([]*model.ImageURL, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["url"]; ok {
		clause += " AND a.url =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM image_urls a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetImageUrlsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetImageURLByID(c context.Context, node sqalx.Node, id int64) (item *model.ImageURL, err error) {
	item = new(model.ImageURL)
	sqlSelect := "SELECT a.* FROM image_urls a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetImageURLByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetImageURLByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ImageURL, err error) {
	item = new(model.ImageURL)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}
	if val, ok := cond["url"]; ok {
		clause += " AND a.url =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM image_urls a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetImageUrlsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error) {
	sqlInsert := "INSERT INTO image_urls( id,target_id,target_type,url,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TargetID, item.TargetType, item.URL, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddImageUrls err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error) {
	sqlUpdate := "UPDATE image_urls SET target_id=?,target_type=?,url=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TargetID, item.TargetType, item.URL, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateImageUrls err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelImageURL(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE image_urls SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelImageUrls err(%+v), item(%+v)", err, id))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelImageURLByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE image_urls SET deleted=1 WHERE target_id=? AND target_type =?"

	if _, err = node.ExecContext(c, sqlDelete, targetID, targetType); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelImageUrls err(%+v), target_id(%d) target_type(%s)", err, targetID, targetType))
		return
	}

	return
}
