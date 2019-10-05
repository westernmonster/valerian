package service

import (
	"context"
	"fmt"
	"time"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/feed/model"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	var err error
	info := new(model.NotifyDiscussionAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onDiscussionAdded Unmarshal failed %#v", err)
		return
	}

	var item *discuss.DiscussionInfo
	if item, err = p.d.GetDiscussion(context.Background(), info.ID); err != nil {
		return
	}

	feed := &model.TopicFeed{
		ID:         gid.NewID(),
		TopicID:    item.TopicID,
		ActionType: model.ActionTypeCreateDiscussion,
		ActionTime: item.CreatedAt,
		ActionText: model.ActionTextCreateDiscussion,
		ActorID:    item.CreatedBy,
		ActorType:  model.ActorTypeUser,
		TargetID:   item.ID,
		TargetType: model.TargetTypeDiscussion,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if err = p.d.AddTopicFeed(context.Background(), p.d.DB(), feed); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
}

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
	var err error
	info := new(model.NotifyDiscussionDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onDiscussionDeleted Unmarshal failed %#v", err)
		return
	}

	fmt.Printf("delete topic_id(%d), target_type(%s), id(%d)\n", info.TopicID, model.TargetTypeDiscussion, info.ID)
	if err = p.d.DelTopicFeedByCond(context.Background(), p.d.DB(), info.TopicID, model.TargetTypeDiscussion, info.ID); err != nil {
		return
	}
}
