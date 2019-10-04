package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicStatForUpdate(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicStat, err error) {
	item = new(model.TopicStat)
	sqlSelect := "SELECT a.* FROM topic_stats a WHERE a.topic_id=? FOR UPDATE"

	if err = node.GetContext(c, item, sqlSelect, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicStatForUpdate err(%+v), topic_id(%+v)", err, topicID))
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO topic_stats( topic_id,member_count,article_count,discussion_count,created_at,updated_at) VALUES ( ?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.TopicID, item.MemberCount, item.ArticleCount, item.DiscussionCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error) {
	sqlUpdate := "UPDATE topic_stats SET member_count=?,article_count=?,discussion_count=?,updated_at=? WHERE topic_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.MemberCount, item.ArticleCount, item.DiscussionCount, item.UpdatedAt, item.TopicID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error) {
	sqlUpdate := "UPDATE topic_stats SET member_count=member_count+?,article_count=article_count+?,discussion_count=discussion_count+?,updated_at=? WHERE topic_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.MemberCount, item.ArticleCount, item.DiscussionCount, item.UpdatedAt, item.TopicID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrTopicStat err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// GetAll get all records
func (p *Dao) GetTopicStatByID(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicStat, err error) {
	item = new(model.TopicStat)
	sqlSelect := "SELECT a.* FROM topic_stats a WHERE a.topic_id=? "

	if err = node.GetContext(c, item, sqlSelect, topicID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicStatByID err(%+v), topic_id(%+v)", err, topicID))
	}

	return
}
