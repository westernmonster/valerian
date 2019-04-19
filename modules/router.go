package modules

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
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
	AuthCtrl *http.AuthCtrl
)

func Configure(p *bootstrap.Bootstrapper) {

	api := p.Group("/api/v1")

	api.POST("/auth/login/email", AuthCtrl.EmailLogin)
	api.POST("/auth/login/mobile", AuthCtrl.MobileLogin)
	api.POST("/auth/registration/email", AuthCtrl.EmailRegister)
	api.POST("/auth/registration/mobile", AuthCtrl.MobileRegister)
	api.PUT("/auth/password/reset", AuthCtrl.ForgetPassword)
	api.PUT("/auth/password/reset/confirm", AuthCtrl.ResetPassword)

	db, node, err := db.InitDatabase()
	if err != nil {
		panic(err)
		return
	}

	mode := viper.Get("MODE")
	// 阿里云API客户端
	accessKeyID := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_id", mode))
	accessKeySecret := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_secret", mode))
	aliClient, err := sdk.NewClientWithAccessKey("cn-hangzhou", accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
		return
	}

	// 阿里云短信
	smsClient := &sms.SMSClient{Client: aliClient}
	// 阿里云邮件
	emailClient := &email.EmailClient{Client: aliClient}

	// 验证码
	valcodeCtrl := http.NewValcodeCtrl(smsClient, emailClient, db, node)
	api.POST("/valcodes/email", valcodeCtrl.RequestEmailValcode)
	api.POST("/valcodes/mobile", valcodeCtrl.RequestMobileValcode)

	// 账户
	accountCtrl := http.NewAccountCtrl(db, node)
	api.PUT("/me/password", AuthCtrl.Auth, accountCtrl.ChangePassword)
	api.GET("/me", AuthCtrl.Auth, accountCtrl.GetProfile)
	api.PUT("/me", AuthCtrl.Auth, accountCtrl.UpdateProfile)

	// 电话区域码
	countryCodeCtrl := http.NewCountryCodeCtrl(db, node)
	api.GET("/country_codes", countryCodeCtrl.GetAll)

	// 语言
	localeCtrl := http.NewLocaleCtrl(db, node)
	api.GET("/locales", localeCtrl.GetAll)

	// 话题分类
	topicCategoryCtrl := http.NewTopicCategoryCtrl(db, node)
	api.GET("/topic_categories", topicCategoryCtrl.GetAll)
	api.GET("/topic_categories/hierarchy", topicCategoryCtrl.GetHierarchyOfAll)
	api.PUT("/topic_categories/:id", topicCategoryCtrl.Update)
	api.POST("/topic_categories", topicCategoryCtrl.Create)
	api.DELETE("/topic_categories/:id", topicCategoryCtrl.Delete)
	api.PATCH("/topic_categories", topicCategoryCtrl.BulkSave)

	fileCtrl := &http.FileCtrl{}
	api.POST("/files/oss_token", fileCtrl.GetOSSToken)
}

func InitAuthCtrl() (err error) {
	db, node, err := db.InitDatabase()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	AuthCtrl = &http.AuthCtrl{
		AuthUsecase: &usecase.AuthUsecase{
			Node:              node,
			DB:                db,
			AccountRepository: &repo.AccountRepository{},
			ValcodeRepository: &repo.ValcodeRepository{},
			SessionRepository: &repo.SessionRepository{},
		},
	}

	return
}
