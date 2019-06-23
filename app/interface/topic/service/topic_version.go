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

func (p *Service) getTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []int64, err error) {
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

func (p *Service) getTopicVersionsResp(c context.Context, node sqalx.Node, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	var data []int64
	items = make([]*model.TopicVersionResp, 0)
	if data, err = p.getTopicVersions(c, node, topicSetID); err != nil {
		return
	}

	for _, v := range data {
		t, e := p.getTopic(c, node, v)
		if e != nil {
			err = e
			return
		}

		version := &model.TopicVersionResp{
			TopicSetID:  t.TopicSetID,
			TopicID:     t.ID,
			Seq:         t.Seq,
			VersionName: t.VersionName,
			TopicName:   t.Name,
		}

		if version.TopicMeta, err = p.GetTopicMeta(c, t); err != nil {
			return
		}

		items = append(items, version)
	}

	return
}

func (p *Service) GetTopicVersions(c context.Context, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	return p.getTopicVersionsResp(c, p.d.DB(), topicSetID)
}

func (p *Service) AddTopicVersion(c context.Context, arg *model.ArgNewVersion) (id int64, err error) {
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
		err = ecode.TopicNotExist
		return
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

	if err = p.copyMembers(c, tx, aid, arg.FromTopicID, t.ID); err != nil {
		return
	}

	if err = p.copyRelations(c, tx, aid, arg.FromTopicID, t.ID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	id = t.ID

	return
}

func (p *Service) copyMembers(c context.Context, node sqalx.Node, aid int64, fromTopicID, toTopicID int64) (err error) {
	var members []*model.TopicMember
	if members, err = p.d.GetAllTopicMembers(c, node, fromTopicID); err != nil {
		return
	}

	for _, v := range members {
		v.ID = gid.NewID()
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()
		v.TopicID = toTopicID
		if v.Role == model.MemberRoleOwner && v.AccountID != aid {
			v.Role = model.MemberRoleAdmin
		}
		if v.AccountID == aid {
			v.Role = model.MemberRoleOwner
		}
		if err = p.d.AddTopicMember(c, node, v); err != nil {
			return
		}
	}

	return
}

func (p *Service) copyRelations(c context.Context, node sqalx.Node, aid int64, fromTopicID, toTopicID int64) (err error) {
	var relations []*model.TopicRelation
	if relations, err = p.d.GetAllTopicRelations(c, node, fromTopicID); err != nil {
		return
	}

	for _, v := range relations {
		v.ID = gid.NewID()
		v.CreatedAt = time.Now().Unix()
		v.UpdatedAt = time.Now().Unix()
		v.FromTopicID = fromTopicID
		if err = p.d.AddTopicRelation(c, node, v); err != nil {
			return
		}
	}
	return
}

func (p *Service) MergeTopicVersions(c context.Context, arg *model.ArgMergeVersion) (err error) {
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

	var fSets []*model.TopicVersionResp
	if fSets, err = p.getTopicVersionsResp(c, tx, arg.FromTopicSetID); err != nil {
		return
	}

	for _, v := range fSets {
		var curMember *model.TopicMember
		if curMember, err = p.d.GetTopicMemberByCondition(c, tx, v.TopicID, aid); err != nil {
			return
		} else if curMember == nil {
			err = ecode.NotBelongToTopic
			return
		} else if curMember.Role != model.MemberRoleAdmin && curMember.Role != model.MemberRoleOwner {
			err = ecode.NotTopicAdmin
			return
		}
	}

	var tSets []*model.TopicVersionResp
	if tSets, err = p.getTopicVersionsResp(c, tx, arg.ToTopicSetID); err != nil {
		return
	}
	for _, v := range tSets {
		var curMember *model.TopicMember
		if curMember, err = p.d.GetTopicMemberByCondition(c, tx, v.TopicID, aid); err != nil {
			return
		} else if curMember == nil {
			err = ecode.NotBelongToTopic
			return
		} else if curMember.Role != model.MemberRoleAdmin && curMember.Role != model.MemberRoleOwner {
			err = ecode.NotTopicAdmin
			return
		}
	}

	for _, v := range tSets {
		var t *model.Topic
		if t, err = p.d.GetTopicByID(c, tx, v.TopicID); err != nil {
			return
		} else if t == nil {
			return ecode.TopicNotExist
		}

		t.TopicSetID = arg.FromTopicSetID
		if err = p.d.UpdateTopic(c, tx, t); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicVersionCache(context.TODO(), arg.FromTopicSetID)
		p.d.DelTopicVersionCache(context.TODO(), arg.ToTopicSetID)

		for _, v := range tSets {
			p.d.DelTopicCache(context.TODO(), v.TopicID)
			p.d.DelTopicCatalogCache(context.TODO(), v.TopicID)
			p.d.DelTopicRelationCache(context.TODO(), v.TopicID)
			p.d.DelTopicMembersCache(context.TODO(), v.TopicID)
		}
	})

	return
}
