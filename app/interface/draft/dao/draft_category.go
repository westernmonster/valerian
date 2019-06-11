package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getUserDraftCategoriesSQL = "SELECT a.id,a.`name`,a.color_id,b.color FROM draft_categories a LEFT JOIN colors b ON a.color_id=b.id WHERE a.deleted=0 ORDER BY a.id"
	_addDraftCategorySQL       = "INSERT INTO draft_categories( id,name,color_id, account_id,deleted,created_at,updated_at) VALUES (?,?,?,?,?,?)"
	_updateDraftCategorySQL    = "UPDATE draft_categories SET name=?,color_id=?,account_id=?,updated_at=? WHERE id=?"
	_delDraftCategorySQL       = "UPDATE draft_categories SET deleted=1 WHERE id=? "
	_getDraftCategorySQL       = "SELECT a.* FROM draft_categories a WHERE a.deleted=0 AND a.id=? "
)

func (p *Dao) GetUserDraftCategories(c context.Context, node sqalx.Node, aid int64) (items []*model.DraftCategoryResp, err error) {
	items = make([]*model.DraftCategoryResp, 0)

	if err = node.SelectContext(c, &items, _getUserDraftCategoriesSQL, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDraftCategories error(%+v)", err))
	}

	return
}

func (p *Dao) AddDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error) {
	if _, err = node.ExecContext(c, _addDraftCategorySQL, item.ID, item.Name, item.ColorID, item.AccountID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDraftCategory error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) UpdateDraftCategory(c context.Context, node sqalx.Node, item *model.DraftCategory) (err error) {
	if _, err = node.ExecContext(c, _updateDraftCategorySQL, item.Name, item.ColorID, item.AccountID, item.UpdatedAt, item.ID); err != nil {
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

func (p *Dao) GetDraftCategory(c context.Context, node sqalx.Node, id int64) (item *model.DraftCategory, err error) {
	item = new(model.DraftCategory)

	if err = node.GetContext(c, item, _getDraftCategorySQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDraftCategory error(%+v), id(%d)", err, id))
	}

	return
}
