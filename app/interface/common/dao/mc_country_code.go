package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/common/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

const (
	key = "country_codes"
)

func (p *Dao) SetCountryCodesCache(c context.Context, items []*model.CountryCode) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: items, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set token cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) CountryCodesCache(c context.Context) (res []*model.CountryCode, err error) {
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
	if err = conn.Scan(item, &res); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelCountryCodesCache(c context.Context, token string) (err error) {
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
