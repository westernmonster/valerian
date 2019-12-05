package dao

import (
	"context"
	"fmt"
	"valerian/app/admin/feedback/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Update update a exist record
func (p *Dao) UpdateFeedbackVerify(c context.Context, node sqalx.Node, item *model.Feedback) (err error) {
	sqlUpdate := "UPDATE feedbacks SET verify_status=?,verify_desc=? WHERE id=?"

	_, err = node.ExecContext(c, sqlUpdate, item.VerifyStatus, item.VerifyDesc, item.ID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeedbacks err(%+v), item(%+v)", err, item))
		return
	}

	return
}