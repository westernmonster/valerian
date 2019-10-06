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

	GetDiscussionFilesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussionFile, err error)
	GetDiscussionFiles(c context.Context, node sqalx.Node) (items []*model.DiscussionFile, err error)
	GetDiscussionFileByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussionFile, err error)
	GetDiscussionFileByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussionFile, err error)
	AddDiscussionFile(c context.Context, node sqalx.Node, item *model.DiscussionFile) (err error)
	UpdateDiscussionFile(c context.Context, node sqalx.Node, item *model.DiscussionFile) (err error)
	DelDiscussionFile(c context.Context, node sqalx.Node, id int64) (err error)
	DelDiscussionFiles(c context.Context, node sqalx.Node, discussionID int64) (err error)

	IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)
	IncrTopicStat(c context.Context, node sqalx.Node, item *model.TopicStat) (err error)

	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetTopicMemberRole(c context.Context, topicID, accountID int64) (resp *topic.MemberRoleReply, err error)

	NotifyDiscussionAdded(c context.Context, id int64) (err error)
	NotifyDiscussionUpdated(c context.Context, id int64) (err error)
	NotifyDiscussionDeleted(c context.Context, id int64, topicID int64) (err error)

	SetDiscussionFilesCache(c context.Context, discussionID int64, m []*model.DiscussionFileResp) (err error)
	DiscussionFilesCache(c context.Context, discussionID int64) (m []*model.DiscussionFileResp, err error)
	DelDiscussionFilesCache(c context.Context, discussionID int64) (err error)

	GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error)
	AddDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error)
	IncrDiscussionStat(c context.Context, node sqalx.Node, item *model.DiscussionStat) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
