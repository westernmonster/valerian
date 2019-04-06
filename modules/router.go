package modules

import (
	"fmt"

	alidm "github.com/denverdino/aliyungo/dm"
	alisms "github.com/denverdino/aliyungo/sms"
	"github.com/spf13/viper"
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/infrastructure/db"
	"git.flywk.com/flywiki/api/infrastructure/email"
	"git.flywk.com/flywiki/api/infrastructure/sms"
	"git.flywk.com/flywiki/api/modules/delivery/http"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
)

var (
	AccountCtrl *http.AccountCtrl
)

func Configure(p *bootstrap.Bootstrapper) {
	api := p.Group("/api/v1")

	api.POST("/auth/login/email", AccountCtrl.EmailLogin)
	api.POST("/auth/login/mobile", AccountCtrl.MobileLogin)
	api.POST("/auth/registration/email", AccountCtrl.EmailRegister)
	api.POST("/auth/registration/mobile", AccountCtrl.MobileRegister)

	api.PUT("/auth/password/reset", AccountCtrl.ForgetPassword)
	api.PUT("/auth/password/reset/confirm", AccountCtrl.ResetPassword)
	api.PUT("/me/password", AccountCtrl.Auth, AccountCtrl.ChangePassword)
	api.PUT("/me", AccountCtrl.Auth, AccountCtrl.UpdateProfile)

	db, node, err := db.InitDatabase()
	if err != nil {
		panic(err)
		return
	}

	mode := viper.Get("MODE")
	accessKeyID := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_id", mode))
	accessKeySecret := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_secret", mode))

	aliSMSClient := alisms.NewDYSmsClient(accessKeyID, accessKeySecret)
	smsClient := &sms.SMSClient{
		Client: aliSMSClient,
	}

	aliDMClient := alidm.NewClient(accessKeyID, accessKeySecret)
	emailClient := &email.EmailClient{
		Client: aliDMClient,
	}

	valcodeCtrl := &http.ValcodeCtrl{
		ValcodeUsecase: &usecase.ValcodeUsecase{
			Node:              node,
			DB:                db,
			SMSClient:         smsClient,
			EmailClient:       emailClient,
			ValcodeRepository: &repo.ValcodeRepository{},
		},
	}

	api.POST("/valcodes/email", valcodeCtrl.RequestEmailValcode)
	api.POST("/valcodes/mobile", valcodeCtrl.RequestMobileValcode)

	countryCodeCtrl := &http.CountryCodeCtrl{
		CountryCodeUsecase: &usecase.CountryCodeUsecase{
			Node:                  node,
			DB:                    db,
			CountryCodeRepository: &repo.CountryCodeRepository{},
		},
	}

	api.GET("/country_codes", countryCodeCtrl.GetAll)
}

func InitAccountCtrl() (err error) {
	db, node, err := db.InitDatabase()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	AccountCtrl = &http.AccountCtrl{
		AccountUsecase: &usecase.AccountUsecase{
			Node:              node,
			DB:                db,
			AccountRepository: &repo.AccountRepository{},
			ValcodeRepository: &repo.ValcodeRepository{},
			SessionRepository: &repo.SessionRepository{},
			AreaRepository:    &repo.AreaRepository{},
		},
	}

	return
}
