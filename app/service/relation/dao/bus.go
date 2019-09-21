package dao

import (
	"context"
)

const (
	BusNotifyFollow   = "notify.follow"
	BusNotifyUnfollow = "notify.unfollow"
)

func (p *Dao) NotifyFollow(c context.Context, aid, fid int64) (err error) {
	msg := &NotifyFollow{AccountID: aid, TargetAccountID: fid}
	var data []byte
	if data, err = msg.Marshal(); err != nil {
		return
	}

	if err = p.sc.Publish(BusNotifyFollow, data); err != nil {
		return
	}

	return
}

func (p *Dao) NotifyUnfollow(c context.Context, aid, fid int64) (err error) {
	msg := &NotifyUnfollow{AccountID: aid, TargetAccountID: fid}
	var data []byte
	if data, err = msg.Marshal(); err != nil {
		return
	}

	if err = p.sc.Publish(BusNotifyUnfollow, data); err != nil {
		return
	}

	return
}
