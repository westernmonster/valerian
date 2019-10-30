package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
	"valerian/app/interface/topic/model"
	account "valerian/app/service/account/api"
	"valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *api.ArgTopicFollow) (status int, err error) {
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
	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": arg.Aid, "topic_id": arg.TopicID}); err != nil {
		return
	} else if member != nil {
		return model.FollowStatusFollowed, nil
	}

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByCond(c, tx, map[string]interface{}{
		"topic_id":   arg.TopicID,
		"account_id": arg.Aid,
	}); err != nil {
		return
	}

	if req != nil {
		switch req.Status {
		case model.FollowRequestStatusCommited:
			return model.FollowStatusApproving, nil
		case model.FollowRequestStatusApproved:
			return model.FollowStatusFollowed, nil
		case model.FollowRequestStatusRejected:
			break
		}
	}

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, arg.TopicID); err != nil {
		return
	} else if t == nil {
		return 0, ecode.TopicNotExist
	}

	var account *account.BaseInfoReply
	if account, err = p.d.GetAccountBaseInfo(c, arg.Aid); err != nil {
		return
	} else if account == nil {
		return 0, ecode.UserNotExist
	}

	item := &model.TopicFollowRequest{
		ID:        gid.NewID(),
		Status:    model.FollowRequestStatusCommited,
		TopicID:   arg.TopicID,
		Reason:    arg.Reason,
		AccountID: arg.Aid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	switch t.JoinPermission {
	case model.JoinPermissionMember:
		status = model.FollowStatusFollowed
		if err = p.addMember(c, tx, arg.TopicID, aid, model.MemberRoleUser); err != nil {
			return
		}
		break
	case model.JoinPermissionMemberApprove:
		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, tx, item); err != nil {
			return
		}
		break
	case model.JoinPermissionCertApprove:
		if !account.IDCert || !account.WorkCert {
			return model.FollowStatusUnfollowed, ecode.NeedWorkCert
		}

		status = model.FollowStatusApproving
		if err = p.d.AddTopicFollowRequest(c, tx, item); err != nil {
			return
		}
		break
	case model.JoinPermissionManualAdd:
		status = model.FollowStatusUnfollowed
		err = ecode.OnlyAllowAdminAdded
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return

}

func (p *Service) FollowedTopics(c context.Context, query string, pn, ps int) (resp *model.JoinedTopicsResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var ids []int64
	if ids, err = p.d.GetFollowedTopicsIDs(c, p.d.DB(), aid); err != nil {
		return
	}

	var data *model.SearchResult
	if data, err = p.d.TopicSearch(c, &model.TopicSearchParams{&model.BasicSearchParams{KW: query, Pn: pn, Ps: ps}}, ids); err != nil {
		err = ecode.SearchTopicFailed
		return
	}

	resp = &model.JoinedTopicsResp{
		Items:  make([]*model.JoinedTopicItem, len(data.Result)),
		Paging: &model.Paging{},
	}

	for i, v := range data.Result {
		t := new(model.ESTopic)
		err = json.Unmarshal(v, t)
		if err != nil {
			return
		}

		item := &model.JoinedTopicItem{
			ID:             t.ID,
			Name:           *t.Name,
			Introduction:   *t.Introduction,
			EditPermission: *t.EditPermission,
			Avatar:         *t.Avatar,
		}

		var stat *model.TopicStat
		if stat, err = p.GetTopicStat(c, t.ID); err != nil {
			return
		}

		item.MemberCount = stat.MemberCount
		item.ArticleCount = stat.ArticleCount
		item.DiscussionCount = stat.DiscussionCount

		resp.Items[i] = item

	}

	if resp.Paging.Prev, err = genURL("/api/v1/topic/list/followed", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn - 1)},
		"ps":    []string{strconv.Itoa(ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/topic/list/followed", url.Values{
		"query": []string{query},
		"pn":    []string{strconv.Itoa(pn + 1)},
		"ps":    []string{strconv.Itoa(ps)},
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

func (p *Service) GetTopicStat(c context.Context, topicID int64) (stat *model.TopicStat, err error) {
	if stat, err = p.d.GetTopicStatByID(c, p.d.DB(), topicID); err != nil {
		return
	} else if stat == nil {
		stat = &model.TopicStat{
			TopicID: topicID,
		}
	}
	return
}

func (p *Service) AuditFollow(c context.Context, arg *model.ArgAuditFollow) (err error) {
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

	var req *model.TopicFollowRequest
	if req, err = p.d.GetTopicFollowRequestByID(c, tx, arg.ID); err != nil {
		return
	} else if req == nil {
		err = ecode.TopicFollowRequestNotExist
		return
	}

	var member *model.TopicMember
	if member, err = p.d.GetTopicMemberByCond(c, tx, map[string]interface{}{"account_id": req.AccountID, "topic_id": req.TopicID}); err != nil {
		return
	} else if member != nil {
		return
	}

	switch req.Status {
	case model.FollowRequestStatusApproved:
	case model.FollowRequestStatusRejected:
		return
	}

	if arg.Approve {
		req.Status = model.FollowRequestStatusApproved
		req.UpdatedAt = time.Now().Unix()

		if err = p.addMember(c, tx, req.TopicID, req.AccountID, model.MemberRoleUser); err != nil {
			return
		}
	} else {
		req.Status = model.FollowRequestStatusRejected
		req.UpdatedAt = time.Now().Unix()
	}

	if err = p.d.UpdateTopicFollowRequest(c, tx, req); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		switch req.Status {
		case model.FollowRequestStatusApproved:
			p.onTopicFollowApproved(c, req.ID, req.TopicID, aid, time.Now().Unix())
			break
		case model.FollowRequestStatusRejected:
			p.onTopicFollowRejected(c, req.ID, req.TopicID, aid, time.Now().Unix())
			break
		}
	})

	return

}
