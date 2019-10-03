package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/article/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func reviseFileKey(reviseID int64) string {
	return fmt.Sprintf("rev_files_%d", reviseID)
}

func (p *Dao) SetReviseFileCache(c context.Context, reviseID int64, m []*model.ReviseFileResp) (err error) {
	key := reviseFileKey(reviseID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set article file cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) ReviseFileCache(c context.Context, reviseID int64) (m []*model.ReviseFileResp, err error) {
	key := reviseFileKey(reviseID)
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

func (p *Dao) DelReviseFileCache(c context.Context, reviseID int64) (err error) {
	key := reviseFileKey(reviseID)
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
