package service

import (
	"context"

	"valerian/app/interface/topic/model"
	"valerian/library/database/sqalx"
)

func (p *MockDao) GetTopicCatalogChildrenCount(c context.Context, node sqalx.Node, topicID, parentID int64) (count int, err error) {
	args := p.Called(c, node, topicID, parentID)
	return args.Int(0), args.Error(1)
}
func (p *MockDao) GetTopicCatalogsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicCatalog, err error) {
	return
}
func (p *MockDao) GetTopicCatalogs(c context.Context, node sqalx.Node) (items []*model.TopicCatalog, err error) {
	return
}
func (p *MockDao) GetTopicCatalogByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicCatalog, err error) {
	return
}
func (p *MockDao) GetTopicCatalogByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicCatalog, err error) {
	return
}
func (p *MockDao) AddTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	return
}
func (p *MockDao) UpdateTopicCatalog(c context.Context, node sqalx.Node, item *model.TopicCatalog) (err error) {
	return
}
func (p *MockDao) DelTopicCatalog(c context.Context, node sqalx.Node, id int64) (err error) {

	return
}

func (p *MockDao) GetAuthTopicsResp(c context.Context, node sqalx.Node, topicID int64) (items []*model.AuthTopicResp, err error) {
	return
}
func (p *MockDao) GetAuthTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.AuthTopic, err error) {
	return
}
func (p *MockDao) GetAuthTopics(c context.Context, node sqalx.Node) (items []*model.AuthTopic, err error) {
	return
}
func (p *MockDao) GetAuthTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.AuthTopic, err error) {
	return
}
func (p *MockDao) GetAuthTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.AuthTopic, err error) {
	return
}
func (p *MockDao) AddAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	return
}
func (p *MockDao) UpdateAuthTopic(c context.Context, node sqalx.Node, item *model.AuthTopic) (err error) {
	return
}
func (p *MockDao) DelAuthTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}

func (p *MockDao) GetTopicsByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.Topic, err error) {
	return
}
func (p *MockDao) GetTopics(c context.Context, node sqalx.Node) (items []*model.Topic, err error) {
	return
}
func (p *MockDao) GetTopicByID(c context.Context, node sqalx.Node, id int64) (item *model.Topic, err error) {
	return
}
func (p *MockDao) GetTopicByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Topic, err error) {
	return
}
func (p *MockDao) AddTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	return
}
func (p *MockDao) UpdateTopic(c context.Context, node sqalx.Node, item *model.Topic) (err error) {
	return
}
func (p *MockDao) DelTopic(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}

func (p *MockDao) GetTopicMembersCount(c context.Context, node sqalx.Node, topicID int64) (count int, err error) {
	return
}
func (p *MockDao) GetTopicMembersPaged(c context.Context, node sqalx.Node, topicID int64, page, pageSize int) (count int, items []*model.TopicMember, err error) {
	return
}
func (p *MockDao) GetTopicMembersByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (items []*model.TopicMember, err error) {
	return
}
func (p *MockDao) GetTopicMembers(c context.Context, node sqalx.Node) (items []*model.TopicMember, err error) {
	return
}
func (p *MockDao) GetTopicMemberByID(c context.Context, node sqalx.Node, id int64) (item *model.TopicMember, err error) {
	return
}
func (p *MockDao) GetTopicMemberByCond(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.TopicMember, err error) {
	return
}
func (p *MockDao) AddTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	return
}
func (p *MockDao) UpdateTopicMember(c context.Context, node sqalx.Node, item *model.TopicMember) (err error) {
	return
}
func (p *MockDao) DelTopicMember(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}

func (p *MockDao) GetTopicFollowRequests(c context.Context, node sqalx.Node, topicID int64, status int) (items []*model.TopicFollowRequest, err error) {
	return
}
func (p *MockDao) GetTopicFollowRequest(c context.Context, node sqalx.Node, topicID, aid int64) (item *model.TopicFollowRequest, err error) {
	return
}
func (p *MockDao) AddTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	return
}
func (p *MockDao) UpdateTopicFollowRequest(c context.Context, node sqalx.Node, item *model.TopicFollowRequest) (err error) {
	return
}
func (p *MockDao) DelTopicFollowRequest(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}

func (p *MockDao) GetAccountByID(c context.Context, node sqalx.Node, id int64) (item *model.Account, err error) {
	return
}
func (p *MockDao) GetAccountTopicSettings(c context.Context, node sqalx.Node, aid int64) (items []*model.AccountTopicSetting, err error) {
	return
}
func (p *MockDao) AddAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	return
}
func (p *MockDao) UpdateAccountTopicSetting(c context.Context, node sqalx.Node, item *model.AccountTopicSetting) (err error) {
	return
}
func (p *MockDao) DelAccountTopicSetting(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}
func (p *MockDao) GetAccountTopicSetting(c context.Context, node sqalx.Node, aid, topicID int64) (item *model.AccountTopicSetting, err error) {
	return
}

func (p *MockDao) AddArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	return
}
func (p *MockDao) GetArticleByID(c context.Context, node sqalx.Node, id int64) (item *model.Article, err error) {
	return
}
func (p *MockDao) UpdateArticle(c context.Context, node sqalx.Node, item *model.Article) (err error) {
	return
}
func (p *MockDao) DelArticle(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}

