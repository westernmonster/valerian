package dao

import (
	"context"
	"fmt"
	"time"

	"valerian/app/conf"
	"valerian/library/cache/memcache"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/stat/prom"

	"gopkg.in/olivere/elastic.v6"
)

// Dao dao struct
type Dao struct {
	db       sqalx.Node
	mc       *memcache.Pool
	mcExpire int32
	c        *conf.Config

	// esPool
	esPool map[string]*elastic.Client
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:        c,
		db:       sqalx.NewMySQL(c.DB.Main),
		mc:       memcache.NewPool(c.Memcache.Main.Config),
		mcExpire: int32(time.Duration(c.Memcache.Main.Expire) / time.Second),
	}
	dao.esPool = newEsPool(c, dao)
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

// newEsCluster cluster action
func newEsPool(c *conf.Config, d *Dao) (esCluster map[string]*elastic.Client) {
	esCluster = make(map[string]*elastic.Client)
	for esName, e := range c.Es {
		if client, err := elastic.NewClient(elastic.SetURL(e.Addr...),
			elastic.SetSniff(false),
		); err == nil {
			esCluster[esName] = client
		} else {
			fmt.Println(esName)
			fmt.Println(e.Addr)
			PromError(context.TODO(), "es:集群连接失败", "cluster: %s, %v", esName, err)
		}
	}
	return
}

// pingESCluster ping es cluster
func (d *Dao) pingESCluster(ctx context.Context) (err error) {
	for name := range d.c.Es {
		client, ok := d.esPool[name]
		if !ok {
			continue
		}
		_, _, err = client.Ping(d.c.Es["replyExternal"].Addr[0]).Do(ctx)
		if err != nil {
			PromError(ctx, "archiveESClient:Ping", "dao.pingESCluster error(%v) ", err)
			return
		}
	}
	return
}
