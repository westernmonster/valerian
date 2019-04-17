package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"git.flywk.com/flywiki/api/modules/repo"
	"git.flywk.com/flywiki/api/modules/usecase"
)

type AccountCtrl struct {
	infrastructure.BaseCtrl

	AccountUsecase interface {
		ChangePassword(bizCtx *biz.BizContext, req *models.ChangePasswordReq) (err error)
		UpdateProfile(bizCtx *biz.BizContext, req *models.UpdateProfileReq) (err error)
		GetProfile(bizCtx *biz.BizContext) (profile *models.ProfileResp, err error)
	}
}

func NewAccountCtrl(db *sqlx.DB, node sqalx.Node) *AccountCtrl {
	return &AccountCtrl{
		AccountUsecase: &usecase.AccountUsecase{
			Node:              node,
			DB:                db,
			AccountRepository: &repo.AccountRepository{},
			AreaRepository:    &repo.AreaRepository{},
		},
	}
}

// @Summary 更改密码
// @Description 更改密码，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateProfileReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/password [put]
func (p *AccountCtrl) ChangePassword(ctx *gin.Context) {

	req := new(models.ChangePasswordReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	err := p.AccountUsecase.ChangePassword(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// UpdateProfile 更改用户资料
// @Summary 更改用户资料
// @Description 更改用户资料，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateProfileReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me [put]
func (p *AccountCtrl) UpdateProfile(ctx *gin.Context) {

	req := new(models.UpdateProfileReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	err := p.AccountUsecase.UpdateProfile(p.GetBizContext(ctx), req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// GetProfile 获取当前用户资料
// @Summary 获取当前用户资料
// @Description 获取当前用户资料，需要登录验证 （JWT Token 或 Cookie）
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object} models.ProfileResp "用户资料"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me [get]
func (p *AccountCtrl) GetProfile(ctx *gin.Context) {

	item, err := p.AccountUsecase.GetProfile(p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, item)

	return
}
