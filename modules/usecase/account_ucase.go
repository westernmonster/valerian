package usecase

import (
	"github.com/asaskevich/govalidator"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/infrastructure/helper"
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

	AreaRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Area, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Area, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Area, exist bool, err error)
	}
}

func (p *AccountUsecase) ChangePassword(ctx *biz.BizContext, req *models.ChangePasswordReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	account, exist, err := p.AccountRepository.GetByID(tx, *ctx.AccountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到当前用户")
		return
	}

	account.Password = req.Password

	err = p.AccountRepository.Update(tx, account)
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

func (p *AccountUsecase) GetLocationString(nodeID int64) (locationString string, err error) {
	arr := []string{}

	id := nodeID
	for {
		item, exist, errInner := p.AreaRepository.GetByID(p.Node, id)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}
		if !exist {
			err = berr.Errorf("未找到该地址")
			return
		}

		arr = append(arr, item.Name)

		if item.Parent == 0 {
			break
		}

		id = item.Parent
	}

	locationString = ""

	for i := len(arr) - 1; i >= 0; i-- {
		locationString += arr[i] + " "
	}

	return
}

func (p *AccountUsecase) GetProfile(ctx *biz.BizContext) (profile *models.ProfileResp, err error) {
	item, exist, err := p.AccountRepository.GetByID(p.Node, *ctx.AccountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未找到当前用户")
		return
	}

	profile = &models.ProfileResp{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if profile.Avatar == "" {
		profile.Avatar = "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png"
	}

	ipStr := helper.InetNtoA(item.IP)
	profile.IP = &ipStr

	if item.Location != nil {
		locationString, errInner := p.GetLocationString(*item.Location)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		profile.LocationString = &locationString
	}

	return
}

func (p *AccountUsecase) GetProfileByID(accountID int64) (profile *models.ProfileResp, err error) {
	item, exist, err := p.AccountRepository.GetByID(p.Node, accountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未找到当前用户")
		return
	}

	profile = &models.ProfileResp{
		ID:           item.ID,
		Mobile:       item.Mobile,
		Email:        item.Email,
		Gender:       item.Gender,
		BirthYear:    item.BirthYear,
		BirthMonth:   item.BirthMonth,
		BirthDay:     item.BirthDay,
		Location:     item.Location,
		Introduction: item.Introduction,
		Avatar:       item.Avatar,
		Source:       item.Source,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if profile.Avatar == "" {
		profile.Avatar = "https://flywiki.oss-cn-hangzhou.aliyuncs.com/765-default-avatar.png"
	}

	ipStr := helper.InetNtoA(item.IP)
	profile.IP = &ipStr

	if item.Location != nil {
		locationString, errInner := p.GetLocationString(*item.Location)
		if errInner != nil {
			err = tracerr.Wrap(errInner)
			return
		}

		profile.LocationString = &locationString
	}

	return
}

func (p *AccountUsecase) UpdateProfile(ctx *biz.BizContext, req *models.UpdateProfileReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	account, exist, err := p.AccountRepository.GetByID(tx, *ctx.AccountID)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	if !exist {
		err = berr.Errorf("未找到当前用户")
		return
	}

	if req.Gender != nil {
		if *req.Gender != models.GenderMale && *req.Gender != models.GenderFemale {
			err = berr.Errorf("性别数据错误")
			return
		}
		account.Gender = req.Gender
	}

	if req.Avatar != nil {
		if !govalidator.IsURL(*req.Avatar) {
			err = berr.Errorf("头像格式不正确")
			return
		}
		account.Avatar = *req.Avatar
	}

	if req.Introduction != nil {
		account.Introduction = req.Introduction
	}

	if req.BirthYear != nil {
		account.BirthYear = req.BirthYear
	}

	if req.BirthMonth != nil {
		account.BirthMonth = req.BirthMonth
	}

	if req.BirthDay != nil {
		account.BirthDay = req.BirthDay
	}

	// TODO: Validate BirthDay

	if req.Password != nil {
		if len(*req.Password) != 32 {
			err = berr.Errorf("密码格式不正确")
			return
		}
		account.Password = *req.Password
	}

	if req.Location != nil {
		_, exist, errGet := p.AreaRepository.GetByID(tx, *req.Location)
		if errGet != nil {
			err = tracerr.Wrap(err)
			return
		}

		if !exist {
			err = berr.Errorf("地址信息错误")
			return
		}

		account.Location = req.Location
	}

	err = p.AccountRepository.Update(tx, account)
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
