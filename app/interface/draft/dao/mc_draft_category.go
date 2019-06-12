package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/draft/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func keyDraftCategories(aid int64) (key string) {
	return fmt.Sprintf("dcates_%d", aid)
}

func keyDraftCategory(aid int64) (key string) {
	return fmt.Sprintf("dcate_%d", aid)
}

func (p *Dao) SetDraftCategoriesCache(c context.Context, aid int64, items []*model.DraftCategoryResp) (err error) {
	keyCategories := keyDraftCategories(aid)
	conn := p.mc.Get(c)
	defer conn.Close()

	ids := make([]int64, 0)
	for _, v := range items {
		ids = append(ids, v.ID)
		if err = p.SetDraftCategoryCache(c, v); err != nil {
			return
		}
	}

	item := &memcache.Item{Key: keyCategories, Object: ids, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set draft categories cache error(%s,%d,%v)", keyDraftCategories, p.mcExpire, err))
	}
	return
}

func (p *Dao) DraftCategoriesCache(c context.Context, aid int64) (res []*model.DraftCategoryResp, err error) {
	key := keyDraftCategories(aid)
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", keyDraftCategories, err))
		return
	}

	var ids []int64
	if err = conn.Scan(item, &ids); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
		return
	}

	res = make([]*model.DraftCategoryResp, 0)
	for _, v := range ids {
		var item *model.DraftCategoryResp
		if item, err = p.DraftCategoryCache(c, v); err != nil {
			return
		} else if item == nil {
			res = nil
			return
		} else {
			res = append(res, item)
		}

	}
	return
}

func (p *Dao) DelDraftCategoriesCache(c context.Context, aid int64) (err error) {
	key := keyDraftCategories(aid)
	conn := p.mc.Get(c)
	defer conn.Close()
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", keyDraftCategories, err))
		return
	}
	return
}

func (p *Dao) SetDraftCategoryCache(c context.Context, m *model.DraftCategoryResp) (err error) {
	key := keyDraftCategory(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topic cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) DraftCategoryCache(c context.Context, id int64) (m *model.DraftCategoryResp, err error) {
	key := keyDraftCategory(id)
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

	m = new(model.DraftCategoryResp)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelDrafCategoryCache(c context.Context, id int64) (err error) {
	key := keyDraftCategory(id)
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
