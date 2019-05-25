package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper" // gin-swagger middleware
	"go.uber.org/zap"

	_ "valerian/docs"
	"valerian/infrastructure/locale"
	"valerian/library/log"
	"valerian/library/net/http/mars"
	"valerian/library/tracing"
	"valerian/modules"
)

var (
	logger *zap.Logger
)

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

	locale.LoadTranslateFile()

	log.Init(nil)
	tracing.Init(nil)

	InitHTTP()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info(fmt.Sprintf("web-interface get a signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("web-interface exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}

}

func InitHTTP() {
	httpConfig := &mars.ServerConfig{
		Network: "tcp",
		Address: "0.0.0.0:7001",
	}
	httpConfig.Timeout.UnmarshalText([]byte("1s"))
	httpConfig.ReadTimeout.UnmarshalText([]byte("1s"))
	httpConfig.WriteTimeout.UnmarshalText([]byte("1s"))
	engine := mars.DefaultServer(httpConfig)
	modules.Configure(engine)

	if err := engine.Start(); err != nil {
		log.Error(fmt.Sprintf("engine.Start error(%v)", err))
		panic(err)
	}
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
