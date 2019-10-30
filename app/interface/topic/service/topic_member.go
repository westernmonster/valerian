package service

import (
	"context"

	"valerian/app/interface/topic/model"
)

// Leave 退出话题
func (p *Service) Leave(c context.Context, topicID int64) (err error) {
	return
}

//  GetTopicMembersPaged 分页获取话题成员
func (p *Service) GetTopicMembersPaged(c context.Context, topicID int64, page, pageSize int) (resp *model.TopicMembersPagedResp, err error) {
	return
}

func (p *Service) BulkSaveMembers(c context.Context, req *model.ArgBatchSavedTopicMember) (err error) {
	return
}

// ChangeOwner 更改主理人
func (p *Service) ChangeOwner(c context.Context, arg *model.ArgChangeOwner) (err error) {
	return
}
