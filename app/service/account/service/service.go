package service

import (
	"context"

	"valerian/app/service/account/conf"
	"valerian/app/service/account/dao"
	"valerian/app/service/account/model"
	certification "valerian/app/service/certification/api"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

type Service struct {
	c *conf.Config
	d interface {
		GetAccounts(c context.Context, node sqalx.Node) (items []*model.Account, err error)
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		AddAccount(c context.Context, node sqalx.Node, item *model.Account) (err error)

		AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)
		GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error)

		IncrMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		UpdateMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		AddMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		GetMessageStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.MessageStat, err error)

		GetArea(ctx context.Context, node sqalx.Node, id int64) (item *model.Area, err error)

		GetAccountSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountSetting, err error)
		AddAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error)
		UpdateAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error)

		GetWorkCertStatus(c context.Context, aid int64) (status int32, err error)
		GetIDCertStatus(c context.Context, aid int64) (status int32, err error)
		GetWorkCert(c context.Context, aid int64) (resp *certification.WorkCertInfo, err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)
		BatchAccountCache(c context.Context, aids []int64) (cached map[int64]*model.Account, missed []int64, err error)
		SetBatchAccountCache(c context.Context, bs []*model.Account) (err error)

		AccountSetLock(c context.Context, node sqalx.Node, accountID int64, isLock bool) (err error)
		AnnulAccount(c context.Context, node sqalx.Node, aid int64) (err error)
		UnAnnulAccount(c context.Context, node sqalx.Node, aid int64, username, password, salt string) (err error)

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
