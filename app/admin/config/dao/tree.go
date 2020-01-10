package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/admin/config/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetTrees(c context.Context, node sqalx.Node) (items []*model.Tree, err error) {
	items = make([]*model.Tree, 0)
	sqlSelect := "SELECT a.id,a.name,a.tree_id,a.platform_id,a.mark,a.deleted,a.created_at,a.updated_at FROM trees a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTrees err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetTreesByCondPaged(c context.Context, node sqalx.Node, cond map[string]interface{}, page, pageSize int32) (count int32, items []*model.Tree, err error) {
	items = make([]*model.Tree, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tree_id"]; ok {
		clause += " AND a.tree_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["platform_id"]; ok {
		clause += " AND a.platform_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}

	sqlCount := fmt.Sprintf("SELECT COUNT(1) AS count FROM trees a WHERE a.deleted=0 %s", clause)
	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.tree_id,a.platform_id,a.mark,a.deleted,a.created_at,a.updated_at FROM trees a WHERE a.deleted=0 %s ORDER BY a.platform_id,a.tree_id LIMIT ?,?", clause)

	if err = node.GetContext(c, &count, sqlCount, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTreesByCondPaged error(%+v), condition(%+v)", err, condition))
	}

	offset := (page - 1) * pageSize
	condition = append(condition, offset)
	condition = append(condition, pageSize)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTreesByCondPaged err(%+v), condition(%+v)", err, condition))
		return
	}

	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTreesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Tree, err error) {
	items = make([]*model.Tree, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tree_id"]; ok {
		clause += " AND a.tree_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["platform_id"]; ok {
		clause += " AND a.platform_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.tree_id,a.platform_id,a.mark,a.deleted,a.created_at,a.updated_at FROM trees a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTreesByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTreeByID(c context.Context, node sqalx.Node, id int64) (item *model.Tree, err error) {
	item = new(model.Tree)
	sqlSelect := "SELECT a.id,a.name,a.tree_id,a.platform_id,a.mark,a.deleted,a.created_at,a.updated_at FROM trees a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTreeByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTreeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Tree, err error) {
	item = new(model.Tree)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["name"]; ok {
		clause += " AND a.name =?"
		condition = append(condition, val)
	}
	if val, ok := cond["tree_id"]; ok {
		clause += " AND a.tree_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["platform_id"]; ok {
		clause += " AND a.platform_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["mark"]; ok {
		clause += " AND a.mark =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.name,a.tree_id,a.platform_id,a.mark,a.deleted,a.created_at,a.updated_at FROM trees a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTreesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTree(c context.Context, node sqalx.Node, item *model.Tree) (err error) {
	sqlInsert := "INSERT INTO trees( id,name,tree_id,platform_id,mark,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.Name, item.TreeID, item.PlatformID, item.Mark, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTrees err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTree(c context.Context, node sqalx.Node, item *model.Tree) (err error) {
	sqlUpdate := "UPDATE trees SET name=?,tree_id=?,platform_id=?,mark=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.Name, item.TreeID, item.PlatformID, item.Mark, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTrees err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTree(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE trees SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTrees err(%+v), item(%+v)", err, id))
		return
	}

	return
}
