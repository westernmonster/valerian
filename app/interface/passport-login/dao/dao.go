package dao

import (
	"context"

	"valerian/app/interface/passport-login/conf"
	account "valerian/app/service/account/api"
	identify "valerian/app/service/identify/api/grpc"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	c           *conf.Config
	accountRPC  account.AccountClient
	identifyRPC identify.IdentifyClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
	}

	if identifyRPC, err := identify.NewClient(c.IdentifyRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial identify service"))
	} else {
		dao.identifyRPC = identifyRPC
	}

	return
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
}
