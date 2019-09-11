package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getTopicFollowRequestsSQL   = "SELECT a.* FROM topic_follow_requests a WHERE a.deleted=0 AND a.topic_id=? AND a.status=? ORDER BY a.id "
	_getTopicFollowRequestSQL    = "SELECT a.* FROM topic_follow_requests a WHERE a.deleted=0 AND a.topic_id=? AND a.account_id=? ORDER BY a.id DESC limit 1"
	_addTopicFollowRequestSQL    = "INSERT INTO topic_follow_requests( id,account_id,topic_id,status,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"
	_updateTopicFollowRequestSQL = "UPDATE topic_follow_requests SET account_id=?,topic_id=?,status=?,updated_at=? WHERE id=?"
	_delTopicFollowRequestSQL    = "UPDATE topic_follow_requests SET deleted=1 WHERE id=? "
)

func (p *Dao) GetTopicFollowRequests(c context.Context, node sqalx.Node, topicID int64, status int) (items []*model.TopicFollowRequest, err error) {
	items = make([]*model.TopicFollowRequest, 0)

	if err = node.SelectContext(c, &items, _getTopicFollowRequestsSQL, topicID, status); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequests error(%+v), topic id(%d) status(%d)", err, topicID, status))
	}
	return
}

func (p *Dao) GetTopicFollowRequest(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicFollowRequest, err error) {
	item = new(model.TopicFollowRequest)

	if err = node.GetContext(c, item, _getTopicFollowRequestSQL, topicID, aid); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFollowRequest error(%+v), topic id(%d) account id(%d)", err, topicID, aid))
	}

	return
}

func (p *Dao) AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	if _, err = node.ExecContext(c, _addTopicFollowRequestSQL, item.ID, item.AccountID, item.TopicID, item.Status, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.addTopicFollowRequest error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	if _, err = node.ExecContext(c, _updateTopicFollowRequestSQL, item.AccountID, item.TopicID, item.Status, item.UpdatedAt, item.ID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicFollowRequest error(%+v), item(%+v)", err, item))
	}

	return
}

func (p *Dao) DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error) {
	if _, err = node.ExecContext(c, _delTopicFollowRequestSQL, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFollowRequest error(%+v), topic follow id(%d)", err, id))
	}

	return
}
