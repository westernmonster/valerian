package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by Account ID
func (p *Dao) GetTopicStatByID(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicResStat, err error) {
	item = new(model.TopicResStat)
	sqlSelect := "SELECT a.* FROM topic_res_stats a WHERE a.topic_id=? AND a.deleted=0"

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

// Insert insert a new record
func (p *Dao) AddTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error) {
	sqlInsert := "INSERT IGNORE INTO topic_res_stats( topic_id,article_count,discussion_count,created_at,updated_at) VALUES ( ?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.TopicID, item.ArticleCount, item.DiscussionCount, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error) {
	sqlUpdate := "UPDATE topic_res_stats SET topic_id=?,article_count=?,discussion_count=?,updated_at=? WHERE topic_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.ArticleCount, item.DiscussionCount, item.UpdatedAt, item.TopicID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}

func (p *Dao) IncrTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error) {
	sqlUpdate := "UPDATE topic_res_stats SET article_count=article_count+?,discussion_count=discussion_count+?,updated_at=? WHERE topic_id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.ArticleCount, item.DiscussionCount, time.Now().Unix(), item.TopicID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IncrTopicStats err(%+v), item(%+v)", err, item))
		return
	}

	return
}
