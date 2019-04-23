package usecase

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"

	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
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
		// QueryListPaged get paged records by condition
		QueryListPaged(node sqalx.Node, page int, pageSize int, cond map[string]string) (total int, items []*repo.Account, err error)
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.Account, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Account, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Account, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Account, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.Account) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.Account) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	ValcodeRepository interface {
		// IsCodeCorrect determine current code's correctness
		// if used return false
		// if could not found in database, return false
		// if found in database and isn't used, return ture
		IsCodeCorrect(node sqalx.Node, identity string, codeType int, code string) (correct bool, item *repo.Valcode, err error)

		// Update update a exist record
		Update(node sqalx.Node, item *repo.Valcode) (err error)
	}

	SessionRepository interface {
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.Session, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Session, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Session, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Session, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.Session) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.Session) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)
		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	OauthAccessTokenRepository interface {
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.OauthAccessToken, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.OauthAccessToken, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.OauthAccessToken, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.OauthAccessToken, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.OauthAccessToken) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.OauthAccessToken) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)

		// Delete logic delete a exist record
		DeleteByCondition(node sqalx.Node, cond map[string]string) (err error)

		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
	}

	OauthClientRepository interface {
		// GetAll get all records
		GetAll(node sqalx.Node) (items []*repo.OauthClient, err error)
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.OauthClient, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.OauthClient, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.OauthClient, exist bool, err error)
		// Insert insert a new record
		Insert(node sqalx.Node, item *repo.OauthClient) (err error)
		// Update update a exist record
		Update(node sqalx.Node, item *repo.OauthClient) (err error)
		// Delete logic delete a exist record
		Delete(node sqalx.Node, id int64) (err error)

		// BatchDelete logic batch delete records
		BatchDelete(node sqalx.Node, ids []int64) (err error)
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
