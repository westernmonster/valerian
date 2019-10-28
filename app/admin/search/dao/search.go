package dao

import (
	"context"
	"fmt"

	"valerian/app/admin/search/model"
	search "valerian/app/service/search/api"
	"valerian/library/log"
)

func (p *Dao) SearchTopic(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error) {
	req := &search.SearchParam{
		KW:       arg.KW,
		KwFields: arg.KwFields,
		Order:    arg.Order,
		Sort:     arg.Sort,
		Pn:       int32(arg.Pn),
		Ps:       int32(arg.Ps),
		Debug:    arg.Debug,
		Source:   arg.Source,
	}

	if info, err = p.searchRPC.SearchTopic(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchAccount(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error) {
	req := &search.SearchParam{
		KW:       arg.KW,
		KwFields: arg.KwFields,
		Order:    arg.Order,
		Sort:     arg.Sort,
		Pn:       int32(arg.Pn),
		Ps:       int32(arg.Ps),
		Debug:    arg.Debug,
		Source:   arg.Source,
	}

	if info, err = p.searchRPC.SearchAccount(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchAccount err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchArticle(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error) {
	req := &search.SearchParam{
		KW:       arg.KW,
		KwFields: arg.KwFields,
		Order:    arg.Order,
		Sort:     arg.Sort,
		Pn:       int32(arg.Pn),
		Ps:       int32(arg.Ps),
		Debug:    arg.Debug,
		Source:   arg.Source,
	}

	if info, err = p.searchRPC.SearchArticle(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchArticle err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchDiscussion(c context.Context, arg *model.BasicSearchParams) (info *search.SearchResult, err error) {
	req := &search.SearchParam{
		KW:       arg.KW,
		KwFields: arg.KwFields,
		Order:    arg.Order,
		Sort:     arg.Sort,
		Pn:       int32(arg.Pn),
		Ps:       int32(arg.Ps),
		Debug:    arg.Debug,
		Source:   arg.Source,
	}

	if info, err = p.searchRPC.SearchDiscussion(c, req); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchDiscussion err(%+v) arg(%+v)", err, arg))
	}
	return
}
