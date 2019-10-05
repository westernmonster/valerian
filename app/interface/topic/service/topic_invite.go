package service

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/net/metadata"
)

func (p *Service) GetMemberFansList(c context.Context, topicID int64, query string, limit, offset int) (resp *model.TopicMemberFansResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	// resp = &model.TopicMemberFansResp{
	// 	Items:  make([]*model.FollowItem, 0),
	// 	Paging: &model.Paging{},
	// }

	// if resp.Items, err = p.d.GetFansPaged(c, p.d.DB(), aid, query, limit, offset); err != nil {
	// 	return
	// }

	// param := url.Values{}
	// param.Set("topic_id", strconv.FormatInt(topicID, 10))
	// param.Set("query", query)
	// param.Set("limit", strconv.Itoa(limit))
	// param.Set("offset", strconv.Itoa(offset-limit))

	// if resp.Paging.Prev, err = genURL("/api/v1/topic/list/member_fans", param); err != nil {
	// 	return
	// }
	// param.Set("offset", strconv.Itoa(offset+limit))
	// if resp.Paging.Next, err = genURL("/api/v1/topic/list/member_fans", param); err != nil {
	// 	return
	// }

	// if len(resp.Items) < limit {
	// 	resp.Paging.IsEnd = true
	// 	resp.Paging.Next = ""
	// }

	// if offset == 0 {
	// 	resp.Paging.Prev = ""
	// }

	// for _, v := range resp.Items {
	// 	if v.IsMember, err = p.isTopicMember(c, p.d.DB(), v.ID, topicID); err != nil {
	// 		return
	// 	}

	// 	if v.Invited, err = p.hasInvited(c, p.d.DB(), v.ID, topicID); err != nil {
	// 		return
	// 	}
	// }

	var data *relation.FansResp
	return
	if data, err = p.d.GetFans(c, aid, limit, offset); err != nil {
	}

	resp = &model.TopicMemberFansResp{
		Items:  make([]*model.FollowItem, len(data.Items)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Items {
		var acc *account.BaseInfoReply
		if acc, err = p.d.GetAccountBaseInfo(c, v.AccountID); err != nil {
			return
		}
		member := &model.FollowItem{
			ID:       v.AccountID,
			Avatar:   acc.Avatar,
			UserName: acc.UserName,
			IDCert:   acc.IDCert,
			WorkCert: acc.WorkCert,
			IsOrg:    acc.IsOrg,
			IsVIP:    acc.IsVIP,
		}

		intro := acc.GetIntroductionValue()
		member.Introduction = &intro

		if acc.Gender != nil {
			gender := int(acc.GetGenderValue())
			member.Gender = &gender
		}

		// TODO: 返回成员信息
		// var stat *relation.StatInfo
		// if stat, err = p.d.Stat(c, v.AccountID, v.AccountID); err != nil {
		// 	return
		// }

		// member.FansCount = int(stat.Fans)
		// member.FollowingCount = int(stat.Following)

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

func (p *Service) hasInvited(c context.Context, node sqalx.Node, accountID, topicID int64) (hasInvited bool, err error) {
	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByCond(c, p.d.DB(), map[string]interface{}{
		"topic_id":   topicID,
		"account_id": accountID,
	}); err != nil {
		return
	} else if req != nil {
		hasInvited = true
	}
	return
}

func (p *Service) Invite(c context.Context, arg *model.ArgTopicInvite) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if aid == arg.AccountID {
		err = ecode.InviteSelfNotAllowed
		return
	}

	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByCond(c, p.d.DB(), map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.AccountID,
	}); err != nil {
		return
	} else if req != nil {
		return
	}

	if req == nil {
		item := &model.TopicInviteRequest{
			ID:        gid.NewID(),
			TopicID:   arg.TopicID,
			AccountID: arg.AccountID,
			Status:    model.InviteStatusSent,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddTopicInviteRequest(c, p.d.DB(), item); err != nil {
			return
		}

		p.addCache(func() {
			p.onTopicInviteSent(c, item.ID)
		})
	}

	return
}
