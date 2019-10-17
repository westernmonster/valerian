package service

import (
	"context"

	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	"valerian/app/service/message/model"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetMessagesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Message, err error)
	GetMessages(c context.Context, node sqalx.Node) (items []*model.Message, err error)
	GetMessageByID(c context.Context, node sqalx.Node, id int64) (item *model.Message, err error)
	GetMessageByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Message, err error)
	AddMessage(c context.Context, node sqalx.Node, item *model.Message) (err error)
	UpdateMessage(c context.Context, node sqalx.Node, item *model.Message) (err error)
	DelMessage(c context.Context, node sqalx.Node, id int64) (err error)

	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (info *topic.TopicInfo, err error)
	GetFansIDs(c context.Context, aid int64) (resp *relation.IDsResp, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetBelongsTopicIDs(c context.Context, aid int64) (resp *topic.IDsResp, err error)
	GetTopicMemberIDs(c context.Context, topicID int64) (resp *topic.IDsResp, err error)
	GetComment(c context.Context, id int64) (info *comment.CommentInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
