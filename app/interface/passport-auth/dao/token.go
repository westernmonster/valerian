package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/passport-login/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getClientSQL             = `SELECT a.* FROM oauth_clients a WHERE a.deleted=0 AND a.client_Id=? `
	_addAccessTokenSQL        = `INSERT INTO oauth_access_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delExpiredAccessTokenSQL = `DELETE oauth_access_tokens WHERE token=?`
	_delAccessTokenSQL        = `DELETE oauth_access_tokens WHERE client_id=? AND account_id=? AND expires_at <= ?`
	_getAccessTokenSQL        = "SELECT * FROM oauth_access_tokens WHERE token=? LIMIT 1"

	_addRefreshTokenSQL = `INSERT INTO oauth_refresh_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delRefreshTokenSQL = `DELETE oauth_refresh_tokens WHERE token=?`
)

func (p *Dao) GetClient(c context.Context, node sqalx.Node, clientID string) (item *model.Client, err error) {
	item = new(model.Client)

	if err = node.GetContext(c, item, _getClientSQL, clientID); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetClient error(%+v), id(%s)", err, clientID))
	}

	return
}

func (p *Dao) AddAccessToken(c context.Context, node sqalx.Node, t *model.AccessToken) (affected int64, err error) {
	r, err := node.ExecContext(c, _addAccessTokenSQL, t.ID, t.ClientID, t.AccountID, t.Token, t.ExpiresAt, t.Scope, t.Deleted, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddAccessToken(%+v), error(%+v)", t, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) GetAccessToken(c context.Context, node sqalx.Node, token string) (item *model.AccessToken, err error) {
	item = new(model.AccessToken)

	if err = node.GetContext(c, item, _getAccessTokenSQL, token); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetAccessToken error(%+v), token(%s)", err, token))
	}

	return
}

func (p *Dao) DelExpiredAccessToken(c context.Context, node sqalx.Node, clientID string, accountID int64, expiresAt int64) (affected int64, err error) {
	r, err := node.ExecContext(c, _delExpiredAccessTokenSQL, clientID, accountID, expiresAt)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelExpiredAccessToken(%s, %d, %d), error(%+v)", clientID, accountID, expiresAt, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) AddRefreshToken(c context.Context, node sqalx.Node, t *model.RefreshToken) (affected int64, err error) {
	r, err := node.ExecContext(c, _addRefreshTokenSQL, t.ID, t.ClientID, t.AccountID, t.Token, t.ExpiresAt, t.Scope, t.Deleted, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRefreshToken(%+v), error(%+v)", t, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) DelRefreshToken(c context.Context, node sqalx.Node, token string) (affected int64, err error) {
	r, err := node.ExecContext(c, _delRefreshTokenSQL, token)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.AddRefreshToken(%+v), error(%+v)", token, err))
		return
	}

	return r.RowsAffected()
}
