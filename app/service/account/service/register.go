package service

import (
	"context"
	"fmt"
	"time"
	"valerian/app/service/account/api"
	"valerian/app/service/account/model"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/log"
)

func (p *Service) AddAccount(c context.Context, item *model.Account) (resp *api.SelfProfile, err error) {
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	var account *model.Account
	var e error
	if item.Mobile != "" {
		if account, e = p.d.GetAccountByMobile(c, tx, item.Mobile); e != nil {
			return nil, e
		} else if account != nil && account.IsAnnul == false {
			err = ecode.AccountExist
			return
		}
	} else {
		if account, e = p.d.GetAccountByEmail(c, tx, item.Email); e != nil {
			return nil, e
		} else if account != nil && account.IsAnnul == false {
			err = ecode.AccountExist
			return
		}
	}
	// 如果是注销用户则设置 is_annul ，用户名 ,新的密码信息 即可
	if account != nil && account.IsAnnul == true {
		if err = p.d.UnAnnulAccount(c, tx, account.ID, item.UserName, item.Password, item.Salt); err != nil {
			return
		}
		item.ID = account.ID
		// 清理缓存
		p.d.DelAccountCache(c, account.ID)
	} else {
		item.CreatedAt = time.Now().Unix()
		item.UpdatedAt = time.Now().Unix()

		if err = p.d.AddAccount(c, tx, item); err != nil {
			return
		}

		if err = p.d.AddAccountStat(c, tx, &model.AccountStat{
			AccountID: item.ID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}); err != nil {
			return
		}

		if err = p.d.AddMessageStat(c, tx, &model.MessageStat{
			AccountID: item.ID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}); err != nil {
			return
		}
	}
	if resp, err = p.getSelfProfile(c, tx, item.ID); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
		return
	}

	return
}
