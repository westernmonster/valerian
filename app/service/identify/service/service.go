package service

import (
	"context"
	api "valerian/app/service/identify/api/grpc"
	"valerian/app/service/identify/conf"
	"valerian/app/service/identify/dao"
	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars/middleware/permit"
)

var (
	_noLogin = &api.AuthReply{
		Login: false,
	}
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error)

		GetAccessToken(c context.Context, node sqalx.Node, token string) (item *model.AccessToken, err error)
		AddAccessToken(c context.Context, node sqalx.Node, t *model.AccessToken) (affected int64, err error)
		DelExpiredAccessToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error)
		GetClientAccessTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error)

		GetRefreshToken(c context.Context, node sqalx.Node, token string) (item *model.RefreshToken, err error)
		AddRefreshToken(c context.Context, node sqalx.Node, t *model.RefreshToken) (affected int64, err error)
		DelRefreshToken(c context.Context, node sqalx.Node, token string) (affected int64, err error)
		DelExpiredRefreshToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error)
		GetClientRefreshTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error)

		SetSession(ctx context.Context, p *permit.Session) (err error)
		Session(ctx context.Context, sid string) (res *permit.Session, err error)

		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)

		AccessTokenCache(c context.Context, token string) (res *model.AccessToken, err error)
		SetAccessTokenCache(c context.Context, m *model.AccessToken) (err error)
		DelAccessTokenCache(c context.Context, token string) (err error)

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
