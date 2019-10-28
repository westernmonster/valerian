package server

import (
	"context"
	"fmt"

	"valerian/app/service/search/api"
	"valerian/app/service/search/model"
	"valerian/app/service/search/service"
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
	api.RegisterSearchServer(w.Server(), &server{s})
	ws, err := w.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	svr *service.Service
}

func (s *server) SearchTopic(ctx context.Context, req *api.SearchParam) (*api.SearchResult, error) {
	ret, err := s.svr.TopicSearch(ctx, &model.BasicSearchParams{
		KW:       req.KW,
		KwFields: req.KwFields,
		Order:    req.Order,
		Sort:     req.Sort,
		Pn:       int(req.Pn),
		Ps:       int(req.Ps),
		Debug:    req.Debug,
		Source:   req.Source,
	})
	if err != nil {
		return nil, err
	}

	resp := &api.SearchResult{
		Order:  ret.Order,
		Sort:   ret.Sort,
		Result: make([][]byte, len(ret.Result)),
		Page: &api.Page{
			Pn:    int32(ret.Page.Pn),
			Ps:    int32(ret.Page.Ps),
			Total: ret.Page.Total,
		},
		Debug: ret.Debug,
	}

	for i, v := range ret.Result {
		resp.Result[i] = []byte(v)
	}

	return resp, nil
}

func (s *server) SearchAccount(ctx context.Context, req *api.SearchParam) (*api.SearchResult, error) {
	ret, err := s.svr.AccountSearch(ctx, &model.BasicSearchParams{
		KW:       req.KW,
		KwFields: req.KwFields,
		Order:    req.Order,
		Sort:     req.Sort,
		Pn:       int(req.Pn),
		Ps:       int(req.Ps),
		Debug:    req.Debug,
		Source:   req.Source,
	})
	if err != nil {
		return nil, err
	}

	resp := &api.SearchResult{
		Order:  ret.Order,
		Sort:   ret.Sort,
		Result: make([][]byte, len(ret.Result)),
		Page: &api.Page{
			Pn:    int32(ret.Page.Pn),
			Ps:    int32(ret.Page.Ps),
			Total: ret.Page.Total,
		},
		Debug: ret.Debug,
	}

	for i, v := range ret.Result {
		resp.Result[i] = []byte(v)
	}

	return resp, nil
}

func (s *server) SearchArticle(ctx context.Context, req *api.SearchParam) (*api.SearchResult, error) {
	ret, err := s.svr.ArticleSearch(ctx, &model.BasicSearchParams{
		KW:       req.KW,
		KwFields: req.KwFields,
		Order:    req.Order,
		Sort:     req.Sort,
		Pn:       int(req.Pn),
		Ps:       int(req.Ps),
		Debug:    req.Debug,
		Source:   req.Source,
	})
	if err != nil {
		return nil, err
	}

	resp := &api.SearchResult{
		Order:  ret.Order,
		Sort:   ret.Sort,
		Result: make([][]byte, len(ret.Result)),
		Page: &api.Page{
			Pn:    int32(ret.Page.Pn),
			Ps:    int32(ret.Page.Ps),
			Total: ret.Page.Total,
		},
		Debug: ret.Debug,
	}

	for i, v := range ret.Result {
		resp.Result[i] = []byte(v)
	}

	return resp, nil
}

func (s *server) SearchDiscussion(ctx context.Context, req *api.SearchParam) (*api.SearchResult, error) {
	ret, err := s.svr.DiscussionSearch(ctx, &model.BasicSearchParams{
		KW:       req.KW,
		KwFields: req.KwFields,
		Order:    req.Order,
		Sort:     req.Sort,
		Pn:       int(req.Pn),
		Ps:       int(req.Ps),
		Debug:    req.Debug,
		Source:   req.Source,
	})
	if err != nil {
		return nil, err
	}

	resp := &api.SearchResult{
		Order:  ret.Order,
		Sort:   ret.Sort,
		Result: make([][]byte, len(ret.Result)),
		Page: &api.Page{
			Pn:    int32(ret.Page.Pn),
			Ps:    int32(ret.Page.Ps),
			Total: ret.Page.Total,
		},
		Debug: ret.Debug,
	}

	for i, v := range ret.Result {
		resp.Result[i] = []byte(v)
	}

	return resp, nil
}
