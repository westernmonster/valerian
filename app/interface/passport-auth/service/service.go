package service

import (
	"context"

	"valerian/app/conf"
	"valerian/app/interface/passport-auth/dao"
	"valerian/app/interface/passport-auth/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

var (
	_noLogin = &model.AuthReply{
		Login: false,
	}
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error)
		AddAccessToken(c context.Context, node sqalx.Node, t *model.AccessToken) (affected int64, err error)
		GetAccessToken(c context.Context, node sqalx.Node, token string) (item *model.AccessToken, err error)
		DelExpiredAccessToken(c context.Context, node sqalx.Node, clientID string, accountID int64, expiresAt int64) (affected int64, err error)
		AddRefreshToken(c context.Context, node sqalx.Node, t *model.RefreshToken) (affected int64, err error)
		DelRefreshToken(c context.Context, node sqalx.Node, token string) (affected int64, err error)
		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)

		RefreshTokenCache(c context.Context, sd string) (item *model.RefreshToken, err error)
		SetAccessTokenCache(c context.Context, m *model.AccessToken) (err error)
		AccessTokenCache(c context.Context, token string) (res *model.AccessToken, err error)
		DelTokenCache(c context.Context, token string) (err error)

		DB() sqalx.Node
		AuthDB() sqalx.Node
		// Ping check db and mc health.
		Ping(c context.Context) (err error)
		// Close close connections of mc, redis, db.
		Close()
	}
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
	}
	go s.cacheproc()
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close dao.
func (s *Service) Close() {
	s.d.Close()
}

func (s *Service) addCache(f func()) {
	select {
	case s.missch <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc is a routine for executing closure.
func (s *Service) cacheproc() {
	for {
		f := <-s.missch
		f()
	}
}
