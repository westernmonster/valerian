package http

import (
	"net/http"

	"git.flywk.com/flywiki/api/infrastructure"
	"git.flywk.com/flywiki/api/infrastructure/biz"
	"git.flywk.com/flywiki/api/models"
	"github.com/gin-gonic/gin"
)

type TopicCategoryCtrl struct {
	infrastructure.BaseCtrl

	TopicCategoryUsecase interface {
		Create(ctx *biz.BizContext, req *models.CreateTopicCategoryReq) (err error)
		Update(ctx *biz.BizContext, id int64, req *models.UpdateTopicCategoryReq) (err error)
		Delete(ctx *biz.BizContext, id int64) (err error)
		BulkSave(ctx *biz.BizContext, req *models.SaveTopicCategoriesReq) (err error)
		GetAll(ctx *biz.BizContext, topicID int64) (items []*models.TopicCategory, err error)
		GetHierarchyOfAll(ctx *biz.BizContext, topicID int64) (resp *models.TopicCategoriesResp, err error)
	}
}

// @Summary 获取分类列表
// @Description 获取分类列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Success 200 {array} models.TopicCategory "分类"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_categories [get]
func (p *TopicCategoryCtrl) GetAll(ctx *gin.Context) {
	topicID := ctx.GetInt64("topic_id")
	if topicID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "请提供话题ID",
		})

		return
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicCategoryUsecase.GetAll(bizCtx, topicID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

}

// @Summary 按层级结构获取分类
// @Description 按层级结构获取分类
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Success 200 {object} models.TopicCategoriesResp "层级分类"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_categories/hierarchy [get]
func (p *TopicCategoryCtrl) GetHierarchyOfAll(ctx *gin.Context) {
	topicID := ctx.GetInt64("topic_id")
	if topicID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "请提供话题ID",
		})

		return
	}

	bizCtx := p.GetBizContext(ctx)
	resp, err := p.TopicCategoryUsecase.GetHierarchyOfAll(bizCtx, topicID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, resp)

	return

}

// @Summary 新增话题分类
// @Description 新增话题分类
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.CreateTopicCategoryReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_categories [post]
func (p *TopicCategoryCtrl) Create(ctx *gin.Context) {
	req := new(models.CreateTopicCategoryReq)

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
	err := p.TopicCategoryUsecase.Create(bizCtx, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 更新话题分类
// @Description 更新话题分类
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.UpdateTopicCategoryReq true "请求"
// @Param id path string true "ID"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_categories/{id} [put]
func (p *TopicCategoryCtrl) Update(ctx *gin.Context) {
	req := new(models.UpdateTopicCategoryReq)

	if e := ctx.Bind(req); e != nil {
		p.HandleError(ctx, e)
		return
	}

	id := ctx.GetInt64("id")

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	bizCtx := p.GetBizContext(ctx)
	err := p.TopicCategoryUsecase.Update(bizCtx, id, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 删除话题分类
// @Description 删除话题分类
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
// @Router /topic_categories/{id} [delete]
func (p *TopicCategoryCtrl) Delete(ctx *gin.Context) {
	id := ctx.GetInt64("id")

	bizCtx := p.GetBizContext(ctx)
	err := p.TopicCategoryUsecase.Delete(bizCtx, id)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 批量新增/更改话题分类
// @Description 批量新增/更改话题分类
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.SaveTopicCategoriesReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_categories [patch]
func (p *TopicCategoryCtrl) BulkSave(ctx *gin.Context) {
	req := new(models.SaveTopicCategoriesReq)

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
	err := p.TopicCategoryUsecase.BulkSave(bizCtx, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}
