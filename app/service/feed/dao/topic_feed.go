package dao

import (
	"context"
	"fmt"

	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

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
func (p *Dao) DelTopicFeed(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_feeds SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicFeeds err(%+v), item(%+v)", err, id))
		return
	}

	return
}
