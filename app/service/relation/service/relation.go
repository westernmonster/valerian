package service

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) onFollow(c context.Context, aid, fid int64) {
	// 关注通知
	if err := p.d.NotifyFollow(c, aid, fid); err != nil {
		log.For(c).Error(fmt.Sprintf("Failed to set recent follow notify aid(%d) fid(%d): %+v", aid, fid, err))
		return
	}
}

func (p *Service) onUnfollow(c context.Context, aid, fid int64) {
	// 取关通知
	if err := p.d.NotifyUnfollow(c, aid, fid); err != nil {
		log.For(c).Error(fmt.Sprintf("Failed to set recent follow notify aid(%d) fid(%d): %+v", aid, fid, err))
		return
	}
}

func (p *Service) initStat(c context.Context, node sqalx.Node, aid int64, fid int64) (err error) {
	var statCurrent, statTarget *model.AccountRelationStat

	if statCurrent, err = p.d.GetStatByID(c, node, aid); err != nil {
		return
	} else if statCurrent == nil {
		if err = p.d.AddStat(c, node, &model.AccountRelationStat{
			AccountID: aid,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}); err != nil {
			return
		}
	}

	if statTarget, err = p.d.GetStatByID(c, node, aid); err != nil {
		return
	} else if statTarget == nil {
		if err = p.d.AddStat(c, node, &model.AccountRelationStat{
			AccountID: fid,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}); err != nil {
			return
		}
	}

	return
}

// Unfollow 取关
func (p *Service) Unfollow(c context.Context, aid int64, fid int64) (err error) {
	if aid <= 0 || fid <= 0 {
		return
	}
	if aid == fid {
		err = ecode.RelFollowSelfBanned
		return
	}

	if err = p.initStat(c, p.d.DB(), aid, fid); err != nil {
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

	var followingAttr uint32
	var friend bool
	var statCurrent = &model.AccountRelationStat{AccountID: aid}
	var statTarget = &model.AccountRelationStat{AccountID: fid}

	// 如果没有关注对方，则设置为AttrNoRelation
	var followingItem *model.AccountFollowing
	if followingItem, err = p.d.GetFollowingByCond(c, tx, map[string]interface{}{
		"account_id":   aid,
		"following_id": fid,
	}); err != nil {
		return
	} else if followingItem == nil {
		followingAttr = model.AttrNoRelation
	} else {
		followingAttr = followingItem.Attribute
	}

	// 如果已经关注，则清空Cache
	switch model.Attr(followingAttr) {
	case model.AttrFriend:
		friend = true
	case model.AttrFollowing:
	default:
		// p.CompareAndDelCache(c, mid, fid, a)
		log.For(c).Warn(fmt.Sprintf("Invalid state between %d and %d with attribute: %d", aid, fid, followingAttr))
		return
	}

	if friend {
		var newFansAttr = model.AttrFollowing
		if err = p.d.DelFollowing(c, tx, aid, fid); err != nil {
			return
		}
		if err = p.d.SetFollowing(c, tx, newFansAttr, fid, aid); err != nil {
			return
		}

		if err = p.d.DelFans(c, tx, fid, aid); err != nil {
			return
		}

		if err = p.d.SetFans(c, tx, newFansAttr, fid, aid); err != nil {
			return
		}
	} else {
		if err = p.d.DelFollowing(c, tx, aid, fid); err != nil {
			return
		}
		if err = p.d.DelFans(c, tx, fid, aid); err != nil {
			return
		}
	}

	statCurrent.Following = -1
	statTarget.Fans = -1

	if !statCurrent.Empty() {
		if err = p.d.IncrStat(c, tx, statCurrent); err != nil {
			return
		}
	}

	if !statTarget.Empty() {
		if err = p.d.IncrStat(c, tx, statTarget); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onFollow(context.Background(), aid, fid)
		p.d.DelFansCache(context.Background(), aid)
		p.d.DelFansCache(context.Background(), fid)
		p.d.DelFollowingsCache(context.Background(), aid)
		p.d.DelFollowingsCache(context.Background(), fid)
	})

	return
}

