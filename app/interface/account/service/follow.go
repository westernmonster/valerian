package service

import (
	"context"
	"net/url"
	"strconv"
	"time"
	"valerian/app/interface/account/model"
	"valerian/library/conf/env"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *model.ArgFollow) (isFollow bool, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var item *model.AccountFollower
	if item, err = p.d.GetAccountFollowerByCond(c, p.d.DB(), map[string]interface{}{"account_id": arg.AccountID, "follower_id": aid}); err != nil {
		return
	} else if item == nil {
		isFollow = true
		af := &model.AccountFollower{
			ID:         gid.NewID(),
			AccountID:  arg.AccountID,
			FollowerID: aid,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}
		if err = p.d.AddAccountFollower(c, p.d.DB(), af); err != nil {
			return
		}
	} else {
		isFollow = false
		if err = p.d.DelAccountFollower(c, p.d.DB(), item.ID); err != nil {
			return
		}
	}

	return
}

func (p *Service) FansPaged(c context.Context, aid int64, query string, limit, offset int) (resp *model.MemberResp, err error) {
	resp = &model.MemberResp{
		Items:  make([]*model.MemberItem, 0),
		Paging: &model.Paging{},
	}

	if resp.Items, err = p.d.GetFansPaged(c, p.d.DB(), aid, query, limit, offset); err != nil {
		return
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
