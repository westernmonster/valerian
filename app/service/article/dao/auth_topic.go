package dao

import (
	"context"
	"fmt"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

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
