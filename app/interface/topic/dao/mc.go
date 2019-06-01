package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"
)

type TopicMemberPagedData struct {
	Count int                      `json:"count"`
	Data  []*model.TopicMemberResp `json:"data"`
}

func topicKey(topicID int64) string {
	return fmt.Sprintf("t_%d", topicID)
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

func (p *Dao) SetTopicCache(c context.Context, m *model.TopicResp) (err error) {
	key := topicKey(m.ID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topic cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) TopicCache(c context.Context, topicID int64) (m *model.TopicResp, err error) {
	key := topicKey(topicID)
	conn := p.mc.Get(c)
	defer conn.Close()
	var item *memcache.Item
	if item, err = conn.Get(key); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", key, err))
		return
	}

	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
	}
	return
}

func (p *Dao) DelTopicCache(c context.Context, topicID int64) (err error) {
	key := topicKey(topicID)
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
