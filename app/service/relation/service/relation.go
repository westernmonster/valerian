package service

import (
	"context"
	"fmt"
	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Unfollow 取关
func (p *Service) Unfollow(c context.Context, aid int64, fid int64) (err error) {
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

	var followingAttr uint32
	var fansAttr uint32
	var friend bool

	// 如果没有关注对方，则设置为AttrNoRelation
	if followingItem, e := p.d.GetAccountRelationByCond(c, tx, map[string]interface{}{
		"account_id":   aid,
		"following_id": fid,
	}); e != nil {
		err = e
		return
	} else if followingItem == nil {
		followingAttr = model.AttrNoRelation
	} else {
		followingAttr = followingItem.Attribute
	}

	// 如果对方没有关注你，则设置为AttrNoRelation
	if fansItem, e := p.d.GetAccountRelationByCond(c, tx, map[string]interface{}{
		"account_id":   fid,
		"following_id": aid,
	}); e != nil {
		err = e
		return
	} else if fansItem == nil {
		fansAttr = model.AttrNoRelation
	} else {
		fansAttr = fansItem.Attribute
	}

	var newFollowingAttr = model.AttrFollowing
	var newFansAttr = fansAttr

	// 如果对方已经关注你，则设置双方的关系为AttrFriend
	switch model.Attr(fansAttr) {
	case model.AttrFollowing:
		friend = true
		newFollowingAttr = model.SetAttr(newFollowingAttr, model.AttrFriend)
		newFansAttr = model.SetAttr(fansAttr, model.AttrFriend)
	}

	at := model.Attr(followingAttr)

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		// p.d.DelTopicCatalogCache(context.TODO(), req.TopicID)
	})
	return
}

// Follow 关注
func (s *Service) Follow(c context.Context, aid int64, fid int64) (err error) {
	return
}
