package dao

import (
	"context"
	"fmt"
	"valerian/library/cache/memcache"
	"valerian/library/log"
	"valerian/library/net/http/mars/middleware/permit"
)

const (
	_sessionlen  = 32
	_sessionLife = 2592000
	_dsbCaller   = "manager-go"
)

func sessionKey(sid string) string {
	return fmt.Sprintf("sess_%d", sid)
}

func (d *Dao) Session(ctx context.Context, sid string) (res *permit.Session, err error) {
	conn := d.mc.Get(ctx)
	defer conn.Close()
	r, err := conn.Get(sessionKey(sid))
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(ctx).Error(fmt.Sprintf("conn.Get(%s) error(%v)", sid, err))
		return
	}
	res = &permit.Session{}
	if err = conn.Scan(r, res); err != nil {
		log.For(ctx).Error(fmt.Sprintf("conn.Scan(%s) error(%v)", string(r.Value), err))
	}
	return
}

func (d *Dao) SetSession(ctx context.Context, p *permit.Session) (err error) {
	conn := d.mc.Get(ctx)
	defer conn.Close()
	item := &memcache.Item{
		Key:        sessionKey(p.Sid),
		Object:     p,
		Flags:      memcache.FlagJSON,
		Expiration: int32(_sessionLife),
	}
	err = conn.Set(item)
	return
}
