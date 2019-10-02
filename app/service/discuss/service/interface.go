package service

import (
	"context"
	"valerian/app/service/discuss/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetUserDiscussionsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.Discussion, err error)

	GetDiscussionsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Discussion, err error)
	GetDiscussions(c context.Context, node sqalx.Node) (items []*model.Discussion, err error)
	GetDiscussionByID(c context.Context, node sqalx.Node, id int64) (item *model.Discussion, err error)
	GetDiscussionByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Discussion, err error)
	AddDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error)
	UpdateDiscussion(c context.Context, node sqalx.Node, item *model.Discussion) (err error)
	DelDiscussion(c context.Context, node sqalx.Node, id int64) (err error)
	GetDiscussionStatByID(c context.Context, node sqalx.Node, discussionID int64) (item *model.DiscussionStat, err error)

	GetDiscussCategoriesByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.DiscussCategory, err error)
	GetDiscussCategories(c context.Context, node sqalx.Node) (items []*model.DiscussCategory, err error)
	GetDiscussCategoryByID(c context.Context, node sqalx.Node, id int64) (item *model.DiscussCategory, err error)
	GetDiscussCategoryByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.DiscussCategory, err error)
	AddDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	UpdateDiscussCategory(c context.Context, node sqalx.Node, item *model.DiscussCategory) (err error)
	DelDiscussCategory(c context.Context, node sqalx.Node, id int64) (err error)

	GetImageUrlsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.ImageURL, err error)
	GetImageUrls(c context.Context, node sqalx.Node) (items []*model.ImageURL, err error)
	GetImageURLByID(c context.Context, node sqalx.Node, id int64) (item *model.ImageURL, err error)
	GetImageURLByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.ImageURL, err error)
	AddImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error)
	UpdateImageURL(c context.Context, node sqalx.Node, item *model.ImageURL) (err error)
	DelImageURL(c context.Context, node sqalx.Node, id int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
