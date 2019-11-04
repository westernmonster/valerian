package service

import (
	"context"
	"fmt"

	"valerian/app/service/feed/def"
	"valerian/library/log"
)

func (p *Service) onCatalogArticleAdded(c context.Context, articleID, articleHistoryID, topicID, aid, actionTime int64) {
	msg := &def.MsgCatalogArticleAdded{ArticleID: articleID, ArticleHistoryID: articleHistoryID, TopicID: topicID, ActorID: aid, ActionTime: actionTime}

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

func (p *Service) onArticleAdded(c context.Context, articleID, aid, actionTime int64) {
	msg := &def.MsgArticleCreated{ArticleID: articleID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onArticleUpdated(c context.Context, articleID, historyID, aid, actionTime int64) {
	msg := &def.MsgArticleUpdated{ArticleID: articleID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleUpdated.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleUpdated, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleUpdated.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onArticleDeleted(c context.Context, articleID, aid, actionTime int64) {
	msg := &def.MsgArticleDeleted{ArticleID: articleID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseAdded(c context.Context, reviseID, aid, actionTime int64) {
	msg := &def.MsgReviseAdded{ReviseID: reviseID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseAdded.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseAdded, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseAdded.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseUpdated(c context.Context, reviseID, aid, actionTime int64) {
	msg := &def.MsgReviseUpdated{ReviseID: reviseID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseUpdated.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseUpdated, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseUpdated.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onReviseDeleted(c context.Context, reviseID, aid, actionTime int64) {
	msg := &def.MsgReviseDeleted{ReviseID: reviseID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseDeleted.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusReviseDeleted, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onReviseDeleted.Publish(), err(%+v)", err))
		return
	}

	return
}

func (p *Service) onArticleViewed(c context.Context, articleID, aid, actionTime int64) {
	msg := &def.MsgArticleViewed{ArticleID: articleID, ActorID: aid, ActionTime: actionTime}

	var data []byte
	var err error

	if data, err = msg.Marshal(); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleViewed.Marshal(), err(%+v)", err))
		return
	}

	if err = p.mq.Publish(def.BusArticleViewed, data); err != nil {
		log.For(c).Error(fmt.Sprintf("onArticleViewed.Publish(), err(%+v)", err))
		return
	}

	return
}
