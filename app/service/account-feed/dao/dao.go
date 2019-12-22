package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/service/account-feed/conf"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"
)

// Dao dao struct
type Dao struct {
	db       sqalx.Node
	mc       *memcache.Pool
	mcExpire int32
	c        *conf.Config

	topicRPC   topic.TopicClient
	discussRPC discuss.DiscussionClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
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
