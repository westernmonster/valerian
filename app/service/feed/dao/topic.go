package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/feed/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	item = new(model.Topic)
	sqlSelect := "SELECT a.id,a.name,a.avatar,a.bg,a.introduction,a.allow_discuss,a.allow_chat,a.is_private,a.view_permission,a.edit_permission,a.join_permission,a.catalog_view_type,a.topic_home,a.created_by,a.deleted,a.created_at,a.updated_at FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}

func (p *Dao) GetTopicMemberIDs(c context.Context, node sqalx.Node, topicID int64) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.account_id FROM topic_members a WHERE a.deleted=0 AND a.topic_id=?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

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
