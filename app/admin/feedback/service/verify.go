package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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

func (s *Service) GetFeedbacksByCondPaged(c *mars.Context, cond map[string]interface{}, limit int, offset int) (resp *model.FeedbackListResp, err error) {

	fbs, err := s.d.GetFeedbacksByCondPaged(c, s.d.DB(), cond, limit, offset)

	resp = &model.FeedbackListResp{
		Items:  []*model.Feedback{},
		Paging: &model.Paging{},
	}
	if len(fbs) > 0 {
		for _, wc := range fbs {
			resp.Items = append(resp.Items, wc)
		}
	}
	prevUrlVal := url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}
	for k, v := range cond {
		prevUrlVal.Add(k, fmt.Sprintf("%s", v))
	}

	if resp.Paging.Prev, err = genURL("/api/v1/admin/admin/feedback/list", prevUrlVal); err != nil {
		return
	}

	nextUrlVal := url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset - limit)},
	}
	for k, v := range cond {
		nextUrlVal.Add(k, fmt.Sprintf("%s", v))
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/admin/feedback/list", nextUrlVal); err != nil {
		return
	}

	return
}
