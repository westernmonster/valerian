package dao

import (
	"context"
	"fmt"

	"gopkg.in/olivere/elastic.v6"
)

func (d *Dao) ArticleSuggest(c context.Context, text string, size int) (res []string, err error) {
	titleSuggester := elastic.NewCompletionSuggester("title-suggester").
		Field("title").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	contentSuggester := elastic.NewCompletionSuggester("content-suggester").
		Field("content_text").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(titleSuggester).Suggester(contentSuggester).FetchSource(false).TrackScores(true)

	searchResult, err := d.esClient.Search().Index("articles").Type("article").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	titleResult := searchResult.Suggest["title-suggester"]
	for _, options := range titleResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	introResult := searchResult.Suggest["content-suggester"]
	for _, options := range introResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) TopicSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("name-suggester").
		Field("name").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	introductionSuggester := elastic.NewCompletionSuggester("introduction-suggester").
		Field("introduction").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).Suggester(introductionSuggester).FetchSource(false).TrackScores(true)

	searchResult, err := d.esClient.Search().Index("topics").Type("topic").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["name-suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	introResult := searchResult.Suggest["introduction-suggester"]
	for _, options := range introResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) AccountSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("name-suggester").
		Field("user_name").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	introductionSuggester := elastic.NewCompletionSuggester("introduction-suggester").
		Field("introduction").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).Suggester(introductionSuggester).FetchSource(false).TrackScores(true)

	searchResult, err := d.esClient.Search().Index("accounts").Type("account").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["name-suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	introResult := searchResult.Suggest["introduction-suggester"]
	for _, options := range introResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}

func (d *Dao) DiscussionSuggest(c context.Context, text string, size int) (res []string, err error) {
	nameSuggester := elastic.NewCompletionSuggester("name-suggester").
		Field("title").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	introductionSuggester := elastic.NewCompletionSuggester("content-suggester").
		Field("content_text").
		Fuzziness(2).
		Text(text).
		Size(size).SkipDuplicates(true)

	searchSource := elastic.NewSearchSource().Suggester(nameSuggester).Suggester(introductionSuggester).FetchSource(false).TrackScores(true)

	searchResult, err := d.esClient.Search().Index("accounts").Type("account").SearchSource(searchSource).Do(c)
	if err != nil {
		PromError(c, fmt.Sprintf("es:执行查询失败"), "%v", err)
		return
	}

	res = make([]string, 0)
	nameResult := searchResult.Suggest["name-suggester"]
	for _, options := range nameResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	introResult := searchResult.Suggest["content-suggester"]
	for _, options := range introResult {
		for _, option := range options.Options {
			fmt.Printf("%v ", option.Text)
			res = append(res, option.Text)
		}
	}

	return
}
