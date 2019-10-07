package service

import (
	"context"

	"valerian/app/interface/account/conf"
	"valerian/app/interface/account/dao"
	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
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

		GetAccountSettingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountSetting, err error)
		AddAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error)
		UpdateAccountSetting(c context.Context, node sqalx.Node, item *model.AccountSetting) (err error)

		GetFollowings(c context.Context, accountID int64, limit, offset int) (resp *relation.FollowingResp, err error)
		GetFans(c context.Context, accountID int64, limit, offset int) (resp *relation.FansResp, err error)
		Follow(c context.Context, accountID, targetAccountID int64) (err error)
		Unfollow(c context.Context, accountID, targetAccountID int64) (err error)
		IsFollowing(c context.Context, aid, targetAccountID int64) (IsFollowing bool, err error)

		GetAccountStat(c context.Context, aid int64) (stat *account.AccountStatInfo, err error)

		GetAccountFeedPaged(c context.Context, accountID int64, limit, offset int) (info *feed.AccountFeedResp, err error)
		GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *discuss.UserDiscussionsResp, err error)
		GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *topic.UserTopicsResp, err error)
		GetUserArticlesPaged(c context.Context, aid int64, limit, offset int) (resp *article.UserArticlesResp, err error)

		GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)

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
