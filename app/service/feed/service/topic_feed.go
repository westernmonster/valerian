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

type FeedConsumer struct {
	Subscriptions []stan.Subscription
}

func (p *FeedConsumer) Unsubscribe() {
	for _, v := range p.Subscriptions {
		if err := v.Unsubscribe(); err != nil {
			log.Errorf("Unsubscribe failed %#v", err)
		}
	}
}

func (p *Service) initFeedConsumer() (consumer *FeedConsumer) {
	consumer = &FeedConsumer{
		Subscriptions: make([]stan.Subscription, 0),
	}
	if sub, e := p.sc.Subscribe(model.BusDiscussionAdded, p.onDiscussionAdded, stan.DeliverAllAvailable(), stan.SetManualAckMode()); e != nil {
		panic(e)
	} else {
		consumer.Subscriptions = append(consumer.Subscriptions, sub)
	}

	if sub, e := p.sc.Subscribe(model.BusDiscussionDeleted, p.onDiscussionDeleted, stan.DeliverAllAvailable(), stan.SetManualAckMode()); e != nil {
		panic(e)
	} else {
		consumer.Subscriptions = append(consumer.Subscriptions, sub)
	}

	p.feedConsumer = consumer

	return
}

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

func (p *Service) onArticleAdded(m *stan.Msg) {
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
}

func (p *Service) onMemberJoined(m *stan.Msg) {
}

func (p *Service) onReviseAdded(m *stan.Msg) {
}

func (p *Service) onReviseUpdated(m *stan.Msg) {
}

func (p *Service) onReviseDeleted(m *stan.Msg) {
}

func (p *Service) GetTopicFeedPaged(c context.Context, topicID int64, limit, offset int) (items []*model.TopicFeed, err error) {
	return p.d.GetTopicFeedPaged(c, p.d.DB(), topicID, limit, offset)
}
