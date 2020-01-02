package dao

import (
	"context"
	"fmt"

	account "valerian/app/service/account/api"
	"valerian/library/log"
)

func (p *Dao) GetAccountBaseInfo(c context.Context, aid int64) (info *account.BaseInfoReply, err error) {
	if info, err = p.accountRPC.BasicInfo(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountBaseInfo err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetMemberInfo(c context.Context, aid int64) (info *account.MemberInfoReply, err error) {
	if info, err = p.accountRPC.MemberInfo(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.MemberInfo err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetSelfProfile(c context.Context, aid int64) (info *account.SelfProfile, err error) {
	if info, err = p.accountRPC.SelfProfileInfo(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetSelfProfile err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetAccountStat(c context.Context, aid int64) (info *account.AccountStatInfo, err error) {
	if info, err = p.accountRPC.AccountStat(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountStat err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) GetAccountByEmail(c context.Context, email string) (info *account.DBAccount, err error) {
	if info, err = p.accountRPC.GetAccountByEmail(c, &account.EmailReq{Email: email, UseMaster: true}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByEmail err(%+v) email(%s)", err, email))
	}
	return
}

func (p *Dao) GetAccountByMobile(c context.Context, prefix, mobile string) (info *account.DBAccount, err error) {
	if info, err = p.accountRPC.GetAccountByMobile(c, &account.MobileReq{Prefix: prefix, Mobile: mobile, UseMaster: true}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountByMobile err(%+v) prefix(%s) email(%s)", err, prefix, mobile))
	}
	return
}
