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

func (p *Dao) GetAccountSetting(c context.Context, aid int64) (info *account.Setting, err error) {
	if info, err = p.accountRPC.SettingInfo(c, &account.AidReq{Aid: aid}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccountSetting err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) UpdateAccountSetting(c context.Context, aid int64, boolVals map[string]bool, language string) (err error) {
	if _, err = p.accountRPC.UpdateSetting(c, &account.SettingReq{Aid: aid, Settings: boolVals, Language: language}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.UpdateSetting err(%+v) aid(%d)", err, aid))
	}
	return
}

func (p *Dao) AnnulAccount(c context.Context, aid int64, password string) (err error) {
	if _, err = p.accountRPC.AnnulAccount(c, &account.AnnulReq{
		Aid:      aid,
		Password: password,
	}); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AnnulAccount err(%+v) aid(%d)", err, aid))
		return
	}
	return
}
