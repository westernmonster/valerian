package http

import (
	"context"
	"strconv"

	"valerian/infrastructure"
	"valerian/infrastructure/biz"
	"valerian/models"
	"valerian/modules/repo"
	"valerian/modules/usecase"

	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

type TopicCtrl struct {
	infrastructure.BaseCtrl

	TopicUsecase interface {
		Create(c context.Context, ctx *biz.BizContext, req *models.CreateTopicReq) (id int64, err error)
		Update(c context.Context, ctx *biz.BizContext, id int64, req *models.UpdateTopicReq) (err error)
		Delete(c context.Context, ctx *biz.BizContext, id int64) (err error)
		Get(c context.Context, ctx *biz.BizContext, topicID int64) (item *models.Topic, err error)
		SeachTopics(c context.Context, ctx *biz.BizContext, query string) (items []*models.TopicSearchResult, err error)
		GetTopicMembersPaged(c context.Context, ctx *biz.BizContext, topicID int64, page, pageSize int) (resp *models.TopicMembersPagedResp, err error)
		BulkSaveMembers(c context.Context, ctx *biz.BizContext, topicID int64, req *models.BatchSavedTopicMemberReq) (err error)
		GetTopicVersions(c context.Context, ctx *biz.BizContext, topicSetID int64) (items []*models.TopicVersion, err error)
		FollowTopic(c context.Context, ctx *biz.BizContext, topicID int64, isFollowed bool) (err error)
		GetAllRelatedTopics(c context.Context, ctx *biz.BizContext, topicID int64) (items []*models.RelatedTopic, err error)
		GetAllTopicTypes(c context.Context, ctx *biz.BizContext) (items []*models.TopicType, err error)
		CreateNewVersion(c context.Context, ctx *biz.BizContext, arg *models.ArgNewVersion) (id int64, err error)
		SaveCatalogs(c context.Context, ctx *biz.BizContext, topicID int64, req *models.ArgSaveTopicCatalog) (err error)
	}
}

func NewTopicCtrl(node sqalx.Node) *TopicCtrl {
	return &TopicCtrl{
		TopicUsecase: &usecase.TopicUsecase{
			Node:                    node,
			TopicRepository:         &repo.TopicRepository{},
			AccountRepository:       &repo.AccountRepository{},
			TopicMemberRepository:   &repo.TopicMemberRepository{},
			TopicCatalogRepository:  &repo.TopicCatalogRepository{},
			TopicTypeRepository:     &repo.TopicTypeRepository{},
			TopicSetRepository:      &repo.TopicSetRepository{},
			TopicRelationRepository: &repo.TopicRelationRepository{},
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
// @Param topic_id query string true "话题ID"
// @Success 200 {array} models.RelatedTopic "话题成员"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/related [get]
func (p *TopicCtrl) GetAllRelatedTopics(ctx *mars.Context) {
	idStr := ctx.Request.Form.Get("topic_id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetAllRelatedTopics(ctx.Context, bizCtx, topicID)

	ctx.JSON(items, err)
}

// @Summary 获取成员列表
// @Description 获取成员列表
// @Tags topic
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param topic_id query string true "话题ID"
// @Param page query integer true "页码"
// @Param page_size query integer true "每页大小"
// @Success 200 {object} models.TopicMembersPagedResp "话题成员"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics/{id}/members [get]
func (p *TopicCtrl) GetTopicMembersPaged(ctx *mars.Context) {
	params := ctx.Request.Form
	idStr := params.Get("topic_id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}
	page := 1
	pageSize := 20

	pStr := params.Get("page")
	if v, e := strconv.Atoi(pStr); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
	} else {
		page = v
	}

	pSize := params.Get("page_size")
	if v, e := strconv.Atoi(pSize); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
	} else {
		page = v
	}

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetTopicMembersPaged(ctx.Context, bizCtx, topicID, page, pageSize)

	ctx.JSON(items, err)

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
// @Success 200 {array} models.TopicSearchResult "话题搜索结果"
// @Failure 400 "验证请求失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /topics [get]
func (p *TopicCtrl) Search(ctx *mars.Context) {
	query := ctx.Request.Form.Get("q")

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.SeachTopics(ctx.Context, bizCtx, query)
	ctx.JSON(items, err)

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
func (p *TopicCtrl) Create(ctx *mars.Context) {
	req := new(models.CreateTopicReq)

	ctx.Bind(req)
	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	id, err := p.TopicUsecase.Create(ctx.Context, bizCtx, req)
	ctx.JSON(strconv.FormatInt(id, 10), err)

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
func (p *TopicCtrl) Update(ctx *mars.Context) {
	req := new(models.UpdateTopicReq)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	idStr := ctx.Request.Form.Get("topic_id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	err = p.TopicUsecase.Update(ctx.Context, bizCtx, topicID, req)
	ctx.JSON(nil, err)

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
func (p *TopicCtrl) Get(ctx *mars.Context) {

	idStr := ctx.Request.Form.Get("id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	item, err := p.TopicUsecase.Get(ctx.Context, bizCtx, topicID)
	ctx.JSON(item, err)
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
func (p *TopicCtrl) Delete(ctx *mars.Context) {
	idStr := ctx.Request.Form.Get("id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	err = p.TopicUsecase.Delete(ctx.Context, bizCtx, topicID)
	ctx.JSON(nil, err)
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
func (p *TopicCtrl) GetAllTopicTypes(ctx *mars.Context) {

	bizCtx := p.GetBizContext(ctx)
	items, err := p.TopicUsecase.GetAllTopicTypes(ctx.Context, bizCtx)
	ctx.JSON(items, err)
}

// @Summary 批量更新话题类目
// @Description 批量更新话题类目
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
// @Router /topics/{id}/catalogs [patch]
func (p *TopicCtrl) BatchSavedTopicCatalogs(ctx *mars.Context) {
	req := new(models.ArgSaveTopicCatalog)

	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	idStr := ctx.Request.Form.Get("topic_id")
	topicID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(nil, ecode.RequestErr)
		return
	}

	bizCtx := p.GetBizContext(ctx)
	err = p.TopicUsecase.SaveCatalogs(ctx.Context, bizCtx, topicID, req)
	ctx.JSON(nil, err)
}
