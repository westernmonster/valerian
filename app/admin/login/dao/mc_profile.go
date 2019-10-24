package dao

import (
	"context"
	"fmt"
	"valerian/app/admin/login/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func profileKey(aid int64) string {
	return fmt.Sprintf("profile_%d", aid)
}

func (p *Dao) SetProfileCache(c context.Context, m *model.Profile) (err error) {
	key := profileKey(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set profile cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) ProfileCache(c context.Context, id int64) (m *model.Profile, err error) {
	key := profileKey(id)
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

	m = new(model.Profile)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelProfileCache(c context.Context, id int64) (err error) {
	key := profileKey(id)
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
