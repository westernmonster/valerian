package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/article/conf"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	fav "valerian/app/service/fav/api"
	like "valerian/app/service/like/api"
	search "valerian/app/service/search/api"
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
	accountRPC account.AccountClient
	articleRPC article.ArticleClient
	topicRPC   topic.TopicClient
	likeRPC    like.LikeClient
	favRPC     fav.FavClient
	searchRPC  search.SearchClient
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

	if articleRPC, err := article.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial article service"))
	} else {
		dao.articleRPC = articleRPC
	}

	if topicRPC, err := topic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if likeRPC, err := like.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial like service"))
	} else {
		dao.likeRPC = likeRPC
	}

	if favRPC, err := fav.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial fav service"))
	} else {
		dao.favRPC = favRPC
	}

	if searchRPC, err := search.NewClient(c.SearchRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial search service"))
	} else {
		dao.searchRPC = searchRPC
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
