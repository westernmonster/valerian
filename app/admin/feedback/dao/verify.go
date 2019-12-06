package dao

import (
	"context"
	"fmt"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Update update a exist record
func (p *Dao) UpdateFeedbackVerify(c context.Context, node sqalx.Node, feedbackID int64, verifyStatus int32, verifyDesc string) (err error) {
	sqlUpdate := "UPDATE feedbacks SET verify_status=?,verify_desc=? WHERE id=?"
	_, err = node.ExecContext(c, sqlUpdate, verifyStatus, verifyDesc, feedbackID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateFeedbackVerify err(%+v)", err))
		return
	}
	return
}
