package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

const (
	_getTopicMembersCountSQL = "SELECT COUNT(1) as count FROM topic_members a WHERE a.deleted=0 AND a.topic_id=?"
	_getTopicMembersPagedSQL = "SELECT a.id,a.topic_id,a.account_id,a.role,a.deleted,a.created_at,a.updated_at FROM topic_members a WHERE a.deleted=0 AND a.topic_id=? ORDER BY a.role,a.id DESC limit ?,?"
)

func (p *Dao) GetManageTopicIDsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int32) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.topic_id FROM topic_members a WHERE a.deleted=0 AND a.account_id=? AND a.role IN ('owner', 'admin') ORDER BY a.topic_id DESC LIMIT ?,?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetManageTopicIDsPaged err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetManageTopicIDsPaged err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) GetFollowedTopicIDsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int32) (items []int64, err error) {
	items = make([]int64, 0)

	sqlSelect := "SELECT a.topic_id FROM topic_members a WHERE a.deleted=0 AND a.account_id=? ORDER BY a.topic_id DESC LIMIT ?,?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicIDsPaged err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicIDsPaged err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) GetFollowedTopicsPaged(c context.Context, node sqalx.Node, aid int64, query string, limit, offset int) (items []*model.Topic, err error) {
	items = make([]*model.Topic, 0)
	sqlSelect := `SELECT b.id,b.name,b.avatar,b.bg,b.introduction,b.allow_discuss,b.allow_chat,b.is_private,b.view_permission,b.edit_permission,b.join_permission,b.catalog_view_type,b.topic_home,b.created_by,b.deleted,b.created_at,b.updated_at
	FROM topic_members a LEFT JOIN topics b ON a.topic_id=b.id
	WHERE a.deleted=0 AND a.account_id=?
	ORDER BY b.id DESC LIMIT ?,?`

	if err = node.SelectContext(c, &items, sqlSelect, aid, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicsPaged err(%+v) aid(%d) limit(%d) offset(%d)", err, aid, limit, offset))
		return
	}
	return
}

func (p *Dao) IsAllowedEditMember(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
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

func (p *Dao) IsAllowedViewMember(c context.Context, node sqalx.Node, aid int64, topicID int64) (ret bool, err error) {
	// 逻辑
	// 1. 用户为当前话题成员
	// 2. 用户为授权话题成员
	sqlSelect := `
   SELECT * FROM topic_members a WHERE a.deleted= 0 AND a.account_id= ?
   AND a.topic_id IN (SELECT a.to_topic_id FROM auth_topics a WHERE a.deleted= 0 AND a.topic_id= ? UNION SELECT ? AS to_topic_id)`

	item := new(model.TopicMember)
	if err = node.GetContext(c, item, sqlSelect, aid, topicID, topicID); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			ret = false
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.IsAllowedViewMember err(%+v), aid(%+v) topic_id(%d)", err, aid, topicID))
		return
	}

	ret = true

	return
}

func (p *Dao) GetFollowedTopicsIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.topic_id FROM topic_members a  WHERE a.deleted=0 AND a.account_id=?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicsIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicsIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) GetSpecificTopicsMemberIDs(c context.Context, node sqalx.Node, topicIDs []int64) (items []int64, err error) {
	items = make([]int64, 0)

	sqlSelect := "SELECT DISTINCT a.account_id FROM topic_members a WHERE a.deleted=0 AND a.topic_id IN(?) ORDER BY a.account_id LIMIT 20"
	query, args, err := sqlx.In(sqlSelect, topicIDs)
	query = node.Rebind(query)

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, query, args...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetSpecificTopicsMemberIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetSpecificTopicsMemberIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
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

// GetTopicMembersCount
func (p *Dao) GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error) {
	if err = node.GetContext(c, &count, _getTopicMembersCountSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersCount error(%+v), topic id(%d)", err, topicID))
	}
	return
}

// GetTopicMembersPaged
func (p *Dao) GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int32) (count int32, items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	offset := (page - 1) * pageSize

	if err = node.GetContext(c, &count, _getTopicMembersCountSQL, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged error(%+v), topic id(%d)", err, topicID))
	}

	if err = node.SelectContext(c, &items, _getTopicMembersPagedSQL, topicID, offset, pageSize); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged error(%+v), topic id(%d) page(%d) pageSize(%d)", err, topicID, page, pageSize))
	}
	return
}

// GetAll get all records
func (p *Dao) GetTopicMembers(c context.Context, node sqalx.Node) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembers err(%+v)", err))
		return
	}
	return
}

func (p *Dao) GetTopicAdminMembers(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.deleted=0 AND a.topic_id=? AND a.role IN('owner', 'admin') ORDER BY a.id"

	if err = node.SelectContext(c, &items, sqlSelect, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicAdminMembers err(%+v), topic_id(%+v)", err, topicID))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetTopicMembersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_members a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetTopicMemberByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicMember, err error) {
	item = new(model.TopicMember)
	sqlSelect := "SELECT a.* FROM topic_members a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMemberByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error) {
	item = new(model.TopicMember)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["topic_id"]; ok {
		clause += " AND a.topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["role"]; ok {
		clause += " AND a.role =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM topic_members a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	sqlInsert := "INSERT INTO topic_members( id,topic_id,account_id,role,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.AccountID, item.Role, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddTopicMembers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	sqlUpdate := "UPDATE topic_members SET topic_id=?,account_id=?,role=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.AccountID, item.Role, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopicMembers err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelTopicMember(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE topic_members SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopicMembers err(%+v), item(%+v)", err, id))
		return
	}

	return
}
