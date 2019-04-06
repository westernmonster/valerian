package usecase

import (
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/berr"
	"git.flywk.com/flywiki/api/infrastructure/gid"
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

	AreaRepository interface {
		// GetAllByCondition get records by condition
		GetAllByCondition(node sqalx.Node, cond map[string]string) (items []*repo.Area, err error)
		// GetByID get a record by ID
		GetByID(node sqalx.Node, id int64) (item *repo.Area, exist bool, err error)
		// GetByCondition get a record by condition
		GetByCondition(node sqalx.Node, cond map[string]string) (item *repo.Area, exist bool, err error)
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

// EmailLogin 登录
func (p *AccountUsecase) EmailLogin(req *models.EmailLoginReq, ip string) (item *repo.Account, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	user, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"email": req.Email,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}

	if !exist {
		err = berr.Errorf("邮件地址不正确")
		return
	}

	if !strings.EqualFold(user.Password, req.Password) {
		err = berr.Errorf("密码不正确")
		return
	}

	item = user

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

// MobileLogin 登录
func (p *AccountUsecase) MobileLogin(req *models.MobileLoginReq, ip string) (item *repo.Account, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	user, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"mobile": req.Prefix + req.Mobile,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}

	if !exist {
		err = berr.Errorf("邮件地址不正确")
		return
	}

	if !strings.EqualFold(user.Password, req.Password) {
		err = berr.Errorf("密码不正确")
		return
	}

	item = user

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

func (p *AccountUsecase) EmailRegister(req *models.EmailRegisterReq, ip string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	ipAddr := helper.InetAtoN(ip)
	item := &repo.Account{
		ID:       id,
		Source:   req.Source,
		Password: req.Password,
		IP:       ipAddr,
	}

	_, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"email": req.Email,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}
	if exist {
		err = berr.Errorf("该邮件地址已经注册")
		return
	}
	item.Email = req.Email

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(tx, req.Email, models.ValcodeRegister, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.AccountRepository.Insert(tx, item)
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

func (p *AccountUsecase) MobileRegister(req *models.MobileRegisterReq, ip string) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	ipAddr := helper.InetAtoN(ip)
	item := &repo.Account{
		ID:       id,
		Source:   req.Source,
		Password: req.Password,
		IP:       ipAddr,
	}

	mobile := req.Prefix + req.Mobile

	_, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
		"mobile": mobile,
	})
	if errGet != nil {
		err = tracerr.Wrap(errGet)
		return
	}
	if exist {
		err = berr.Errorf("该手机号已经注册")
		return
	}
	item.Mobile = mobile

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(tx, mobile, models.ValcodeRegister, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	err = p.AccountRepository.Insert(tx, item)
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

func (p *AccountUsecase) ForgetPassword(req *models.ForgetPasswordReq) (sessionID int64, err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	var accountID int64

	if govalidator.IsEmail(req.Identity) {
		item, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
			"email": req.Identity,
		})
		if errGet != nil {
			err = tracerr.Wrap(errGet)
			return
		}
		if !exist {
			err = berr.Errorf("该邮件未注册")
			return
		}

		accountID = item.ID

	} else {
		item, exist, errGet := p.AccountRepository.GetByCondition(tx, map[string]string{
			"mobile": req.Identity,
		})
		if errGet != nil {
			err = tracerr.Wrap(errGet)
			return
		}
		if !exist {
			err = berr.Errorf("该手机未注册")
			return
		}
		accountID = item.ID
	}

	// Valcode
	correct, valcodeItem, errValcode := p.ValcodeRepository.IsCodeCorrect(tx, req.Identity, models.ValcodeForgetPassword, req.Valcode)
	if errValcode != nil {
		err = tracerr.Wrap(errValcode)
		return
	}
	if !correct {
		err = berr.Errorf("验证码不正确或已经使用")
		return
	}
	valcodeItem.Used = 1

	err = p.ValcodeRepository.Update(tx, valcodeItem)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	session := &repo.Session{
		ID:          id,
		SessionType: models.SessionTypeResetPassword,
		Used:        0,
		AccountID:   accountID,
	}

	err = p.SessionRepository.Insert(tx, session)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	sessionID = id

	err = tx.Commit()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	return

}

func (p *AccountUsecase) ResetPassword(req *models.ResetPasswordReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	idStr, err := helper.Base64Decode(req.SessionID)
	if err != nil {
		err = berr.Errorf("错误的 Session ID")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = berr.Errorf("错误的 Session ID")
		return
	}

	session, exist, err := p.SessionRepository.GetByID(tx, id)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if !exist {
		err = berr.Errorf("未获取到当前Session")
		return
	}

	if session.Used == 1 {
		err = berr.Errorf("当前Session已失效")
		return
	}

	if time.Now().Sub(time.Unix(session.CreatedAt, 0)) > 2*time.Minute {
		err = berr.Errorf("当前Session已失效")
		return
	}

	account, exist, err := p.AccountRepository.GetByID(tx, session.AccountID)
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

func (p *AccountUsecase) ChangePassword(accountID int64, req *models.ChangePasswordReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	account, exist, err := p.AccountRepository.GetByID(tx, accountID)
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

func (p *AccountUsecase) GetProfile(accountID int64) (profile *models.ProfileResp, err error) {
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

func (p *AccountUsecase) UpdateProfile(accountID int64, req *models.UpdateProfileReq) (err error) {
	tx, err := p.Node.Beginx()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	defer tx.Rollback()

	account, exist, err := p.AccountRepository.GetByID(tx, accountID)
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
