package dao

import (
	"context"

	"valerian/app/interface/certification/conf"
	account "valerian/app/service/account/api"
	certification "valerian/app/service/certification/api"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	c                *conf.Config
	certificationRPC certification.CertificationClient
	accountRPC       account.AccountClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
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

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
}
