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

func (p *Service) getTopicVersionsResp(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicVersionResp, err error) {
	var addCache = true

	if items, err = p.d.TopicVersionCache(c, topicID); err != nil {
		addCache = false
	} else if items != nil {
		return
	}

	if items, err = p.d.GetTopicVersions(c, node, topicID); err != nil {
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetTopicVersionCache(context.TODO(), topicID, items)
		})
	}

	return
}

func (p *Service) GetTopicVersions(c context.Context, topicSetID int64) (items []*model.TopicVersionResp, err error) {
	return p.getTopicVersionsResp(c, p.d.DB(), topicSetID)
}

func (p *Service) AddTopicVersion(c context.Context, arg *model.ArgNewTopicVersion) (id int64, err error) {
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
	if t, err = p.d.GetTopicByID(c, tx, arg.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	if m, e := p.d.GetTopicVersionByName(c, tx, arg.TopicID, arg.Name); e != nil {
		return 0, e
	} else if m != nil {
		err = ecode.TopicVersionNameExist
		return
	}

	var maxSeq int
	if maxSeq, err = p.d.GetTopicVersionMaxSeq(c, tx, t.ID); err != nil {
		return
	}

	item := &model.TopicVersion{
		ID:        gid.NewID(),
		TopicID:   t.ID,
		Name:      arg.Name,
		Seq:       maxSeq + 1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTopicVersion(c, tx, item); err != nil {
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

	var t *model.Topic
	if t, err = p.d.GetTopicByID(c, tx, arg.TopicID); err != nil {
		return
	} else if t == nil {
		err = ecode.TopicNotExist
		return
	}

	for _, v := range arg.Items {
		var ver *model.TopicVersion
		if ver, err = p.d.GetTopicVersion(c, tx, v.ID); err != nil {
			return
		} else if ver == nil {
			return ecode.TopicVersionNotExit
		}

		if m, e := p.d.GetTopicVersionByName(c, tx, t.ID, v.Name); e != nil {
			return e
		} else if m != nil && m.TopicID != ver.ID {
			err = ecode.TopicVersionNameExist
			return
		}

		ver.Name = v.Name
		ver.Seq = v.Seq

		if err = p.d.UpdateTopicVersion(c, tx, ver); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.d.DelTopicVersionCache(context.TODO(), arg.TopicID)
	})

	return
}
