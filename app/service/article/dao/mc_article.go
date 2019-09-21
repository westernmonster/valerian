package dao

import (
	"context"
	"fmt"
	"valerian/app/service/article/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func articleKey(articleID int64) string {
	return fmt.Sprintf("article_%d", articleID)
}

func (p *Dao) SetArticleCache(c context.Context, m *model.ArticleResp) (err error) {
	key := articleKey(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set article cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) ArticleCache(c context.Context, articleID int64) (m *model.ArticleResp, err error) {
	key := articleKey(articleID)
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

	m = new(model.ArticleResp)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelArticleCache(c context.Context, articleID int64) (err error) {
	key := articleKey(articleID)
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
