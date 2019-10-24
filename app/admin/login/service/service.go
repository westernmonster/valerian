package service

import (
	"context"
	"net/url"

	"valerian/app/admin/login/conf"
	"valerian/app/admin/login/dao"
	"valerian/app/admin/login/model"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/http/mars/middleware/permit"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error)

		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		SetPassword(c context.Context, node sqalx.Node, password, salt string, aid int64) (err error)
		UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)

		MobileValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelMobileValcodeCache(c context.Context, vtype int, mobile string) (err error)
		EmailValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelEmailValcideCache(c context.Context, vtype int, mobile string) (err error)

		SetSession(ctx context.Context, p *permit.Session) (err error)
		Session(ctx context.Context, sid string) (res *permit.Session, err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
		AuthDB() sqalx.Node
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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}