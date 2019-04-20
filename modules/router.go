package modules

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/viper"

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/infrastructure/db"
	"git.flywk.com/flywiki/api/infrastructure/email"
	"git.flywk.com/flywiki/api/infrastructure/sms"
	"git.flywk.com/flywiki/api/modules/delivery/http"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
)

func Configure(p *bootstrap.Bootstrapper) {

	api := p.Group("/api/v1")
	{
		// api.GET("/oauth/token", authserver.HandleTokenRequest)
		// api.GET("/oauth/authorize", authserver.HandleAuthorizeRequest)

		db, node, err := db.InitDatabase()
		if err != nil {
			panic(err)
			return
		}

		authCtrl := &http.AuthCtrl{
			AuthUsecase: &usecase.AuthUsecase{
				Node:              node,
				DB:                db,
				AccountRepository: &repo.AccountRepository{},
				ValcodeRepository: &repo.ValcodeRepository{},
				SessionRepository: &repo.SessionRepository{},
			},
		}

		api.POST("/oauth/login/email", authCtrl.EmailLogin)
		api.POST("/oauth/login/mobile", authCtrl.MobileLogin)
		api.POST("/oauth/registration/email", authCtrl.EmailRegister)
		api.POST("/oauth/registration/mobile", authCtrl.MobileRegister)
		api.PUT("/oauth/password/reset", authCtrl.ForgetPassword)
		api.PUT("/oauth/password/reset/confirm", authCtrl.ResetPassword)

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
		// accountCtrl := http.NewAccountCtrl(db, node)
		// api.PUT("/me/password", authserver.HandleTokenVerify(), accountCtrl.ChangePassword)
		// api.GET("/me", authserver.HandleTokenVerify(), accountCtrl.GetProfile)
		// api.PUT("/me", authserver.HandleTokenVerify(), accountCtrl.UpdateProfile)

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
}
