package dao

import (
	"context"
	"fmt"
	"valerian/app/service/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetTopicDiscussCategories(c context.Context, node sqalx.Node, topicID int64) (items []*model.DiscussCategory, err error) {
	items = make([]*model.DiscussCategory, 0)
	sqlSelect := "SELECT a.id,a.topic_id,a.seq,a.name,a.deleted,a.created_at,a.updated_at FROM discuss_categories a WHERE a.deleted=0 AND a.topic_id=? ORDER BY a.id "

	if err = node.SelectContext(c, &items, sqlSelect, topicID); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicDiscussCategories err(%+v) topic_id(%d)", err, topicID))
		return
	}
	return
}