// Follow 关注
func (p *Service) Follow(c context.Context, aid int64, fid int64) (err error) {
	if aid <= 0 || fid <= 0 {
		return
	}
	if aid == fid {
		err = ecode.RelFollowSelfBanned
		return
	}

	if err = p.initStat(c, p.d.DB(), aid, fid); err != nil {
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

	log.For(c).Info(fmt.Sprintf("aid(%d) Follow fid(%d)", aid, fid))

	var followingAttr uint32
	var fansAttr uint32
	var friend bool
	var statCurrent = &model.AccountRelationStat{AccountID: aid}
	var statTarget = &model.AccountRelationStat{AccountID: fid}

	// 如果没有关注对方，则设置为AttrNoRelation
	var followingItem *model.AccountFollowing
	if followingItem, err = p.d.GetFollowingByCond(c, tx, map[string]interface{}{
		"account_id":   aid,
		"following_id": fid,
	}); err != nil {
		return
	} else if followingItem == nil {
		followingAttr = model.AttrNoRelation
	} else {
		followingAttr = followingItem.Attribute
	}

	// 如果对方没有关注你，则设置为AttrNoRelation
	var fansItem *model.AccountFans
	if fansItem, err = p.d.GetFansByCond(c, tx, map[string]interface{}{
		"account_id":   fid,
		"following_id": aid,
	}); err != nil {
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
		friend = true // 如果对方已经关注了你，则设置为friend
		newFollowingAttr = model.SetAttr(newFollowingAttr, model.AttrFriend)
		newFansAttr = model.SetAttr(fansAttr, model.AttrFriend)
	}

	at := model.Attr(followingAttr)
	switch at {
	case model.AttrBlack:
		// 已经把对方加入了黑名单
		err = ecode.RelFollowAlreadyBlack
		return
	case model.AttrFriend, model.AttrFollowing:
		// 已经关注或已经是好友
		return
	case model.AttrWhisper:
		// 悄悄关注，暂时不做这个逻辑
		return
	case model.AttrNoRelation:
		if friend {
			// 对方已经关注了你时候逻辑处理

			// 关注列表中中添加 你->目标用户
			if err = p.d.AddFollowing(c, tx, &model.AccountFollowing{
				ID:              gid.NewID(),
				AccountID:       aid,
				TargetAccountID: fid,
				Attribute:       newFollowingAttr,
				CreatedAt:       time.Now().Unix(),
				UpdatedAt:       time.Now().Unix(),
			}); err != nil {
				return
			}

			// 关注列表中 目标用户->你 属性加上朋友
			if err = p.d.SetFollowing(c, tx, newFansAttr, fid, aid); err != nil {
				return
			}

			// 目标用户粉丝列表中 添加上你
			if err = p.d.AddFans(c, tx, &model.AccountFans{
				ID:              gid.NewID(),
				AccountID:       fid,
				TargetAccountID: aid,
				Attribute:       newFollowingAttr,
				CreatedAt:       time.Now().Unix(),
				UpdatedAt:       time.Now().Unix(),
			}); err != nil {
				return
			}
			// 你的粉丝列表中设置你和目标用户的关系为好友
			if err = p.d.SetFans(c, tx, newFansAttr, aid, fid); err != nil {
				return
			}

		} else {
			// 关注列表中中添加 你->目标用户
			if err = p.d.AddFollowing(c, tx, &model.AccountFollowing{
				ID:              gid.NewID(),
				AccountID:       aid,
				TargetAccountID: fid,
				Attribute:       newFollowingAttr,
				CreatedAt:       time.Now().Unix(),
				UpdatedAt:       time.Now().Unix(),
			}); err != nil {
				return
			}

			// 目标用户粉丝列表中 添加上你
			if err = p.d.AddFans(c, tx, &model.AccountFans{
				ID:              gid.NewID(),
				AccountID:       fid,
				TargetAccountID: aid,
				Attribute:       newFollowingAttr,
				CreatedAt:       time.Now().Unix(),
				UpdatedAt:       time.Now().Unix(),
			}); err != nil {
				return
			}
		}

		statCurrent.Following = 1
		statTarget.Fans = 1
	}

	if !statCurrent.Empty() {
		if err = p.d.IncrStat(c, tx, statCurrent); err != nil {
			return
		}
	}

	if !statTarget.Empty() {
		if err = p.d.IncrStat(c, tx, statTarget); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	p.addCache(func() {
		p.onUnfollow(context.Background(), aid, fid)
		p.d.DelFansCache(context.Background(), aid)
		p.d.DelFansCache(context.Background(), fid)
		p.d.DelFollowingsCache(context.Background(), aid)
		p.d.DelFollowingsCache(context.Background(), fid)
	})

	return
}
