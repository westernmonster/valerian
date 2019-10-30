package service

import (
	"context"

	"valerian/app/interface/topic/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) Follow(c context.Context, arg *model.ArgTopicFollow) (status int32, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var resp *topic.StatusResp
	if resp, err = p.d.FollowTopic(c, &topic.ArgTopicFollow{TopicID: arg.TopicID, AllowViewCert: arg.AllowViewCert, Reason: arg.Reason, Aid: aid}); err != nil {
		return
	}

	status = resp.Status
	return
}

func (p *Service) FollowedTopics(c context.Context, query string, pn, ps int) (resp *model.JoinedTopicsResp, err error) {
	return
}

func (p *Service) AuditFollow(c context.Context, arg *model.ArgAuditFollow) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	item := &topic.ArgAuditFollow{
		Aid:     aid,
		ID:      arg.ID,
		Reason:  arg.Reason,
		Approve: arg.Approve,
	}
	if err = p.d.AuditFollow(c, item); err != nil {
		return
	}

	return
}
