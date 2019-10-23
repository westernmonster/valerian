package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"valerian/app/interface/topic/model"
	relation "valerian/app/service/relation/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) GetMemberFansList(c context.Context, topicID int64, query string, pn, ps int) (resp *model.TopicMemberFansResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var idsResp *relation.IDsResp
	if idsResp, err = p.d.GetFansIDs(c, aid); err != nil {
		return
	}

	var data *model.SearchResult
	if data, err = p.d.AccountSearch(c, &model.AccountSearchParams{&model.BasicSearchParams{KW: query, Pn: pn, Ps: ps}}, idsResp.IDs); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	resp = &model.TopicMemberFansResp{
		Items:  make([]*model.FollowItem, len(data.Result)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Result {
		t := new(model.ESAccount)
		err = json.Unmarshal(v, t)
		if err != nil {
			return
		}

		member := &model.FollowItem{
			ID:           t.ID,
			Avatar:       *t.Avatar,
			UserName:     *t.UserName,
			IDCert:       *t.IDCert,
			WorkCert:     *t.WorkCert,
			IsOrg:        *t.IsOrg,
			IsVIP:        *t.IsVIP,
			Introduction: *t.Introduction,
			Gender:       *t.Gender,
		}

		var stat *model.AccountStat
		if stat, err = p.d.GetAccountStatByID(c, p.d.DB(), t.ID); err != nil {
			return
		}

		member.FansCount = int(stat.Fans)
		member.FollowingCount = int(stat.Following)

		if member.IsMember, err = p.isTopicMember(c, p.d.DB(), t.ID, topicID); err != nil {
			return
		}

		if member.Invited, err = p.hasInvited(c, p.d.DB(), t.ID, topicID); err != nil {
			return
		}
		resp.Items[i] = member
	}

	if resp.Paging.Prev, err = genURL("/api/v1/topic/list/member_fans", url.Values{
		"topic_id": []string{strconv.FormatInt(topicID, 10)},
		"query":    []string{query},
		"pn":       []string{strconv.Itoa(pn - 1)},
		"ps":       []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/topic/list/member_fans", url.Values{
		"topic_id": []string{strconv.FormatInt(topicID, 10)},
		"query":    []string{query},
		"pn":       []string{strconv.Itoa(pn + 1)},
		"ps":       []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if len(resp.Items) < ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if pn == 1 {
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
			ID:            gid.NewID(),
			TopicID:       arg.TopicID,
			AccountID:     arg.AccountID,
			FromAccountID: aid,
			Status:        model.InviteStatusSent,
			CreatedAt:     time.Now().Unix(),
			UpdatedAt:     time.Now().Unix(),
		}

		if err = p.d.AddTopicInviteRequest(c, p.d.DB(), item); err != nil {
			return
		}

		p.addCache(func() {
			p.onTopicInviteSent(c, item.ID, item.TopicID, aid, item.CreatedAt)
		})
	}

	return
}

func (p *Service) ProcessInvite(c context.Context, arg *model.ArgProcessInvite) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var req *model.TopicInviteRequest
	if req, err = p.d.GetTopicInviteRequestByID(c, tx, arg.ID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicInviteRequestNotExist
		return
	}

	if req.AccountID != aid {
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": aid, "topic_id": req.TopicID}); err != nil {
		return
	} else if member != nil {
		return
	}

	switch req.Status {
	case model.InviteStatusJoined:
	case model.InviteStatusRejected:
		return
	}

	if arg.Result {
		req.Status = model.InviteStatusJoined
		req.UpdatedAt = time.Now().Unix()

		if err = p.addMember(c, tx, req.TopicID, req.AccountID, model.MemberRoleUser); err != nil {
			return
		}
	} else {
		req.Status = model.InviteStatusRejected
		req.UpdatedAt = time.Now().Unix()
	}

	if err = p.d.UpdateTopicInviteRequest(c, tx, req); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return

}
