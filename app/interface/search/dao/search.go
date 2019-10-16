package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/search/model"
	"valerian/library/conf/env"

	"gopkg.in/olivere/elastic.v6"
)

func (p *Dao) AccountSearch(c context.Context, arg *model.AccountSearchParams) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }
	indexName := fmt.Sprintf("%s_accounts", env.DeployEnv)

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg.BasicSearchParams); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) TopicSearch(c context.Context, arg *model.TopicSearchParams) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }
	indexName := fmt.Sprintf("%s_topics", env.DeployEnv)

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg.BasicSearchParams); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) ArticleSearch(c context.Context, arg *model.ArticleSearchParams) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)
	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }

	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg.BasicSearchParams); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}

func (p *Dao) DiscussionSearch(c context.Context, arg *model.DiscussionSearchParams) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }

	indexName := fmt.Sprintf("%s_accounts", env.DeployEnv)
	if arg.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.KW, arg.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, indexName, query, arg.BasicSearchParams); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg), "%v", err)
		return
	}

	return
}
