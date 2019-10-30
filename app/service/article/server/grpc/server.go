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
	"valerian/library/xstr"

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

	article, err := s.svr.GetArticle(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	changeDesc, err := s.svr.GetArticleLastChangeDesc(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := s.svr.GetArticleStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := s.svr.GetArticleImageUrls(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := s.svr.GetAccountBaseInfo(ctx, article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp := &api.ArticleInfo{
		ID:             article.ID,
		Title:          article.Title,
		Excerpt:        xstr.Excerpt(article.ContentText),
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		ImageUrls:      urls,
		DisableRevise:  bool(article.DisableRevise),
		DisableComment: bool(article.DisableComment),
		Stat: &api.ArticleStat{
			ReviseCount:  int32(stat.ReviseCount),
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ChangeDesc: changeDesc,
	}

	inc := includeParam(req.Include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return resp, nil
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

	revise, err := s.svr.GetRevise(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	article, err := s.svr.GetArticle(ctx, revise.ArticleID)
	if err != nil {
		return nil, err
	}

	stat, err := s.svr.GetReviseStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	urls, err := s.svr.GetReviseImageUrls(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := s.svr.GetAccountBaseInfo(ctx, article.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp := &api.ReviseInfo{
		ID:        revise.ID,
		Title:     article.Title,
		Excerpt:   xstr.Excerpt(revise.ContentText),
		CreatedAt: revise.CreatedAt,
		UpdatedAt: revise.UpdatedAt,
		ImageUrls: urls,
		Stat: &api.ReviseStat{
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:           m.ID,
			UserName:     m.UserName,
			Avatar:       m.Avatar,
			Introduction: m.Introduction,
		},
		ArticleID: revise.ArticleID,
	}

	inc := includeParam(req.Include)

	if inc["content"] {
		resp.Content = article.Content
	}

	if inc["content_text"] {
		resp.ContentText = article.ContentText
	}

	return resp, nil
}

func (s *server) GetUserArticlesPaged(c context.Context, req *api.UserArticlesReq) (*api.UserArticlesResp, error) {
	items, err := s.svr.GetUserArticlesPaged(c, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	resp := &api.UserArticlesResp{
		Items: make([]*api.ArticleInfo, len(items)),
	}

	for i, v := range items {
		stat, err := s.svr.GetArticleStat(c, v.ID)
		if err != nil {
			return nil, err
		}

		urls, err := s.svr.GetArticleImageUrls(c, v.ID)
		if err != nil {
			return nil, err
		}

		m, err := s.svr.GetAccountBaseInfo(c, v.CreatedBy)
		if err != nil {
			return nil, err
		}

		info := &api.ArticleInfo{
			ID:        v.ID,
			Title:     v.Title,
			Excerpt:   xstr.Excerpt(v.ContentText),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			ImageUrls: urls,
			Stat: &api.ArticleStat{
				ReviseCount:  int32(stat.ReviseCount),
				CommentCount: int32(stat.CommentCount),
				LikeCount:    int32(stat.LikeCount),
				DislikeCount: int32(stat.DislikeCount),
			},
			Creator: &api.Creator{
				ID:           m.ID,
				UserName:     m.UserName,
				Avatar:       m.Avatar,
				Introduction: m.Introduction,
			},
		}

		resp.Items[i] = info

	}

	return resp, nil
}

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}
