package dao

import (
	"context"
	"fmt"
	"valerian/library/cache/redis"
	"valerian/library/log"
)

const (
	_preLock        = "lk_"
	_preExpireCount = "ec_"

	_cronLock       = "cron_lock"
	_cronLockExpire = 2 * 3600
)

func lockKey(key string) string {
	return _preLock + key
}

//SetNxLock redis lock.
func (p *Dao) SetNxLock(c context.Context, k string, times int64) (res bool, err error) {
	var (
		key = lockKey(k)
	)

	var conn redis.Conn
	if conn, err = p.redis.GetContext(c); err != nil {
		return
	}

	defer conn.Close()
	if res, err = redis.Bool(conn.Do("SETNX", key, "1")); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Do(SETNX(%d)) error(%v)", key, err))
		return
	}
	//spew.Dump(res, err )
	if res {
		if _, err = redis.Bool(conn.Do("EXPIRE", key, times)); err != nil {
			log.For(c).Error(fmt.Sprintf("conn.Do(EXPIRE, %s, %d) error(%v)", key, times, err))
			return
		}
		log.For(c).Info(fmt.Sprintf("conn.Do(EXPIRE, %s, %d) ", key, times))
	}
	return
}

// PaLock redis定时任务锁，保证只有一个地方在执行定时任务
func (p *Dao) PaLock(c context.Context, key string) (incr int, err error) {
	var (
		r interface{}
	)

	var conn redis.Conn
	if conn, err = p.redis.GetContext(c); err != nil {
		return
	}

	key = _cronLock + key
	defer conn.Close()
	if err = conn.Send("INCRBY", key, 1); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Send(INCRBY,%s) error(%v)", key, err))
		return
	}
	if err = conn.Send("EXPIRE", key, _cronLockExpire); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Send(EXPIRE key(%s) expire(%d)) error(%v)", key, _cronLockExpire, err))
		return
	}
	if err = conn.Flush(); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Flush error(%v)", err))
		return
	}
	if r, err = conn.Receive(); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Receive error(%v)", err))
		return
	}
	if incr, err = redis.Int(r, err); err != nil {
		log.For(c).Error(fmt.Sprintf("redis.Int error(%v)", err))
		return
	}
	if _, err = conn.Receive(); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Receive error(%v)", err))
	}
	return
}
