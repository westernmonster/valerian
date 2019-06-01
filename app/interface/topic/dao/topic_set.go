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

	_getTopicVersionsSQL = "SELECT a.id AS topic_set_id,b.name as topic_name, b.id AS topic_id,b.version_name FROM topic_sets a LEFT JOIN topics b ON a.id=b.topic_set_id WHERE a.id=?"
)

func (p *Dao) AddTopicSet(c context.Context, node sqalx.Node, item *model.TopicSet) (err error) {
	if _, err = node.ExecContext(c, _addTopicSetSQL, item.ID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicSet error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) GetTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	items = make([]*model.TopicVersionResp, 0)
	if err = node.SelectContext(c, &items, _getTopicVersionsSQL, topicSetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicVersions error(%+v), topic set id(%d)", err, topicSetID))
	}
	return
}
