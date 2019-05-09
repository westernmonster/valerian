package http

import (
	"net/http"
	"strconv"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"

	"github.com/gin-gonic/gin"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"
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

func NewTopicCategoryCtrl(db *sqlx.DB, node sqalx.Node) *TopicCategoryCtrl {
	return &TopicCategoryCtrl{
		TopicCategoryUsecase: &usecase.TopicCategoryUsecase{
			Node:                    node,
			DB:                      db,
			TopicCategoryRepository: &repo.TopicCategoryRepository{},
		},
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
// @Router /topics/:id/categories [get]
func (p *TopicCategoryCtrl) GetAll(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		p.SuccessResp(ctx, nil)
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
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
// @Router /topics/:id/categories/hierarchy [get]
func (p *TopicCategoryCtrl) GetHierarchyOfAll(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		p.SuccessResp(ctx, nil)
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}
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
// @Router /topics/:id/categories [patch]
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
