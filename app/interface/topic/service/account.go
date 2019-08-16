package service

import (
	"context"
	"encoding/json"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
)

func (p *Service) getAccountByID(c context.Context, node sqalx.Node, aid int64) (account *model.Account, err error) {
	var (
		addCache = true
	)

	if account, err = p.d.AccountCache(c, aid); err != nil {
		addCache = false
	} else if account != nil {
		return
	}

	if account, err = p.d.GetAccountByID(c, node, aid); err != nil {
		return
	} else if account == nil {
		err = ecode.UserNotExist
		return
	}

	if addCache {
		p.addCache(func() {
			p.d.SetAccountCache(context.TODO(), account)
		})
	}
	return
}

func (p *Service) getBasicAccountResp(c context.Context, node sqalx.Node, aid int64) (resp *model.BasicAccountResp, err error) {
	acc, err := p.getAccountByID(c, node, aid)
	if err != nil {
		return
	}
	resp = &model.BasicAccountResp{
		Avatar:    acc.Avatar,
		UserName:  acc.UserName,
		AccountID: acc.ID,
	}

	return
}

func (p *Service) AccountSearch(c context.Context, arg *model.AccountSearchParams) (res *model.AccountSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.AccountSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
	}

	res = &model.AccountSearchResult{
		Order: data.Order,
		Sort:  data.Sort,
		Page:  data.Page,
		Debug: data.Debug,
		Data:  make([]*model.ESAccount, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESAccount)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		res.Data = append(res.Data, acc)
	}
	return
}
