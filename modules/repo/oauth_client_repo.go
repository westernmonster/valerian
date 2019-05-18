package repo

import (
	"context"
	"database/sql"
	"fmt"
	"valerian/library/database/sqalx"

	tracerr "github.com/ztrue/tracerr"
)

type OauthClient struct {
	ID           int64   `db:"id" json:"id,string"`                        // ID ID
	ClientID     string  `db:"client_id" json:"client_id"`                 // ClientID Client ID
	ClientSecret string  `db:"client_secret" json:"client_secret"`         // ClientSecret Client Secret
	Extra        string  `db:"extra" json:"extra"`                         // Extra Extra
	RedirectURI  *string `db:"redirect_uri" json:"redirect_uri,omitempty"` // RedirectURI Redirect URI
	Deleted      int     `db:"deleted" json:"deleted"`                     // Deleted 是否删除
	CreatedAt    int64   `db:"created_at" json:"created_at"`               // CreatedAt 创建时间
	UpdatedAt    int64   `db:"updated_at" json:"updated_at"`               // UpdatedAt 更新时间
}

type OauthClientRepository struct{}

// GetByCondition get record by condition
func (p *OauthClientRepository) GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *OauthClient, exist bool, err error) {
	item = new(OauthClient)
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
	if val, ok := cond["client_secret"]; ok {
		clause += " AND a.client_secret =?"
		condition = append(condition, val)
	}
	if val, ok := cond["extra"]; ok {
		clause += " AND a.extra =?"
		condition = append(condition, val)
	}
	if val, ok := cond["redirect_uri"]; ok {
		clause += " AND a.redirect_uri =?"
		condition = append(condition, val)
	}

	sqlSelect := fmt.Sprintf("SELECT a.* FROM oauth_clients a WHERE a.deleted=0 %s ", clause)

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
