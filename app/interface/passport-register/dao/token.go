package dao

import (
	"context"
	"database/sql"
	"fmt"

	"valerian/app/interface/passport-register/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getClientSQL              = `SELECT a.* FROM clients a WHERE a.deleted=0 AND a.client_Id=? `
	_addAccessTokenSQL         = `INSERT INTO access_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delExpiredAccessTokenSQL  = `DELETE FROM access_tokens WHERE client_id=? AND account_id=? `
	_delExpiredRefreshTokenSQL = `DELETE FROM refresh_tokens WHERE client_id=? AND account_id=? `
	_getClientAccessTokensSQL  = "SELECT * FROM access_tokens WHERE client_id=? AND account_id=? AND deleted=0"
	_getClientRefreshTokensSQL = "SELECT * FROM refresh_tokens WHERE client_id=? AND account_id=? AND deleted=0"
	_delAccessTokenSQL         = `DELETE FROM access_tokens WHERE token=?`
	_getAccessTokenSQL         = "SELECT * FROM access_tokens WHERE token=? AND deleted=0 LIMIT 1"

	_addRefreshTokenSQL = `INSERT INTO refresh_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delRefreshTokenSQL = `DELETE refresh_tokens WHERE token=?`
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

func (p *Dao) GetClientAccessTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error) {
	tokens = make([]string, 0)
	items := make([]*model.AccessToken, 0)

	if err = node.SelectContext(c, &items, _getClientAccessTokensSQL, clientID, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetClientAccessTokens error(%+v), aid(%d)", err, aid))
		return
	}

	for _, v := range items {
		tokens = append(tokens, v.Token)
	}

	return
}

func (p *Dao) DelExpiredAccessToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error) {
	r, err := node.ExecContext(c, _delExpiredAccessTokenSQL, clientID, accountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelExpiredAccessToken(%s, %d), error(%+v)", clientID, accountID, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) DelExpiredRefreshToken(c context.Context, node sqalx.Node, clientID string, accountID int64) (affected int64, err error) {
	r, err := node.ExecContext(c, _delExpiredRefreshTokenSQL, clientID, accountID)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelExpiredRefreshToken(%s, %d), error(%+v)", clientID, accountID, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) GetClientRefreshTokens(c context.Context, node sqalx.Node, aid int64, clientID string) (tokens []string, err error) {
	tokens = make([]string, 0)
	items := make([]*model.RefreshToken, 0)

	if err = node.SelectContext(c, &items, _getClientRefreshTokensSQL, clientID, aid); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetClientRefreshTokens error(%+v), aid(%d)", err, aid))
		return
	}

	for _, v := range items {
		tokens = append(tokens, v.Token)
	}

	return
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
