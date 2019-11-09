package dao

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

const (
	mobileExpires = 60 * 5
	emailExpires  = 60 * 10
)

// pingMC ping memcache.
func (p *Dao) pingMC(c context.Context) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Set(&memcache.Item{
		Key:        "ping",
		Value:      []byte{1},
		Expiration: p.mcExpire,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.pingMC error(%+v)", err))
	}
	return
}

func (p *Dao) SetMobileValcodeCache(c context.Context, vtype int32, mobile, code string) (err error) {
	key := def.MobileValcodeKey(vtype, mobile)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Value: []byte(code), Flags: memcache.FlagRAW, Expiration: int32(mobileExpires)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set valcode cache error(%s,%d,%v)", key, mobileExpires, err))
	}
	return
}

func (p *Dao) MobileValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error) {
	key := def.MobileValcodeKey(vtype, mobile)
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

	if err = conn.Scan(item, &code); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelMobileCache(c context.Context, vtype int32, mobile string) (err error) {
	key := def.MobileValcodeKey(vtype, mobile)
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

func (p *Dao) SetEmailValcodeCache(c context.Context, vtype int32, mobile, code string) (err error) {
	key := def.EmailValcodeKey(vtype, mobile)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Value: []byte(code), Flags: memcache.FlagRAW, Expiration: int32(emailExpires)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set valcode cache error(%s,%d,%v)", key, emailExpires, err))
	}
	return
}

func (p *Dao) EmailValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error) {
	key := def.EmailValcodeKey(vtype, mobile)
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

	if err = conn.Scan(item, &code); err != nil {

		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelEmailCache(c context.Context, vtype int32, mobile string) (err error) {
	key := def.EmailValcodeKey(vtype, mobile)
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
