package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) IsTopicExist(c context.Context, node sqalx.Node, id int64) (exist bool, err error) {
	var tid int64
	sqlSelect := "SELECT a.id FROM topics a WHERE a.id=? AND a.deleted=0"
	if err = node.GetContext(c, &tid, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			exist = false
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.IsTopicExist err(%+v), id(%+v)", err, id))
		return
	}

	exist = true
	return
}

func (p *Dao) IsTopicMember(c context.Context, node sqalx.Node, aid, topicID int64) (isMember bool, role string, err error) {
	sqlSelect := "SELECT a.role FROM topic_members a WHERE a.deleted=0 AND a.topic_id=? AND a.account_id=?"
	if err = node.GetContext(c, &role, sqlSelect, topicID, aid); err != nil {
		if err == sql.ErrNoRows {
			isMember = false
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.IsTopicMember err(%+v), aid(%d) topic_id(%d)", err, aid, topicID))
		return
	}

	isMember = true
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
