package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/conf"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Dao dao struct
type Dao struct {
	authDB       sqalx.Node
	db           sqalx.Node
	authMC       *memcache.Pool
	authMCExpire int32
	c            *conf.Config
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:            c,
		db:           sqalx.NewMySQL(c.DB.Main),
		authDB:       sqalx.NewMySQL(c.DB.Auth),
		authMC:       memcache.NewPool(c.Memcache.Auth.Config),
		authMCExpire: int32(time.Duration(c.Memcache.Auth.Expire) / time.Second),
	}
	return
}

func (d *Dao) DB() sqalx.Node {
	return d.db
}

func (d *Dao) AuthDB() sqalx.Node {
	return d.authDB
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}
	if err = d.authDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.authDB.Ping() error(%v)", err))
	}
	return
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
