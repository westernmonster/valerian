package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/account-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetDiscussionByID(c context.Context, node sqalx.Node, id int64) (item *model.Discussion, err error) {
	item = new(model.Discussion)
	sqlSelect := "SELECT a.id,a.topic_id,a.category_id,a.created_by,a.title,a.content,a.content_text,a.deleted,a.created_at,a.updated_at FROM discussions a WHERE a.deleted=0 AND a.id=?"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetDiscussionByID err(%+v), id(%+v)", err, id))
	}

	return
}
