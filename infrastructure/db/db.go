package db

import (
	"fmt"
	"strings"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"

	"github.com/hashicorp/go-multierror"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/ztrue/tracerr"
)

var mustKeys = []string{
	"dialect",
	"database",
	"host",
	"port",
	"username",
	"password",
}

func ConstructDBConnStr(config map[string]string) (connStr string, err error) {
	var result error
	for _, v := range mustKeys {
		if _, exist := config[v]; !exist {
			multierror.Append(result, tracerr.Errorf("database config key \"%s\" is not existed.", v))
			continue
		} else {
			if strings.TrimSpace(config[v]) == "" {
				multierror.Append(result, tracerr.Errorf("database config  \"%s\" is empty.", v))
			}
		}

	}

	if result != nil {
		err = result
		return
	}

	if config["dialect"] == "mysql" {
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
			config["username"],
			config["password"],
			config["host"],
			config["port"],
			config["database"])
		return
	}

	if config["dialect"] == "postgres" {
		connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config["username"],
			config["password"],
			config["host"],
			config["port"],
			config["database"])
		return
	}

	err = tracerr.New(`unsupported database`)

	return
}

func InitDatabase(tracer opentracing.Tracer) (db *sqlx.DB, node sqalx.Node, err error) {
	mode := viper.Get("MODE")
	dbConfig := viper.GetStringMapString(fmt.Sprintf("%s.db.flywiki", mode))

	connStr, err := ConstructDBConnStr(dbConfig)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	db, err = sqlx.Open(dbConfig["dialect"], connStr, tracer)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	node, err = sqalx.New(db)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	return
}
