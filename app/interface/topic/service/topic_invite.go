package service

import (
	"context"

	"valerian/app/interface/topic/model"
)

func (p *Service) GetMemberFansList(c context.Context, topicID int64, query string, pn, ps int) (resp *model.TopicMemberFansResp, err error) {
	return
}

func (p *Service) Invite(c context.Context, arg *model.ArgTopicInvite) (err error) {
	return
}

func (p *Service) ProcessInvite(c context.Context, arg *model.ArgProcessInvite) (err error) {
	return
}
