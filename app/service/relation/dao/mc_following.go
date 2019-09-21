package dao

import (
	"context"
	"fmt"
	"valerian/app/service/relation/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"

	uuid "github.com/satori/go.uuid"
)

type FollowingPagedData struct {
	Data []*model.AccountFollowing `json:"data"`
}

func followingsKey(aid int64, page, pageSize int, version string) string {
	return fmt.Sprintf("fll_%d_%d_%d_%s", aid, page, pageSize, version)
}

func followingVersionKey(aid int64) string {
	return fmt.Sprintf("fllv_%d", aid)
}

func (p *Dao) setFollowingVersionCache(c context.Context, aid int64, version string) (err error) {
	key := followingVersionKey(aid)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Value: []byte(version), Flags: memcache.FlagRAW, Expiration: p.mcExpire}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topic member version cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) followingVersionCache(c context.Context, aid int64) (version string, err error) {
	key := followingVersionKey(aid)
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

	if err = conn.Scan(item, &version); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) SetFollowingsCache(c context.Context, aid int64, page, pageSize int, data []*model.AccountFollowing) (err error) {
	var version string
	if version, err = p.followingVersionCache(c, aid); err != nil {
		return
	} else if version == "" {
		version = uuid.NewV4().String()
		if err = p.setFollowingVersionCache(c, aid, version); err != nil {
			return
		}
	}

	key := followingsKey(aid, page, pageSize, version)
	conn := p.mc.Get(c)
	defer conn.Close()

	m := &FollowingPagedData{
		Data: data,
	}

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set followings cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) FollowingsCache(c context.Context, aid int64, page, pageSize int) (data []*model.AccountFollowing, err error) {
	var version string
	if version, err = p.followingVersionCache(c, aid); err != nil {
		return
	} else if version == "" {
		version = uuid.NewV4().String()
		if err = p.setFollowingVersionCache(c, aid, version); err != nil {
			return
		}
	}

	key := followingsKey(aid, page, pageSize, version)
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

	m := new(FollowingPagedData)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
		return
	}

	data = m.Data
	return
}

func (p *Dao) DelFollowingsCache(c context.Context, aid int64) (err error) {
	key := followingVersionKey(aid)
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
