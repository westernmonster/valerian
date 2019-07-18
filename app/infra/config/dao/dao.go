package dao

import (
	"context"
	"time"
	"valerian/app/infra/config/conf"
	"valerian/library/cache/redis"
	"valerian/library/database/sqalx"
)

// Dao dao.
type Dao struct {
	// redis
	redis  *redis.Pool
	expire time.Duration
	// cache
	pathCache string
	//DB
	DB *sqalx.Node
}

// New new a dao.
func New(c *conf.Config) *Dao {
	d := &Dao{
		// redis
		redis:  redis.NewPool(c.Redis),
		expire: time.Duration(c.PollTimeout),
		// cache
		pathCache: c.PathCache,
		// orm
		Node: sqalx.NewSQL(c.DB),
	}
	return d
}

// Ping ping is ok.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.pingRedis(c); err != nil {
		return
	}
	return d.DB.DB().PingContext(c)
}

// Close close resuouces.
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}
