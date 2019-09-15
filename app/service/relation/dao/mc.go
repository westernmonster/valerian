package dao

import (
	"context"
	"fmt"
	"strconv"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func fansKey(mid int64) string {
	return "fans_" + strconv.FormatInt(mid, 10)
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
