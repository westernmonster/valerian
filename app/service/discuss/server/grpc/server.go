package server

import (
	"context"
	"fmt"
	"strings"

	"valerian/app/service/discuss/api"
	"valerian/app/service/discuss/service"
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
	api.RegisterDiscussionServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) GetDiscussionStat(ctx context.Context, req *api.IDReq) (*api.DiscussionStat, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	stat, err := s.svr.GetDiscussionStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp := &api.DiscussionStat{
		LikeCount:    int32(stat.LikeCount),
		DislikeCount: int32(stat.DislikeCount),
		CommentCount: int32(stat.CommentCount),
	}
	return resp, err
}

func (s *server) GetDiscussionInfo(ctx context.Context, req *api.IDReq) (*api.DiscussionInfo, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	v, imgs, err := s.svr.GetDiscussion(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	stat, err := s.svr.GetDiscussionStat(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	m, err := s.svr.GetAccountBaseInfo(ctx, v.CreatedBy)
	if err != nil {
		return nil, err
	}

	c, err := s.svr.GetDiscussCategory(ctx, v.CategoryID)
	if err != nil {
		return nil, err
	}

	resp := &api.DiscussionInfo{
		ID:         v.ID,
		TopicID:    v.TopicID,
		CategoryID: v.CategoryID,
		// CreatedBy:   v.CreatedBy,
		Excerpt:   xstr.Excerpt(v.ContentText),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		ImageUrls: imgs,
		Title:     v.Title,
	}

	inc := includeParam(req.Include)
	if inc["content"] {
		resp.Content = v.Content
	}

	if inc["content_text"] {
		resp.ContentText = v.ContentText
	}

	resp.Stat = &api.DiscussionStat{
		DislikeCount: int32(stat.DislikeCount),
		LikeCount:    int32(stat.LikeCount),
		CommentCount: int32(stat.CommentCount),
	}

	resp.CategoryInfo = &api.CategoryInfo{
		ID:      c.ID,
		TopicID: c.TopicID,
		Name:    c.Name,
		Seq:     c.Seq,
	}

	resp.Creator = &api.Creator{
		ID:           m.ID,
		UserName:     m.UserName,
		Avatar:       m.Avatar,
		Introduction: m.Introduction,
	}

	return resp, err
}

func (s *server) GetDiscussionCategories(ctx context.Context, req *api.CategoriesReq) (*api.CategoriesResp, error) {
	if req.UseMaster {
		ctx = sqalx.NewContext(ctx, true)
		defer func() {
			ctx = sqalx.NewContext(ctx, false)
		}()
	}
	data, err := s.svr.GetDiscussCategories(ctx, req.TopicID)
	if err != nil {
		return nil, err
	}

	resp := &api.CategoriesResp{
		Items: make([]*api.CategoryInfo, len(data)),
	}

	for i, v := range data {
		item := api.FromCategory(v)
		resp.Items[i] = item
	}

	return resp, nil
}

func (s *server) GetUserDiscussionsPaged(ctx context.Context, req *api.UserDiscussionsReq) (*api.UserDiscussionsResp, error) {
	items, err := s.svr.GetUserDiscussionsPaged(ctx, req.AccountID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	resp := &api.UserDiscussionsResp{
		Items: items,
	}

	return resp, nil
}

func (s *server) GetAllDiscussions(ctx context.Context, req *api.EmptyStruct) (*api.AllDiscussionsResp, error) {
	items, err := s.svr.GetAllDiscussions(ctx)
	if err != nil {
		return nil, err
	}

	resp := &api.AllDiscussionsResp{
		Items: items,
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
