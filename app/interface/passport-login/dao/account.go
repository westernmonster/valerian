package dao

import (
	"context"

	"valerian/app/interface/passport-login/model"
)

func (p *Dao) GetAccountByEmail(ctx context.Context, email string) (account model.Account, exist bool, err error) {
	return
}

func (p *Dao) GetAccountByMobile(ctx context.Context, mobile string) (account model.Account, exist bool, err error) {
	return
}
