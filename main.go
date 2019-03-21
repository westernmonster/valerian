package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/modules"

	_ "git.flywk.com/flywiki/api/docs"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("flywk.com", "admin@flywk.com")
	app.Bootstrap()

	err := modules.InitAccountCtrl()
	if err != nil {
		panic(err)
	}

	app.Configure(
		modules.Configure,
	)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return app
}

// @title
// @version 1.0
// @description 飞行百科 API
//
// @host www.flywk.com
// @BasePath /api/v1
func main() {
	setupConfig()

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
