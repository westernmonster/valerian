package service

import (
	"context"

	"valerian/app/interface/topic/model"
)

func (p *Service) Follow(c context.Context, arg *model.ArgTopicFollow) (status int, err error) {
	return
}

func (p *Service) FollowedTopics(c context.Context, query string, pn, ps int) (resp *model.JoinedTopicsResp, err error) {
	return
}

func (p *Service) GetTopicStat(c context.Context, topicID int64) (stat *model.TopicStat, err error) {
	return
}

func (p *Service) AuditFollow(c context.Context, arg *model.ArgAuditFollow) (err error) {
	return
}
