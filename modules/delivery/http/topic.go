package http

import (
	"net/http"
	"strconv"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/westernmonster/sqalx"
)

type TopicCtrl struct {
	infrastructure.BaseCtrl

	TopicUsecase interface {
		Create(ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error)
		Update(ctx *biz.BizContext, id int64, req *models.UpdateTopicReq) (err error)
		Delete(ctx *biz.BizContext, id int64) (err error)
		Get(ctx *biz.BizContext, topicID int64) (item *models.Topic, err error)
		SeachTopics(ctx *biz.BizContext, topicID int64, query string) (items []*models.TopicSearchResult, err error)
	}
}

func NewTopicCtrl(db *sqlx.DB, node sqalx.Node) *TopicCtrl {
	return &TopicCtrl{
		TopicUsecase: &usecase.TopicUsecase{
			Node:                    node,
			DB:                      db,
			TopicRepository:         &repo.TopicRepository{},
			TopicMemberRepository:   &repo.TopicMemberRepository{},
			TopicCategoryRepository: &repo.TopicCategoryRepository{},
			TopicTypeRepository:     &repo.TopicTypeRepository{},
			TopicSetRepository:      &repo.TopicSetRepository{},
			TopicRelationRepository: &repo.TopicRelationRepository{},
		},
	}
}

// @Summary 搜索话题
// @Description 搜索话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param q query string true  "查询条件"
// @Param from_topic_id query string true  "排除的话题ID"
// @Success 200 "成功,返回topic_id"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics [get]
func (p *TopicCtrl) Search(ctx *gin.Context) {
	query := ctx.Query("q")
	fromTopicID := ctx.Query("from_topic_id")

	topicID, err := strconv.ParseInt(fromTopicID, 10, 64)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.SeachTopics(bizCtx, topicID, query)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

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
// @Success 200 "成功,返回topic_id"
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
	id, err := p.TopicUsecase.Create(bizCtx, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, strconv.FormatInt(id, 10))

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

	if e := req.Validate(); e != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

		return
	}

	id, exist := ctx.Params.Get("id")
	if !exist {
		p.SuccessResp(ctx, nil)
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	err = p.TopicUsecase.Update(bizCtx, topicID, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return
}

// @Summary 获取话题
// @Description 获取话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true "ID"
// @Success 200 {object} models.Topic "话题"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id} [get]
func (p *TopicCtrl) Get(ctx *gin.Context) {

	id, exist := ctx.Params.Get("id")
	if !exist {
		p.SuccessResp(ctx, nil)
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	item, err := p.TopicUsecase.Get(bizCtx, topicID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, item)

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
