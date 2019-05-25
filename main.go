package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"

	_ "valerian/docs"
	"valerian/infrastructure/bootstrap"
	"valerian/infrastructure/locale"
	"valerian/library/net/http/mars"
	"valerian/modules"
)

var (
	logger *zap.Logger
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("flywk.com", "admin@flywk.com")
	app.Bootstrap()

	locale.LoadTranslateFile()
	app.Configure(
		modules.Configure,
	)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return app
}

// @title 飞行百科 API
// @version 1.0
// @description 飞行百科 API
// @description 所有返回结果以如下JSON对象返回
// @description <pre>
// @description {
// @description    "code": 0,           // 如果没有另行约定，一般为200
// @description    "message": "string", // 一般用户发生错误或者验证失败时候返回的消息
// @description    "result": {},        // 返回结果，所有文档所示的返回结果处于这个字段
// @description    "success": true      // 用于判断是否执行成功
// @description }
// @description </pre>
// @description
// @description
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @host dev.flywk.com
// @BasePath /api/v1
func main() {
	setupConfig()

	httpConfig := &mars.ServerConfig{
		Address: "0.0.0.0:7001",
	}

	httpConfig.Timeout.UnmarshalText([]byte("1s"))
	httpConfig.ReadTimeout.UnmarshalText([]byte("1s"))
	httpConfig.WriteTimeout.UnmarshalText([]byte("1s"))
	engine := mars.DefaultServer(httpConfig)

	app := newApp()

	app.Run(":7001")
}

func setupConfig() {
	viper.SetEnvPrefix("FLYWIKI")
	viper.BindEnv("MODE")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: :%s \n", err))

	}
	if viper.Get("MODE") == "" {
		panic(fmt.Errorf("[main] Env Var \"Mode\" is empty "))
	}
}
