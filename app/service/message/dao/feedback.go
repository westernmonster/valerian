package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/message/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// GetByID get a record by ID
func (p *Dao) GetFeedbackByID(c context.Context, node sqalx.Node, id int64) (item *model.Feedback, err error) {
	item = new(model.Feedback)
	sqlSelect := "SELECT a.id,a.target_id,a.target_type,a.feedback_type,a.feedback_desc,a.created_by,a.deleted,a.created_at,a.updated_at FROM feedbacks a WHERE a.id=? AND a.deleted=0"

	if err = node.GetContext(c, item, sqlSelect, id); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("dao.GetFeedbackByID err(%+v), id(%+v)", err, id))
	}

	return
}