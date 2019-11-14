package dao

import (
	"context"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

func (p *Dao) GetUserCanEditArticleIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	items = make([]int64, 0)
	// 授权为edit的话题
	// 1. 编辑权限为管理员，且用户是管理员
	// 2. 编辑权限为成员，且用户为成员
	// 3. 授权话题，授权为可编辑，用户是授权话题成员
	// 4. 授权话题，授权为管理员可编辑，用户是授权话题管理员
	sqlSelect := `
SELECT a.ref_id FROM topic_catalogs a WHERE a.deleted =0 AND a.type = 'article' AND a.permission='edit' AND a.topic_id IN (
SELECT v1.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.role IN ( 'owner', 'admin' ) ) v1
	LEFT JOIN topics b ON v1.topic_id = b.id
WHERE b.deleted = 0 AND b.edit_permission = 'admin'
UNION
SELECT v2.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? ) v2
    LEFT JOIN topics b ON v2.topic_id = b.id
WHERE b.deleted = 0 AND b.edit_permission = 'member'
UNION
SELECT v3.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.role IN ( 'owner', 'admin' ) ) v3
	LEFT JOIN auth_topics b ON v3.topic_id = b.to_topic_id
WHERE b.deleted = 0 AND b.permission = 'admin_edit' UNION
SELECT v4.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? ) v4
	LEFT JOIN auth_topics b ON v4.topic_id = b.to_topic_id
WHERE b.deleted = 0 AND b.permission = 'edit'
    )
`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, aid, aid, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditArticleIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditArticleIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) IsAllowedEditMember(c context.Context, node sqalx.Node, aid, articleID int64) (canEdit bool, err error) {
	items := make([]int64, 0)
	// 授权为edit的话题
	// 1. 编辑权限为管理员，且用户是管理员
	// 2. 编辑权限为成员，且用户为成员
	// 3. 授权话题，授权为可编辑，用户是授权话题成员
	// 4. 授权话题，授权为管理员可编辑，用户是授权话题管理员
	sqlSelect := `
SELECT a.ref_id FROM topic_catalogs a WHERE a.deleted =0 AND a.type = 'article' AND a.ref_id=? AND a.permission='edit' AND a.topic_id IN (
SELECT v1.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id=? AND a.role IN ( 'owner', 'admin' ) ) v1
	LEFT JOIN topics b ON v1.topic_id = b.id
WHERE b.deleted = 0 AND b.edit_permission = 'admin'
UNION
SELECT v2.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id=? ) v2
    LEFT JOIN topics b ON v2.topic_id = b.id
WHERE b.deleted = 0 AND b.edit_permission = 'member'
UNION
SELECT v3.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id=? AND a.role IN ( 'owner', 'admin' ) ) v3
	LEFT JOIN auth_topics b ON v3.topic_id = b.to_topic_id
WHERE b.deleted = 0 AND b.permission = 'admin_edit' UNION
SELECT v4.topic_id FROM ( SELECT a.topic_id FROM topic_members a WHERE a.deleted = 0 AND a.account_id=? ) v4
	LEFT JOIN auth_topics b ON v4.topic_id = b.to_topic_id
WHERE b.deleted = 0 AND b.permission = 'edit'
    )
`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, articleID, aid, aid, aid, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CanEdit err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditArticleIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if len(items) > 0 {
		canEdit = true
	}
	return
}

func (p *Dao) IsAllowedViewMember(c context.Context, node sqalx.Node, aid, articleID int64) (canView bool, err error) {
	items := make([]int64, 0)
	// 授权话题成员
	// 授权话题的关联话题成员
	sqlSelect := `
	SELECT x.account_id  FROM topic_members x WHERE x.deleted = 0 AND x.account_id = ? AND x.topic_id  IN (
		SELECT a.to_topic_id FROM auth_topics a WHERE a.deleted= 0 AND a.topic_id IN(
			SELECT a.topic_id FROM topic_catalogs a WHERE a.deleted= 0 AND a.type= 'article' AND a.ref_id= ?
			)
		UNION
		SELECT a.topic_id FROM topic_catalogs a WHERE a.deleted= 0 AND a.type= 'article' AND a.ref_id= ?
	)
	`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, articleID, articleID); err != nil {
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

func (p *Dao) IsCanEditTopic(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
	// 逻辑
	// 1. 当前话题编辑权限为管理员，且用户是管理员
	// 2. 当前话题编辑权限为成员，且用户为成员
	// 3. 授权话题，授权为可编辑，用户是授权话题成员
	// 4. 授权话题，授权为管理员可编辑，用户是授权话题管理员
	sqlSelect := `
SELECT a.id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.role IN ( 'owner', 'admin' ) AND a.topic_id IN (
	SELECT a.id FROM topics a WHERE a.id = ? AND a.deleted = 0 AND a.edit_permission = 'admin' )
UNION
SELECT a.id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.topic_id IN (
	SELECT a.id FROM topics a WHERE a.id = ? AND a.deleted = 0 AND a.edit_permission = 'member' )
UNION
SELECT a.id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.topic_id IN (
	SELECT a.to_topic_id FROM auth_topics a WHERE a.deleted = 0 AND a.topic_id = ? AND a.permission = 'edit' )
UNION
SELECT a.id FROM topic_members a WHERE a.deleted = 0 AND a.account_id = ? AND a.role IN ( 'owner', 'admin' ) AND a.topic_id IN (
	SELECT a.to_topic_id FROM auth_topics a WHERE a.deleted = 0 AND a.topic_id = ? AND a.permission = 'admin_edit')
`
	ids := make([]int64, 0)

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, topicID, aid, topicID, aid, topicID, aid, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsAllowedEditMember err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.IsAllowedEditMember err(%+v)", err))
			return
		}
		ids = append(ids, targetID)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if len(ids) > 0 {
		ret = true
	}

	return
}
