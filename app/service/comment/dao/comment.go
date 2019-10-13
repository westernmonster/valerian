package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by ID
func (p *Dao) GetCommentByID(c context.Context, node sqalx.Node, id int64) (item *model.Comment, err error) {
	item = new(model.Comment)
	sqlSelect := "SELECT a.* FROM comments a WHERE a.id=? "

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetCommentByID err(%+v), id(%+v)", err, id))
	}

	return
}
