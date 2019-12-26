package dao

import (
	"context"
	"fmt"
	"strconv"
	"valerian/app/service/feed/def"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

const (
	sessionExpires = 60 * 5
)

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

func (p *Dao) SetSessionResetPasswordCache(c context.Context, sessionID string, accountID int64) (err error) {
	key := def.ResetPasswordKey(sessionID)
	conn := p.authMC.Get(c)
	defer conn.Close()

	aid := strconv.FormatInt(accountID, 10)
	item := &memcache.Item{Key: key, Value: []byte(aid), Flags: memcache.FlagRAW, Expiration: int32(sessionExpires)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set session cache error(%s,%d,%v)", key, sessionExpires, err))
	}
	return
}

func (p *Dao) SessionResetPasswordCache(c context.Context, sessionID string) (aid int64, err error) {
	key := def.ResetPasswordKey(sessionID)
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

	var idStr string
	if err = conn.Scan(item, &idStr); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}

	if aid, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		log.For(c).Error(fmt.Sprintf("ParseInt(%v) error(%v)", idStr, err))
	}

	return
}

func (p *Dao) DelResetPasswordCache(c context.Context, sessionID string) (err error) {
	key := def.ResetPasswordKey(sessionID)
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

func (p *Dao) MobileValcodeCache(c context.Context, vtype int32, mobile string) (code string, err error) {
	key := def.MobileValcodeKey(vtype, mobile)
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

func (p *Dao) DelMobileCache(c context.Context, vtype int32, mobile string) (err error) {
	key := def.MobileValcodeKey(vtype, mobile)
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
	key := def.EmailValcodeKey(vtype, mobile)
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

func (p *Dao) DelEmailCache(c context.Context, vtype int32, mobile string) (err error) {
	key := def.EmailValcodeKey(vtype, mobile)
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

func (p *Dao) DelAccessTokenCache(c context.Context, token string) (err error) {
	key := def.RefreshTokenKey(token)
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
