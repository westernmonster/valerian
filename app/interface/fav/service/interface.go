package service

import (
	"context"

	"valerian/app/interface/fav/model"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetFavsPaged(c context.Context, node sqalx.Node, aid int64, targetType string, limit, offset int) (items []*model.Fav, err error)
	GetFavsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Fav, err error)
	GetFavs(c context.Context, node sqalx.Node) (items []*model.Fav, err error)
	GetFavByID(c context.Context, node sqalx.Node, id int64) (item *model.Fav, err error)
	GetFavByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Fav, err error)
	AddFav(c context.Context, node sqalx.Node, item *model.Fav) (err error)
	UpdateFav(c context.Context, node sqalx.Node, item *model.Fav) (err error)
	DelFav(c context.Context, node sqalx.Node, id int64) (err error)

	GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error)
	GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error)
	GetArticle(c context.Context, id int64) (info *article.ArticleInfo, err error)
	GetRevise(c context.Context, id int64) (info *article.ReviseInfo, err error)
	GetDiscussion(c context.Context, id int64) (info *discuss.DiscussionInfo, err error)
	GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
