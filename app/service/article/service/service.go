package service

import (
	"context"
	"strings"
	"valerian/app/service/article/api"

	"valerian/app/service/article/conf"
	"valerian/app/service/article/dao"
	"valerian/library/conf/env"
	"valerian/library/log"
	"valerian/library/mq"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"
	"github.com/pkg/errors"
)

// Service struct of service
type Service struct {
	c         *conf.Config
	d         *dao.Dao
	mq        *mq.MessageQueue
	immClient *imm.Client
	missch    chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		mq:     mq.New(env.Hostname, c.Nats),
		missch: make(chan func(), 1024),
	}

	if client, err := imm.NewClientWithAccessKey(c.Aliyun.RegionID, c.Aliyun.AccessKeyID, c.Aliyun.AccessKeySecret); err != nil {
		panic(errors.WithMessage(err, "Failed to init Aliyun IMM Client"))
	} else {
		s.immClient = client
	}
	go s.cacheproc()
	return
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return s.d.Ping(c)
}

// Close dao.
func (s *Service) Close() {
	s.d.Close()
}

func (s *Service) addCache(f func()) {
	select {
	case s.missch <- f:
	default:
		log.Warn("cacheproc chan full")
	}
}

// cacheproc is a routine for executing closure.
func (s *Service) cacheproc() {
	for {
		f := <-s.missch
		f()
	}
}

func (s *Service) PullArticleAppCache(ctx context.Context, req *api.IdUpdatedReq) (results []*api.DBArticle, err error) {
	var ids []int64
	for _, idUpdated := range req.Items {
		ids = append(ids, idUpdated.ID)
	}
	items, err := s.d.GetArticlesByIDs(ctx, s.d.DB(), ids)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		for _, idUpdate := range req.Items {
			if item.ID == idUpdate.ID && item.UpdatedAt != idUpdate.UpdatedAt {
				article := api.DBArticle{
					ID:             item.ID,
					Title:          item.Title,
					Content:        item.Content,
					ContentText:    item.ContentText,
					DisableRevise:  bool(item.DisableRevise),
					DisableComment: bool(item.DisableComment),
					CreatedBy:      item.CreatedBy,
					CreatedAt:      item.CreatedAt,
					UpdatedAt:      item.UpdatedAt,
				}
				results = append(results, &article)
			}
		}
	}
	return
}

func (s *Service) PullReviseAppCache(ctx context.Context, req *api.IdUpdatedReq) (results []*api.ReviseDetail, err error) {
	var ids []int64
	for _, idUpdated := range req.Items {
		ids = append(ids, idUpdated.ID)
	}
	items, err := s.d.GetRevisesByIDs(ctx, s.d.DB(), ids)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		for _, idUpdate := range req.Items {
			if item.ID == idUpdate.ID && item.UpdatedAt != idUpdate.UpdatedAt {
				revise := api.ReviseDetail{
					ID:          item.ID,
					Title:       item.Title,
					CreatedAt:   item.CreatedAt,
					UpdatedAt:   item.UpdatedAt,
					ArticleID:   item.ArticleID,
					Content:     item.Content,
					ContentText: item.ContentText,
				}
				results = append(results, &revise)
			}
		}
	}
	return
}

func includeParam(include string) (dic map[string]bool) {
	arr := strings.Split(include, ",")
	dic = make(map[string]bool)
	for _, v := range arr {
		dic[v] = true
	}

	return
}
