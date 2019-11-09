package dao

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/topic/model"
	"valerian/library/cache/memcache"
	"valerian/library/log"

	uuid "github.com/satori/go.uuid"
)

type TopicMemberPagedData struct {
	Count int32                `json:"count"`
	Data  []*model.TopicMember `json:"data"`
}

func (p *Dao) setTopicMemberVersionCache(c context.Context, topicID int64, version string) (err error) {
	key := def.TopicMemberVersionKey(topicID)
	conn := p.mc.Get(c)
	defer conn.Close()

	item := &memcache.Item{Key: key, Value: []byte(version), Flags: memcache.FlagRAW, Expiration: p.mcExpire}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topic member version cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) topicMemberVersionCache(c context.Context, topicID int64) (version string, err error) {
	key := def.TopicMemberVersionKey(topicID)
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

func (p *Dao) SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int32, data []*model.TopicMember) (err error) {
	var version string
	if version, err = p.topicMemberVersionCache(c, topicID); err != nil {
		return
	} else if version == "" {
		version = uuid.NewV4().String()
		if err = p.setTopicMemberVersionCache(c, topicID, version); err != nil {
			return
		}
	}

	key := def.TopicMembersKey(topicID, page, pageSize, version)
	conn := p.mc.Get(c)
	defer conn.Close()

	m := &TopicMemberPagedData{
		Data:  data,
		Count: count,
	}

	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
	if err = conn.Set(item); err != nil {
		log.For(c).Error(fmt.Sprintf("set topicMembers cache error(%s,%d,%v)", key, p.mcExpire, err))
	}
	return
}

func (p *Dao) TopicMembersCache(c context.Context, topicID int64, page, pageSize int32) (count int32, data []*model.TopicMember, err error) {
	var version string
	if version, err = p.topicMemberVersionCache(c, topicID); err != nil {
		return
	} else if version == "" {
		version = uuid.NewV4().String()
		if err = p.setTopicMemberVersionCache(c, topicID, version); err != nil {
			return
		}
	}

	key := def.TopicMembersKey(topicID, page, pageSize, version)
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

	m := new(TopicMemberPagedData)
	if err = conn.Scan(item, m); err != nil {
		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
		return
	}

	count = m.Count
	data = m.Data
	return
}

func (p *Dao) DelTopicMembersCache(c context.Context, topicID int64) (err error) {
	key := def.TopicMemberVersionKey(topicID)
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
