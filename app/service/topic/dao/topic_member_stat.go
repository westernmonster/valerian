package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) AddTopicMemberStat(c context.Context, node sqalx.Node, item *model.TopicMemberStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO topic_member_stats( topic_id,member_count,created_at,updated_at) VALUES (?, ?, ?, ?)"

	if _, err = node.ExecContext(c, sqlInsert, item.TopicID, item.MemberCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicMemberStat err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) GetTopicMemberStatForUpdate(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicMemberStat, err error) {
	item = new(model.TopicMemberStat)
	sqlSelect := "SELECT a.* FROM topic_member_stats a WHERE a.topic_id=? FOR UPDATE"

	if err = node.GetContext(c, item, sqlSelect, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberStatForUpdate err(%+v), topic_id(%+v)", err, topicID))
	}

	return
}

func (p *Dao) GetTopicMemberStat(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicMemberStat, err error) {
	item = new(model.TopicMemberStat)
	sqlSelect := "SELECT a.* FROM topic_member_stats a WHERE a.topic_id=?"

	if err = node.GetContext(c, item, sqlSelect, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberStat err(%+v), topic_id(%+v)", err, topicID))
	}

	return
}

func (p *Dao) UpdateTopicMemberStat(c context.Context, node sqalx.Node, item *model.TopicMemberStat) (err error) {
	sqlInsert := "UPDATE topic_member_stats SET member_count=?,updated_at=? WHERE topic_id=?"

	if _, err = node.ExecContext(c, sqlInsert, item.MemberCount, item.UpdatedAt, item.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicMemberStat err(%+v), item(%+v)", err, item))
		return
	}

	return
}
