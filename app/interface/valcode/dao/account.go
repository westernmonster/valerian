package dao

import (
	"context"
	"fmt"

	account "valerian/app/service/account/api"
	"valerian/library/log"
)

func (p *Dao) IsEmailExistAndNotAnnul(c context.Context, email string) (exist bool, err error) {
	var info *account.ExistResp
	if info, err = p.accountRPC.EmailExistAndNotAnnul(c, &account.EmailReq{Email: email}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsEmailExist err(%+v) email(%s)", err, email))
		return
	}

	return info.Exist, nil
}

func (p *Dao) IsMobileExistAndNotAnnul(c context.Context, prefix, mobile string) (exist bool, err error) {
	var info *account.ExistResp
	if info, err = p.accountRPC.MobileExistAndNotAnnul(c, &account.MobileReq{Prefix: prefix, Mobile: mobile}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.IsMobileExist err(%+v) prefix(%s) mobile(%s)", err, prefix, mobile))
		return
	}

	return info.Exist, nil
}
