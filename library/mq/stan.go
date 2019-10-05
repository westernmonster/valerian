package mq

import (
	"fmt"
	"strings"
	"time"
	"valerian/library/log"

	xtime "valerian/library/time"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
)

type Config struct {
	Nodes        []string
	ClusterID    string
	AckTimeout   xtime.Duration
	MaxInflight  int
	PingInterval int
	PingMaxOut   int
}

type MessageQueue struct {
	c             *Config
	subscriptions map[string]stan.Subscription
	conn          stan.Conn
}

func (p *MessageQueue) Init(clientID string, config *Config) (err error) {
	p.c = config
	servers := strings.Join(p.c.Nodes, ",")
	if p.conn, err = stan.Connect(p.c.ClusterID, clientID,
		stan.Pings(p.c.PingInterval, p.c.PingMaxOut),
		stan.NatsURL(servers),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Errorf("Nats Connection lost, reason: %v", reason)
			panic(reason)
		}),
	); err != nil {
		log.Errorf("connect to servers failed %#v\n", err)
		panic(err)
	}

	return
}

func (p *MessageQueue) Publish(subject string, data []byte) (err error) {
	if err = p.conn.Publish(subject, data); err != nil {
		log.Error(fmt.Sprintf("mq.Publish(), subject(%s), error(%+v)", subject, err))
	}
	return
}

// QueueSubscribe 分组订阅 手动Ack模式
// 每个组都会收到消息，但是组内成员是随机分配一个接收
func (p *MessageQueue) QueueSubscribe(subject string, qgroup string, cb stan.MsgHandler) (err error) {
	return p.QueueSubscribeWithOpts(subject, qgroup, cb,
		stan.SetManualAckMode(),
		stan.MaxInflight(p.c.MaxInflight),
		stan.AckWait(time.Duration(p.c.AckTimeout)))
}

// QueueSubscribe 分组订阅
// 每个组都会收到消息，但是组内成员是随机分配一个接收
func (p *MessageQueue) QueueSubscribeWithOpts(subject string, qgroup string, cb stan.MsgHandler, options ...stan.SubscriptionOption) (err error) {
	var sub stan.Subscription
	if sub, err = p.conn.QueueSubscribe(subject, qgroup, cb, options...); err != nil {
		log.Error(fmt.Sprintf("mq.QueueSubscribe(), subject(%s) ,qgroup(%s), error(%+v)", subject, qgroup, err))
	}

	key := fmt.Sprintf("%s_%s", subject, qgroup)

	if _, ok := p.subscriptions[key]; ok {
		err = errors.Errorf("subscription already exist, key(%s)", key)
		return
	}

	p.subscriptions[key] = sub

	return
}

func (p *MessageQueue) Close() (err error) {
	if err = p.conn.Close(); err != nil {
		log.Error(fmt.Sprintf("mq.Close(), error(%+v)", err))
	}
	return
}
