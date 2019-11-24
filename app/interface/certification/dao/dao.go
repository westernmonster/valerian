package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/certification/conf"
	account "valerian/app/service/account/api"
	certification "valerian/app/service/certification/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	authDB           sqalx.Node
	db               sqalx.Node
	mc               *memcache.Pool
	mcExpire         int32
	c                *conf.Config
	certificationRPC certification.CertificationClient
	accountRPC       account.AccountClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
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
