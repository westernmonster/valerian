package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/models"
)

type dicItem struct {
	Done bool
	Item *model.TopicCatalog
}

func (p *Service) GetCatalogHierarchyOfAll(c context.Context, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	return p.getCatalogHierarchyOfAll(c, p.d.DB(), topicID)
}

func (p *Service) getCatalogHierarchyOfAll(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicLevel1Catalog, err error) {
	items = make([]*model.TopicLevel1Catalog, 0)

	parents, err := p.d.GetTopicCatalogsByCondition(c, node, topicID, 0)
	if err != nil {
		return
	}

	for _, lvl1 := range parents {
		parent := &model.TopicLevel1Catalog{
			ID:       &lvl1.ID,
			Name:     lvl1.Name,
			Seq:      lvl1.Seq,
			Type:     lvl1.Type,
			RefID:    lvl1.RefID,
			Children: make([]*model.TopicLevel2Catalog, 0),
		}

		children, eInner := p.d.GetTopicCatalogsByCondition(c, node, topicID, lvl1.ID)
		if eInner != nil {
			err = eInner
			return
		}

		for _, lvl2 := range children {
			child := &model.TopicLevel2Catalog{
				ID:       &lvl2.ID,
				Name:     lvl2.Name,
				Seq:      lvl2.Seq,
				Type:     lvl2.Type,
				RefID:    lvl2.RefID,
				Children: make([]*model.TopicChildCatalog, 0),
			}

			sub, eInner := p.d.GetTopicCatalogsByCondition(c, node, topicID, lvl2.ID)
			if eInner != nil {
				err = eInner
				return
			}

			for _, lvl3 := range sub {
				subItem := &model.TopicChildCatalog{
					ID:    &lvl3.ID,
					Name:  lvl3.Name,
					Seq:   lvl3.Seq,
					Type:  lvl3.Type,
					RefID: lvl3.RefID,
				}
				child.Children = append(child.Children, subItem)
			}

			parent.Children = append(parent.Children, child)
		}

		items = append(items, parent)

	}

	return
}

func (p *Service) bulkCreateCatalogs(c context.Context, node sqalx.Node, topicID int64, catalogs []*model.TopicLevel1Catalog) (err error) {
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

	for _, v := range catalogs {
		var idLvl1 int64
		if idLvl1, err = p.createCatalog(c, tx, topicID, v.Name, v.Seq, v.Type, v.RefID, 0); err != nil {
			return
		}
		for _, x := range v.Children {
			var idLvl2 int64
			if idLvl2, err = p.createCatalog(c, tx, topicID, x.Name, x.Seq, x.Type, x.RefID, idLvl1); err != nil {
				return
			}

			for _, y := range x.Children {
				if _, err = p.createCatalog(c, tx, topicID, y.Name, y.Seq, y.Type, y.RefID, idLvl2); err != nil {
					return
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}

func (p *Service) createCatalog(c context.Context, node sqalx.Node, topicID int64, name string, seq int, rtype string, refID *int64, parentID int64) (id int64, err error) {
	item := &model.TopicCatalog{
		ID:        gid.NewID(),
		Name:      name,
		Seq:       seq,
		Type:      rtype,
		ParentID:  parentID,
		RefID:     refID,
		TopicID:   topicID,
		Deleted:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err = p.d.AddTopicCatalog(c, node, item); err != nil {
		return
	}

	return item.ID, nil

}

func (p *Service) updateCatalog(c context.Context, node sqalx.Node, id, topicID int64, name string, seq int, rtype string, refID *int64, parentID int64) (err error) {
	var item *model.TopicCatalog
	if item, err = p.d.GetTopicCatalogByCondition(c, node, map[string]string{
		"topic_id":  strconv.FormatInt(topicID, 10),
		"id":        strconv.FormatInt(id, 10),
		"type":      rtype,
		"parent_id": strconv.FormatInt(parentID, 10),
	}); err != nil {
		return
	} else if item == nil {
		return ecode.TopicCatalogNotExist
	}

	item.Name = name
	item.Seq = seq
	item.ParentID = parentID
	item.RefID = refID
	item.UpdatedAt = time.Now().Unix()

	err = p.d.UpdateTopicCatalog(c, node, item)

	return

}

func (p *Service) SaveCatalogs(c context.Context, topicID int64, req *models.ArgSaveTopicCatalog) (err error) {
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

	var dbCatalogs []*model.TopicCatalog
	if dbCatalogs, err = p.d.GetTopicCatalogsByCondition(c, tx, topicID, req.ParentID); err != nil {
		return
	}

	dic := make(map[int64]dicItem)
	for _, v := range dbCatalogs {
		dic[v.ID] = dicItem{
			Done: false,
			Item: v,
		}
	}

	for _, v := range req.Items {
		if v.ID == nil {
			if _, err = p.createCatalog(c, tx, topicID, v.Name, v.Seq, v.Type, v.RefID, req.ParentID); err != nil {
				return
			}
			continue
		}

		dic[*v.ID] = dicItem{Done: true}
		if err = p.updateCatalog(c, tx, *v.ID, topicID, v.Name, v.Seq, v.Type, v.RefID, req.ParentID); err != nil {
			return
		}
	}

	for k, v := range dic {
		if v.Done {
			continue
		}

		if v.Item.Type == models.TopicCatalogTaxonomy {
			var childrenCount int
			if childrenCount, err = p.d.GetTopicCatalogChildrenCount(c, tx, topicID, k); err != nil {
				return
			}
			if childrenCount > 0 {
				err = ecode.MustDeleteChildrenCatalogFirst
				return
			}
		}

		if err = p.d.DelTopicCatalog(c, tx, k); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}
	return
}
