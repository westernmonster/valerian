package dao

import (
	"context"
	"fmt"

	"valerian/app/interface/passport-login/model"
)

const (
	_addAccessTokenSQL = `INSERT INTO oauth_access_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delAccessTokenSQL = `DELETE oauth_access_tokens WHERE token=?`

	_addRefreshTokenSQL = `INSERT INTO oauth_refresh_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delRefreshTokenSQL = `DELETE oauth_refresh_tokens WHERE token=?`
)

func (p *Dao) AddAccessToken(c context.Context, t *model.OauthAccessToken) (affected int64, err error) {
	r, err := p.node.ExecContext(c, _addAccessTokenSQL, t.ID, t.ClientID, t.AccountID, t.Token, t.ExpiresAt, t.Scope, t.Deleted, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		p.logger.For(c).Error(fmt.Sprintf("failed to add access_token(%+v), dao.authDB.Exec() error(%+v)", t, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) DelAccessToken(c context.Context, token string) (affected int64, err error) {
	r, err := p.node.ExecContext(c, _delAccessTokenSQL, token)
	if err != nil {
		p.logger.For(c).Error(fmt.Sprintf("failed to delete access_token(%+v), dao.authDB.Exec() error(%+v)", token, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) AddRefreshToken(c context.Context, t *model.OauthRefreshToken) (affected int64, err error) {
	r, err := p.node.ExecContext(c, _addRefreshTokenSQL, t.ID, t.ClientID, t.AccountID, t.Token, t.ExpiresAt, t.Scope, t.Deleted, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		p.logger.For(c).Error(fmt.Sprintf("failed to add refresh_token(%+v), dao.authDB.Exec() error(%+v)", t, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) DelRefreshToken(c context.Context, token string) (affected int64, err error) {
	r, err := p.node.ExecContext(c, _delRefreshTokenSQL, token)
	if err != nil {
		p.logger.For(c).Error(fmt.Sprintf("failed to delete refresh_token(%+v), dao.authDB.Exec() error(%+v)", token, err))
		return
	}

	return r.RowsAffected()
}
