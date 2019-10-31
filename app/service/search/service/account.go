package service

import (
	"context"
	"fmt"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
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
		return
	} else if v == nil {
		m.Ack()
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
	item.IDCert = &idCert
	workCert := bool(v.WorkCert)
	item.WorkCert = &workCert
	isVIP := bool(v.IsVip)
	item.IsVIP = &isVIP
	isOrg := bool(v.IsOrg)
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
		return
	} else if v == nil {
		m.Ack()
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
	item.IDCert = &idCert
	workCert := bool(v.WorkCert)
	item.WorkCert = &workCert
	isVIP := bool(v.IsVip)
	item.IsVIP = &isVIP
	isOrg := bool(v.IsOrg)
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
