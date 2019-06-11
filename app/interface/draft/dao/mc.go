package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

const (
	keyColors          = "colors"
	keyDraftCategories = "draft_categories"
)

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

func (p *Dao) SetColorsCache(c context.Context, items []*model.Color) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: keyColors, Object: items, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set colors cache error(%s,%d,%v)", keyColors, p.mcExpire, err))
	}
	return
}

func (p *Dao) ColorsCache(c context.Context) (res []*model.Color, err error) {
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(keyColors); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", keyColors, err))
		return
	}
	if err = conn.Scan(item, &res); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelColorsCache(c context.Context) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Delete(keyColors); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", keyColors, err))
		return
	}
	return
}

func (p *Dao) SetDraftCategoriesCache(c context.Context, items []*model.DraftCategory) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: keyDraftCategories, Object: items, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set draft categories cache error(%s,%d,%v)", keyDraftCategories, p.mcExpire, err))
	}
	return
}

func (p *Dao) DraftCategoriesCache(c context.Context) (res []*model.DraftCategory, err error) {
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(keyDraftCategories); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", keyDraftCategories, err))
		return
	}
	if err = conn.Scan(item, &res); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelDraftCategoriesCache(c context.Context) (err error) {
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Delete(keyDraftCategories); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", keyDraftCategories, err))
		return
	}
	return
}
