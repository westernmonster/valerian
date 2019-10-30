package service

import (
	"context"
	"valerian/app/interface/topic/model"
)

func (p *Service) GetCatalogsHierarchy(c context.Context, topicID int64) (items []*model.TopicRootCatalog, err error) {
	return
}

func (p *Service) GetCatalogTaxonomiesHierarchy(c context.Context, topicID int64) (items []*model.TopicRootCatalog, err error) {

	return
}

func (p *Service) SaveCatalogs(c context.Context, req *model.ArgSaveTopicCatalog) (err error) {
	return
}
