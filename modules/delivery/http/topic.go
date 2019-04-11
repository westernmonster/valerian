package http

import (
	"net/http"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type TopicCtrl struct {
	infrastructure.BaseCtrl

	TopicUsecase interface {
		Create(ctx *biz.BizContext, req *models.CreateTopicReq) (err error)
		Update(ctx *biz.BizContext, id int64, req *models.UpdateTopicReq) (err error)
		Delete(ctx *biz.BizContext, id int64) (err error)
	}
}

// @Summary 新增话题
// @Description 新增话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.CreateTopicReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics [post]
func (p *TopicCtrl) Create(ctx *gin.Context) {
	req := new(models.CreateTopicReq)

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

	bizCtx := p.GetBizContext(ctx)
	err := p.TopicUsecase.Create(bizCtx, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 更新话题
// @Description 更新话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateTopicReq true "请求"
// @Param id path string true "ID"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id} [put]
func (p *TopicCtrl) Update(ctx *gin.Context) {
	req := new(models.UpdateTopicReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	id := ctx.GetInt64("id")

	bizCtx := p.GetBizContext(ctx)
	err := p.TopicUsecase.Update(bizCtx, id, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 删除话题
// @Description 删除话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true "ID"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id} [delete]
func (p *TopicCtrl) Delete(ctx *gin.Context) {
	id := ctx.GetInt64("id")

	bizCtx := p.GetBizContext(ctx)
	err := p.TopicUsecase.Delete(bizCtx, id)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}
