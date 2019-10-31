package service

import (
	"context"
	"valerian/app/admin/topic/model"
	topic "valerian/app/service/topic/api"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

func (p *Service) GetCatalogsHierarchy(c context.Context, topicID int64) (items []*model.TopicRootCatalog, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var resp *topic.CatalogsResp
	if resp, err = p.d.GetCatalogsHierarchy(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}

	items = p.FromCatalogs(resp.Items)

	return
}

func (p *Service) GetCatalogTaxonomiesHierarchy(c context.Context, topicID int64) (items []*model.TopicRootCatalog, err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var resp *topic.CatalogsResp
	if resp, err = p.d.GetCatalogTaxonomiesHierarchy(c, &topic.IDReq{ID: topicID, Aid: aid}); err != nil {
		return
	}

	items = p.FromCatalogs(resp.Items)

	return
}

func (p *Service) SaveCatalogs(c context.Context, req *model.ArgSaveTopicCatalog) (err error) {
	aid, ok := metadata.Value(c, metadata.Uid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	item := &topic.ArgSaveCatalogs{
		TopicID:  req.TopicID,
		Aid:      aid,
		ParentID: req.ParentID,
		Items:    make([]*topic.ArgTopicCatalog, 0),
	}

	for _, v := range req.Items {
		x := &topic.ArgTopicCatalog{
			Name:  v.Name,
			Seq:   v.Seq,
			Type:  v.Type,
			RefID: v.RefID,
		}

		if v.ID != nil {
			x.ID = &topic.ArgTopicCatalog_IDValue{*v.ID}
		}

		item.Items = append(item.Items, x)

	}

	if err = p.d.SaveCatalogs(c, item); err != nil {
		return
	}

	return
}
