package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/common/conf"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	comment "valerian/app/service/comment/api"
	discuss "valerian/app/service/discuss/api"
	like "valerian/app/service/like/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"

	"github.com/pkg/errors"
)

// Dao dao struct
type Dao struct {
	db         sqalx.Node
	mc         *memcache.Pool
	mcExpire   int32
	c          *conf.Config
	likeRPC    like.LikeClient
	accountRPC account.AccountClient
	discussRPC discuss.DiscussionClient
	articleRPC article.ArticleClient
	topicRPC   topic.TopicClient
	commentRPC comment.CommentClient
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
	}

	if accountRPC, err := account.NewClient(c.AccountRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial account service"))
	} else {
		dao.accountRPC = accountRPC
	}
	if discussRPC, err := discuss.NewClient(c.DiscussRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial discuss service"))
	} else {
		dao.discussRPC = discussRPC
	}

	if articleRPC, err := article.NewClient(c.ArticleRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial article service"))
	} else {
		dao.articleRPC = articleRPC
	}

	if topicRPC, err := topic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if commentRPC, err := comment.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial comment service"))
	} else {
		dao.commentRPC = commentRPC
	}

	if likeRPC, err := like.NewClient(c.LikeRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial like service"))
	} else {
		dao.likeRPC = likeRPC
	}

	return
}

func (d *Dao) DB() sqalx.Node {
	return d.db
}

// Ping check db and mc health.
func (d *Dao) Ping(c context.Context) (err error) {
	if err = d.db.Ping(c); err != nil {
		log.Info(fmt.Sprintf("dao.db.Ping() error(%v)", err))
	}
	if err = d.pingMC(c); err != nil {
		log.Info(fmt.Sprintf("dao.mc.Ping() error(%v)", err))
		return
	}
	return
}

// Close close connections of mc, redis, db.
func (d *Dao) Close() {
	if d.mc != nil {
		d.mc.Close()
	}
	if d.db != nil {
		d.db.Close()
	}
}

// PromError prometheus error count.
func PromError(c context.Context, name, format string, args ...interface{}) {
	prom.BusinessErrCount.Incr(name)
	log.For(c).Error(fmt.Sprintf(format, args...))
}
