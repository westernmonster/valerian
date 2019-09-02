package dao

import (
	"context"
	"fmt"
	"valerian/app/service/msm/conf"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars"
)

// Dao dao.
type Dao struct {
	client     *mars.Client
	db         sqalx.Node
	authDB     sqalx.Node
	apmDB      sqalx.Node
	treeHost   string
	platformID string
}

// New new dao.
func New(c *conf.Config) *Dao {
	d := &Dao{
		db:         sqalx.NewMySQL(c.DB.Main),
		authDB:     sqalx.NewMySQL(c.DB.Auth),
		apmDB:      sqalx.NewMySQL(c.DB.Apm),
		client:     mars.NewClient(c.HTTPClient),
		treeHost:   c.Tree.Host,
		platformID: c.Tree.PlatformID,
	}
	return d
}

// Close close mysql resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}

	if d.authDB != nil {
		d.authDB.Close()
	}

	if d.apmDB != nil {
		d.apmDB.Close()
	}
}

func (d *Dao) DB() sqalx.Node {
	return d.db
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}

	if err = d.authDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.authDB.Ping() error(%v)", err))
	}

	if err = d.apmDB.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.apmDB.Ping() error(%v)", err))
	}
	return
}
