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
		GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicVersionID, parentID int64) (count int, err error)
		GetTopicCatalogsByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error)
		GetTopicCatalogByCondition(ctx context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error)
		GetTopicCatalogMaxChildrenSeq(c context.Context, node sqalx.Node, topicVersionID, parentID int64) (seq int, err error)
		GetTopicCatalogsCountByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (total int, err error)

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
		DelTopicRelation(c context.Context, node sqalx.Node, id int64) (err error)

		GetTopicVersion(c context.Context, node sqalx.Node, id int64) (item *model.TopicVersion, err error)
		AddTopicVersion(c context.Context, node sqalx.Node, item *model.TopicVersion) (err error)
		UpdateTopicVersion(c context.Context, node sqalx.Node, item *model.TopicVersion) (err error)
		GetTopicVersions(c context.Context, node sqalx.Node, topicID int64) (items []*model.TopicVersionResp, err error)
		GetTopicVersionByName(c context.Context, node sqalx.Node, topicID int64, versionName string) (item *model.TopicVersionResp, err error)
		GetTopicVersionMaxSeq(c context.Context, node sqalx.Node, topicID int64) (seq int, err error)

		GetAllTopicTypes(c context.Context, node sqalx.Node) (items []*model.TopicType, err error)
		GetTopicType(c context.Context, node sqalx.Node, id int) (item *model.TopicType, err error)

		GetTopicFollowRequests(c context.Context, node sqalx.Node, topicID int64, status int) (items []*model.TopicFollowRequest, err error)
		GetTopicFollowRequest(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicFollowRequest, err error)
		AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
		UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error)
		DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error)

		GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error)
		GetAccountTopicSettings(c context.Context, node sqalx.Node, aid int64) (items []*model.AccountTopicSetting, err error)
		AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
		UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error)
		DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error)
		GetAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error)

		AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error)
		UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error)
		DelArticle(c context.Context, node sqalx.Node, id int64) (err error)

		GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error)
		AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error)
		DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error)
		GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error)

		AddArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error)
		UpdateArticleVersion(c context.Context, node sqalx.Node, item *model.ArticleVersion) (err error)
		DelArticleVersion(c context.Context, node sqalx.Node, id int64) (err error)
		GetArticleVersion(c context.Context, node sqalx.Node, id int64) (item *model.ArticleVersion, err error)
		GetArticleVersionByName(c context.Context, node sqalx.Node, articleID int64, versionName string) (item *model.ArticleVersionResp, err error)
		GetArticleVersions(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleVersionResp, err error)
		GetArticleVersionsMaxSeq(c context.Context, node sqalx.Node, articleID int64) (seq int, err error)

		AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error)
		GetArticleHistories(c context.Context, node sqalx.Node, articleVersionID int64) (items []*model.ArticleHistory, err error)
		GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error)
		GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleVersionID int64) (seq int, err error)

		AddAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
		UpdateAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error)
		DelAccountArticleAttr(c context.Context, node sqalx.Node, id int64) (err error)
		GetAccountArticleAttr(c context.Context, node sqalx.Node, aid, articleID int64) (item *model.AccountArticleAttr, err error)
		GetArticleFavCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)
		GetArticleLikeCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error)

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

		SetTopicVersionCache(c context.Context, topicID int64, m []*model.TopicVersionResp) (err error)
		TopicVersionCache(c context.Context, topicID int64) (m []*model.TopicVersionResp, err error)
		DelTopicVersionCache(c context.Context, topicID int64) (err error)

		SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int, data []*model.TopicMember) (err error)
		TopicMembersCache(c context.Context, topicID int64, page, pageSize int) (count int, data []*model.TopicMember, err error)
		DelTopicMembersCache(c context.Context, topicID int64) (err error)

		SetArticleCache(c context.Context, m *model.ArticleResp) (err error)
		ArticleCache(c context.Context, articleID int64) (m *model.ArticleResp, err error)
		DelArticleCache(c context.Context, articleID int64) (err error)

		SetArticleFileCache(c context.Context, articleID int64, m []*model.ArticleFileResp) (err error)
		ArticleFileCache(c context.Context, articleID int64) (m []*model.ArticleFileResp, err error)
		DelArticleFileCache(c context.Context, articleID int64) (err error)

		SetArticleVersionsCache(c context.Context, articleID int64, m []*model.ArticleVersionResp) (err error)
		ArticleVersionsCache(c context.Context, articleID int64) (m []*model.ArticleVersionResp, err error)
		DelArticleVersionsCache(c context.Context, articleID int64) (err error)

		SetArticleHistoryCache(c context.Context, articleVersionID int64, m []*model.ArticleHistoryResp) (err error)
		ArticleHistoryCache(c context.Context, articleVersionID int64) (m []*model.ArticleHistoryResp, err error)
		DelArticleHistoryCache(c context.Context, articleVersionID int64) (err error)

		SetAccountCache(c context.Context, m *model.Account) (err error)
		AccountCache(c context.Context, accountID int64) (m *model.Account, err error)
		DelAccountCache(c context.Context, accountID int64) (err error)

		SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error)
		AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error)
		DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error)

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
