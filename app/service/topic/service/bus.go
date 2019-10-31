package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/app/service/topic/model"
	"valerian/library/log"
)

func (p *Service) onTopicAdded(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicAdded{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicUpdated(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicUpdated{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicUpdated.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicUpdated, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicUpdated.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicDeleted(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicDeleted{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowed(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicFollowed{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowed.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicFollowed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowed.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicViewed(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicViewed{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicViewed.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicViewed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicViewed.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicLeaved(c context.Context, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicLeaved{TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicLeaved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicLeaved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicLeaved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowRequested(c context.Context, id, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicFollowRequested{RequestID: id, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRequested.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicFollowRequested, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRequested.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowRejected(c context.Context, id, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicFollowRejected{RequestID: id, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRejected.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicFollowRejected, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowRejected.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicFollowApproved(c context.Context, id, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicFollowApproved{RequestID: id, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowApproved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicFollowApproved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicFollowApproved.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicInviteSent(c context.Context, id, topicID, aid, actionTime int64) {
	msg := &def.MsgTopicInviteSent{InviteID: id, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicInviteSent.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicInviteSent, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicInviteSent.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onCatalogArticleAdded(c context.Context, articleID, topicID, aid, actionTime int64) {
	msg := &def.MsgCatalogArticleAdded{ArticleID: articleID, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onCatalogArticleAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusCatalogArticleAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onCatalogArticleAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onCatalogArticleDeleted(c context.Context, articleID, topicID, aid, actionTime int64) {
	msg := &def.MsgCatalogArticleDeleted{ArticleID: articleID, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onCatalogArticleDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusCatalogArticleDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onCatalogArticleDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicTaxonomyCatalogAdded(c context.Context, item *model.NewTaxonomyItem, aid, actionTime int64) {
	msg := &def.MsgTopicTaxonomyCatalogAdded{
		TopicID:    item.TopicID,
		CatalogID:  item.ID,
		Name:       item.Name,
		ActorID:    aid,
		ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicTaxonomyCatalogAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicTaxonomyCatalogDeleted(c context.Context, item *model.DelTaxonomyItem, aid, actionTime int64) {
	msg := &def.MsgTopicTaxonomyCatalogDeleted{
		TopicID:    item.TopicID,
		CatalogID:  item.ID,
		Name:       item.Name,
		ActorID:    aid,
		ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicTaxonomyCatalogDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicTaxonomyCatalogRenamed(c context.Context, item *model.RenamedTaxonomyItem, aid, actionTime int64) {
	msg := &def.MsgTopicTaxonomyCatalogRenamed{
		TopicID:    item.TopicID,
		CatalogID:  item.ID,
		OldName:    item.OldName,
		NewName:    item.NewName,
		ActorID:    aid,
		ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogRenamed.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicTaxonomyCatalogRenamed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogRenamed.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onTopicTaxonomyCatalogMoved(c context.Context, item *model.MovedTaxonomyItem, aid, actionTime int64) {
	msg := &def.MsgTopicTaxonomyCatalogMoved{
		TopicID:     item.TopicID,
		CatalogID:   item.ID,
		OldParentID: item.OldParentID,
		NewParentID: item.NewParentID,
		Name:        item.Name,
		ActorID:     aid,
		ActionTime:  actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogMoved.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusTopicTaxonomyCatalogMoved, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onTopicTaxonomyCatalogMoved.Publish(), err(%+v)", err))
		return
	}

	return
}
