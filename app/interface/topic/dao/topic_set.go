package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

const (
	_addTopicSetSQL = "INSERT INTO topic_sets( id,deleted,created_at,updated_at) VALUES ( ?,?,?,?)"

	_getTopicVersionsSQL      = "SELECT a.id FROM  topics a  WHERE a.topic_set_id=? AND a.deleted=0"
	_getTopicVersionByNameSQL = "SELECT a.topic_set_id,a.name as topic_name, a.id AS topic_id,a.version_name FROM  topics a  WHERE a.topic_set_id=? AND a.version_name=? AND a.deleted=0"
)

func (p *Dao) AddTopicSet(c context.Context, node sqalx.Node, item *model.TopicSet) (err error) {
	if _, err = node.ExecContext(c, _addTopicSetSQL, item.ID, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicSet error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) GetTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []int64, err error) {
	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, _getTopicVersionsSQL, topicSetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicVersions error(%+v), topic set id(%d)", err, topicSetID))
		return
	}

	defer rows.Close()

	items = make([]int64, 0)

	for rows.Next() {
		var tid int64
		err = rows.Scan(&tid)
		if err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicVersions scan error(%+v), topic set id(%d)", err, topicSetID))
			return nil, err
		}

		items = append(items, tid)
	}

	return
}

func (p *Dao) GetTopicVersionByName(c context.Context, node sqalx.Node, topicSetID int64, versionName string) (item *model.TopicVersionResp, err error) {
	item = new(model.TopicVersionResp)

	if err = node.GetContext(c, item, _getTopicVersionByNameSQL, topicSetID, versionName); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicVersionByName error(%+v), topic_set_id(%d) version_name(%s)", err, topicSetID, versionName))
	}

	return
}
