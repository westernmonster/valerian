package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/account-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

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
