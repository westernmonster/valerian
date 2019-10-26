package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/valcode/conf"
	account "valerian/app/service/account/api"
	"valerian/library/cache/memcache"
	"valerian/library/log"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	mc         *memcache.Pool
	mcExpire   int32
	c          *conf.Config
	accountRPC account.AccountClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		mc:       memcache.NewPool(c.Memcache.Auth.Config),
		mcExpire: int32(time.Duration(c.Memcache.Auth.Expire) / time.Second),
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
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
