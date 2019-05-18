package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/passport-login/model"
	"valerian/library/cache/memcache"
)

func akKey(token string) string {
	return fmt.Sprintf("ak_%s", token)
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
		p.logger.For(c).Error(fmt.Sprintf("dao.pingMC error(%+v)", err))
	}
	return
}

func (p *Dao) RefreshTokenCache(c context.Context, sd string) (item *model.OauthRefreshToken, err error) {
	key := akKey(sd)
	conn := p.mc.Get(c)
	defer conn.Close()
	r, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		p.logger.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", key, err))
		return
	}
	item = new(model.OauthRefreshToken)
	if err = conn.Scan(r, item); err != nil {
		p.logger.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(r.Value), err))
	}
	return
}
