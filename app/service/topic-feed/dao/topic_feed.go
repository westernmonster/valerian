package dao

import (
	"context"
	"fmt"

	"valerian/app/service/topic-feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicFeedPaged(c context.Context, node sqalx.Node, topicID int64, limit, offset int) (items []*model.TopicFeed, err error) {
	items = make([]*model.TopicFeed, 0)

	sql := "SELECT a.id,a.topic_id,a.action_type,a.action_time,a.action_text,a.actor_id,a.actor_type,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM topic_feeds a WHERE a.deleted=0 AND a.topic_id=? ORDER BY a.id DESC limit ?,?"

	if err = node.SelectContext(c, &items, sql, topicID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicFeedPaged error(%+v), topic_id(%d) limit(%d) offset(%d)", err, topicID, limit, offset))
	}
	return
}

func (p *Dao) AddTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error) {
	sqlInsert := "INSERT INTO topic_feeds( id,topic_id,action_type,action_time,action_text,actor_id,actor_type,target_id,target_type,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicFeed(c context.Context, node sqalx.Node, item *model.TopicFeed) (err error) {
	sqlUpdate := "UPDATE topic_feeds SET topic_id=?,action_type=?,action_time=?,action_text=?,actor_id=?,actor_type=?,target_id=?,target_type=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.ActionType, item.ActionTime, item.ActionText, item.ActorID, item.ActorType, item.TargetID, item.TargetType, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicFeeds err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicFeedByTopicID(c context.Context, node sqalx.Node, topicID int64) (err error) {
	sqlDelete := "UPDATE topic_feeds SET deleted=1 WHERE topic_id=? "

	if _, err = node.ExecContext(c, sqlDelete, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFeeds err(%+v), topic_id(%+v)", err, topicID))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicFeedByCond(c context.Context, node sqalx.Node, topicID int64, targetType string, targetID int64) (err error) {
	sqlDelete := "UPDATE topic_feeds SET deleted=1 WHERE topic_id=? AND target_type=? AND target_id=?"

	if _, err = node.ExecContext(c, sqlDelete, topicID, targetType, targetID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFeeds err(%+v), topic_id(%+v), target_type(%+v), target_id(%+v)", err, topicID, targetType, targetID))
		return
	}

	return
}
