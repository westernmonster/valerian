package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/interface/topic/conf"
	account "valerian/app/service/account/api"
	article "valerian/app/service/article/api"
	discuss "valerian/app/service/discuss/api"
	fav "valerian/app/service/fav/api"
	relation "valerian/app/service/relation/api"
	search "valerian/app/service/search/api"
	topicFeed "valerian/app/service/topic-feed/api"
	stopic "valerian/app/service/topic/api"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v6"
)

// Dao dao struct
type Dao struct {
	db           sqalx.Node
	mc           *memcache.Pool
	mcExpire     int32
	c            *conf.Config
	accountRPC   account.AccountClient
	topicFeedRPC topicFeed.TopicFeedClient
	discussRPC   discuss.DiscussionClient
	articleRPC   article.ArticleClient
	topicRPC     stopic.TopicClient
	relationRPC  relation.RelationClient
	searchRPC    search.SearchClient
	favRPC       fav.FavClient

	esClient *elastic.Client
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
	}

	if relationRPC, err := relation.NewClient(c.RelationRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial relation service"))
	} else {
		dao.relationRPC = relationRPC
	}

	if favRPC, err := fav.NewClient(c.FavRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial fav service"))
	} else {
		dao.favRPC = favRPC
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

	if topicFeedRPC, err := topicFeed.NewClient(c.TopicFeedRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic feed service"))
	} else {
		dao.topicFeedRPC = topicFeedRPC
	}

	if topicRPC, err := stopic.NewClient(c.TopicRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial topic service"))
	} else {
		dao.topicRPC = topicRPC
	}

	if searchRPC, err := search.NewClient(c.SearchRPC); err != nil {
		panic(errors.WithMessage(err, "Failed to dial search service"))
	} else {
		dao.searchRPC = searchRPC
	}

	if client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("elastic", "^EIj7UIjd"),
		elastic.SetURL(c.Es.Addr...),
	); err == nil {
		dao.esClient = client
	} else {
		PromError(context.TODO(), "es:集群连接失败", "cluster:  %v", err)
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

// pingESCluster ping es cluster
func (d *Dao) pingESCluster(ctx context.Context) (err error) {
	_, _, err = d.esClient.Ping(d.c.Es.Addr[0]).Do(ctx)
	if err != nil {
		PromError(ctx, "archiveESClient:Ping", "dao.pingESCluster error(%v) ", err)
		return
	}
	return
}
