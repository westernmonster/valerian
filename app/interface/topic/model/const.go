package model

const (
	BusDiscussionAdded   = "discussion.added"
	BusDiscussionDeleted = "discussion.deleted"

	BusArticleAdded   = "article.added"
	BusArticleDeleted = "article.deleted"

	BusReviseAdded   = "revise.added"
	BusReviseDeleted = "revise.deleted"

	BusTopicAdded           = "topic.added"
	BusTopicDeleted         = "topic.deleted"
	BusTopicFollowed        = "topic.followed"
	BusTopicLeaved          = "topic.leaved"
	BusTopicInviteSent      = "topic.invite.sent"
	BusTopicFollowRequested = "topic.follow.requested"
	BusTopicFollowRejected  = "topic.follow.rejected"
	BusTopicFollowApproved  = "topic.follow.approved"

	BusMemberFollowed = "member.followed"

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
	TopicCatalogTaxonomy = "taxonomy"
	TopicCatalogArticle  = "article"
	TopicCatalogTestSet  = "test_set"
)

const (
	JoinPermissionMember        = "member"
	JoinPermissionMemberApprove = "member_approve"
	JoinPermissionCertApprove   = "cert_approve"
	JoinPermissionManualAdd     = "manual_add"
)

const (
	ViewPermissionPublic = "public"
	ViewPermissionJoin   = "join"
)

const (
	EditPermissionMember = "member"
	EditPermissionAdmin  = "admin"
)

const (
	CatalogViewTypeSection = "section"
	CatalogViewTypeColumn  = "column"
)

const (
	TopicHomeFeed     = "feed"
	TopicHomeCataglog = "catalog"
	TopicHomeDiscuss  = "discuss"
)

const (
	MemberRoleOwner = "owner"
	MemberRoleAdmin = "admin"
	MemberRoleUser  = "user"
)

const (
	AuthPermissionView      = "view"
	AuthPermissionEdit      = "edit"
	AuthPermissionAdminEdit = "admin_edit"
)

const (
	FollowStatusUnfollowed = 0
	FollowStatusApproving  = 1
	FollowStatusFollowed   = 2
)

const (
	InviteStatusSent = 1
)

const (
	FollowRequestStatusCommited = 0
	FollowRequestStatusApproved = 1
	FollowRequestStatusRejected = 2
)

const (
	ReportTypeSpam       = 1 // 垃圾广告
	ReportTypeCopyRight  = 2 // 涉嫌侵权
	ReportTypeDiscomfort = 3 // 内容引起不适
	ReportTypeIncorrect  = 4 // 内容有误
	ReportTypeOther      = 5 // 其他
)

const (
	TargetTypeDiscussion = "discussion"
)
