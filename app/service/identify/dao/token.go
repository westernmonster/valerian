package dao

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/app/service/identify/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

const (
	_getClientSQL              = `SELECT a.id,a.client_id,a.client_secret,a.extra,a.redirect_uri,a.deleted,a.created_at,a.updated_at FROM clients a WHERE a.deleted=0 AND a.client_Id=? `
	_addAccessTokenSQL         = `INSERT INTO access_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delExpiredAccessTokenSQL  = `DELETE FROM access_tokens WHERE client_id=? AND account_id=? `
	_getClientAccessTokensSQL  = "SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM access_tokens a WHERE a.client_id=? AND a.account_id=? AND a.deleted=0"
	_getClientRefreshTokensSQL = "SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM refresh_tokens a WHERE a.client_id=? AND a.account_id=? AND a.deleted=0"
	_delAccessTokenSQL         = `DELETE FROM access_tokens WHERE token=?`
	_getAccessTokenSQL         = "SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM access_tokens a WHERE a.token=? AND a.deleted=0 LIMIT 1"

	_delExpiredRefreshTokenSQL = `DELETE FROM refresh_tokens WHERE client_id=? AND account_id=? `
	_addRefreshTokenSQL        = `INSERT INTO refresh_tokens(id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)`
	_delRefreshTokenSQL        = `DELETE refresh_tokens WHERE token=?`
	_getRefreshTokenSQL        = "SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM refresh_tokens a WHERE a.token=? AND a.deleted=0 LIMIT 1"
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

func (p *Dao) GetRefreshToken(c context.Context, node sqalx.Node, token string) (item *model.RefreshToken, err error) {
	item = new(model.RefreshToken)

	if err = node.GetContext(c, item, _getRefreshTokenSQL, token); err != nil {
		if err == sql.ErrNoRows {
			item = nil
			err = nil
			return
		}

		log.For(c).Error(fmt.Sprintf("dao.GetRefreshToken error(%+v), token(%s)", err, token))
	}

	return
}

func (p *Dao) DelAllAccessToken(c context.Context, node sqalx.Node, aid int64) (affected int64, err error) {
	sqlDelete := "DELETE access_tokens WHERE account_id=?"
	r, err := node.ExecContext(c, sqlDelete, aid)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAllAccessToken(%+v), error(%+v)", aid, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) DelAllRefreshToken(c context.Context, node sqalx.Node, aid int64) (affected int64, err error) {
	sqlDelete := "DELETE refresh_tokens WHERE account_id=?"
	r, err := node.ExecContext(c, sqlDelete, aid)
	if err != nil {
		log.For(c).Error(fmt.Sprintf("dao.DelAllRefreshToken(%+v), error(%+v)", aid, err))
		return
	}

	return r.RowsAffected()
}

func (p *Dao) GetAccessTokensByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AccessToken, err error) {
	items = make([]*model.AccessToken, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =?"
		condition = append(condition, val)
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =?"
		condition = append(condition, val)
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM access_tokens a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetAccessTokensByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}

func (p *Dao) GetRefreshTokensByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.RefreshToken, err error) {
	items = make([]*model.RefreshToken, 0)
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["id"]; ok {
		clause += " AND a.id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}
	if val, ok := cond["token"]; ok {
		clause += " AND a.token =?"
		condition = append(condition, val)
	}
	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at =?"
		condition = append(condition, val)
	}
	if val, ok := cond["scope"]; ok {
		clause += " AND a.scope =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.id,a.client_id,a.account_id,a.token,a.expires_at,a.scope,a.deleted,a.created_at,a.updated_at FROM refresh_tokens a WHERE a.deleted=0 %s ORDER BY a.id DESC", clause)

	if err = node.SelectContext(c, &items, sqlSelect, condition...); err != nil {
		log.For(c).Error(fmt.Sprintf("dao.GetRefreshTokensByCond err(%+v), condition(%+v)", err, cond))
		return
	}
	return
}
