package usecase

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
)

type AccountUsecase struct {
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
}

func (p *AccountUsecase) GetByID(userID int64) (item *repo.Account, err error) {
	item, exist, err := p.AccountRepository.GetByID(p.Node, userID)

	if !exist {
		err = tracerr.Errorf("获取用户信息失败")
		return
	}

	return
}

func (p *AccountUsecase) DoLogin(req *models.LoginReq, ip string) (item *repo.Account, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	user, existed, err := p.AccountRepository.GetByCondition(tx, map[string]string{
		"email": req.Identity,
	})
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !existed {
		err = tracerr.Errorf("user is not exist.")
		return
	}

	if !strings.EqualFold(user.Password, req.Password) {
		err = tracerr.Errorf("invalid password")
		return
	}

	item = user

	// ipAddr := infrastructure.InetAtoN(ip)

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

func (p *AccountUsecase) Register(req *models.LoginReq, ip, agent string) (item *repo.Account, err error) {
	return
}
