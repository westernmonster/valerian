package model

const (
	ActionTypeCreateDiscussion = "MEMBER_CREATE_DISCUSSION"
	ActionTypeUpdateDiscussion = "MEMBER_UPDATE_DISCUSSION"
	ActionTypeDeleteDiscussion = "MEMBER_DELETE_DISCUSSION"

	ActionTypeCreateArticle = "MEMBER_CREATE_ARTICLE"
	ActionTypeUpdateArticle = "MEMBER_UPDATE_ARTICLE"
	ActionTypeDeleteArticle = "MEMBER_DELETE_ARTICLE"
)

const (
	ActionTextCreateDiscussion = "发布了讨论"
	ActionTextUpdateDiscussion = "更新了讨论"
	ActionTextDeleteDiscussion = "删除了讨论"

	ActionTextCreateArticle = "发布了文章"
	ActionTextUpdateArticle = "更新了文章"
	ActionTextDeleteArticle = "删除了文章"
)

const (
	ActorTypeUser = "user"
)

const (
	BusDiscussionAdded   = "discussion.added"
	BusDiscussionDeleted = "discussion.deleted"

	BusArticleAdded   = "article.added"
	BusArticleDeleted = "article.deleted"

	BusReviseAdded   = "revise.added"
	BusReviseDeleted = "revise.deleted"

	BusTopicAdded   = "topic.added"
	BusTopicDeleted = "topic.deleted"

	BusMemberFollowed   = "member.followed"
	BusMemberUnfollowed = "member.unfollowed"

	BusDiscussionLiked     = "discussion.liked"
	BusDiscussionCommented = "discussion.commented"

	BusArticleLiked     = "article.liked"
	BusArticleCommented = "article.commented"

	BusReviseLiked     = "revise.liked"
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
