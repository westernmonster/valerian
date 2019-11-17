package service

import (
	"context"

	"valerian/app/interface/topic/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// Leave 退出话题
func (p *Service) Leave(c context.Context, topicID int64) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.Leave(c, &topic.TopicReq{Aid: aid, ID: topicID}); err != nil {
		return
	}
	return
}

//  GetTopicMembersPaged 分页获取话题成员
func (p *Service) GetTopicMembersPaged(c context.Context, topicID int64, page, pageSize int32) (resp *model.TopicMembersPagedResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	// 检测查看权限
	if err = p.checkViewPermission(c, aid, topicID); err != nil {
		return
	}

	return p.getTopicMembersPaged(c, aid, topicID, page, pageSize)
}

func (p *Service) getTopicMembersPaged(c context.Context, aid, topicID int64, page, pageSize int32) (resp *model.TopicMembersPagedResp, err error) {
	var data *topic.TopicMembersPagedResp
	if data, err = p.d.GetTopicMembersPaged(c, &topic.ArgTopicMembers{TopicID: topicID, Page: page, PageSize: pageSize}); err != nil {
		return
	}

	resp = &model.TopicMembersPagedResp{
		Count:    data.Count,
		Data:     make([]*model.TopicMemberResp, 0),
		PageSize: data.PageSize,
	}

	if data.Data != nil {
		for _, v := range data.Data {
			resp.Data = append(resp.Data, &model.TopicMemberResp{
				AccountID: v.AccountID,
				Role:      v.Role,
				Avatar:    v.Avatar,
				UserName:  v.UserName,
			})
		}
	}

	return
}

func (p *Service) BulkSaveMembers(c context.Context, req *model.ArgBatchSavedTopicMember) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &topic.ArgBatchSavedTopicMember{
		TopicID: req.TopicID,
		Aid:     aid,
		Members: make([]*topic.ArgTopicMember, 0),
	}

	for _, v := range req.Members {
		item.Members = append(item.Members, &topic.ArgTopicMember{
			AccountID: v.AccountID,
			Opt:       v.Opt,
			Role:      v.Role,
		})
	}

	if err = p.d.BulkSaveMembers(c, item); err != nil {
		return
	}

	return
}

// ChangeOwner 更改主理人
func (p *Service) ChangeOwner(c context.Context, arg *model.ArgChangeOwner) (err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	if err = p.d.ChangeOwner(c, &topic.ArgChangeOwner{TopicID: arg.TopicID, ToAccountID: arg.ToAccountID, Aid: aid}); err != nil {
		return
	}
	return
}
