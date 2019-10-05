package dao

import (
	"context"
	"fmt"
	"strings"
	"time"

	"valerian/app/interface/topic/conf"
	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	relation "valerian/app/service/relation/api"
	stopic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	db          sqalx.Node
	mc          *memcache.Pool
	mcExpire    int32
	c           *conf.Config
	accountRPC  account.AccountClient
	feedRPC     feed.FeedClient
	discussRPC  discuss.DiscussionClient
	topicRPC    stopic.TopicClient
	relationRPC relation.RelationClient
	sc          stan.Conn
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
	}

	servers := strings.Join(c.Nats.Nodes, ",")
	if sc, err := stan.Connect("valerian",
		env.Hostname,
		stan.Pings(10, 5),
		stan.NatsURL(servers),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Errorf("Nats Connection lost, reason: %v", reason)
			panic(reason)
		}),
	); err != nil {
		log.Errorf("connect to servers failed %#v\n", err)
		panic(err)
	} else {
		dao.sc = sc
	}

	if relationRPC, err := relation.NewClient(c.RelationRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial relation service"))
	} else {
		dao.relationRPC = relationRPC
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
	}
	if discussRPC, err := discuss.NewClient(c.DiscussRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial discuss service"))
	} else {
		dao.discussRPC = discussRPC
	}

	if feedRPC, err := feed.NewClient(c.FeedRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial feed service"))
	} else {
		dao.feedRPC = feedRPC
	}

	if topicRPC, err := stopic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	return
}

func (d *Dao) DB() sqalx.Node {
	return d.db
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}
	if err = d.pingMC(c); err != nil {
		log.Info(fmt.Sprintf("dao.mc.Ping() error(%v)", err))
		return
	}
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.mc != nil {
		d.mc.Close()
	}
	if d.db != nil {
		d.db.Close()
	}
}

// PromError prometheus error count.
func PromError(c context.Context, name, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.For(c).Error(fmt.Sprintf(format, args...))
}
