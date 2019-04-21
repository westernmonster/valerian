package usecase

import (
	"time"

	gopher_utils "github.com/felipeweb/gopher-utils"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/gid"
	"git.flywk.com/flywiki/api/infrastructure/osin"
	"git.flywk.com/flywiki/api/modules/repo"
)

type OAuthStorage struct {
	sqalx.Node
	*sqlx.DB
	AuthClientRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.AuthClient, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.AuthClient, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.AuthClient, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.AuthClient, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.AuthClient, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.AuthClient) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.AuthClient) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	AuthExpiresRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.AuthExpires, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.AuthExpires, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.AuthExpires, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.AuthExpires, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.AuthExpires, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.AuthExpires) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.AuthExpires) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	AuthAccessRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.AuthAccess, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.AuthAccess, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.AuthAccess, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.AuthAccess, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.AuthAccess, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.AuthAccess) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.AuthAccess) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	AuthRefreshRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.AuthRefresh, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.AuthRefresh, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.AuthRefresh, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.AuthRefresh, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.AuthRefresh, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.AuthRefresh) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.AuthRefresh) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	AuthAuthorizeRepository interface {
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.AuthAuthorize, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.AuthAuthorize, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.AuthAuthorize, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.AuthAuthorize, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.AuthAuthorize, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.AuthAuthorize) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.AuthAuthorize) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}
}

func (p *OAuthStorage) Clone() osin.Storage {
	return p
}

func (p *OAuthStorage) Close() {
	return
}

