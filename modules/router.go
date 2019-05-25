package modules

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"valerian/infrastructure/db"
	"valerian/library/email"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/sms"
	"valerian/modules/delivery/http"
	"valerian/modules/middleware"
	"valerian/modules/repo"
	"valerian/modules/usecase"
)

var (
	logger *zap.Logger
)

func Configure(p *mars.Engine) {
	node, err := db.InitDatabase()
	if err != nil {
		panic(err)
		return
	}

	auth := middleware.New()

	api := p.Group("/api/v1")
	{

		// api.GET("/oauth/token", HandleTokenRequest)
		// api.GET("/oauth/authorize", HandleAuthorizeRequest)

		authCtrl := &http.AuthCtrl{
			OauthUsecase: &usecase.OauthUsecase{
				Node:                       node,
				AccountRepository:          &repo.AccountRepository{},
				ValcodeRepository:          &repo.ValcodeRepository{},
				SessionRepository:          &repo.SessionRepository{},
				OauthClientRepository:      &repo.OauthClientRepository{},
				OauthAccessTokenRepository: &repo.OauthAccessTokenRepository{},
			},
			AccountUsecase: &usecase.AccountUsecase{
				Node:              node,
				AccountRepository: &repo.AccountRepository{},
				AreaRepository:    &repo.AreaRepository{},
			},
		}

		api.POST("/oauth/logout", authCtrl.Logout)
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
		valcodeCtrl := http.NewValcodeCtrl(smsClient, emailClient, node)
		api.POST("/valcodes/email", valcodeCtrl.RequestEmailValcode)
		api.POST("/valcodes/mobile", valcodeCtrl.RequestMobileValcode)

		// 账户
		accountCtrl := http.NewAccountCtrl(node)
		api.PUT("/me/password", auth.User, accountCtrl.ChangePassword)
		api.GET("/me", auth.User, accountCtrl.GetProfile)
		api.PUT("/me/edit", auth.User, accountCtrl.UpdateProfile)

		// 电话区域码
		countryCodeCtrl := http.NewCountryCodeCtrl(node, log.NewFactory())
		api.GET("/country_codes", countryCodeCtrl.GetAll)

		// 语言
		localeCtrl := http.NewLocaleCtrl(node)
		api.GET("/locales", localeCtrl.GetAll)

		// 话题
		topicCtrl := http.NewTopicCtrl(node)
		api.POST("/topics/add", auth.User, topicCtrl.Create)
		api.PUT("/topics/edit", auth.User, topicCtrl.Update)
		api.GET("/topics", auth.User, topicCtrl.Get)
		api.DELETE("/topics/delete", auth.User, topicCtrl.Delete)
		api.GET("/topics/search", auth.User, topicCtrl.Search)
		api.GET("/topics_members", auth.User, topicCtrl.GetTopicMembersPaged)
		api.POST("/topics_members/edit", auth.User, topicCtrl.BatchSavedTopicMember)
		api.GET("/topic_sets/versions", auth.User, topicCtrl.GetTopicVersions)
		api.POST("/topic_sets/versions/add", auth.User, topicCtrl.CreateNewVersion)
		api.GET("/topics/related", auth.User, topicCtrl.GetAllRelatedTopics)
		api.GET("/topic_types", auth.User, topicCtrl.GetAllTopicTypes)
		api.POST("/topics_catalogs", auth.User, topicCtrl.BatchSavedTopicCatalogs)

		fileCtrl := &http.FileCtrl{}
		api.POST("/files/oss_token", auth.User, fileCtrl.GetOSSToken)
	}
}
