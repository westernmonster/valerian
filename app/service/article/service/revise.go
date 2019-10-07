package service

import (
	"context"

	"valerian/app/service/article/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) GetReviseStat(c context.Context, reviseID int64) (item *model.ReviseStat, err error) {
	return p.d.GetReviseStatByID(c, p.d.DB(), reviseID)
}

func (p *Service) GetRevise(c context.Context, reviseID int64) (item *model.Revise, err error) {
	return p.getRevise(c, p.d.DB(), reviseID)
}

func (p *Service) getRevise(c context.Context, node sqalx.Node, reviseID int64) (item *model.Revise, err error) {
	var addCache = true
	if item, err = p.d.ReviseCache(c, reviseID); err != nil {
		addCache = false
	} else if item != nil {
		return
	}

	if item, err = p.d.GetReviseByID(c, p.d.DB(), reviseID); err != nil {
		return
	} else if item == nil {
		err = ecode.ReviseNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetReviseCache(context.TODO(), item)
		})
	}
	return
}

func (p *Service) GetReviseImageUrls(c context.Context, reviseID int64) (urls []string, err error) {
	urls = make([]string, 0)
	var imgs []*model.ImageURL
	if imgs, err = p.d.GetImageUrlsByCond(c, p.d.DB(), map[string]interface{}{
		"target_type": model.TargetTypeRevise,
		"target_id":   reviseID,
	}); err != nil {
		return
	}

	for _, v := range imgs {
		urls = append(urls, v.URL)
	}

	return
}
