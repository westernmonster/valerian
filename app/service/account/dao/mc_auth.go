package dao

import (
	"context"
	"fmt"
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
