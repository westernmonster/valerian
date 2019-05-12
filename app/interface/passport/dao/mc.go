package dao

import (
	"context"
	"go-common/library/log"
	"valerian/library/cache/memcache"
)

// pingMC ping memcache.
func (d *Dao) pingMC(c context.Context) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	if err = conn.Set(&memcache.Item{
		Key:        "ping",
		Value:      []byte{1},
		Expiration: d.mcExpire,
	}); err != nil {
		log.Error("conn.Set(ping, 1) error(%v)", err)
	}
	return
}
