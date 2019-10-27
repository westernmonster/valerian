package service

import (
	"context"

	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/search/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	CreateAccountIndices(c context.Context) (err error)
	CreateArticleIndices(c context.Context) (err error)
	CreateDiscussionIndices(c context.Context) (err error)
	CreateTopicIndices(c context.Context) (err error)

	BulkAccount2ES(c context.Context, items []*model.ESAccount) (err error)
	PutAccount2ES(c context.Context, item *model.ESAccount) (err error)

	PutArticle2ES(c context.Context, item *model.ESArticle) (err error)
	BulkArticle2ES(c context.Context, items []*model.ESArticle) (err error)

	PutDiscussion2ES(c context.Context, item *model.ESDiscussion) (err error)
	PutTopic2ES(c context.Context, item *model.ESTopic) (err error)

	DelESDiscussion(c context.Context, id int64) (err error)
	DelESTopic(c context.Context, id int64) (err error)
	DelESArticle(c context.Context, id int64) (err error)
	DelESAccount(c context.Context, id int64) (err error)

	GetAllArticles(c context.Context) (items []*article.DBArticle, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)

	GetAllAccounts(c context.Context) (items []*account.DBAccount, err error)
	GetAccountInfo(c context.Context, aid int64) (info *account.DBAccount, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
	GetSelfProfile(c context.Context, aid int64) (info *account.SelfProfile, err error)
	GetAccountStat(c context.Context, aid int64) (info *account.AccountStatInfo, err error)

	GetAllDiscussions(c context.Context) (items []*discuss.DiscussionInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)

	GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)
	GetAllTopics(c context.Context) (items []*topic.TopicInfo, err error)

	AccountSearch(c context.Context, arg *model.AccountSearchParams) (res *model.SearchResult, err error)
	TopicSearch(c context.Context, arg *model.TopicSearchParams) (res *model.SearchResult, err error)
	ArticleSearch(c context.Context, arg *model.ArticleSearchParams) (res *model.SearchResult, err error)
	DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (res *model.SearchResult, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
