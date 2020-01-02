package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

// IsAllowedViewMember 用户是否是允许查看的成员
func (p *Dao) IsAllowedViewMember(c context.Context, node sqalx.Node, aid, discussionID int64) (canView bool, err error) {
	items := make([]int64, 0)
	// 话题成员
	// 话题的关联话题成员
	sqlSelect := `
	SELECT x.account_id  FROM topic_members x WHERE x.deleted = 0 AND x.account_id = ? AND x.topic_id  IN (
		SELECT a.to_topic_id FROM auth_topics a WHERE a.deleted= 0 AND a.topic_id IN(SELECT a.topic_id FROM discussions a WHERE a.deleted=0 AND a.id = ?)
		UNION
		SELECT a.topic_id FROM discussions a WHERE a.deleted= 0 AND a.id=?
	)
	`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, discussionID, discussionID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsAllowedViewMember err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.IsAllowedViewMember err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if len(items) > 0 {
		canView = true
	}
	return
}

// GetTopicViewPermissionByID 获取话题的查看权限
func (p *Dao) GetTopicViewPermissionByID(c context.Context, node sqalx.Node, id int64) (viewPermission string, err error) {
	sqlSelect := "SELECT a.a.view_permission FROM topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, &viewPermission, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			viewPermission = ""
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicViewPermissionByID err(%+v), id(%+v)", err, id))
		return
	}

	return
}
