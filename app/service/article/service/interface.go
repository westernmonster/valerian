package service

import (
	"context"
	"valerian/library/database/sqalx"
)

type IDao interface {
	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
