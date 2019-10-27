package service

import (
	"context"
	"fmt"
	account "valerian/app/service/account/api"
	"valerian/app/service/feed/def"
	"valerian/app/service/search/model"
	"valerian/library/log"

	retry "github.com/kamilsk/retry/v4"
	strategy "github.com/kamilsk/retry/v4/strategy"
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

	var v *account.DBAccount
	action := func(c context.Context, _ uint) error {
		acc, e := p.d.GetAccountInfo(c, info.AccountID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
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
		IDCert:       &v.IDCert,
		WorkCert:     &v.WorkCert,
		IsVIP:        &v.IsVIP,
		IsOrg:        &v.IsOrg,
	}

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

	var v *account.DBAccount
	action := func(c context.Context, _ uint) error {
		acc, e := p.d.GetAccountInfo(c, info.AccountID)
		if e != nil {
			return e
		}

		v = acc
		return nil
	}

	if err := retry.TryContext(c, action, strategy.Limit(3)); err != nil {
		m.Ack()
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
		IDCert:       &v.IDCert,
		WorkCert:     &v.WorkCert,
		IsVIP:        &v.IsVIP,
		IsOrg:        &v.IsOrg,
	}

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
