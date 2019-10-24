package dao

import (
	"context"
	"fmt"

	"valerian/app/service/certification/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func accountKey(aid int64) string {
	return fmt.Sprintf("account_%d", aid)
}

func (p *Dao) SetAccountCache(c context.Context, m *model.Account) (err error) {
	key := accountKey(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set account cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) AccountCache(c context.Context, accountID int64) (m *model.Account, err error) {
	key := accountKey(accountID)
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

	m = new(model.Account)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelAccountCache(c context.Context, accountID int64) (err error) {
	key := accountKey(accountID)
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
