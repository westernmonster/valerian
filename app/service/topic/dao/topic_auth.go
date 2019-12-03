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

// GetAll get all records
func (p *Dao) GetUserCanEditTopicIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	items = make([]int64, 0)
	// 1. 编辑权限为管理员，且用户是管理员
	// 2. 编辑权限为成员，且用户为成员
	// 3. 授权话题，授权为可编辑，用户是授权话题成员
	// 4. 授权话题，授权为管理员可编辑，用户是授权话题管理员
	sqlSelect := `
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
`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid, aid, aid, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditTopicIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditTopicIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

func (p *Dao) GetAuthTopicIDs(c context.Context, node sqalx.Node, topicIDs []int64) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT DISTINCT a.to_topic_id FROM auth_topics a WHERE a.deleted=0 AND a.topic_id IN(?) limit 20"
	query, args, err := sqlx.In(sqlSelect, topicIDs)
	query = node.Rebind(query)

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, args...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicIDs err(%+v)", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicIDs err(%+v)", err))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}

// GetAll get all records
func (p *Dao) GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error) {
	items = make([]*model.AuthTopic, 0)
	sqlSelect := "SELECT a.id,a.topic_id,a.to_topic_id,a.permission,a.deleted,a.created_at,a.updated_at  FROM auth_topics a WHERE a.deleted=0 ORDER BY a.id DESC "

	if err = node.SelectContext(c, &items, sqlSelect); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopics err(%+v)", err))
		return
	}
	return
}

// GetAllByCondition get records by condition
func (p *Dao) GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error) {
	items = make([]*model.AuthTopic, 0)
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
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.topic_id,a.to_topic_id,a.permission,a.deleted,a.created_at,a.updated_at FROM auth_topics a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error) {
	item = new(model.AuthTopic)
	sqlSelect := "SELECT a.id,a.topic_id,a.to_topic_id,a.permission,a.deleted,a.created_at,a.updated_at FROM auth_topics a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicByID err(%+v), id(%+v)", err, id))
	}

	return
}

// GetByCondition get a record by condition
func (p *Dao) GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error) {
	item = new(model.AuthTopic)
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
	if val, ok := cond["to_topic_id"]; ok {
		clause += " AND a.to_topic_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["permission"]; ok {
		clause += " AND a.permission =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.topic_id,a.to_topic_id,a.permission,a.deleted,a.created_at,a.updated_at FROM auth_topics a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}

// Insert insert a new record
func (p *Dao) AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	sqlInsert := "INSERT INTO auth_topics( id,topic_id,to_topic_id,permission,deleted,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?)"

	if _, err = node.ExecContext(c, sqlInsert, item.ID, item.TopicID, item.ToTopicID, item.Permission, item.Deleted, item.CreatedAt, item.UpdatedAt); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAuthTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Update update a exist record
func (p *Dao) UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	sqlUpdate := "UPDATE auth_topics SET topic_id=?,to_topic_id=?,permission=?,updated_at=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.TopicID, item.ToTopicID, item.Permission, item.UpdatedAt, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateAuthTopics err(%+v), item(%+v)", err, item))
		return
	}

	return
}

// Delete logic delete a exist record
func (p *Dao) DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	sqlDelete := "UPDATE auth_topics SET deleted=1 WHERE id=? "

	if _, err = node.ExecContext(c, sqlDelete, id); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAuthTopics err(%+v), item(%+v)", err, id))
		return
	}

	return
}

func (p *Dao) GetAuthed2CurrentTopicIDsPaged(c context.Context, node sqalx.Node, topicID int64, limit, offset int32) (items []int64, err error) {
	items = make([]int64, 0)
	sqlSelect := "SELECT a.topic_id FROM auth_topics a WHERE a.deleted=0 AND a.to_topic_id=? ORDER BY a.topic_id DESC LIMIT ?,?"

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, topicID, offset, limit); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthed2CurrentTopicIDsPaged err(%+v) to_topic_id(%d), limit(%d), offset(%d)", err, topicID, limit, offset))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			targetID int64
		)
		if err = rows.Scan(&targetID); err != nil {
			log.For(c).Error(fmt.Sprintf("dao.GetAuthed2CurrentTopicIDsPaged err(%+v) to_topic_id(%d), limit(%d), offset(%d)", err, topicID, limit, offset))
			return
		}
		items = append(items, targetID)
	}

	err = rows.Err()
	return
}
