package dao

import (
	"context"
	"fmt"
	"valerian/app/interface/topic/model"

	"gopkg.in/olivere/elastic.v6"
)

func (p *Dao) AccountSearch(c context.Context, arg *model.AccountSearchParams) (res *model.SearchResult, err error) {
	var (
		query = elastic.NewBoolQuery()
	)

	// if len(arg.Query) > 0 {
	// 	query = query.Must(elastic.NewTermQuery("deleted", false))
	// }

	if arg.Bsp.KW != "" {
		query = query.Must(elastic.NewMultiMatchQuery(arg.Bsp.KW, arg.Bsp.KwFields...).Type("best_fields").TieBreaker(0.6))
	}

	if res, err = p.searchResult(c, "external", "accounts", query, arg.Bsp); err != nil {
		PromError(c, fmt.Sprintf("es:%+v ", arg.Bsp), "%v", err)
		return
	}

	return
}
