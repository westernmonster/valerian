package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"

	"github.com/ztrue/tracerr"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"
	"valerian/models"
	"valerian/modules/repo"
)

type OauthUsecase struct {
	sqalx.Node
	*sqlx.DB
	SMSClient interface {
		SendRegisterValcode(mobile string, valcode string) (err error)
		SendResetPasswordValcode(mobile string, valcode string) (err error)
	}

	EmailClient interface {
		SendActiveEmail(email string, valcode string) (err error)
		SendResetPasswordValcode(email string, valcode string) (err error)
	}
	AccountRepository interface {
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int64) (item *repo.Account, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.Account, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.Account) (err error)
		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.Account) (err error)
	}

	ValcodeRepository interface {
		// IsCodeCorrect determine current code's correctness
		// if used return false
		// if could not found in database, return false
		// if found in database and isn't used, return ture
		IsCodeCorrect(ctx context.Context, node sqalx.Node, identity string, codeType int, code string) (correct bool, item *repo.Valcode, err error)

		// Update update a exist record
		Update(ctx context.Context, node sqalx.Node, item *repo.Valcode) (err error)
	}

	SessionRepository interface {
		// GetByID get a record by ID
		GetByID(ctx context.Context, node sqalx.Node, id int64) (item *repo.Session, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.Session) (err error)
	}

	OauthAccessTokenRepository interface {
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.OauthAccessToken, exist bool, err error)
		// Insert insert a new record
		Insert(ctx context.Context, node sqalx.Node, item *repo.OauthAccessToken) (err error)

		// Delete logic delete a exist record
		DeleteByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (err error)
	}

	OauthClientRepository interface {
		// GetByCondition get a record by condition
		GetByCondition(ctx context.Context, node sqalx.Node, cond map[string]string) (item *repo.OauthClient, exist bool, err error)
	}
}

func comparePassword(password, passwordHash, salt string) (identical bool, err error) {
	hash, err := hashPassword(password, salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if hash == passwordHash {
		identical = true
	}

	return
}

func hashPassword(password string, salt string) (passwordHash string, err error) {
	mac := hmac.New(sha1.New, []byte(models.PasswordPepper))
	_, err = mac.Write([]byte(password))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	sha := hex.EncodeToString(mac.Sum(nil))

	h := sha1.New()
	_, err = h.Write([]byte(sha + salt))
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	passwordHash = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}

func generateSalt(n uint32) (salt string, err error) {
	b := make([]byte, n)
	_, err = rand.Read(b)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	salt = base64.URLEncoding.EncodeToString(b)
	return
}
