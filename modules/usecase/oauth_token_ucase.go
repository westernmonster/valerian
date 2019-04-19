package usecase

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/gid"
	"git.flywk.com/flywiki/api/infrastructure/oauth2"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type OAUTHTokenUsecase struct {
	sqalx.Node
	*sqlx.DB
	OAUTHTokenRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.OAUTHToken, err error)
		// GetByID get record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.OAUTHToken, exist bool, err error)

		// GetByCondition get record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.OAUTHToken, exist bool, err error)

		// HasSentRecordsInDuration determine current identity has sent records in specified duration
		HasSentRecordsInDuration(node sqalx.Node, identity string, codeType int, duration time.Duration) (has bool, err error)

		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.OAUTHToken) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.OAUTHToken) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
		DeleteExpired(node sqalx.Node, expiredAt int64) (err error)
		GetExpiredCount(node sqalx.Node, expiredAt int64) (total int, err error)
	}
}

func (p *OAUTHTokenUsecase) Create(info oauth2.TokenInfo) (err error) {
	jv, err := json.Marshal(info)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	item := &repo.OAUTHToken{
		ID:   id,
		Data: string(jv),
	}

	if code := info.GetCode(); code != "" {
		item.Code = code
		item.ExpiredAt = info.GetCodeCreateAt().Add(info.GetCodeExpiresIn()).Unix()
	} else {
		item.Access = info.GetAccess()
		item.ExpiredAt = info.GetAccessCreateAt().Add(info.GetAccessExpiresIn()).Unix()

		if refresh := info.GetRefresh(); refresh != "" {
			item.Refresh = info.GetRefresh()
			item.ExpiredAt = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Unix()
		}
	}

	err = p.OAUTHTokenRepository.Insert(p.Node, item)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}

func (p *OAUTHTokenUsecase) RemoveByCode(code string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.OAUTHTokenRepository.GetByCondition(tx, map[string]string{
		"code": code,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	item.Code = ""

	err = p.OAUTHTokenRepository.Update(tx, item)
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

func (p *OAUTHTokenUsecase) RemoveByAccess(access string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.OAUTHTokenRepository.GetByCondition(tx, map[string]string{
		"access": access,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	item.Code = ""

	err = p.OAUTHTokenRepository.Update(tx, item)
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

func (p *OAUTHTokenUsecase) RemoveByRefresh(refresh string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	item, exist, err := p.OAUTHTokenRepository.GetByCondition(tx, map[string]string{
		"refresh": refresh,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	item.Code = ""

	err = p.OAUTHTokenRepository.Update(tx, item)
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

func (p *OAUTHTokenUsecase) toTokenInfo(data string) oauth2.TokenInfo {
	var tm models.Token
	jsoniter.Unmarshal([]byte(data), &tm)
	return &tm
}

// GetByAccess use the access token for token information data
func (p *OAUTHTokenUsecase) GetByAccess(access string) (token oauth2.TokenInfo, err error) {
	if access == "" {
		return
	}

	item, exist, err := p.OAUTHTokenRepository.GetByCondition(p.Node, map[string]string{
		"access": access,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	return p.toTokenInfo(item.Data), nil
}

// GetByRefresh use the refresh token for token information data
func (p *OAUTHTokenUsecase) GetByRefresh(refresh string) (token oauth2.TokenInfo, err error) {
	if refresh == "" {
		return
	}

	item, exist, err := p.OAUTHTokenRepository.GetByCondition(p.Node, map[string]string{
		"refresh": refresh,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		return
	}

	return p.toTokenInfo(item.Data), nil
}

// GetByRefresh use the refresh token for token information data
func (p *OAUTHTokenUsecase) Clean(expiredAt int64) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	total, err := p.OAUTHTokenRepository.GetExpiredCount(tx, expiredAt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if total <= 0 {
		return
	}

	err = p.OAUTHTokenRepository.DeleteExpired(tx, expiredAt)
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
