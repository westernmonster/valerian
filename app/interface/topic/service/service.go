package service

import (
	"context"
	"valerian/app/conf"
	"valerian/app/interface/topic/dao"
	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
	"valerian/library/log"
)

// Service struct of service
type Service struct {
	c *conf.Config
	d interface {
		GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error)
		AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
		UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error)
		DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error)
		GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error)
		GetTopicCatalogsByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error)
		GetTopicCatalogByCondition(ctx context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error)
		GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicID, parentID int64) (seq int, err error)

		GetAllTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error)
		GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error)
		AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)
		DelTopic(c context.Context, node sqalx.Node, topicID int64) (err error)
		UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error)

		GetTopicMemberByCondition(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicMember, err error)
		GetAllTopicMembers(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicMember, err error)
		GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error)
		GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*model.TopicMember, err error)
		AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
		UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error)
		DeleteTopicMember(c context.Context, node sqalx.Node, id int64) (err error)

		GetTopicRelationByCondition(c context.Context, node sqalx.Node, fromTopicID, toTopicID int64) (item *model.TopicRelation, err error)
		GetAllTopicRelations(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicRelation, err error)
		AddTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error)
		UpdateTopicRelation(c context.Context, node sqalx.Node, item *model.TopicRelation) (err error)
		DeleteTopicRelation(c context.Context, node sqalx.Node, id int64) (err error)

		AddTopicSet(c context.Context, node sqalx.Node, item *model.TopicSet) (err error)
		GetTopicVersions(c context.Context, node sqalx.Node, topicSetID int64) (items []int64, err error)
		GetTopicVersionByName(c context.Context, node sqalx.Node, topicSetID int64, versionName string) (item *model.TopicVersionResp, err error)
		GetAllTopicTypes(c context.Context, node sqalx.Node) (items []*model.TopicType, err error)
		GetTopicType(c context.Context, node sqalx.Node, id int) (item *model.TopicType, err error)

		GetTopicFollowRequests(c context.Context, node sqalx.Node, topicID int64, status int) (items []*model.TopicFollowRequest, err error)
		GetTopicFollowRequest(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicFollowRequest, err error)
		AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
		UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
		DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error)

		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)

		AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
		UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

		GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error)
		AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error)
		GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error)

		AddArticleSet(c context.Context, node sqalx.Node, item *model.ArticleSet) (err error)
		DelArticleSet(c context.Context, node sqalx.Node, id int64) (err error)
		GetArticleVersionByName(c context.Context, node sqalx.Node, articleSetID int64, versionName string) (item *model.ArticleVersionResp, err error)

		AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error)
		GetArticleHistories(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleHistory, err error)
		GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error)
		GetArticleHistoryMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error)
		GetOrderMemberArticleHistoriesCount(c context.Context, node sqalx.Node, articleID int64, aid int64) (count int, err error)

		GetLocaleByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Locale, err error)

		SetTopicCache(c context.Context, m *model.TopicResp) (err error)
		TopicCache(c context.Context, topicID int64) (m *model.TopicResp, err error)
		DelTopicCache(c context.Context, topicID int64) (err error)

		SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error)
		TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error)
		DelTopicCatalogCache(c context.Context, topicID int64) (err error)

		SetTopicRelationCache(c context.Context, topicID int64, m []*model.TopicRelation) (err error)
		TopicRelationCache(c context.Context, topicID int64) (m []*model.TopicRelation, err error)
		DelTopicRelationCache(c context.Context, topicID int64) (err error)

		SetTopicVersionCache(c context.Context, topicSetID int64, m []int64) (err error)
		TopicVersionCache(c context.Context, topicSetID int64) (m []int64, err error)
		DelTopicVersionCache(c context.Context, topicSetID int64) (err error)

		SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int, data []*model.TopicMember) (err error)
		TopicMembersCache(c context.Context, topicID int64, page, pageSize int) (count int, data []*model.TopicMember, err error)
		DelTopicMembersCache(c context.Context, topicID int64) (err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)

		Ping(c context.Context) (err error)
		Close()
		DB() sqalx.Node
	}
	missch chan func()
}

// New create new service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		d:      dao.New(c),
		missch: make(chan func(), 1024),
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
