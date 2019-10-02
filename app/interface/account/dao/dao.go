package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/account/conf"
	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	mc           *memcache.Pool
	mcExpire     int32
	authMC       *memcache.Pool
	authMCExpire int32
	db           sqalx.Node
	c            *conf.Config
	accountRPC   account.AccountClient
	relationRPC  relation.RelationClient
	topicRPC     topic.TopicClient
	discussRPC   discuss.DiscussionClient
	feedRPC      feed.FeedClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:            c,
		db:           sqalx.NewMySQL(c.DB.Main),
		authMC:       memcache.NewPool(c.Memcache.Auth.Config),
		authMCExpire: int32(time.Duration(c.Memcache.Auth.Expire) / time.Second),
		mc:           memcache.NewPool(c.Memcache.Main.Config),
		mcExpire:     int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
	}

	if relationRPC, err := relation.NewClient(c.RelationRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial relation service"))
	} else {
		dao.relationRPC = relationRPC
	}

	if feedRPC, err := feed.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial feed service"))
	} else {
		dao.feedRPC = feedRPC
	}

	if topicRPC, err := topic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if discussRPC, err := discuss.NewClient(c.DiscussRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.discussRPC = discussRPC
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
