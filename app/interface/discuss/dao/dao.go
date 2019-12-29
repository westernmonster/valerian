package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/discuss/conf"
	account "valerian/app/service/account/api"
	discuss "valerian/app/service/discuss/api"
	fav "valerian/app/service/fav/api"
	like "valerian/app/service/like/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
	"valerian/library/stat/prom"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	c             *conf.Config
	accountRPC    account.AccountClient
	topicRPC      topic.TopicClient
	likeRPC       like.LikeClient
	favRPC        fav.FavClient
	discussionRPC discuss.DiscussionClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
	}

	if topicRPC, err := topic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if likeRPC, err := like.NewClient(c.LikeRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial like service"))
	} else {
		dao.likeRPC = likeRPC
	}

	if favRPC, err := fav.NewClient(c.FavRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial fav service"))
	} else {
		dao.favRPC = favRPC
	}

	if discussionRPC, err := discuss.NewClient(c.DiscussionRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial discuss service"))
	} else {
		dao.discussionRPC = discussionRPC
	}

	return
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
}

// PromError prometheus error count.
func PromError(c context.Context, name, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.For(c).Error(fmt.Sprintf(format, args...))
}
