package dao

import (
	"context"

	"valerian/app/interface/passport-auth/conf"
	identify "valerian/app/service/identify/api/grpc"
	"valerian/library/database/sqalx"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	authDB      sqalx.Node
	db          sqalx.Node
	c           *conf.Config
	identifyRPC identify.IdentifyClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
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
