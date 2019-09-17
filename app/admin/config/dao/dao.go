package dao

import (
	"context"
	"fmt"

	"valerian/app/admin/config/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Dao dao struct
type Dao struct {
	apmDB    sqalx.Node
	configDB sqalx.Node
	c        *conf.Config
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		apmDB:    sqalx.NewMySQL(c.DB.Apm),
		configDB: sqalx.NewMySQL(c.DB.Config),
	}
	return
}

func (d *Dao) ApmDB() sqalx.Node {
	return d.apmDB
}

func (d *Dao) ConfigDB() sqalx.Node {
	return d.configDB
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.apmDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}

	if err = d.configDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.apmDB != nil {
		d.apmDB.Close()
	}

	if d.configDB != nil {
		d.configDB.Close()
	}
}
