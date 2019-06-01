package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getAllTopicTypesSQL = "SELECT a.* FROM topic_types a WHERE a.deleted=0 ORDER BY a.id DESC "
	_getTopicTypeSQL     = "SELECT a.* FROM topic_types a WHERE a.id=? AND a.deleted=0"
)

func (p *Dao) GetAllTopicTypes(c context.Context, node sqalx.Node) (items []*model.TopicType, err error) {
	items = make([]*model.TopicType, 0)
	if err = node.SelectContext(c, &items, _getAllTopicTypesSQL); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAllTopicTypes error(%+v)", err))

	}
	return
}

func (p *Dao) GetTopicType(c context.Context, node sqalx.Node, id int) (item *model.TopicType, err error) {
	item = new(model.TopicType)

	if err = node.GetContext(c, item, _getTopicTypeSQL, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetTopicType error(%+v) id(%d)", err, id))
	}

	return
}
