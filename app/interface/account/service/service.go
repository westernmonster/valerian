package service

import (
	"context"
	"valerian/app/conf"
	"valerian/app/interface/account/dao"
	"valerian/app/interface/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetArea(ctx context.Context, node sqalx.Node, id int64) (item *model.Area, err error)
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		SetPassword(c context.Context, node sqalx.Node, password, salt string, aid int64) (err error)
		UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error)

		GetAccountFollowersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFollower, err error)
		GetAccountFollowers(c context.Context, node sqalx.Node) (items []*model.AccountFollower, err error)
		GetAccountFollowerByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFollower, err error)
		GetAccountFollowerByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFollower, err error)
		AddAccountFollower(c context.Context, node sqalx.Node, item *model.AccountFollower) (err error)
		UpdateAccountFollower(c context.Context, node sqalx.Node, item *model.AccountFollower) (err error)
		DelAccountFollower(c context.Context, node sqalx.Node, id int64) (err error)

		GetFansCount(c context.Context, node sqalx.Node, aid int64) (count int, err error)
		GetFansPaged(c context.Context, node sqalx.Node, aid int64, page, pageSize int) (count int, items []*model.AccountFollower, err error)
		GetFollowingCount(c context.Context, node sqalx.Node, aid int64) (count int, err error)
		GetFollowingPaged(c context.Context, node sqalx.Node, aid int64, page, pageSize int) (count int, items []*model.AccountFollower, err error)

		ProfileCache(c context.Context, id int64) (m *model.Profile, err error)
		SetProfileCache(c context.Context, m *model.Profile) (err error)
		DelProfileCache(c context.Context, id int64) (err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)

		MobileValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelMobileCache(c context.Context, vtype int, mobile string) (err error)
		EmailValcodeCache(c context.Context, vtype int, mobile string) (code string, err error)
		DelEmailCache(c context.Context, vtype int, mobile string) (err error)
		SetSessionResetPasswordCache(c context.Context, sessionID string, accountID int64) (err error)
		SessionResetPasswordCache(c context.Context, sessionID string) (aid int64, err error)
		DelResetPasswordCache(c context.Context, sessionID string) (err error)
		DelAccessTokenCache(c context.Context, token string) (err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
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
