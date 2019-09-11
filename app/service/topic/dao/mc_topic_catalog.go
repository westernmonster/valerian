package dao

// func topicCatalogKey(topicID int64) string {
// 	return fmt.Sprintf("t_catalog_%d", topicID)
// }

// func (p *Dao) SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error) {
// 	key := topicCatalogKey(topicID)
// 	conn := p.mc.Get(c)
// 	defer conn.Close()

// 	item := &memcache.Item{Key: key, Object: m, Flags: memcache.FlagJSON, Expiration: int32(p.mcExpire)}
// 	if err = conn.Set(item); err != nil {
// 		log.For(c).Error(fmt.Sprintf("set topic cache error(%s,%d,%v)", key, p.mcExpire, err))
// 	}
// 	return
// }

// func (p *Dao) TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error) {
// 	key := topicCatalogKey(topicID)
// 	conn := p.mc.Get(c)
// 	defer conn.Close()
// 	var item *memcache.Item
// 	if item, err = conn.Get(key); err != nil {
// 		if err == memcache.ErrNotFound {
// 			err = nil
// 			return
// 		}
// 		log.For(c).Error(fmt.Sprintf("conn.Get(%s) error(%v)", key, err))
// 		return
// 	}

// 	if err = conn.Scan(item, &m); err != nil {
// 		log.For(c).Error(fmt.Sprintf("conn.Scan(%v) error(%v)", string(item.Value), err))
// 	}
// 	return
// }

// func (p *Dao) DelTopicCatalogCache(c context.Context, topicID int64) (err error) {
// 	key := topicCatalogKey(topicID)
// 	conn := p.mc.Get(c)
// 	defer conn.Close()
// 	if err = conn.Delete(key); err != nil {
// 		if err == memcache.ErrNotFound {
// 			err = nil
// 			return
// 		}
// 		log.For(c).Error(fmt.Sprintf("conn.Delete(%s) error(%v)", key, err))
// 		return
// 	}
// 	return
// }
