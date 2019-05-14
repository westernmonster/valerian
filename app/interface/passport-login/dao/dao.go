package dao

import (
	"context"
	"valerian/app/interface/passport-login/conf"

	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Dao dao struct
type Dao struct {
	node     sqalx.Node
	mc       *memcache.Pool
	mcExpire int32
	logger   log.Factory
	c        *conf.Config
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
	}
	return
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	// if err = d.Node.Ping(c); err != nil {
	// 	return
	// }
	// return d.pingMC(c)
	return nil
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.mc != nil {
		d.mc.Close()
	}
	// if d.db != nil {
	// 	d.db.Close()
	// }
}
