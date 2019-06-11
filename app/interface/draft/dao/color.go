package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getColorsSQL   = "SELECT a.* FROM colors a WHERE a.deleted=0 ORDER BY a.id "
	_addColorSQL    = "INSERT INTO colors( id,name,color,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"
	_updateColorSQL = "UPDATE colors SET name=?,color=?,updated_at=? WHERE id=?"
	_delColorSQL    = "UPDATE colors SET deleted=1 WHERE id=? "
)

func (p *Dao) GetAllColors(c context.Context, node sqalx.Node) (items []*model.Color, err error) {
	items = make([]*model.Color, 0)

	if err = node.SelectContext(c, &items, _getColorsSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllColors error(%+v)", err))
	}

	return
}

func (p *Dao) AddColor(c context.Context, node sqalx.Node, item *model.Color) (err error) {
	if _, err = node.ExecContext(c, _addColorSQL, item.ID, item.Name, item.Color, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddColor error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) UpdateColor(c context.Context, node sqalx.Node, item *model.Color) (err error) {
	if _, err = node.ExecContext(c, _updateColorSQL, item.Name, item.Color, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateColor error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) DelColor(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delColorSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelColor error(%+v), id(%d)", err, id))
	}
	return
}
