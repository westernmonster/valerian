package dao

import (
	"context"
	"fmt"
	"valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func (p *Dao) SetArticleHistoryCache(c context.Context, articleID int64, m []*api.ArticleHistoryResp) (err error) {
	key := def.ArticleHistoryKey(articleID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set article histories cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) ArticleHistoryCache(c context.Context, articleID int64) (m []*api.ArticleHistoryResp, err error) {
	key := def.ArticleHistoryKey(articleID)
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", key, err))
		return
	}

	if err = conn.Scan(item, &m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelArticleHistoryCache(c context.Context, articleID int64) (err error) {
	key := def.ArticleHistoryKey(articleID)
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", key, err))
		return
	}
	return
}
