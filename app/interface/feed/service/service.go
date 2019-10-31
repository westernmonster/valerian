package service

import (
	"context"
	"net/url"

	"valerian/app/interface/feed/conf"
	"valerian/app/interface/feed/dao"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	feed "valerian/app/service/feed/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/conf/env"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
		GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
		GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
		GetComment(c context.Context, id int64) (info *comment.CommentInfo, err error)
		GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
		GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
		GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)
		GetFeedPaged(c context.Context, accountID int64, limit, offset int) (info *feed.FeedResp, err error)
		IsFollowing(c context.Context, aid, targetAccountID int64) (IsFollowing bool, err error)

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

func genURL(path string, param url.Values) (uri string, err error) {
	u, err := url.Parse(env.SiteURL + path)
	if err != nil {
		return
	}
	u.RawQuery = param.Encode()

	return u.String(), nil
}
