package service

import (
	"context"
	"valerian/app/interface/article/model"
	"valerian/library/database/sqalx"
)

func (p *Service) editPermission(c context.Context, node sqalx.Node, aid, articleID int64) (canEdit bool, err error) {
	var catalogs []*model.TopicCatalog
	if catalogs, err = p.d.GetTopicCatalogsByCond(c, node, map[string]interface{}{
		"ref_id": articleID,
		"type":   model.TopicCatalogArticle,
	}); err != nil {
		return
	}

	for _, v := range catalogs {
		if v.IsPrimary {
		} else {
		}
	}

	return
}
