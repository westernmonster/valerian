package service

import (
	"fmt"
	"net/url"
	"strconv"
	"valerian/app/admin/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/metadata"
)

func (s *Service) WorkCert(c *mars.Context, arg *model.ArgWorkCert) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if _, err := s.d.SetWorkCert(c, arg, aid); err != nil {
		return err
	}
	return
}

func (s *Service) GetWorkCertsByCondPaged(c *mars.Context, cond map[string]interface{}, limit, offset int) (resp *model.WorkCertListResp, err error) {
	workCerts, err := s.d.GetWorkCertificationsByCond(c, s.d.DB(), cond, limit, offset)

	resp = &model.WorkCertListResp{
		Items:  []*model.WorkCertification{},
		Paging: &model.Paging{},
	}

	if len(workCerts) > 0 {
		for _, wc := range workCerts {
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

	if resp.Paging.Prev, err = genURL("/api/v1/admin/account/workcert/list", prevUrlVal); err != nil {
		return
	}

	nextUrlVal := url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
	}
	for k, v := range cond {
		nextUrlVal.Add(k, fmt.Sprintf("%s", v))
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/account/workcert/list", nextUrlVal); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}
	return
}

func (s *Service) GetWorkCertHistorysByAccount(c *mars.Context, cond map[string]interface{}, limit, offset int) (resp *model.WorkCertHistoryResp, err error) {
	workCerts, err := s.d.GetWorkCertHistorysByAccount(c, s.d.DB(), cond, limit, offset)

	resp = &model.WorkCertHistoryResp{
		Items:  []*model.WorkCertHistory{},
		Paging: &model.Paging{},
	}

	if len(workCerts) > 0 {
		for _, wc := range workCerts {
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

	if resp.Paging.Prev, err = genURL("/api/v1/admin/account/workcert/list", prevUrlVal); err != nil {
		return
	}

	nextUrlVal := url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset + limit)},
	}
	for k, v := range cond {
		nextUrlVal.Add(k, fmt.Sprintf("%s", v))
	}

	if resp.Paging.Next, err = genURL("/api/v1/admin/account/workcert/list", nextUrlVal); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}
	return

	return
}
