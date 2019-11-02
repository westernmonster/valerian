package server

import (
	"context"
	"fmt"
	"strings"
	"valerian/app/service/article/api"
	"valerian/app/service/article/service"
	"valerian/library/database/sqalx"
	"valerian/library/log"
	"valerian/library/net/metadata"
	"valerian/library/net/rpc/warden"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// New Identify warden rpc server
func New(cfg *warden.ServerConfig, s *service.Service) *warden.Server {
	w := warden.NewServer(cfg)
	w.Use(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if resp, err = handler(ctx, req); err != nil {
			log.For(ctx).Info("rpc call",
				zap.String("path", info.FullMethod),
				zap.String("caller", metadata.String(ctx, metadata.Caller)),
				zap.String("args", fmt.Sprintf("%v", req)),
				zap.String("error", fmt.Sprintf("%+v", err)))
		}
		return
	})
	api.RegisterArticleServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetArticleInfo(ctx context.Context, req *api.IDReq) (*api.ArticleInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}

	return s.svr.GetArticleInfo(ctx, req)
}

func (s *server) GetArticleDetail(ctx context.Context, req *api.IDReq) (*api.ArticleDetail, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}

	return s.svr.GetArticleDetail(ctx, req)
}

func (s *server) GetAllArticles(ctx context.Context, req *api.EmptyStruct) (*api.ArticlesResp, error) {
	resp := &api.ArticlesResp{
		Items: make([]*api.DBArticle, 0),
	}

	items, err := s.svr.GetAllArticles(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range items {
		item := &api.DBArticle{
			ID:             v.ID,
			Title:          v.Title,
			Content:        v.Content,
			ContentText:    v.ContentText,
			DisableRevise:  bool(v.DisableRevise),
			DisableComment: bool(v.DisableComment),
			CreatedBy:      v.CreatedBy,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		}

		resp.Items = append(resp.Items, item)
	}

	return resp, nil
}

func (s *server) GetReviseStat(ctx context.Context, req *api.IDReq) (*api.ReviseStat, error) {
	stat, err := s.svr.GetReviseStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.ReviseStat{
		CommentCount: int32(stat.CommentCount),
		LikeCount:    int32(stat.LikeCount),
		DislikeCount: int32(stat.DislikeCount),
	}

	return resp, nil
}

func (s *server) AddArticle(ctx context.Context, req *api.ArgAddArticle) (*api.IDResp, error) {
	id, err := s.svr.AddArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.IDResp{
		ID: id,
	}

	return resp, nil
}

func (s *server) UpdateArticle(ctx context.Context, req *api.ArgUpdateArticle) (*api.EmptyStruct, error) {
	err := s.svr.UpdateArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) DelArticle(ctx context.Context, req *api.IDReq) (*api.EmptyStruct, error) {
	err := s.svr.DelArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) GetArticleFiles(ctx context.Context, req *api.IDReq) (*api.ArticleFilesResp, error) {
	data, err := s.svr.GetArticleFiles(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.ArticleFilesResp{Items: data}, nil
}

func (s *server) SaveArticleFiles(ctx context.Context, req *api.ArgSaveArticleFiles) (*api.EmptyStruct, error) {
	err := s.svr.SaveArticleFiles(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) GetArticleHistoriesPaged(ctx context.Context, req *api.ArgArticleHistoriesPaged) (*api.ArticleHistoryListResp, error) {
	resp, err := s.svr.GetArticleHistoriesPaged(ctx, req.ArticleID, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) GetArticleHistory(ctx context.Context, req *api.IDReq) (*api.ArticleHistoryResp, error) {
	resp, err := s.svr.GetArticleHistory(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) GetArticleRelations(ctx context.Context, req *api.IDReq) (*api.ArticleRelationsResp, error) {
	data, err := s.svr.GetArticleRelations(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &api.ArticleRelationsResp{Items: data}, nil
}

func (s *server) UpdateArticleRelation(ctx context.Context, req *api.ArgUpdateArticleRelation) (*api.EmptyStruct, error) {
	err := s.svr.UpdateArticleRelation(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) SetPrimary(ctx context.Context, req *api.ArgSetPrimaryArticleRelation) (*api.EmptyStruct, error) {
	err := s.svr.SetPrimary(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) AddArticleRelation(ctx context.Context, req *api.ArgAddArticleRelation) (*api.EmptyStruct, error) {
	err := s.svr.AddArticleRelation(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) DelArticleRelation(ctx context.Context, req *api.ArgDelArticleRelation) (*api.EmptyStruct, error) {
	err := s.svr.DelArticleRelation(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) GetArticleStat(ctx context.Context, req *api.IDReq) (*api.ArticleStat, error) {
	stat, err := s.svr.GetArticleStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.ArticleStat{
		CommentCount: int32(stat.CommentCount),
		ReviseCount:  int32(stat.ReviseCount),
		LikeCount:    int32(stat.LikeCount),
		DislikeCount: int32(stat.DislikeCount),
	}

	return resp, nil
}

func (s *server) GetReviseInfo(ctx context.Context, req *api.IDReq) (*api.ReviseInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}

	return s.svr.GetReviseInfo(ctx, req)
}

func (s *server) GetUserArticlesPaged(c context.Context, req *api.UserArticlesReq) (*api.UserArticlesResp, error) {
	return s.svr.GetUserArticlesPaged(c, req)
}

func (s *server) AddRevise(ctx context.Context, req *api.ArgAddRevise) (*api.IDResp, error) {
	id, err := s.svr.AddRevise(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.IDResp{
		ID: id,
	}

	return resp, nil
}

func (s *server) UpdateRevise(ctx context.Context, req *api.ArgUpdateRevise) (*api.EmptyStruct, error) {
	err := s.svr.UpdateRevise(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) DelRevise(ctx context.Context, req *api.IDReq) (*api.EmptyStruct, error) {
	err := s.svr.DelRevise(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &api.EmptyStruct{}

	return resp, nil
}

func (s *server) GetArticleRevisesPaged(ctx context.Context, req *api.ArgArticleRevisesPaged) (*api.ReviseListResp, error) {
	resp, err := s.svr.GetArticleRevisesPaged(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *server) GetReviseFiles(ctx context.Context, req *api.IDReq) (*api.ReviseFilesResp, error) {
	data, err := s.svr.GetReviseFiles(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.ReviseFilesResp{Items: data}, nil
}

func (s *server) SaveReviseFiles(ctx context.Context, req *api.ArgSaveReviseFiles) (*api.EmptyStruct, error) {
	err := s.svr.SaveReviseFiles(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.EmptyStruct{}, nil
}

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}
