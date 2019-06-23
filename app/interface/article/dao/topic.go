package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getTopicSQL = "SELECT a.* FROM topics a WHERE a.id=? AND a.deleted=0"
)

// GetByID get record by ID
func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)

	if err = node.GetContext(c, item, _getTopicSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicByID error(%+v), id(%d)", err, id))
	}

	return
}
