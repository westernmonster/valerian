package service

import (
	"valerian/app/admin/feedback/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func (s *Service) VerifyFeedback(c *mars.Context, feedback *model.ArgVerifyFeedback) (err error) {
	if feedback.VerifyStatus < 0 && feedback.VerifyStatus > 2 {
		err = ecode.RequestErr
		return
	}
	if err = s.d.UpdateFeedbackVerify(c, s.d.DB(), feedback.FeedbackID, feedback.VerifyStatus, feedback.VerifyDesc); err != nil {
		return err
	}
	return
}
