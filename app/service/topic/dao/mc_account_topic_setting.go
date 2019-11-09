package dao

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

func (p *Dao) SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error) {
	key := def.AccountTopicSettingKey(m.AccountID, m.TopicID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topic cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error) {
	key := def.AccountTopicSettingKey(aid, topicID)
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

	m = new(model.AccountTopicSetting)
	if err = conn.Scan(item, &m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error) {
	key := def.AccountTopicSettingKey(aid, topicID)
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
