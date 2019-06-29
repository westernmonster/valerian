package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_addTopicVersionSQL       = "INSERT INTO topic_versions( id,topic_id,name,seq,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"
	_updateTopicVersionSQL    = "UPDATE topic_versions SET topic_id=?,name=?,seq=?,updated_at=? WHERE id=?"
	_getTopicVersionsSQL      = "SELECT a.* FROM topic_versions a WHERE a.deleted=0 AND a.topic_id=? ORDER BY a.seq "
	_getTopicVersionByNameSQL = "SELECT a.id, a.topic_id,a.name, b.name as topic_name FROM  topic_versions a LEFT JOIN topics b ON a.topic_id = b.id  WHERE a.topic_id=? AND a.name=? AND a.deleted=0"
)

func (p *Dao) AddTopicVersion(c context.Context, node sqalx.Node, item *model.TopicVersion) (err error) {
	if _, err = node.ExecContext(c, _addTopicVersionSQL, item.ID, item.TopicID, item.Name, item.Seq, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicVersion error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateTopicVersion(c context.Context, node sqalx.Node, item *model.TopicVersion) (err error) {
	if _, err = node.ExecContext(c, _updateTopicVersionSQL, item.TopicID, item.Name, item.Seq, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicVersion error(%+v), item(%+v)", err, item))
	}
	return
}

func (p *Dao) GetTopicVersions(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicVersionResp, err error) {
	items = make([]*model.TopicVersionResp, 0)
	if err = node.SelectContext(c, &items, _getTopicVersionsSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicVersions error(%+v)", err))

	}
	return
}

func (p *Dao) GetTopicVersionByName(c context.Context, node sqalx.Node, topicID int64, versionName string) (item *model.TopicVersionResp, err error) {
	item = new(model.TopicVersionResp)

	if err = node.GetContext(c, item, _getTopicVersionByNameSQL, topicID, versionName); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicVersionByName error(%+v), topic_id(%d) version_name(%s)", err, topicID, versionName))
	}

	return
}
