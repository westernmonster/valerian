package dao

import (
	"context"
	"fmt"
	"valerian/app/service/search/model"
	"valerian/library/conf/env"
	"valerian/library/xstr"

	"gopkg.in/olivere/elastic.v6"
)

func (p *Dao) AccountSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	if ids != nil && len(ids) > 0 {
		query = query.Filter(elastic.NewIdsQuery("account").Ids(xstr.Int64Array2StringArray(ids)...))
	}

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }
	indexName := fmt.Sprintf("%s_accounts", env.DeployEnv)

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) TopicSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	if ids != nil && len(ids) > 0 {
		query = query.Filter(elastic.NewIdsQuery("topic").Ids(xstr.Int64Array2StringArray(ids)...))
	}

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }
	indexName := fmt.Sprintf("%s_topics", env.DeployEnv)

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) ArticleSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	if ids != nil && len(ids) > 0 {
		query = query.Filter(elastic.NewIdsQuery("article").Ids(xstr.Int64Array2StringArray(ids)...))
	}
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) DiscussionSearch(c context.Context, arg *model.BasicSearchParams, ids []int64) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)
	if ids != nil && len(ids) > 0 {
		query = query.Filter(elastic.NewIdsQuery("discussion").Ids(xstr.Int64Array2StringArray(ids)...))
	}

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }

	indexName := fmt.Sprintf("%s_discussions", env.DeployEnv)
	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}