// GetClient loads the client by id
func (p *OAuthStorage) GetClient(id string) (client osin.Client, err error) {
	item, exist, err := p.AuthClientRepository.GetByCondition(p.Node, map[string]string{"client_id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	var c = new(osin.DefaultClient)
	c.Id = item.ClientID
	c.Secret = item.ClientSecret
	c.RedirectUri = item.RedirectURI
	c.UserData = item.Extra

	return c, nil
}

// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
func (p *OAuthStorage) UpdateClient(c osin.Client) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()
	data := gopher_utils.ToStr(c.GetUserData())

	item, exist, err := p.AuthClientRepository.GetByCondition(tx, map[string]string{"client_id": c.GetId()})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	item.ClientSecret = c.GetSecret()
	item.Extra = data
	item.RedirectURI = c.GetRedirectUri()

	err = p.AuthClientRepository.Update(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// CreateClient stores the client in the database and returns an error, if something went wrong.
func (p *OAuthStorage) CreateClient(c osin.Client) (err error) {

	data := gopher_utils.ToStr(c.GetUserData())

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := new(repo.AuthClient)
	item.ID = id
	item.ClientID = c.GetId()
	item.ClientSecret = c.GetSecret()
	item.RedirectURI = c.GetRedirectUri()
	item.Extra = data

	err = p.AuthClientRepository.Insert(p.Node, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
func (p *OAuthStorage) RemoveClient(id string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()
	item, exist, err := p.AuthClientRepository.GetByCondition(tx, map[string]string{"client_id": id})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	err = p.AuthClientRepository.Delete(tx, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

// SaveAuthorize saves authorize data.
func (p *OAuthStorage) SaveAuthorize(data *osin.AuthorizeData) (err error) {
	extra := gopher_utils.ToStr(data.UserData)

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := &repo.AuthAuthorize{
		ID:          id,
		ClientID:    data.Client.GetId(),
		Code:        data.Code,
		ExpiredIn:   int64(data.ExpiresIn),
		Scope:       data.Scope,
		RedirectURI: data.RedirectUri,
		State:       data.State,
		Extra:       extra,
	}

	err = p.AuthAuthorizeRepository.Insert(p.Node, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (p *OAuthStorage) LoadAuthorize(code string) (data *osin.AuthorizeData, err error) {
	item, exist, err := p.AuthAuthorizeRepository.GetByCondition(p.Node, map[string]string{
		"code": code,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	client, err := p.GetClient(item.ClientID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	data = &osin.AuthorizeData{
		Code:        item.Code,
		ExpiresIn:   int32(item.ExpiredIn),
		Scope:       item.Scope,
		RedirectUri: item.RedirectURI,
		State:       item.State,
		UserData:    item.Extra,
		CreatedAt:   time.Unix(item.CreatedAt, 0),
		Client:      client,
	}

	return

}

// RemoveAuthorize revokes or deletes the authorization code.
func (p *OAuthStorage) RemoveAuthorize(code string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.AuthAuthorizeRepository.GetByCondition(tx, map[string]string{"code": code})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	err = p.AuthAuthorizeRepository.Delete(tx, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (p *OAuthStorage) SaveAccess(data *osin.AccessData) (err error) {
	prev := ""

	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	extra := gopher_utils.ToStr(data.UserData)

	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	if data.RefreshToken != "" {
		if e := p.saveRefresh(tx, data.RefreshToken, data.AccessToken); e != nil {
			err = tracerr.Wrap(e)
			return
		}
	}

	if data.Client == nil {
		err = berr.Errorf("data.Client must not be nil")
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := &repo.AuthAccess{
		ID:           id,
		ClientID:     data.Client.GetId(),
		Authorize:    authorizeData.Code,
		Previous:     prev,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiredIn:    int64(data.ExpiresIn),
		Scope:        data.Scope,
		RedirectURI:  data.RedirectUri,
		Extra:        extra,
	}

	err = p.AuthAccessRepository.Insert(tx, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (p *OAuthStorage) LoadAccess(code string) (data *osin.AccessData, err error) {
	item, exist, err := p.AuthAccessRepository.GetByCondition(p.Node, map[string]string{
		"access_token": code,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	data = &osin.AccessData{
		AccessToken:  item.AccessToken,
		RefreshToken: item.RefreshToken,
		ExpiresIn:    int32(item.ExpiredIn),
		Scope:        item.Scope,
		RedirectUri:  item.RedirectURI,
		CreatedAt:    time.Unix(item.CreatedAt, 0),
		UserData:     item.Extra,
	}

	client, err := p.GetClient(item.ClientID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	data.Client = client
	data.AuthorizeData, _ = p.LoadAuthorize(item.Authorize)
	prevAccess, _ := p.LoadAccess(item.AccessToken)
	data.AccessData = prevAccess

	return
}

// RemoveAccess revokes or deletes an AccessData.
func (p *OAuthStorage) RemoveAccess(code string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.AuthAccessRepository.GetByCondition(p.Node, map[string]string{
		"access_token": code,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	err = p.AuthAccessRepository.Delete(tx, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (p *OAuthStorage) LoadRefresh(code string) (data *osin.AccessData, err error) {
	item, exist, err := p.AuthRefreshRepository.GetByCondition(p.Node, map[string]string{
		"token": code,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	return p.LoadAccess(item.Access)
}

// RemoveRefresh revokes or deletes refresh AccessData.
func (p *OAuthStorage) RemoveRefresh(code string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.AuthRefreshRepository.GetByCondition(tx, map[string]string{"token": code})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	err = p.AuthRefreshRepository.Delete(tx, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// CreateClientWithInformation Makes easy to create a osin.DefaultClient
func (p *OAuthStorage) CreateClientWithInformation(id string, secret string, redirectURI string, userData interface{}) (client osin.Client) {
	return &osin.DefaultClient{
		Id:          id,
		Secret:      secret,
		RedirectUri: redirectURI,
		UserData:    userData,
	}
}

// AddExpireAtData add info in expires table
func (p *OAuthStorage) AddExpireAtData(code string, expireAt time.Time) (err error) {
	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := &repo.AuthExpires{
		ID:        id,
		Token:     code,
		ExpiresAt: expireAt.Unix(),
	}

	err = p.AuthExpiresRepository.Insert(p.Node, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

// RemoveExpireAtData remove info in expires table
func (p *OAuthStorage) RemoveExpireAtData(code string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.AuthExpiresRepository.GetByCondition(tx, map[string]string{"token": code})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = osin.ErrNotFound
		return
	}

	err = p.AuthExpiresRepository.Delete(tx, item.ID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}

func (p *OAuthStorage) saveRefresh(node sqalx.Node, refresh, access string) (err error) {
	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := &repo.AuthRefresh{
		ID:     id,
		Token:  refresh,
		Access: access,
	}

	err = p.AuthRefreshRepository.Insert(p.Node, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return
}
