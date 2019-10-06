package service

import (
	"context"

	"valerian/app/service/relation/model"
	"valerian/library/database/sqalx"
)

type IDao interface {
	GetFollowingsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFollowing, err error)
	GetFollowings(c context.Context, node sqalx.Node) (items []*model.AccountFollowing, err error)
	GetFollowingByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFollowing, err error)
	GetFollowingByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFollowing, err error)
	AddFollowing(c context.Context, node sqalx.Node, item *model.AccountFollowing) (err error)
	UpdateFollowing(c context.Context, node sqalx.Node, item *model.AccountFollowing) (err error)
	DelFollowing(c context.Context, node sqalx.Node, aid, fid int64) (err error)
	SetFollowing(c context.Context, node sqalx.Node, attr uint32, aid, fid int64) (err error)

	GetFansListByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccountFans, err error)
	GetFansList(c context.Context, node sqalx.Node) (items []*model.AccountFans, err error)
	GetFansByID(c context.Context, node sqalx.Node, id int64) (item *model.AccountFans, err error)
	GetFansByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AccountFans, err error)
	AddFans(c context.Context, node sqalx.Node, item *model.AccountFans) (err error)
	UpdateFans(c context.Context, node sqalx.Node, item *model.AccountFans) (err error)
	DelFans(c context.Context, node sqalx.Node, aid, fid int64) (err error)
	SetFans(c context.Context, node sqalx.Node, attr uint32, aid, fid int64) (err error)

	GetFansPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountFans, err error)
	GetFollowingsPaged(c context.Context, node sqalx.Node, aid int64, limit, offset int) (items []*model.AccountFollowing, err error)

	GetAccountStatByID(c context.Context, node sqalx.Node, aid int64) (item *model.AccountStat, err error)
	IncrAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)
	AddAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)
	UpdateAccountStat(c context.Context, node sqalx.Node, item *model.AccountStat) (err error)

	SetFollowingsCache(c context.Context, aid int64, page, pageSize int, data []*model.AccountFollowing) (err error)
	FollowingsCache(c context.Context, aid int64, page, pageSize int) (data []*model.AccountFollowing, err error)
	DelFollowingsCache(c context.Context, aid int64) (err error)

	SetFansCache(c context.Context, aid int64, page, pageSize int, data []*model.AccountFans) (err error)
	FansCache(c context.Context, aid int64, page, pageSize int) (data []*model.AccountFans, err error)
	DelFansCache(c context.Context, aid int64) (err error)

	NotifyFollow(c context.Context, aid, fid int64) (err error)
	NotifyUnfollow(c context.Context, aid, fid int64) (err error)

	Ping(c context.Context) (err error)
	Close()
	DB() sqalx.Node
}
