package service

import (
	"context"

	"valerian/app/interface/passport-register/conf"
	"valerian/app/interface/passport-register/dao"
	"valerian/app/interface/passport-register/model"
	account "valerian/app/service/account/api"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/mq"
)

// Service struct of service
type Service struct {
	c *conf.Config

	mq *mq.MessageQueue
	d  interface {
		GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error)
		GetArea(ctx context.Context, node sqalx.Node, id int64) (item *model.Area, err error)

		GetAccountRolesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountRole, err error)
		GetAccountRoles(c context.Context, node sqalx.Node) (items []*model.AccountRole, err error)
		GetAccountRoleByID(c context.Context, node sqalx.Node, id string) (item *model.AccountRole, err error)
		GetAccountRoleByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountRole, err error)
		AddAccountRole(c context.Context, node sqalx.Node, item *model.AccountRole) (err error)
		UpdateAccountRole(c context.Context, node sqalx.Node, item *model.AccountRole) (err error)
		DelAccountRole(c context.Context, node sqalx.Node, id string) (err error)

		GetAccessToken(c context.Context, node sqalx.Node, token string) (item *model.AccessToken, err error)
		AddAccessToken(c context.Context, node sqalx.Node, t *model.AccessToken) (affected int64, err error)
		DelExpiredAccessToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error)
		GetClientAccessTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error)

		AddRefreshToken(c context.Context, node sqalx.Node, t *model.RefreshToken) (affected int64, err error)
		DelRefreshToken(c context.Context, node sqalx.Node, token string) (affected int64, err error)
		DelExpiredRefreshToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error)
		GetClientRefreshTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error)

		AddAccount(c context.Context, v *account.AddAccountReq) (resp *account.SelfProfile, err error)

		IncrMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		UpdateMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		AddMessageStat(c context.Context, node sqalx.Node, item *model.MessageStat) (err error)
		GetMessageStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.MessageStat, err error)

		AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)

		GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
		GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
		GetSelfProfile(c context.Context, aid int64) (info *account.SelfProfile, err error)
		GetAccountStat(c context.Context, aid int64) (info *account.AccountStatInfo, err error)

		MobileValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error)
		DelMobileValcodeCache(c context.Context, vtype int32, mobile string) (err error)
		EmailValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error)
		DelEmailValcodeCache(c context.Context, vtype int32, mobile string) (err error)

		AccessTokenCache(c context.Context, token string) (res *model.AccessToken, err error)
		SetAccessTokenCache(c context.Context, m *model.AccessToken) (err error)
		DelAccessTokenCache(c context.Context, token string) (err error)

		SetProfileCache(c context.Context, m *model.Profile) (err error)

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
		mq:     mq.New(env.Hostname, c.Nats),
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
