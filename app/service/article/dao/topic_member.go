package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetMemberBelongsTopicIDs(c context.Context, node sqalx.Node, accountID int64) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.topic_id FROM topic_members a WHERE a.deleted=0 AND a.account_id=?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, accountID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetMemberBelongsTopicIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetMemberBelongsTopicIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) IsTopicAdmin(c context.Context, node sqalx.Node, topicID, aid int64) (isAdmin bool, err error) {
	item := new(model.TopicMember)
	sqlSelect := "SELECT a.id,a.topic_id,a.account_id,a.role,a.deleted,a.created_at,a.updated_at FROM topic_members a WHERE a.deleted=0 AND a.account_id=? AND a.topic_id=? AND a.role IN(?,?) LIMIT 1"

	if err = node.GetContext(c, item, sqlSelect, aid, topicID, model.MemberRoleAdmin, model.MemberRoleOwner); err != nil {
		if err == sql.ErrNoRows {
			isAdmin = false
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.IsTopicAdmin err(%+v), aid(%d) topic_id(%d)", err, aid, topicID))
		return
	}

	if item != nil {
		isAdmin = true
	}

	return
}

func (p *Dao) IsTopicMember(c context.Context, node sqalx.Node, topicID, aid int64) (isMember bool, err error) {
	item := new(model.TopicMember)
	sqlSelect := "SELECT a.id,a.topic_id,a.account_id,a.role,a.deleted,a.created_at,a.updated_at FROM topic_members a WHERE a.deleted=0 AND a.account_id=? AND a.topic_id=? LIMIT 1"

	if err = node.GetContext(c, item, sqlSelect, aid, topicID, model.MemberRoleAdmin, model.MemberRoleOwner); err != nil {
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
