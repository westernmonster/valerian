package dao

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/identify/conf"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"
)

const (
	_tokenURI  = "/intranet/auth/tokenInfo"
	_cookieURI = "/intranet/auth/cookieInfo"
)

var (
	errorsCount = prom.BusinessErrCount
	cachedCount = prom.CacheHit
	missedCount = prom.CacheMiss
)

// PromError prom error
func PromError(name string) {
	errorsCount.Incr(name)
}

// Dao struct info of Dao
type Dao struct {
	c *conf.Config

	db     sqalx.Node
	authDB sqalx.Node

	authMC       *memcache.Pool
	authMCExpire int32

	// mcLogin *memcache.Pool
	// client  *mars.Client
}

// New new a Dao and return.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:            c,
		authMC:       memcache.NewPool(c.AuthMC.Config),
		authMCExpire: int32(time.Duration(c.AuthMC.Expire) / time.Second),
		db:           sqalx.NewMySQL(c.DB.Main),
		authDB:       sqalx.NewMySQL(c.DB.Auth),
	}
	return
}

func (d *Dao) DB() sqalx.Node {
	return d.db
}

func (d *Dao) AuthDB() sqalx.Node {
	return d.authDB
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.authMC != nil {
		d.authMC.Close()
	}
	if d.db != nil {
		d.db.Close()
	}

	if d.authDB != nil {
		d.authDB.Close()
	}
}

// Ping ping health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}
	if err = d.authDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.authDB.Ping() error(%v)", err))
	}

	if err = d.pingMC(c); err != nil {
		PromError("mc:Ping")
	}
	return
}
