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
)

func (p *Service) getTopicVersionsResp(c context.Context, node sqalx.Node, topicSetID int64) (items []*model.TopicVersionResp, err error) {
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
	return p.getTopicVersionsResp(c, p.d.DB(), topicSetID)
}

func (p *Service) AddTopicVersion(c context.Context, arg *model.ArgNewTopicVersion) (id int64, err error) {
	// aid, ok := metadata.Value(c, metadata.Aid).(int64)
	// if !ok {
	// 	err = ecode.AcquireAccountIDFailed
	// 	return
	// }
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

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	id = t.ID

	p.addCache(func() {
		p.d.DelTopicVersionCache(context.TODO(), t.ID)
	})

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

func (p *Service) SaveTopicVersions(c context.Context, arg *model.ArgSaveTopicVersions) (err error) {
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

	// for _, v := range arg.Items {
	// var t *model.TopicVersion
	// if t, err = p.d.GetTopicVersionByID(c, tx, v.TopicID); err != nil {
	// 	return
	// } else if t == nil {
	// 	return ecode.TopicNotExist
	// }

	// if m, e := p.d.GetTopicVersionByName(c, tx, arg.TopicSetID, v.VersionName); e != nil {
	// 	return e
	// } else if m != nil && m.TopicID != t.ID {
	// 	err = ecode.TopicVersionNameExist
	// 	return
	// }

	// t.Name = v.VersionName
	// t.Seq = v.Seq

	// if err = p.d.UpdateTopic(c, tx, t); err != nil {
	// 	return
	// }
	// }

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicVersionCache(context.TODO(), arg.TopicSetID)
		for _, v := range arg.Items {
			p.d.DelTopicCache(context.TODO(), v.TopicID)
		}
	})

	return
}
