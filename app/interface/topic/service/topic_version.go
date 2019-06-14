package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) getTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	var addCache = true

	if items, err = p.d.TopicVersionCache(c, topicSetID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetTopicVersions(c, node, topicSetID); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicVersionCache(context.TODO(), topicSetID, items)
		})
	}

	return
}

func (p *Service) GetTopicVersions(c context.Context, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	return p.getTopicVersions(c, p.d.DB(), topicSetID)
}

func (p *Service) addTopicVersion(c context.Context, arg *model.ArgNewVersion) (err error) {
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

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, arg.FromTopicID); err != nil {
		return
	} else if t == nil {
		return ecode.TopicNotExist
	}

	var curMember *model.TopicMember
	if curMember, err = p.d.GetTopicMemberByCondition(c, tx, arg.FromTopicID, aid); err != nil {
		return
	} else if curMember == nil {
		err = ecode.NotBelongToTopic
		return
	} else if curMember.Role != model.MemberRoleAdmin && curMember.Role != model.MemberRoleOwner {
		err = ecode.NotTopicAdmin
		return
	}

	if v, e := p.d.GetTopicVersionByName(c, tx, t.TopicSetID, arg.VersionName); e != nil {
		err = e
		return
	} else if v != nil {
		err = ecode.TopicVersionNameExist
		return
	}

	t.ID = gid.NewID()
	t.VersionName = arg.VersionName
	t.CreatedBy = aid
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = time.Now().Unix()

	if err = p.d.AddTopic(c, tx, t); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
