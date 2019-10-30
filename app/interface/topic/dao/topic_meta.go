package dao

import (
	"context"
	"fmt"

	"valerian/app/service/topic/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetTopicMeta(c context.Context, aid, topicID int64) (info *topic.TopicMetaInfo, err error) {
	if info, err = p.topicRPC.GetTopicMeta(c, &topic.TopicMetaReq{AccountID: aid, TopicID: topicID}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMeta err(%+v)", err))
	}
	return
}

func (p *Dao) CreateTopic(c context.Context, arg *topic.ArgCreateTopic) (id int64, err error) {
	var idRet *topic.IDResp
	if idRet, err = p.topicRPC.CreateTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.CreateTopic err(%+v) arg(%+v)", err, arg))
		return
	}
	return idRet.ID, nil
}

func (p *Dao) UpdateTopic(c context.Context, arg *topic.ArgUpdateTopic) (err error) {
	if _, err = p.topicRPC.UpdateTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) DelTopic(c context.Context, arg *topic.IDReq) (err error) {
	if _, err = p.topicRPC.DelTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetTopicResp(c context.Context, arg *topic.IDReq) (resp *topic.TopicResp, err error) {
	if resp, err = p.topicRPC.GetTopicResp(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicResp err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) ChangeOwner(c context.Context, arg *topic.ArgChangeOwner) (err error) {
	if _, err = p.topicRPC.ChangeOwner(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.ChangeOwner err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SaveAuthTopics(c context.Context, arg *topic.ArgSaveAuthTopics) (err error) {
	if _, err = p.topicRPC.SaveAuthTopics(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SaveAuthTopics err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetAuthTopics(c context.Context, arg *api.IDReq) (resp *api.AuthTopicsResp, err error) {
	if resp, err = p.topicRPC.GetAuthTopics(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopics err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) FollowTopic(c context.Context, arg *api.ArgTopicFollow) (resp *api.StatusResp, err error) {
	if resp, err = p.topicRPC.FollowTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAuthTopics err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) AuditFollow(c context.Context, arg *api.ArgAuditFollow) (err error) {
	if _, err = p.topicRPC.AuditFollow(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AuditFollow err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetCatalogTaxonomiesHierarchy(c context.Context, arg *api.IDReq) (resp *api.CatalogsResp, err error) {
	if resp, err = p.topicRPC.GetCatalogTaxonomiesHierarchy(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCatalogTaxonomiesHierarchy err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetCatalogsHierarchy(c context.Context, arg *api.IDReq) (resp *api.CatalogsResp, err error) {
	if resp, err = p.topicRPC.GetCatalogsHierarchy(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetCatalogsHierarchy err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SaveCatalogs(c context.Context, arg *api.ArgSaveCatalogs) (err error) {
	if _, err = p.topicRPC.SaveCatalogs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SaveCatalogs err(%+v) arg(%+v)", err, arg))
	}
	return
}
