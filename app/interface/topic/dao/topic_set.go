package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addTopicSetSQL = "INSERT INTO topic_sets( id,deleted,created_at,updated_at) VALUES ( ?,?,?,?)"
)

func (p *Dao) AddTopicSet(c context.Context, node sqalx.Node, item *model.TopicSet) (err error) {
	if _, err = node.ExecContext(c, _addTopicSetSQL, item.ID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicSet error(%+v), item(%+v)", err, item))
	}

	return
}
