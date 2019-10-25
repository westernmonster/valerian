package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	"valerian/library/ecode"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onAccountAdded(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgAccountAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onAccountAdded Unmarshal failed %#v", err))
		return
	}

	var v *model.Account
	if v, err = p.d.GetAccountByID(c, p.d.DB(), info.AccountID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onAccountAdded GetAccountByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESAccount{
		ID:           v.ID,
		Mobile:       &v.Mobile,
		Email:        &v.Email,
		UserName:     &v.UserName,
		Role:         v.Role,
		Gender:       &v.Gender,
		BirthYear:    &v.BirthYear,
		BirthMonth:   &v.BirthMonth,
		BirthDay:     &v.BirthDay,
		Location:     &v.Location,
		Introduction: &v.Introduction,
		Avatar:       &v.Avatar,
		Source:       &v.Source,
		CreatedAt:    &v.CreatedAt,
		UpdatedAt:    &v.UpdatedAt,
	}

	idCert := bool(v.IDCert)
	workCert := bool(v.WorkCert)
	isOrg := bool(v.IsOrg)
	isVIP := bool(v.IsVIP)
	item.IDCert = &idCert
	item.WorkCert = &workCert
	item.IsVIP = &isVIP
	item.IsOrg = &isOrg

	if err = p.d.PutAccount2ES(c, item); err != nil {
		return
	}
	m.Ack()
}

func (p *Service) onAccountUpdated(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgAccountUpdated)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onAccountUpdated Unmarshal failed %#v", err))
		return
	}

	var v *model.Account
	if v, err = p.d.GetAccountByID(c, p.d.DB(), info.AccountID); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onAccountUpdated GetAccountByID failed %#v", err))
		return
	} else if v == nil {
		return
	}

	item := &model.ESAccount{
		ID:           v.ID,
		Mobile:       &v.Mobile,
		Email:        &v.Email,
		UserName:     &v.UserName,
		Role:         v.Role,
		Gender:       &v.Gender,
		BirthYear:    &v.BirthYear,
		BirthMonth:   &v.BirthMonth,
		BirthDay:     &v.BirthDay,
		Location:     &v.Location,
		Introduction: &v.Introduction,
		Avatar:       &v.Avatar,
		Source:       &v.Source,
		CreatedAt:    &v.CreatedAt,
		UpdatedAt:    &v.UpdatedAt,
	}

	idCert := bool(v.IDCert)
	workCert := bool(v.WorkCert)
	isOrg := bool(v.IsOrg)
	isVIP := bool(v.IsVIP)
	item.IDCert = &idCert
	item.WorkCert = &workCert
	item.IsVIP = &isVIP
	item.IsOrg = &isOrg

	if err = p.d.PutAccount2ES(c, item); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) onAccountDeleted(m *stan.Msg) {
	var err error
	c := context.Background()
	info := new(def.MsgAccountDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.For(c).Error(fmt.Sprintf("service.onAccountDeleted Unmarshal failed %#v", err))
		return
	}

	if err = p.d.DelESAccount(c, info.AccountID); err != nil {
		return
	}

	m.Ack()
}

func (p *Service) AccountSearch(c context.Context, arg *model.AccountSearchParams) (resp *model.AccountSearchResult, err error) {
	var data *model.SearchResult
	if data, err = p.d.AccountSearch(c, arg); err != nil {
		err = ecode.SearchAccountFailed
		return
	}

	resp = &model.AccountSearchResult{
		Paging: &model.Paging{},
		Debug:  data.Debug,
		Data:   make([]*model.ESAccount, 0),
	}

	for _, v := range data.Result {
		acc := new(model.ESAccount)
		err = json.Unmarshal(v, acc)
		if err != nil {
			return
		}

		var stat *model.AccountStat
		if stat, err = p.d.GetAccountStatByID(c, p.d.DB(), acc.ID); err != nil {
			return
		}

		acc.FansCount = int(stat.Fans)
		acc.FollowingCount = int(stat.Following)

		resp.Data = append(resp.Data, acc)
	}

	resp.Paging.Total = data.Page.Total

	if resp.Paging.Prev, err = genURL("/api/v1/search/accounts", url.Values{
		"kw":        []string{arg.KW},
		"kw_fields": arg.KwFields,
		"order":     arg.Order,
		"sort":      arg.Sort,
		"debug":     []string{strconv.FormatBool(arg.Debug)},
		"source":    arg.Source,
		"pn":        []string{strconv.Itoa(arg.Pn)},
		"ps":        []string{strconv.Itoa(arg.Ps)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/search/accounts", url.Values{
		"kw":        []string{arg.KW},
		"kw_fields": arg.KwFields,
		"order":     arg.Order,
		"sort":      arg.Sort,
		"debug":     []string{strconv.FormatBool(arg.Debug)},
		"source":    arg.Source,
		"pn":        []string{strconv.Itoa(arg.Pn + 1)},
		"ps":        []string{strconv.Itoa(arg.Ps)},
	}); err != nil {
		return
	}

	if len(resp.Data) < arg.Ps {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if arg.Pn == 1 {
		resp.Paging.Prev = ""
	}

	return
}
