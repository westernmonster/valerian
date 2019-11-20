package dao

import (
	"context"
	"fmt"

	"valerian/app/service/topic/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/log"
)

func (p *Dao) GetTopic(c context.Context, id int64) (resp *topic.TopicInfo, err error) {
	if resp, err = p.topicRPC.GetTopicInfo(c, &topic.TopicReq{ID: id}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserTopicsPaged error(%+v), id(%d) ", err, id))
	}

	return
}

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

func (p *Dao) GetUserCanEditTopicIDs(c context.Context, arg *api.AidReq) (resp *api.IDsResp, err error) {
	if resp, err = p.topicRPC.GetUserCanEditTopicIDs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetUserCanEditTopicIDs err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetFollowedTopicsIDs(c context.Context, arg *api.AidReq) (resp *api.IDsResp, err error) {
	if resp, err = p.topicRPC.GetFollowedTopicsIDs(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetFollowedTopicsIDs err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) GetTopicStat(c context.Context, arg *api.TopicReq) (info *topic.TopicStat, err error) {
	if info, err = p.topicRPC.GetTopicStat(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicStat err(%+v)", err))
	}
	return
}

func (p *Dao) HasTaxonomy(c context.Context, arg *api.TopicReq) (has bool, err error) {
	var resp *topic.BoolResp
	if resp, err = p.topicRPC.HasTaxonomy(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.HasTaxonomy err(%+v)", err))
	}

	has = resp.Result
	return
}

func (p *Dao) IsTopicMember(c context.Context, arg *api.ArgIsTopicMember) (has bool, err error) {
	var resp *topic.BoolResp
	if resp, err = p.topicRPC.IsTopicMember(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.HasTaxonomy err(%+v)", err))
	}

	has = resp.Result
	return
}

func (p *Dao) HasInvite(c context.Context, arg *api.ArgHasInvite) (has bool, err error) {
	var resp *topic.BoolResp
	if resp, err = p.topicRPC.HasInvite(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.HasInvite err(%+v)", err))
	}

	has = resp.Result
	return
}

func (p *Dao) Invite(c context.Context, arg *api.ArgTopicInvite) (err error) {
	if _, err = p.topicRPC.Invite(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Invite err(%+v)", err))
	}

	return
}

func (p *Dao) ProcessInvite(c context.Context, arg *api.ArgProcessInvite) (err error) {
	if _, err = p.topicRPC.ProcessInvite(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.ProcessInvite err(%+v)", err))
	}

	return
}

func (p *Dao) Leave(c context.Context, arg *api.TopicReq) (err error) {
	if _, err = p.topicRPC.Leave(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.Leave err(%+v)", err))
	}

	return
}

func (p *Dao) BulkSaveMembers(c context.Context, arg *api.ArgBatchSavedTopicMember) (err error) {
	if _, err = p.topicRPC.BulkSaveMembers(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.BulkSaveMembers err(%+v)", err))
	}

	return
}

func (p *Dao) GetTopicMembersPaged(c context.Context, arg *api.ArgTopicMembers) (resp *api.TopicMembersPagedResp, err error) {
	if resp, err = p.topicRPC.GetTopicMembersPaged(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetTopicMembersPaged err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) AddRecommendTopic(c context.Context, arg *api.TopicReq) (err error) {
	if _, err = p.topicRPC.AddRecommendTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRecommendTopic err(%+v) req(%+v)", err, arg))
	}

	return
}

func (p *Dao) DelRecommendTopic(c context.Context, arg *api.TopicReq) (err error) {
	if _, err = p.topicRPC.DelRecommendTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelRecommendTopic err(%+v) arg(%+v)", err, arg))
	}

	return
}

func (p *Dao) GetRecommendTopicsIDs(c context.Context) (resp *api.IDsResp, err error) {
	if resp, err = p.topicRPC.GetRecommendTopicsIDs(c, &api.EmptyStruct{}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRecommendTopicsIDs err(%+v)", err))
	}

	return
}
