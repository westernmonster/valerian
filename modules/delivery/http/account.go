package http

import (
	"context"

	"valerian/library/database/sqalx"
	"valerian/library/net/http/mars"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"
)

type AccountCtrl struct {
	infrastructure.BaseCtrl

	AccountUsecase interface {
		ChangePassword(c context.Context, ctx *biz.BizContext, req *models.ChangePasswordReq) (err error)
		UpdateProfile(c context.Context, ctx *biz.BizContext, req *models.UpdateProfileReq) (err error)
		GetProfile(c context.Context, ctx *biz.BizContext) (profile *models.ProfileResp, err error)
	}
}

func NewAccountCtrl(node sqalx.Node) *AccountCtrl {
	return &AccountCtrl{
		AccountUsecase: &usecase.AccountUsecase{
			Node:              node,
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
func (p *AccountCtrl) ChangePassword(ctx *mars.Context) {
	req := new(models.ChangePasswordReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, e)
		return
	}

	err := p.AccountUsecase.ChangePassword(ctx.Request.Context(), p.GetBizContext(ctx), req)

	ctx.JSON(nil, err)

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
func (p *AccountCtrl) UpdateProfile(ctx *mars.Context) {

	req := new(models.UpdateProfileReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	err := p.AccountUsecase.UpdateProfile(ctx.Request.Context(), p.GetBizContext(ctx), req)
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
func (p *AccountCtrl) GetProfile(ctx *mars.Context) {
	item, err := p.AccountUsecase.GetProfile(ctx.Request.Context(), p.GetBizContext(ctx))
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, item)

	return
}
