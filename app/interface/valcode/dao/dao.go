package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/valcode/conf"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

// Dao dao struct
type Dao struct {
	mc       *memcache.Pool
	mcExpire int32
	c        *conf.Config
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		mc:       memcache.NewPool(c.Memcache.Auth.Config),
		mcExpire: int32(time.Duration(c.Memcache.Auth.Expire) / time.Second),
	}
	return
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
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
}
