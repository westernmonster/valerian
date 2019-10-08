package model

const (
	ActionTypeCreateDiscussion = "MEMBER_CREATE_DISCUSS"
	ActionTypeUpdateDiscussion = "MEMBER_UPDATE_DISCUSS"
	ActionTypeDeleteDiscussion = "MEMBER_DELETE_DISCUSS"
)

const (
	ActionTextCreateDiscussion = "发布了讨论"
	ActionTextUpdateDiscussion = "更新了讨论"
	ActionTextDeleteDiscussion = "删除了讨论"
)

const (
	ActorTypeUser = "user"
)

const (
	TargetTypeDiscussion = "discussion"
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
