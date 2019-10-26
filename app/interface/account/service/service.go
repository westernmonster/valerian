package service

import (
	"context"
	"net/url"

	"valerian/app/interface/account/conf"
	"valerian/app/interface/account/dao"
	"valerian/app/interface/account/model"
	feed "valerian/app/service/account-feed/api"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	recent "valerian/app/service/recent/api"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/mq"
)

// Service struct of service
type Service struct {
	c  *conf.Config
	mq *mq.MessageQueue
	d  interface {
		GetArea(ctx context.Context, node sqalx.Node, id int64) (item *model.Area, err error)
		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
		GetAccountByEmail(c context.Context, node sqalx.Node, email string) (item *model.Account, err error)
		GetAccountByMobile(c context.Context, node sqalx.Node, mobile string) (item *model.Account, err error)
		SetPassword(c context.Context, node sqalx.Node, password, salt string, aid int64) (err error)
		UpdateAccount(c context.Context, node sqalx.Node, item *model.Account) (err error)

		GetFollowings(c context.Context, accountID int64, limit, offset int) (resp *relation.FollowingResp, err error)
		GetFans(c context.Context, accountID int64, limit, offset int) (resp *relation.FansResp, err error)
		Follow(c context.Context, accountID, targetAccountID int64) (err error)
		Unfollow(c context.Context, accountID, targetAccountID int64) (err error)
		IsFollowing(c context.Context, aid, targetAccountID int64) (IsFollowing bool, err error)

		AddImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error)
		DelImageURLByCond(c context.Context, node sqalx.Node, targetType string, targetID int64) (err error)
		GetImageUrlsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ImageURL, err error)

		GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error)

		GetAccountFeedPaged(c context.Context, accountID int64, limit, offset int) (info *feed.AccountFeedResp, err error)
		GetUserDiscussionsPaged(c context.Context, aid int64, limit, offset int) (resp *discuss.UserDiscussionsResp, err error)
		GetUserTopicsPaged(c context.Context, aid int64, limit, offset int) (resp *topic.UserTopicsResp, err error)
		GetUserArticlesPaged(c context.Context, aid int64, limit, offset int) (resp *article.UserArticlesResp, err error)
		GetRecentPubsPaged(c context.Context, aid int64, targetType string, limit, offset int) (info *recent.RecentPubsResp, err error)

		GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
		GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)
		GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
		GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)

		GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
		GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
		GetSelfProfile(c context.Context, aid int64) (info *account.SelfProfile, err error)
		GetAccountStat(c context.Context, aid int64) (info *account.AccountStatInfo, err error)
		GetAccountSetting(c context.Context, aid int64) (info *account.Setting, err error)
		UpdateAccountSetting(c context.Context, aid int64, boolVals map[string]bool, language string) (err error)

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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}
