package service

import (
	"context"

	"valerian/app/interface/discuss/model"
	account "valerian/app/service/account/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetUserDiscussionsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Discussion, err error)
	GetTopicDiscussionsPaged(c context.Context, node sqalx.Node, topicID, categoryID int64, limit, offset int) (items []*model.Discussion, err error)
	GetDiscussionsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Discussion, err error)
	GetDiscussions(c context.Context, node sqalx.Node) (items []*model.Discussion, err error)
	GetDiscussionByID(c context.Context, node sqalx.Node, id int64) (item *model.Discussion, err error)
	GetDiscussionByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Discussion, err error)
	AddDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error)
	UpdateDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error)
	DelDiscussion(c context.Context, node sqalx.Node, id int64) (err error)

	GetDiscussCategoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussCategory, err error)
	GetDiscussCategories(c context.Context, node sqalx.Node) (items []*model.DiscussCategory, err error)
	GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error)
	GetDiscussCategoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussCategory, err error)
	AddDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	UpdateDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	DelDiscussCategory(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountResStat, err error)
	AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error)
	UpdateAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error)
	IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountResStat) (err error)

	GetTopicStatByID(c context.Context, node sqalx.Node, topicID int64) (item *model.TopicResStat, err error)
	AddTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error)
	UpdateTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error)
	IncrTopicStat(c context.Context, node sqalx.Node, item *model.TopicResStat) (err error)

	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	CheckTopicManager(c context.Context, topicID, accountID int64) (err error)

	NotifyDiscussionAdded(c context.Context, id int64) (err error)
	NotifyDiscussionUpdated(c context.Context, id int64) (err error)
	NotifyDiscussionDeleted(c context.Context, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
