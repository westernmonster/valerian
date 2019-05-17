package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
)

type OauthAccessToken struct {
	ID        int64  `db:"id" json:"id,string"`                 // ID ID
	ClientID  string `db:"client_id" json:"client_id"`          // ClientID Client ID
	AccountID int64  `db:"account_id" json:"account_id,string"` // AccountID Account ID
	Token     string `db:"token" json:"token"`                  // Token Token
	ExpiresAt int64  `db:"expires_at" json:"expires_at,string"` // ExpiresAt 过期时间
	Scope     string `db:"scope" json:"scope"`                  // Scope Scope
	Deleted   int    `db:"deleted" json:"deleted"`              // Deleted 是否删除
	CreatedAt int64  `db:"created_at" json:"created_at"`        // CreatedAt 创建时间
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`        // UpdatedAt 更新时间
}

type OauthAccessTokenRepository struct{}

// GetByCondition get record by condition
func (p *OauthAccessTokenRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *OauthAccessToken, exist bool, err error) {
	item = new(OauthAccessToken)
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

	sqlSelect := fmt.Sprintf("SELECT a.* FROM oauth_access_tokens a WHERE a.deleted=0 %s", clause)

	if e := node.GetContext(ctx, item, sqlSelect, condition...); e != nil {
		if e == sql.ErrNoRows {
			item = nil
			return
		}
		err = tracerr.Wrap(e)
		return
	}

	exist = true
	return
}

// Insert insert a new record
func (p *OauthAccessTokenRepository) Insert(ctx context.Context, node sqalx.Node, item *OauthAccessToken) (err error) {
	sqlInsert := "INSERT INTO oauth_access_tokens( id, client_id, account_id, token, expires_at, scope, deleted, created_at, updated_at) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	_, err = node.ExecContext(ctx, sqlInsert,
		item.ID,
		item.ClientID,
		item.AccountID,
		item.Token,
		item.ExpiresAt,
		item.Scope,
		item.Deleted,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// Delete delete records  by condition
func (p *OauthAccessTokenRepository) DeleteByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (err error) {
	condition := make([]interface{}, 0)
	clause := ""

	if val, ok := cond["client_id"]; ok {
		clause += " AND a.client_id ="
		condition = append(condition, val)
	}
	if val, ok := cond["account_id"]; ok {
		clause += " AND a.account_id =?"
		condition = append(condition, val)
	}

	if val, ok := cond["expires_at"]; ok {
		clause += " AND a.expires_at <=?"
		condition = append(condition, val)
	}

	sqlDelete := fmt.Sprintf("UPDATE oauth_access_tokens a SET a.deleted=1 WHERE 1=1 %s", clause)

	_, err = node.ExecContext(ctx, sqlDelete, condition...)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
