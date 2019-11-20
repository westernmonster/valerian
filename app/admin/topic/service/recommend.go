package service

import (
	"context"
	"valerian/app/admin/topic/model"
	"valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) FromTopic(v *api.TopicInfo) (item *model.TargetTopic) {
	item = &model.TargetTopic{
		ID:              v.ID,
		Name:            v.Name,
		Introduction:    v.Introduction,
		MemberCount:     (v.Stat.MemberCount),
		DiscussionCount: (v.Stat.DiscussionCount),
		ArticleCount:    (v.Stat.ArticleCount),
		Creator: &model.Creator{
			ID:           v.Creator.ID,
			Avatar:       v.Creator.Avatar,
			UserName:     v.Creator.UserName,
			Introduction: v.Creator.Introduction,
		},
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Avatar:    v.Avatar,
	}

	return
}

func (p *Service) AddRecommendTopic(c context.Context, arg *model.ArgAddRecommendTopic) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.AddRecommendTopic(c, &api.TopicReq{Aid: aid, ID: arg.TopicID}); err != nil {
		return
	}

	return
}

func (p *Service) DelRecommendTopic(c context.Context, arg *model.ArgDelete) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	if err = p.d.DelRecommendTopic(c, &api.TopicReq{Aid: aid, ID: arg.ID}); err != nil {
		return
	}

	return
}

func (p *Service) GetRecommendTopics(c context.Context) (resp *model.RecommendTopicListResp, err error) {
	resp = &model.RecommendTopicListResp{
		Items: make([]*model.TargetTopic, 0),
	}

	var idsResp *api.IDsResp
	if idsResp, err = p.d.GetRecommendTopicsIDs(c); err != nil {
		return
	}

	if idsResp.IDs == nil || len(idsResp.IDs) == 0 {
		return
	}

	for _, v := range idsResp.IDs {
		var topic *api.TopicInfo
		if topic, err = p.d.GetTopic(c, v); err != nil {
			return
		}

		item := p.FromTopic(topic)
		resp.Items = append(resp.Items, item)
	}

	return
}
