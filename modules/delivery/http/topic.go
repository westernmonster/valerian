package http

import (
	"net/http"
	"strconv"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"

	"valerian/library/database/sqalx"
	"valerian/library/database/sqlx"

	"github.com/gin-gonic/gin"
)

type TopicCtrl struct {
	infrastructure.BaseCtrl

	TopicUsecase interface {
		Create(ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error)
		Update(ctx *biz.BizContext, id int64, req *models.UpdateTopicReq) (err error)
		Delete(ctx *biz.BizContext, id int64) (err error)
		Get(ctx *biz.BizContext, topicID int64) (item *models.Topic, err error)
		SeachTopics(ctx *biz.BizContext, topicID int64, query string) (items []*models.TopicSearchResult, err error)
		GetTopicMembersPaged(ctx *biz.BizContext, topicID int64, page, pageSize int) (resp *models.TopicMembersPagedResp, err error)
		BulkSaveMembers(ctx *biz.BizContext, topicID int64, req *models.BatchSavedTopicMemberReq) (err error)
		GetTopicVersions(ctx *biz.BizContext, topicSetID int64) (items []*models.TopicVersion, err error)
		FollowTopic(ctx *biz.BizContext, topicID int64, isFollowed bool) (err error)
		GetAllRelatedTopics(ctx *biz.BizContext, topicID int64) (items []*models.RelatedTopic, err error)
		GetAllTopicTypes(ctx *biz.BizContext) (items []*models.TopicType, err error)
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
			TopicFollowerRepository: &repo.TopicFollowerRepository{},
		},
	}

}

// @Summary 获取关联话题
// @Description 获取关联话题
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true "话题ID"
// @Success 200 {array} models.RelatedTopic "话题成员"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id}/related [get]
func (p *TopicCtrl) GetAllRelatedTopics(ctx *gin.Context) {
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
	items, err := p.TopicUsecase.GetAllRelatedTopics(bizCtx, topicID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

}

// @Summary 获取成员列表
// @Description 获取成员列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true "话题ID"
// @Param page query integer true "页码"
// @Param page_size query integer true "每页大小"
// @Success 200 {object} models.TopicMembersPagedResp "话题成员"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id}/members [get]
func (p *TopicCtrl) GetTopicMembersPaged(ctx *gin.Context) {
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

	page := 1
	pageSize := 20

	if p, ok := ctx.GetQuery("page"); ok {
		if v, e := strconv.Atoi(p); e != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "页码不是整数格式",
			})
		} else {
			page = v
		}
	}

	if p, ok := ctx.GetQuery("page_size"); ok {
		if v, e := strconv.Atoi(p); e != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "每页大小不是整数格式",
			})
		} else {
			pageSize = v
		}
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetTopicMembersPaged(bizCtx, topicID, page, pageSize)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

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

// @Summary 关注/取关话题
// @Description 关注/取关话题
// @Tags account
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.TopicFollower true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /me/followed/topics [post]
func (p *TopicCtrl) FollowTopic(ctx *gin.Context) {
	req := new(models.TopicFollower)

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
	err := p.TopicUsecase.FollowTopic(bizCtx, req.TopicID, req.IsFollowed)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

	return

}

// @Summary 批量更新话题成员
// @Description 批量更新话题成员
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body models.BatchSavedTopicMemberReq true "请求"
// @Success 200 "成功"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id}/members [patch]
func (p *TopicCtrl) BatchSavedTopicMember(ctx *gin.Context) {
	req := new(models.BatchSavedTopicMemberReq)

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
	err = p.TopicUsecase.BulkSaveMembers(bizCtx, topicID, req)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, nil)

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

// @Summary 获取话题版本
// @Description 获取话题版本
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param id path string true  "话题集合ID"
// @Success 200 {array} models.TopicVersion "话题成员"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_sets/{id}/versions [get]
func (p *TopicCtrl) GetTopicVersions(ctx *gin.Context) {
	id := ctx.Param("id")
	topicSetID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.RespCommon{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "请提供话题集合ID",
		})

		return
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetTopicVersions(bizCtx, topicSetID)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

}

// @Summary 获取话题类型
// @Description 获取话题类型
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {array} models.TopicType "话题类型"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topic_types [get]
func (p *TopicCtrl) GetAllTopicTypes(ctx *gin.Context) {

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetAllTopicTypes(bizCtx)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, items)

	return

}
