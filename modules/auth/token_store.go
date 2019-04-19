package auth

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"git.flywk.com/flywiki/api/infrastructure/oauth2"
	"git.flywk.com/flywiki/api/modules/usecase"
)

type OAUTHTokenStore struct {
	OAUTHTokenUsecase interface {
		Create(info oauth2.TokenInfo) (err error)
		RemoveByCode(code string) (err error)
		RemoveByAccess(access string) (err error)
		RemoveByRefresh(refresh string) (err error)
		// GetByAccess use the access token for token information data
		GetByAccess(access string) (token oauth2.TokenInfo, err error)
		// GetByRefresh use the refresh token for token information data
		GetByRefresh(refresh string) (token oauth2.TokenInfo, err error)
		// GetByRefresh use the refresh token for token information data
		Clean(expiredAt int64) (err error)
	}

	ticker *time.Ticker
}

func NewStore(gcInterval int) *OAUTHTokenStore {
	store := &OAUTHTokenStore{
		OAUTHTokenUsecase: &usecase.OAUTHTokenUsecase{},
	}
	interval := 600
	if gcInterval > 0 {
		interval = gcInterval
	}
	store.ticker = time.NewTicker(time.Second * time.Duration(interval))

	go store.gc()
	return store
}

// Close close the store
func (p *OAUTHTokenStore) Close() {
	p.ticker.Stop()
}

func (p *OAUTHTokenStore) gc() {
	for range p.ticker.C {
		now := time.Now().Unix()
		err := p.OAUTHTokenUsecase.Clean(now)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"prefix": "oauth_store",
				"method": "gc",
			}).Error(fmt.Sprintf("OAUTH gc failed: %v", err))
			return
		}
	}
}

func (p *OAUTHTokenStore) Create(info oauth2.TokenInfo) (err error) {
	return p.OAUTHTokenUsecase.Create(info)
}

func (p *OAUTHTokenStore) RemoveByCode(code string) (err error) {
	return p.OAUTHTokenUsecase.RemoveByCode(code)
}

func (p *OAUTHTokenStore) RemoveByAccess(access string) (err error) {
	return p.OAUTHTokenUsecase.RemoveByAccess(access)
}

func (p *OAUTHTokenStore) RemoveByRefresh(refresh string) (err error) {
	return p.OAUTHTokenUsecase.RemoveByRefresh(refresh)
}

func (p *OAUTHTokenStore) GetByAccess(access string) (token oauth2.TokenInfo, err error) {
	return p.OAUTHTokenUsecase.GetByAccess(access)
}

func (p *OAUTHTokenStore) GetByRefresh(refresh string) (token oauth2.TokenInfo, err error) {
	return p.OAUTHTokenUsecase.GetByRefresh(refresh)
}
