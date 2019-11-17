package service

import (
	"context"

	"valerian/app/interface/init/model"
	account "valerian/app/service/account/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
)

func (p *Service) FromTopic(v *topic.TopicInfo) (item *model.TargetTopic) {
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

func (p *Service) GetRecommendTopics(c context.Context) (resp *model.MajorListResp, err error) {
	resp = &model.MajorListResp{
		Items: make([]*model.TargetTopic, 0),
	}

	var idsResp *topic.IDsResp
	if idsResp, err = p.d.GetRecommendTopicsIDs(c); err != nil {
		return
	}

	if idsResp.IDs == nil {
		return
	}

	for _, v := range idsResp.IDs {
		var topic *topic.TopicInfo
		if topic, err = p.d.GetTopic(c, v); err != nil {
			if ecode.IsNotExistEcode(err) {
				continue
			}
			return
		}

		resp.Items = append(resp.Items, p.FromTopic(topic))
	}

	return
}

func (p *Service) GetRecommendAuthTopics(c context.Context, ids []int64) (resp *model.RelatedListResp, err error) {
	resp = &model.RelatedListResp{
		Items: make([]*model.TargetTopic, 0),
	}

	var idsResp *topic.IDsResp
	if idsResp, err = p.d.GetRecommendAuthTopicsIDs(c, &topic.IDsReq{IDs: ids}); err != nil {
		return
	}

	if idsResp.IDs == nil {
		return
	}

	for _, v := range idsResp.IDs {
		var topic *topic.TopicInfo
		if topic, err = p.d.GetTopic(c, v); err != nil {
			if ecode.IsNotExistEcode(err) {
				continue
			}
			return
		}

		resp.Items = append(resp.Items, p.FromTopic(topic))
	}

	return
}

func (p *Service) GetRecommendMembers(c context.Context, majorIDs []int64, relatedIDs []int64) (resp *model.MemberListResp, err error) {
	resp = &model.MemberListResp{
		Items: make([]*model.MemberInfo, 0),
	}

	dic := make(map[int64]bool)
	for _, v := range majorIDs {
		dic[v] = true
	}
	for _, v := range relatedIDs {
		dic[v] = true
	}

	ids := make([]int64, 0)
	for k, _ := range dic {
		ids = append(ids, k)
	}

	var idsResp *topic.IDsResp
	if idsResp, err = p.d.GetRecommendMemberIDs(c, &topic.IDsReq{IDs: ids}); err != nil {
		return
	}

	if idsResp.IDs == nil {
		return
	}

	for _, v := range idsResp.IDs {
		var info *model.MemberInfo
		if info, err = p.GetMemberInfo(c, v); err != nil {
			if ecode.IsNotExistEcode(err) {
				continue
			}
			return
		}

		resp.Items = append(resp.Items, info)
	}

	return
}

func (p *Service) GetMemberInfo(c context.Context, targetID int64) (resp *model.MemberInfo, err error) {
	var f *account.MemberInfoReply
	if f, err = p.d.GetMemberInfo(c, targetID); err != nil {
		return
	}

	resp = &model.MemberInfo{
		ID:             f.ID,
		UserName:       f.UserName,
		Gender:         (f.Gender),
		Location:       f.Location,
		LocationString: f.LocationString,
		Introduction:   f.Introduction,
		Avatar:         f.Avatar,
		IDCert:         f.IDCert,
		WorkCert:       f.WorkCert,
		IsOrg:          f.IsOrg,
		IsVIP:          f.IsVIP,
		Company:        f.Company,
		Position:       f.Position,
	}

	resp.Stat = &model.MemberInfoStat{
		FansCount:       (f.Stat.FansCount),
		FollowingCount:  (f.Stat.FollowingCount),
		TopicCount:      (f.Stat.TopicCount),
		ArticleCount:    (f.Stat.ArticleCount),
		DiscussionCount: (f.Stat.DiscussionCount),
	}
	return
}
