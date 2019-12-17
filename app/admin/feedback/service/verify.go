package service

import (
	"context"
	"time"
	"valerian/app/admin/feedback/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func (s *Service) VerifyFeedback(c context.Context, feedback *model.ArgVerifyFeedback) (err error) {
	if feedback.VerifyStatus < 0 || feedback.VerifyStatus > 2 {
		return ecode.Error(ecode.RequestErr, "状态超出范围")
	}
	if err = s.d.UpdateFeedbackVerify(c, s.d.DB(), feedback.FeedbackID, feedback.VerifyStatus, feedback.VerifyDesc); err != nil {
		return err
	}

	s.addCache(func() {
		// 1-审核通过，2 审核不通过
		if feedback.VerifyStatus == 1 {
			s.emitFeedBackAccuseSuit(c, feedback.FeedbackID, time.Now().Unix())
		} else if feedback.VerifyStatus == 2 {
			s.emitFeedBackAccuseNotSuit(c, feedback.FeedbackID, time.Now().Unix())
		}
	})

	return
}

func (s *Service) GetFeedbacksByCondPaged(c *mars.Context, cond map[string]interface{}, limit int, offset int) (items []*model.Feedback, err error) {
	return s.d.GetFeedbacksByCondPaged(c, s.d.DB(), cond, limit, offset)
}
