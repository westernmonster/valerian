package service

import (
	"context"
	"valerian/app/interface/topic/model"

	"github.com/jinzhu/copier"
)

func (p *Service) GetTopicTypes(c context.Context) (items []*model.TopicTypeResp, err error) {
	items = make([]*model.TopicTypeResp, 0)
	data, err := p.d.GetAllTopicTypes(c, p.d.DB())
	if err != nil {
		return
	}

	copier.Copy(&items, &data)

	return

}
