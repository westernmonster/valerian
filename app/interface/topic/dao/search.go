package dao

import (
	"context"
	"fmt"

	search "valerian/app/service/search/api"
	"valerian/library/log"
)

func (p *Dao) SearchTopic(c context.Context, arg *search.SearchParam) (info *search.SearchResult, err error) {
	if info, err = p.searchRPC.SearchTopic(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchTopic err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchAccount(c context.Context, arg *search.SearchParam) (info *search.SearchResult, err error) {
	if info, err = p.searchRPC.SearchAccount(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchAccount err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchArticle(c context.Context, arg *search.SearchParam) (info *search.SearchResult, err error) {
	if info, err = p.searchRPC.SearchArticle(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchArticle err(%+v) arg(%+v)", err, arg))
	}
	return
}

func (p *Dao) SearchDiscussion(c context.Context, arg *search.SearchParam) (info *search.SearchResult, err error) {
	if info, err = p.searchRPC.SearchDiscussion(c, arg); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.SearchDiscussion err(%+v) arg(%+v)", err, arg))
	}
	return
}
