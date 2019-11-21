package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/account/conf"
	accountFeed "valerian/app/service/account-feed/api"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	certification "valerian/app/service/certification/api"
	discuss "valerian/app/service/discuss/api"
	message "valerian/app/service/message/api"
	recent "valerian/app/service/recent/api"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	mc               *memcache.Pool
	mcExpire         int32
	authMC           *memcache.Pool
	authMCExpire     int32
	db               sqalx.Node
	c                *conf.Config
	accountRPC       account.AccountClient
	relationRPC      relation.RelationClient
	articleRPC       article.ArticleClient
	topicRPC         topic.TopicClient
	messageRPC       message.MessageClient
	discussRPC       discuss.DiscussionClient
	accountFeedRPC   accountFeed.AccountFeedClient
	recentRPC        recent.RecentClient
	certificationRPC certification.CertificationClient
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

	if certificationRPC, err := certification.NewClient(c.CertificationRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial certification service"))
	} else {
		dao.certificationRPC = certificationRPC
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

	if articleRPC, err := article.NewClient(c.ArticleRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial article service"))
	} else {
		dao.articleRPC = articleRPC
	}

	if accountFeedRPC, err := accountFeed.NewClient(c.AccountFeedRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial accountFeed service"))
	} else {
		dao.accountFeedRPC = accountFeedRPC
	}

	if topicRPC, err := topic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if messageRPC, err := message.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial message service"))
	} else {
		dao.messageRPC = messageRPC
	}

	if discussRPC, err := discuss.NewClient(c.DiscussRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.discussRPC = discussRPC
	}

	if recentRPC, err := recent.NewClient(c.RecentRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial recent service"))
	} else {
		dao.recentRPC = recentRPC
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
