package dao

import (
	"context"
	"fmt"

	account "valerian/app/service/account/api"
	"valerian/library/log"
)

func (p *Dao) IsEmailExist(c context.Context, email string) (exist bool, err error) {
	var info *account.ExistResp
	if info, err = p.accountRPC.EmailExist(c, &account.EmailReq{Email: email}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsEmailExist err(%+v) email(%s)", err, email))
		return
	}

	return info.Exist, nil
}

func (p *Dao) IsMobileExist(c context.Context, prefix, mobile string) (exist bool, err error) {
	var info *account.ExistResp
	if info, err = p.accountRPC.MobileExist(c, &account.MobileReq{Prefix: prefix, Mobile: mobile}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsMobileExist err(%+v) prefix(%s) mobile(%s)", err, prefix, mobile))
		return
	}

	return info.Exist, nil
}
