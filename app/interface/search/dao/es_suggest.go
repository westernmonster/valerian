package dao

import (
	"context"
	"fmt"
	"valerian/library/conf/env"

	"gopkg.in/olivere/elastic.v6"
)

func (d *Dao) ArticleSuggest(c context.Context, text string, size int) (res []string, err error) {
	titleSuggester := elastic.NewCompletionSuggester("suggester").
		Field("suggest").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(titleSuggester).FetchSource(false).TrackScores(true)

	indexName := fmt.Sprintf("%s_articles", env.DeployEnv)
	searchResult, err := d.esClient.Search().Index(indexName).Type("article").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	titleResult := searchResult.Suggest["suggester"]
	for _, options := range titleResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) TopicSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("suggester").
		Field("suggest").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).FetchSource(false).TrackScores(true)

	indexName := fmt.Sprintf("%s_topics", env.DeployEnv)
	searchResult, err := d.esClient.Search().Index(indexName).Type("topic").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) AccountSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("suggester").
		Field("suggest").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).FetchSource(false).TrackScores(true)

	indexName := fmt.Sprintf("%s_accounts", env.DeployEnv)
	searchResult, err := d.esClient.Search().Index(indexName).Type("account").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) DiscussionSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("suggester").
		Field("suggest").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).FetchSource(false).TrackScores(true)

	indexName := fmt.Sprintf("%s_discussions", env.DeployEnv)
	searchResult, err := d.esClient.Search().Index(indexName).Type("discussion").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}
