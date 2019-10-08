package service

import (
	"context"
	"fmt"
	"valerian/app/service/relation/model"
	"valerian/library/log"
)

// func (p *Dao) NotifyFollow(c context.Context, aid, fid int64) (err error) {
// 	msg := &NotifyFollow{AccountID: aid, TargetAccountID: fid}
// 	var data []byte
// 	if data, err = msg.Marshal(); err != nil {
// 		return
// 	}

// 	if err = p.sc.Publish(BusNotifyFollow, data); err != nil {
// 		return
// 	}

// 	return
// }

// func (p *Dao) NotifyUnfollow(c context.Context, aid, fid int64) (err error) {
// 	msg := &NotifyUnfollow{AccountID: aid, TargetAccountID: fid}
// 	var data []byte
// 	if data, err = msg.Marshal(); err != nil {
// 		return
// 	}

// 	if err = p.sc.Publish(BusNotifyUnfollow, data); err != nil {
// 		return
// 	}

// 	return
// }
func (p *Service) onFollow(c context.Context, aid, fid int64) {
	msg := &model.NotifyFollow{AccountID: aid, TargetAccountID: fid}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onFollow.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusMemberFollowed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onFollow.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onUnfollow(c context.Context, aid, fid int64) {
	msg := &model.NotifyUnfollow{AccountID: aid, TargetAccountID: fid}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onUnfollow.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(model.BusMemberUnfollowed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onUnfollow.Publish(), err(%+v)", err))
		return
	}

	return
}
