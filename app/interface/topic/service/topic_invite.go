package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"valerian/app/interface/topic/model"
	relation "valerian/app/service/relation/api"
	search "valerian/app/service/search/api"
	"valerian/library/ecode"
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

	var data *search.SearchResult
	if data, err = p.d.SearchAccount(c, &search.SearchParam{KW: query, Pn: int32(pn), Ps: int32(ps), IDs: idsResp.IDs}); err != nil {
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

		// var stat *account.AccountStatInfo
		// if stat, err = p.d.GetAccountStat(c, t.ID); err != nil {
		// 	return
		// }

		// member.FansCount = (stat.FansCount)
		// member.FollowingCount = (stat.FollowingCount)

		// if member.IsMember, err = p.isTopicMember(c, p.d.DB(), t.ID, topicID); err != nil {
		// 	return
		// }

		// if member.Invited, err = p.hasInvited(c, p.d.DB(), t.ID, topicID); err != nil {
		// 	return
		// }
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

func (p *Service) Invite(c context.Context, arg *model.ArgTopicInvite) (err error) {
	// if err = p.d.Invite(c, &topic.ArgTopicInvite{}); err != nil {
	// 	return
	// }
	return
}

func (p *Service) ProcessInvite(c context.Context, arg *model.ArgProcessInvite) (err error) {
	return
}
