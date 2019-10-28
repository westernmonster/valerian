package dao

import (
	"context"
	"fmt"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func akKey(token string) string {
	return fmt.Sprintf("ak_%s", token)
}

func vcMobileKey(vtype int32, mobile string) string {
	return fmt.Sprintf("rc_%d_%s", vtype, mobile)
}

func vcEmailKey(vtype int32, email string) string {
	return fmt.Sprintf("rc_%d_%s", vtype, email)
}

// pingMC ping memcache.
func (p *Dao) pingAuthMC(c context.Context) (err error) {
	conn := p.authMC.Get(c)
	defer conn.Close()
	if err = conn.Set(&memcache.Item{
		Key:        "ping",
		Value:      []byte{1},
		Expiration: p.authMCExpire,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.pingMC error(%+v)", err))
	}
	return
}

func (p *Dao) MobileValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error) {
	key := vcMobileKey(vtype, mobile)
	conn := p.authMC.Get(c)
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

func (p *Dao) DelMobileValcodeCache(c context.Context, vtype int32, mobile string) (err error) {
	key := vcMobileKey(vtype, mobile)
	conn := p.authMC.Get(c)
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

func (p *Dao) EmailValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error) {
	key := vcEmailKey(vtype, mobile)
	conn := p.authMC.Get(c)
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

func (p *Dao) DelEmailValcideCache(c context.Context, vtype int32, mobile string) (err error) {
	key := vcEmailKey(vtype, mobile)
	conn := p.authMC.Get(c)
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
