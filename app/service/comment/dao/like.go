package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/service/comment/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

func (p *Dao) GetLikeByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Like, err error) {
	item = new(model.Like)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_id"]; ok {
		clause += " AND a.target_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["target_type"]; ok {
		clause += " AND a.target_type =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.account_id,a.target_id,a.target_type,a.deleted,a.created_at,a.updated_at FROM likes a WHERE a.deleted=0 %s", clause)

	if err = node.GetContext(c, item, sqlSelect, condition...); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetLikesByCond err(%+v), condition(%+v)", err, cond))
		return
	}

	return
}
