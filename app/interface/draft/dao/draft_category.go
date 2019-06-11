package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getDraftCategoriesSQL  = "SELECT a.* FROM draft_categories a WHERE a.deleted=0 ORDER BY a.id"
	_addDraftCategorySQL    = "INSERT INTO draft_categories( id,name,color_id,deleted,created_at,updated_at) VALUES (?,?,?,?,?,?)"
	_updateDraftCategorySQL = "UPDATE draft_categories SET name=?,color_id=?,updated_at=? WHERE id=?"
	_delDraftCategorySQL    = "UPDATE draft_categories SET deleted=1 WHERE id=? "
)

func (p *Dao) GetAllDraftCategories(c context.Context, node sqalx.Node) (items []*model.DraftCategory, err error) {
	items = make([]*model.DraftCategory, 0)

	if err = node.SelectContext(c, &items, _getDraftCategoriesSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllDraftCategories error(%+v)", err))
	}

	return
}

func (p *Dao) AddDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error) {
	if _, err = node.ExecContext(c, _addDraftCategorySQL, item.ID, item.Name, item.ColorID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDraftCategory error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) UpdateDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error) {
	if _, err = node.ExecContext(c, _updateDraftCategorySQL, item.Name, item.ColorID, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDraftCategory error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) DelDraftCategory(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delDraftCategorySQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDraftCategory error(%+v), id(%d)", err, id))
	}
	return
}
