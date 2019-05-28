package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/passport-login/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func vcMobileKey(vtype int, mobile string) string {
	return fmt.Sprintf("rc_%d_%s", vtype, mobile)
}

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

func (p *Dao) SetValcodeCache(c context.Context, m *model.AccessToken) (err error) {
	key := akKey(m.Token)
	conn := p.mc.Get(c)
	defer conn.Close()

	if m.ExpiresAt < 0 {
		log.For(c).Error(fmt.Sprintf("auth expire error(expires:%d)", m.ExpiresAt))
		return
	}

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagProtobuf, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set token cache error(%s,%d,%v)", key, m.ExpiresAt, err))
	}
	return
}

func (p *Dao) AccessTokenCache(c context.Context, token string) (res *model.AccessToken, err error) {
	key := akKey(token)
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
	res = new(model.AccessToken)
	if err = conn.Scan(item, res); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelTokenCache(c context.Context, token string) (err error) {
	key := akKey(token)
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
