package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func keyColors(aid int64) (key string) {
	return fmt.Sprintf("colors_%d", aid)
}

func (p *Dao) SetColorsCache(c context.Context, aid int64, items []*model.Color) (err error) {
	key := keyColors(aid)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: items, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set colors cache error(%s,%d,%v)", keyColors, p.mcExpire, err))
	}
	return
}

func (p *Dao) ColorsCache(c context.Context, aid int64) (res []*model.Color, err error) {
	key := keyColors(aid)
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", keyColors, err))
		return
	}
	if err = conn.Scan(item, &res); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelColorsCache(c context.Context, aid int64) (err error) {
	key := keyColors(aid)
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", keyColors, err))
		return
	}
	return
}
