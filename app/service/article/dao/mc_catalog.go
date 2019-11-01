package dao

import (
	"context"
	"fmt"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func topicCatalogKey(topicID int64) string {
	return fmt.Sprintf("t_catalog_%d", topicID)
}

func (p *Dao) DelTopicCatalogCache(c context.Context, topicID int64) (err error) {
	key := topicCatalogKey(topicID)
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
