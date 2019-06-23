package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx/types"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
)

func (p *Service) bulkCreateCatalogs(c context.Context, node sqalx.Node, articleID int64, title string, relations []*model.AddArticleRelation) (err error) {
	var tx sqalx.Node
	if tx, err = node.Beginx(c); err != nil {
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

	if err = p.checkRelations(relations); err != nil {
		return
	}

	for _, v := range relations {
		if err = p.checkCatalog(c, tx, v.TopicID, v.ParentID); err != nil {
			return
		}

		var maxSeq int
		if maxSeq, err = p.d.GetTopicCatalogMaxChildrenSeq(c, tx, v.TopicID, v.ParentID); err != nil {
			return
		}

		item := &model.TopicCatalog{
			ID:        gid.NewID(),
			Name:      title,
			Seq:       maxSeq + 1,
			Type:      model.TopicCatalogArticle,
			ParentID:  v.ParentID,
			TopicID:   v.TopicID,
			IsPrimary: types.BitBool(v.Primary),
			RefID:     &articleID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if err = p.d.AddTopicCatalog(c, tx, item); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}

func (p *Service) checkRelations(items []*model.AddArticleRelation) (err error) {
	if len(items) == 0 {
		return ecode.NeedPrimaryTopic
	}

	primary := false
	dic := make(map[int64]bool)
	for _, v := range items {
		if primary == true {
			return ecode.OnlyAllowOnePrimaryTopic
		}
		if v.Primary {
			primary = true
		}

		if _, ok := dic[v.TopicID]; ok {
			return ecode.DuplicateTopicID
		}
	}

	return nil
}

func (p *Service) checkCatalog(c context.Context, node sqalx.Node, topicID, parentID int64) (err error) {
	if parentID == 0 {
		return
	}

	var catalog *model.TopicCatalog
	if catalog, err = p.d.GetTopicCatalogByID(c, node, parentID); err != nil {
		return
	} else if catalog == nil {
		err = ecode.TopicCatalogNotExist
		return
	} else if catalog.Type != model.TopicCatalogTaxonomy {
		err = ecode.InvalidCatalog
		return
	} else if catalog.TopicID != topicID {
		err = ecode.InvalidCatalog
		return
	}
	return nil
}
