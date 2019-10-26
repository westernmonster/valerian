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

func (p *Dao) AddAccount(c context.Context, v *account.AddAccountReq) (resp *account.SelfProfile, err error) {
	if resp, err = p.accountRPC.AddAccount(c, v); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccount err(%+v) item(%+v)", err, v))
	}

	return
}
