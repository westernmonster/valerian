package dao

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/app/service/topic/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func (p *Dao) SetAccountCache(c context.Context, m *model.Account) (err error) {
	key := def.AccountKey(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set account cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) AccountCache(c context.Context, accountID int64) (m *model.Account, err error) {
	key := def.AccountKey(accountID)
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

	m = new(model.Account)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelAccountCache(c context.Context, accountID int64) (err error) {
	key := def.AccountKey(accountID)
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

func (d *Dao) BatchAccountCache(c context.Context, aids []int64) (cached map[int64]*model.Account, missed []int64, err error) {
	cached = make(map[int64]*model.Account, len(aids))
	if len(aids) == 0 {
		return
	}
	keys := make([]string, 0, len(aids))
	aidmap := make(map[string]int64, len(aids))
	for _, aid := range aids {
		k := def.AccountKey(aid)
		keys = append(keys, k)
		aidmap[k] = aid
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	bases, err := conn.GetMulti(keys)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Gets(%v) error(%v)", keys, err))
		return
	}
	for _, base := range bases {
		b := &model.Account{}
		if err = conn.Scan(base, b); err != nil {
			log.For(c).Error(fmt.Sprintf("json.Unmarshal(%v) error(%v)", base.Value, err))
			return
		}
		cached[aidmap[base.Key]] = b
		delete(aidmap, base.Key)
	}
	missed = make([]int64, 0, len(aidmap))
	for _, bid := range aidmap {
		missed = append(missed, bid)
	}
	return

}

func (d *Dao) SetBatchAccountCache(c context.Context, bs []*model.Account) (err error) {
	for _, info := range bs {
		d.SetAccountCache(c, info)
	}
	return
}
