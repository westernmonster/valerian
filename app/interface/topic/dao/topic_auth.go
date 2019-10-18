package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/library/log"
)

// GetAll get all records
func (p *Dao) GetUserCanEditTopicIDs(c context.Context, node sqalx.Node, aid int64) (items []int64, err error) {
	items = make([]int64, 0)
	// 当前用户是管理员的话题
	// 用户所在话题角色是普通用户，该话题被授权为edit，找出授权话题
	// 用户所在话题角色是管理员，该话题被授权为admin_edit，找出授权话题
	sqlSelect := `
SELECT topic_id
FROM topic_members
WHERE account_id = ? AND role IN ( 'owner', 'admin' )
UNION
SELECT b.topic_id
FROM topic_members a JOIN auth_topics b ON a.topic_id = b.to_topic_id
WHERE a.account_id = ? AND a.role = 'user' AND b.permission = 'edit'
UNION
SELECT b.topic_id
FROM topic_members a JOIN auth_topics b ON a.topic_id = b.to_topic_id
WHERE a.account_id = ? AND a.role IN ( 'owner', 'admin' ) AND b.permission = 'admin_edit'
`

	var rows *sqlx.Rows
	if rows, err = node.QueryxContext(c, sqlSelect, aid); err != nil {
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

// GetAll get all records
func (p *Dao) GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error) {
	items = make([]*model.AuthTopic, 0)
	sqlSelect := "SELECT a.* FROM auth_topics a WHERE a.deleted=0 ORDER BY a.id DESC "

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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM auth_topics a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopicsByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

// GetByID get a record by ID
func (p *Dao) GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error) {
	item = new(model.AuthTopic)
	sqlSelect := "SELECT a.* FROM auth_topics a WHERE a.id=? AND a.deleted=0"

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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM auth_topics a WHERE a.deleted=0 %s", clause)

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