func (p *MockDao) GetArticleFiles(c context.Context, node sqalx.Node, articleID int64) (items []*model.ArticleFile, err error) {
	return
}
func (p *MockDao) AddArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	return
}
func (p *MockDao) UpdateArticleFile(c context.Context, node sqalx.Node, item *model.ArticleFile) (err error) {
	return
}
func (p *MockDao) DelArticleFile(c context.Context, node sqalx.Node, id int64) (err error) {
	return
}
func (p *MockDao) GetArticleFileByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleFile, err error) {
	return
}

func (p *MockDao) AddArticleHistory(c context.Context, node sqalx.Node, item *model.ArticleHistory) (err error) {
	return
}
func (p *MockDao) GetArticleHistories(c context.Context, node sqalx.Node, articleVersionID int64) (items []*model.ArticleHistory, err error) {
	return
}
func (p *MockDao) GetArticleHistoryByID(c context.Context, node sqalx.Node, id int64) (item *model.ArticleHistory, err error) {
	return
}
func (p *MockDao) GetArticleHistoriesMaxSeq(c context.Context, node sqalx.Node, articleVersionID int64) (seq int, err error) {
	return
}

func (p *MockDao) AddAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error) {
	return
}
func (p *MockDao) UpdateAccountArticleAttr(c context.Context, node sqalx.Node, item *model.AccountArticleAttr) (err error) {
	return
}
func (p *MockDao) DelAccountArticleAttr(c context.Context, node sqalx.Node, id int64) (err error) {

	return
}
func (p *MockDao) GetAccountArticleAttr(c context.Context, node sqalx.Node, aid, articleID int64) (item *model.AccountArticleAttr, err error) {
	return
}
func (p *MockDao) GetArticleFavCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error) {
	return
}
func (p *MockDao) GetArticleLikeCount(c context.Context, node sqalx.Node, articleID int64) (count int, err error) {
	return
}

func (p *MockDao) GetLocaleByCondition(c context.Context, node sqalx.Node, cond map[string]interface{}) (item *model.Locale, err error) {
	return
}

func (p *MockDao) SetTopicCache(c context.Context, m *model.TopicResp) (err error) {

	return
}
func (p *MockDao) TopicCache(c context.Context, topicID int64) (m *model.TopicResp, err error) {
	return
}
func (p *MockDao) DelTopicCache(c context.Context, topicID int64) (err error) {
	return
}

func (p *MockDao) SetTopicCatalogCache(c context.Context, topicID int64, m []*model.TopicLevel1Catalog) (err error) {
	return
}
func (p *MockDao) TopicCatalogCache(c context.Context, topicID int64) (m []*model.TopicLevel1Catalog, err error) {
	return
}
func (p *MockDao) DelTopicCatalogCache(c context.Context, topicID int64) (err error) {
	return
}

func (p *MockDao) SetTopicMembersCache(c context.Context, topicID int64, count, page, pageSize int, data []*model.TopicMember) (err error) {
	return
}
func (p *MockDao) TopicMembersCache(c context.Context, topicID int64, page, pageSize int) (count int, data []*model.TopicMember, err error) {
	return
}
func (p *MockDao) DelTopicMembersCache(c context.Context, topicID int64) (err error) {
	return
}

func (p *MockDao) SetAuthTopicsCache(c context.Context, topicID int64, m []*model.AuthTopic) (err error) {
	return
}
func (p *MockDao) AuthTopicsCache(c context.Context, topicID int64) (m []*model.AuthTopic, err error) {
	return
}
func (p *MockDao) DelAuthTopicsCache(c context.Context, topicID int64) (err error) {
	return
}

func (p *MockDao) SetArticleCache(c context.Context, m *model.ArticleResp) (err error) {
	return
}
func (p *MockDao) ArticleCache(c context.Context, articleID int64) (m *model.ArticleResp, err error) {
	return
}
func (p *MockDao) DelArticleCache(c context.Context, articleID int64) (err error) {
	return
}

func (p *MockDao) SetArticleFileCache(c context.Context, articleID int64, m []*model.ArticleFileResp) (err error) {
	return
}
func (p *MockDao) ArticleFileCache(c context.Context, articleID int64) (m []*model.ArticleFileResp, err error) {
	return
}
func (p *MockDao) DelArticleFileCache(c context.Context, articleID int64) (err error) {
	return
}

func (p *MockDao) SetArticleHistoryCache(c context.Context, articleVersionID int64, m []*model.ArticleHistoryResp) (err error) {
	return
}
func (p *MockDao) ArticleHistoryCache(c context.Context, articleVersionID int64) (m []*model.ArticleHistoryResp, err error) {
	return
}
func (p *MockDao) DelArticleHistoryCache(c context.Context, articleVersionID int64) (err error) {
	return
}

func (p *MockDao) SetAccountCache(c context.Context, m *model.Account) (err error) {
	return
}
func (p *MockDao) AccountCache(c context.Context, accountID int64) (m *model.Account, err error) {
	return
}
func (p *MockDao) DelAccountCache(c context.Context, accountID int64) (err error) {
	return
}

func (p *MockDao) SetAccountTopicSettingCache(c context.Context, m *model.AccountTopicSetting) (err error) {
	return
}
func (p *MockDao) AccountTopicSettingCache(c context.Context, aid, topicID int64) (m *model.AccountTopicSetting, err error) {
	return
}
func (p *MockDao) DelAccountTopicSettingCache(c context.Context, aid, topicID int64) (err error) {
	return
}

func (p *MockDao) Ping(c context.Context) (err error) {
	return
}
func (p *MockDao) Close() {
	return
}
func (p *MockDao) DB() sqalx.Node {
	return nil
}
