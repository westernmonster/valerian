package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

func (p *Service) GetTopicMembersPaged(c context.Context, topicID int64, page, pageSize int) (resp *model.TopicMembersPagedResp, err error) {
	resp = new(model.TopicMembersPagedResp)
	resp.PageSize = pageSize
	resp.Data = make([]*model.TopicMemberResp, 0)

	count, data, err := p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, page, pageSize)
	if err != nil {
		return
	}

	for _, v := range data {
		account, e := p.getAccountByID(c, p.d.DB(), v.AccountID)
		if e != nil {
			return
		}
		resp.Data = append(resp.Data, &model.TopicMemberResp{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    account.Avatar,
			UserName:  account.UserName,
		})

	}

	resp.Count = count
	return
}

func (p *Service) getTopicMembers(c context.Context, node sqalx.Node, topicID int64, limit int) (count int, resp []*model.TopicMemberResp, err error) {
	resp = make([]*model.TopicMemberResp, 0)

	count, items, err := p.d.GetTopicMembersPaged(c, node, topicID, 1, limit)
	if err != nil {
		return
	}

	for _, v := range items {
		account, e := p.getAccountByID(c, node, v.AccountID)
		if e != nil {
			return
		}
		resp = append(resp, &model.TopicMemberResp{
			AccountID: v.AccountID,
			Role:      v.Role,
			Avatar:    account.Avatar,
			UserName:  account.UserName,
		})

	}

	return
}
