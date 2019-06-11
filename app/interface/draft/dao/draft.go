package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getUserDraftsSQL = "SELECT a.* FROM drafts a WHERE a.deleted=0 AND a.created_by=? ORDER BY a.id DESC "
	_addDraftSQL      = "INSERT INTO drafts( id,title,content,created_by,category_id,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?)"
	_updateDraftSQL   = "UPDATE drafts SET title=?,content=?,created_by=?,category_id=?,updated_at=? WHERE id=?"
	_delDraftSQL      = "UPDATE drafts SET deleted=1 WHERE id=? "
)

func (p *Dao) GetUserDrafts(c context.Context, node sqalx.Node, aid int64) (items []*model.Draft, err error) {
	items = make([]*model.Draft, 0)

	if err = node.SelectContext(c, &items, _getUserDraftsSQL, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserDrafts error(%+v) aid(%d)", err, aid))
	}

	return
}

func (p *Dao) AddDraft(c context.Context, node sqalx.Node, item *model.Draft) (err error) {
	if _, err = node.ExecContext(c, _addDraftSQL, item.ID, item.Title, item.Content, item.CreatedBy, item.CategoryID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddDraft error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) UpdateDraft(c context.Context, node sqalx.Node, item *model.Draft) (err error) {
	if _, err = node.ExecContext(c, _updateDraftSQL, item.Title, item.Content, item.CreatedBy, item.CategoryID, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateDraft error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) DelDraft(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delDraftSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelDraft error(%+v), id(%d)", err, id))
	}
	return
}
