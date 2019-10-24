package dao

import (
	"context"
	"fmt"

	"valerian/app/service/certification/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func workcertKey(aid int64) string {
	return fmt.Sprintf("workcert_%d", aid)
}

func (p *Dao) SetWorkCertCache(c context.Context, m *model.WorkCertification) (err error) {
	key := workcertKey(m.AccountID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set account cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) WorkCertCache(c context.Context, aid int64) (m *model.WorkCertification, err error) {
	key := workcertKey(aid)
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

	m = new(model.WorkCertification)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelWorkCertCache(c context.Context, aid int64) (err error) {
	key := workcertKey(aid)
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
