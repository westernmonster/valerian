package modules

import (
	"github.com/ztrue/tracerr"

	"git.flywk.com/flywiki/api/infrastructure/bootstrap"
	"git.flywk.com/flywiki/api/infrastructure/db"
	"git.flywk.com/flywiki/api/modules/delivery/http"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
)

var (
	AccountCtrl *http.AccountCtrl
)

func Configure(p *bootstrap.Bootstrapper) {
	api := p.Group("/api/v1")

	api.POST("/session", AccountCtrl.Login)
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
		},
	}

	return
}
