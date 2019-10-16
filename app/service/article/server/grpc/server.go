package server

import (
	"context"
	"fmt"
	"valerian/app/service/article/api"
	"valerian/app/service/article/service"
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
	article, err := s.svr.GetArticle(ctx, req.ID)
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
		ID:        article.ID,
		Title:     article.Title,
		Excerpt:   xstr.Excerpt(article.ContentText),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		ImageUrls: urls,
		Stat: &api.ArticleStat{
			ReviseCount:  int32(stat.ReviseCount),
			CommentCount: int32(stat.CommentCount),
			LikeCount:    int32(stat.LikeCount),
			DislikeCount: int32(stat.DislikeCount),
		},
		Creator: &api.Creator{
			ID:       m.ID,
			UserName: m.UserName,
			Avatar:   m.Avatar,
		},
	}

	if m.Introduction != nil {
		resp.Creator.Introduction = &api.Creator_IntroductionValue{m.GetIntroductionValue()}
	}

	return resp, nil
}

func (s *server) GetReviseInfo(ctx context.Context, req *api.IDReq) (*api.ReviseInfo, error) {
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
			ID:       m.ID,
			UserName: m.UserName,
			Avatar:   m.Avatar,
		},
	}

	if m.Introduction != nil {
		resp.Creator.Introduction = &api.Creator_IntroductionValue{m.GetIntroductionValue()}
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
				ID:       m.ID,
				UserName: m.UserName,
				Avatar:   m.Avatar,
			},
		}

		fmt.Printf("%+v\n", resp)

		if m.Introduction != nil {
			info.Creator.Introduction = &api.Creator_IntroductionValue{m.GetIntroductionValue()}
		}

		resp.Items[i] = info

	}

	return resp, nil
}
