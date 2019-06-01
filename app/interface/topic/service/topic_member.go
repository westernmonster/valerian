package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/models"
	"valerian/modules/repo"

	"github.com/ztrue/tracerr"
)

func (p *Service) GetTopicMembersPaged(c context.Context, topicID int64, page, pageSize int) (resp *model.TopicMembersPagedResp, err error) {
	resp = new(model.TopicMembersPagedResp)
	resp.PageSize = pageSize
	resp.Data = make([]*model.TopicMemberResp, 0)

	var (
		needCache bool
		count     int
		items     []*model.TopicMember
	)

	if count, items, err = p.d.TopicMembersCache(c, topicID, page, pageSize); err != nil {
		needCache = true
	}

	if items == nil {
		if count, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, page, pageSize); err != nil {
			return
		}
	}

	if needCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), topicID, count, page, pageSize, items)
		})
	}

	for _, v := range items {
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

func (p *Service) getTopicMembers(c context.Context, node sqalx.Node, topicID int64, limit int) (total int, resp []*model.TopicMemberResp, err error) {
	resp = make([]*model.TopicMemberResp, 0)

	var (
		needCache bool
		count     int
		items     []*model.TopicMember
	)

	if count, items, err = p.d.TopicMembersCache(c, topicID, 1, 10); err != nil {
		needCache = true
	}

	if items == nil {
		if count, items, err = p.d.GetTopicMembersPaged(c, p.d.DB(), topicID, 1, 10); err != nil {
			return
		}
	}

	if needCache {
		p.addCache(func() {
			p.d.SetTopicMembersCache(context.TODO(), topicID, count, 1, 10, items)
		})
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

func (p *Service) bulkCreateMembers(c context.Context, node sqalx.Node, aid, topicID int64, req *model.ArgCreateTopic) (err error) {
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

	for _, v := range req.Members {
		if v.AccountID == aid {
			continue
		}

		if v.Role == models.MemberRoleOwner {
			return ecode.OnlyAllowOneOwner
		}

		member, exist, errInner := p.TopicMemberRepository.GetByCondition(c, tx, map[string]string{
			"topic_id":   strconv.FormatInt(topicID, 10),
			"account_id": strconv.FormatInt(v.AccountID, 10),
		})
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		if exist {
			member.Role = v.Role
			errInner = p.TopicMemberRepository.Update(c, tx, member)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
		} else {
			id, errInner := gid.NextID()
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}
			item := &repo.TopicMember{
				ID:        id,
				AccountID: v.AccountID,
				Role:      v.Role,
				TopicID:   topicID,
			}
			errInner = p.TopicMemberRepository.Insert(c, tx, item)
			if errInner != nil {
				err = tracerr.Wrap(errInner)
				return
			}

			break
		}
	}

	id, errInner := gid.NextID()
	if errInner != nil {
		err = tracerr.Wrap(errInner)
		return
	}
	item := &repo.TopicMember{
		ID:        id,
		AccountID: accountID,
		Role:      models.MemberRoleOwner,
		TopicID:   topicID,
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	errInner = p.TopicMemberRepository.Insert(c, tx, item)
	if errInner != nil {
		err = tracerr.Wrap(errInner)
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
}
