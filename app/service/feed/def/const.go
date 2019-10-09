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

	ActionTypeCreateRevise = "MEMBER_CREATE_REVISE"
	ActionTypeUpdateRevise = "MEMBER_UPDATE_REVISE"
	ActionTypeDeleteRevise = "MEMBER_DELETE_REVISE"
	ActionTypeLikeRevise   = "MEMBER_LIKE_REVISE"
	ActionTypeFavRevise    = "MEMBER_FAV_REVISE"

	ActionTypeFollowMember = "MEMBER_FOLLOW_MEMBER"
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
	ActionTextFollowTopic = "关注了话题"

	ActionTextCreateRevise = "添加了补充"
	ActionTextUpdateRevise = "更新了补充"
	ActionTextDeleteRevise = "删除了补充"
	ActionTextLikeRevise   = "喜欢了补充"
	ActionTextFavRevise    = "收藏了补充"

	ActionTextFollowMember = "关注了用户"
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

	BusArticleAdded   = "article.added"
	BusArticleUpdated = "article.updated"
	BusArticleDeleted = "article.deleted"
	BusArticleLiked   = "article.liked"
	BusArticleFaved   = "article.faved"

	BusCatalogArticleAdded   = "catalog.article.added"
	BusCatalogArticleDeleted = "catalog.article.deleted"

	BusReviseAdded   = "revise.added"
	BusReviseUpdated = "revise.updated"
	BusReviseDeleted = "revise.deleted"
	BusReviseLiked   = "revise.liked"
	BusReviseFaved   = "revise.faved"

	BusTopicAdded    = "topic.added"
	BusTopicFollowed = "topic.followed"
	BusTopicDeleted  = "topic.deleted"

	BusMemberFollowed   = "member.followed"
	BusMemberUnfollowed = "member.unfollowed"

	BusDiscussionCommented = "discussion.commented"

	BusArticleCommented = "article.commented"

	BusReviseCommented = "revise.commented"

	BusCommentLiked   = "comment.liked"
	BusCommentReplied = "comment.replied"
)

const (
	TargetTypeTopic      = "topic"
	TargetTypeDiscussion = "discussion"
	TargetTypeRevise     = "revise"
	TargetTypeArticle    = "article"
	TargetTypeMember     = "member"
)
