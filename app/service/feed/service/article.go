package service

import (
	"context"
	"fmt"
	"time"

	article "valerian/app/service/article/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/feed/model"
	relation "valerian/app/service/relation/api"
	topic "valerian/app/service/topic/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onCatalogArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgCatalogArticleAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCatalogArticleAdded Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCatalogArticleAdded GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var membersResp *topic.IDsResp
	if membersResp, err = p.d.GetTopicMemberIDs(c, info.TopicID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onCatalogArticleAdded GetFansIDs failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range membersResp.IDs {
		var data *model.Feed
		if data, err = p.d.GetFeedByCond(c, tx, map[string]interface{}{
			"account_id":  v,
			"action_type": def.ActionTypeCreateArticle,
			"target_id":   article.ID,
		}); err != nil {
			return
		} else if data != nil {
			continue
		}

		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCreateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateArticle,
			ActorID:    article.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onCatalogArticleAdded() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onArticleAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleCreated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleAdded GetFansIDs failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeCreateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextCreateArticle,
			ActorID:    article.Creator.ID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleAdded() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onArticleUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleUpdated Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleUpdated GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleUpdated GetFansIDs failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeUpdateArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextUpdateArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleUpdated() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onArticleLiked(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleLiked)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleLiked GetFansIDs failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeLikeArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextLikeArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleLiked() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onArticleFaved(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgArticleFaved)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleFaved Unmarshal failed %#v", err))
		return
	}

	var article *article.ArticleInfo
	if article, err = p.d.GetArticle(c, info.ArticleID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleFaved GetArticle failed %#v", err))
		if ecode.Cause(err) == ecode.ArticleNotExist {
			m.Ack()
		}
		return
	}

	var fansResp *relation.IDsResp
	if fansResp, err = p.d.GetFansIDs(c, info.ActorID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onArticleFaved GetFansIDs failed %#v", err))
		return
	}

	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	for _, v := range fansResp.IDs {
		feed := &model.Feed{
			ID:         gid.NewID(),
			AccountID:  v,
			ActionType: def.ActionTypeFavArticle,
			ActionTime: time.Now().Unix(),
			ActionText: def.ActionTextFavArticle,
			ActorID:    info.ActorID,
			ActorType:  def.ActorTypeUser,
			TargetID:   article.ID,
			TargetType: def.TargetTypeArticle,
			CreatedAt:  time.Now().Unix(),
			UpdatedAt:  time.Now().Unix(),
		}

		if err = p.d.AddFeed(c, tx, feed); err != nil {
			log.For(c).Error(fmt.Sprintf("service.onArticleFaved() failed %#v", err))
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	m.Ack()
}

func (p *Service) onArticleDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgArticleDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("service.onArticleDeleted Unmarshal failed %#v", err)
		return
	}

	if err = p.d.DelFeedByCond(context.Background(), p.d.DB(), def.TargetTypeArticle, info.ArticleID); err != nil {
		log.Errorf("service.onArticleDeleted() failed %#v", err)
		return
	}

	m.Ack()
}
