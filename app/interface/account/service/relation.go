package service

import (
	"context"
	"net/url"
	"strconv"

	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	relation "valerian/app/service/relation/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// Follow 关注用户
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

// Unfollow 取关用户
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

// FansPaged 获取指定用户粉丝列表
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
		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.AccountID); err != nil {
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

		member.FansCount = (acc.Stat.FansCount)
		member.FollowingCount = (acc.Stat.FollowingCount)

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

// FollowPaged 获取指定用户关注列表
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
		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.AccountID); err != nil {
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

		member.FansCount = (acc.Stat.FansCount)
		member.FollowingCount = (acc.Stat.FollowingCount)
		resp.Items[i] = member
	}

	param := url.Values{}
	param.Set("account_id", strconv.FormatInt(aid, 10))
	param.Set("query", query)
	param.Set("limit", strconv.Itoa(limit))
	param.Set("offset", strconv.Itoa(offset-limit))

	if resp.Paging.Prev, err = genURL("/api/v1/account/list/followings", param); err != nil {
		return
	}
	param.Set("offset", strconv.Itoa(offset+limit))
	if resp.Paging.Next, err = genURL("/api/v1/account/list/followings", param); err != nil {
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
