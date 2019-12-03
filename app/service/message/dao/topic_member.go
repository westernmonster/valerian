package dao

import (
	"context"
	"fmt"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetAdminTopicMembers(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicMember, err error) {
	items = make([]*model.TopicMember, 0)

	sqlSelect := `SELECT a.id,a.topic_id,a.account_id,a.role,a.deleted,a.created_at,a.updated_at FROM topic_members a WHERE a.deleted=0 AND a.role IN('owner', 'admin') AND a.topic_id=? ORDER BY a.id DESC`

	if err = node.SelectContext(c, &items, sqlSelect, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAdminTopicMembers err(%+v), condition(%+v)", err, topicID))
		return
	}
	return
}
