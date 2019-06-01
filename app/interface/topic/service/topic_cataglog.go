package service

import (
	"context"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

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
