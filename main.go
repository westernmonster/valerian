// Package classification FLYWIKI API.
//
//
//     Schemes: https
//     Host: dev.flywiki.com
//     BasePath: /api/v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Authorization
//          in: header
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/modules"
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

	app.StaticFile("/swagger", "./docs/swagger.json")
	return app
}

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
