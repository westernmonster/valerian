package def

const (
	ActionTypeCreateDiscussion = "MEMBER_CREATE_DISCUSSION"
	ActionTypeUpdateDiscussion = "MEMBER_UPDATE_DISCUSSION"
	ActionTypeDeleteDiscussion = "MEMBER_DELETE_DISCUSSION"
	ActionTypeLikeDiscussion   = "MEMBER_LIKE_DISCUSSION"
	ActionTypeFavDiscussion    = "MEMBER_FAV_DISCUSSION"

	ActionTypeCreateArticle = "MEMBER_CREATE_ARTICLE"
	ActionTypeUpdateArticle = "MEMBER_UPDATE_ARTICLE"
	ActionTypeDeleteArticle = "MEMBER_DELETE_ARTICLE"
	ActionTypeLikeArticle   = "MEMBER_LIKE_ARTICLE"
	ActionTypeFavArticle    = "MEMBER_FAV_ARTICLE"

	ActionTypeCreateTopic = "MEMBER_CREATE_TOPIC"
	ActionTypeDeleteTopic = "MEMBER_DELETE_TOPIC"
	ActionTypeFollowTopic = "MEMBER_FOLLOW_TOPIC"
	ActionTypeUpdateTopic = "MEMBER_UPDATE_TOPIC"

	ActionTypeCreateRevise = "MEMBER_CREATE_REVISE"
	ActionTypeUpdateRevise = "MEMBER_UPDATE_REVISE"
	ActionTypeDeleteRevise = "MEMBER_DELETE_REVISE"
	ActionTypeLikeRevise   = "MEMBER_LIKE_REVISE"
	ActionTypeFavRevise    = "MEMBER_FAV_REVISE"

	ActionTypeFollowMember = "MEMBER_FOLLOW_MEMBER"

	ActionTypeTopicTaxonomyCatalogAdded   = "MEMBER_ADD_TOPIC_TAXONOMY"
	ActionTypeTopicTaxonomyCatalogRenamed = "MEMBER_RENAME_TOPIC_TAXONOMY"
	ActionTypeTopicTaxonomyCatalogDeleted = "MEMBER_DELETE_TOPIC_TAXONOMY"
	ActionTypeTopicTaxonomyCatalogMoved   = "MEMBER_MOVE_TOPIC_TAXONOMY"

	ActionTypeCommentArticle    = "MEMBER_COMMENT_ARTICLE"
	ActionTypeCommentRevise     = "MEMBER_COMMENT_REVISE"
	ActionTypeCommentDiscussion = "MEMBER_COMMENT_DISCUSSION"
)

const (
	ActionTextCreateDiscussion = "发布了讨论"
	ActionTextUpdateDiscussion = "更新了讨论"
	ActionTextDeleteDiscussion = "删除了讨论"
	ActionTextLikeDiscussion   = "点赞了讨论"
	ActionTextFavDiscussion    = "收藏了讨论"

	ActionTextCreateArticle = "发布了文章"
	ActionTextUpdateArticle = "更新了文章"
	ActionTextDeleteArticle = "删除了文章"
	ActionTextLikeArticle   = "点赞了文章"
	ActionTextFavArticle    = "收藏了文章"

	ActionTextCreateTopic = "创建了话题"
	ActionTextDeleteTopic = "删除了话题"
	ActionTextFollowTopic = "加入了话题"
	ActionTextUpdateTopic = "更新了话题属性"

	ActionTextCreateRevise = "添加了补充"
	ActionTextUpdateRevise = "更新了补充"
	ActionTextDeleteRevise = "删除了补充"
	ActionTextLikeRevise   = "喜欢了补充"
	ActionTextFavRevise    = "收藏了补充"

	ActionTextFollowMember = "关注了用户"

	ActionTextTopicTaxonomyCatalogAdded   = "新增了分类 %s"
	ActionTextTopicTaxonomyCatalogDeleted = "删除了分类 %s"
	ActionTextTopicTaxonomyCatalogRenamed = "重名了分类 %s, 新名称 %s"
	ActionTextTopicTaxonomyCatalogMoved   = "移动了分类 %s"

	ActionTextCommentArticle    = "评论了文章"
	ActionTextCommentRevise     = "评论了补充"
	ActionTextCommentDiscussion = "评论了讨论"
)

const (
	ActorTypeUser = "user"
)

const (
	BusDiscussionAdded   = "discussion.added"
	BusDiscussionUpdated = "discussion.updated"
	BusDiscussionDeleted = "discussion.deleted"
	BusDiscussionLiked   = "discussion.liked"
	BusDiscussionFaved   = "discussion.faved"
	// BusDiscussionViewed  = "discussion.viewed"
	//
	BusAccountAdded   = "account.added"
	BusAccountUpdated = "account.updated"
	BusAccountDeleted = "account.deleted"

	BusArticleAdded   = "article.added"
	BusArticleUpdated = "article.updated"
	BusArticleDeleted = "article.deleted"
	BusArticleLiked   = "article.liked"
	BusArticleFaved   = "article.faved"
	BusArticleViewed  = "article.viewed"

	BusCatalogArticleAdded   = "catalog.article.added"
	BusCatalogArticleDeleted = "catalog.article.deleted"

	BusReviseAdded   = "revise.added"
	BusReviseUpdated = "revise.updated"
	BusReviseDeleted = "revise.deleted"
	BusReviseLiked   = "revise.liked"
	BusReviseFaved   = "revise.faved"

	BusTopicAdded           = "topic.added"
	BusTopicUpdated         = "topic.updated"
	BusTopicFollowed        = "topic.followed"
	BusTopicDeleted         = "topic.deleted"
	BusTopicLeaved          = "topic.leaved"
	BusTopicFaved           = "topic.faved"
	BusTopicFollowRequested = "topic.follow.requested"
	BusTopicFollowApproved  = "topic.follow.approved"
	BusTopicFollowRejected  = "topic.follow.rejected"
	BusTopicInviteSent      = "topic.invite.sent"
	BusTopicViewed          = "topic.viewed"

	BusTopicTaxonomyCatalogAdded   = "topic.taxonomy.catalog.added"
	BusTopicTaxonomyCatalogDeleted = "topic.taxonomy.catalog.deleted"
	BusTopicTaxonomyCatalogRenamed = "topic.taxonomy.catalog.renamed"
	BusTopicTaxonomyCatalogMoved   = "topic.taxonomy.catalog.moved"

	BusMemberFollowed   = "member.followed"
	BusMemberUnfollowed = "member.unfollowed"

	BusDiscussionCommented = "discussion.commented"
	BusArticleCommented    = "article.commented"
	BusReviseCommented     = "revise.commented"

	BusCommentLiked   = "comment.liked"
	BusCommentReplied = "comment.replied"
)

const (
	TargetTypeTopic        = "topic"
	TargetTypeDiscussion   = "discussion"
	TargetTypeRevise       = "revise"
	TargetTypeArticle      = "article"
	TargetTypeMember       = "member"
	TargetTypeTopicCatalog = "catalog"
	TargetTypeComment      = "comment"
)
