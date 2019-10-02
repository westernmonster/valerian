package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/conf/env"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *model.ArgFollow) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.Follow(c, aid, arg.AccountID); err != nil {
		return
	}

	return
}

func (p *Service) Unfollow(c context.Context, arg *model.ArgUnfollow) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.Unfollow(c, aid, arg.AccountID); err != nil {
		return
	}

	return
}

func (p *Service) FansPaged(c context.Context, aid int64, query string, limit, offset int) (resp *model.MemberResp, err error) {
	var data *relation.FansResp
	if data, err = p.d.GetFans(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.MemberResp{
		Items:  make([]*model.MemberItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		var acc *model.Account
		if acc, err = p.getAccountByID(c, v.AccountID); err != nil {
			return
		}
		member := &model.MemberItem{
			ID:           v.AccountID,
			Introduction: acc.Introduction,
			Avatar:       acc.Avatar,
			UserName:     acc.UserName,
			Gender:       acc.Gender,
			IDCert:       acc.IDCert,
			WorkCert:     acc.WorkCert,
			IsOrg:        acc.IsOrg,
			IsVIP:        acc.IsVIP,
		}

		var stat *relation.StatInfo
		if stat, err = p.d.Stat(c, v.AccountID); err != nil {
			return
		}

		member.FansCount = int(stat.Fans)
		member.FollowingCount = int(stat.Following)

		resp.Items[i] = member
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("query", query)
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/fans", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/fans", param); err != nil {
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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}

func (p *Service) FollowPaged(c context.Context, aid int64, query string, limit, offset int) (resp *model.MemberResp, err error) {
	var data *relation.FollowingResp
	if data, err = p.d.GetFollowings(c, aid, limit, offset); err != nil {
		return
	}

	resp = &model.MemberResp{
		Items:  make([]*model.MemberItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		var acc *model.Account
		if acc, err = p.getAccountByID(c, v.AccountID); err != nil {
			return
		}
		member := &model.MemberItem{
			ID:           v.AccountID,
			Introduction: acc.Introduction,
			Avatar:       acc.Avatar,
			UserName:     acc.UserName,
			Gender:       acc.Gender,
			IDCert:       acc.IDCert,
			WorkCert:     acc.WorkCert,
			IsOrg:        acc.IsOrg,
			IsVIP:        acc.IsVIP,
		}

		resp.Items[i] = member
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("query", query)
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/follow", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/follow", param); err != nil {
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
