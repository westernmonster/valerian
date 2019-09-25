package dao

import (
	"context"
	"valerian/app/interface/discuss/model"
)

const (
	BusNotifyDiscussionAdded   = "notify.discussion.added"
	BusNotifyDiscussionUpdated = "notify.discussion.updated"
	BusNotifyDiscussionDeleted = "notify.discussion.deleted"
)

func (p *Dao) NotifyDiscussionAdded(c context.Context, id int64) (err error) {
	msg := &model.NotifyDiscussionAdded{ID: id}
	var data []byte
	if data, err = msg.Marshal(); err != nil {
		return
	}

	if err = p.sc.Publish(BusNotifyDiscussionAdded, data); err != nil {
		return
	}

	return
}

func (p *Dao) NotifyDiscussionUpdated(c context.Context, id int64) (err error) {
	msg := &model.NotifyDiscussionUpdated{ID: id}
	var data []byte
	if data, err = msg.Marshal(); err != nil {
		return
	}

	if err = p.sc.Publish(BusNotifyDiscussionUpdated, data); err != nil {
		return
	}

	return
}

func (p *Dao) NotifyDiscussionDeleted(c context.Context, id, topicID int64) (err error) {
	msg := &model.NotifyDiscussionDeleted{ID: id, TopicID: topicID}
	var data []byte
	if data, err = msg.Marshal(); err != nil {
		return
	}

	if err = p.sc.Publish(BusNotifyDiscussionDeleted, data); err != nil {
		return
	}

	return
}
