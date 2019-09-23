package service

import (
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

const (
	BusNotifyDiscussionAdded = "feed.discussion.added"
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

func (p *Service) feedConsumer() (consumer *FeedConsumer) {
	consumer = &FeedConsumer{
		Subscriptions: make([]stan.Subscription, 0),
	}
	if sub, e := p.sc.Subscribe(BusNotifyDiscussionAdded, p.onDiscussionAdded, stan.DeliverAllAvailable(), stan.SetManualAckMode()); e != nil {
		panic(e)
	} else {
		consumer.Subscriptions = append(consumer.Subscriptions, sub)
	}

	return
}

func (p *Service) onDiscussionAdded(m *stan.Msg) {
	// feed := &model.TopicFeed{
	// 	ID: gid.NewID(),
	// }
}

func (p *Service) onDiscussionUpdated(m *stan.Msg) {
}

func (p *Service) onDiscussionDeleted(m *stan.Msg) {
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
